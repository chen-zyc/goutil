package btree

// 保存在B树中的关键字，同时也可以保存其他数据
type Key interface {
	CompareTo(other Key) int
}

type BTree struct {
	root          *BTreeNode
	t             int // 最小度数，除根节点外的内部节点至少有t个孩子，至多2t个孩子
	diskReadFunc  func(*BTreeNode)
	diskWriteFunc func(*BTreeNode)
}

type BTreeNode struct {
	keys     []Key        // 关键字
	children []*BTreeNode // 孩子节点
	isLeaf   bool         // 是否是叶子节点
}

func (node BTreeNode) Keys() []Key {
	return node.keys
}

// NewBTree 创建一颗空的B树，只有根节点
// @minDegree 最小度数，除根节点外的内部节点至少有minDegree个孩子，至多minDegreet个孩子
func NewBTree(minDegree int, diskRead, diskWrite func(*BTreeNode)) *BTree {
	t := &BTree{
		t:             minDegree,
		diskReadFunc:  diskRead,
		diskWriteFunc: diskWrite,
	}

	x := t.allocateNode()
	x.isLeaf = true

	t.diskWrite(x)
	t.root = x

	return t
}

// allocateNode 创建一个空节点
func (t *BTree) allocateNode() *BTreeNode {
	return &BTreeNode{
		keys:     make([]Key, 0),        // 关键字最多为2t-1个
		children: make([]*BTreeNode, 0), // 最多2t个孩子
	}
}

// diskWrite 将节点n写入到磁盘
func (t *BTree) diskWrite(n *BTreeNode) {
	if t.diskWriteFunc != nil {
		t.diskWriteFunc(n)
	}
}

// 从磁盘上读取n节点下的child节点
func (t *BTree) diskRead(n *BTreeNode) {
	if t.diskReadFunc != nil {
		t.diskReadFunc(n)
	}
}

/************** 插入关键字 ******************/

// Insert插入key
func (t *BTree) Insert(key Key) {
	r := t.root
	if len(r.keys) == 2*t.t-1 { // 根节点满了
		s := t.allocateNode() // 新的根节点
		t.root = s
		s.isLeaf = false
		s.children = append(s.children, r)
		t.splitChild(s, 0) // 分裂r
		t.insertNotFull(s, key)
	} else {
		t.insertNotFull(r, key)
	}
}

// splitChild分裂x的第i个子节点，x是非满的内部节点，x的children[i]已满，现在要分裂children[i]节点
func (t *BTree) splitChild(x *BTreeNode, i int) {
	z := t.allocateNode() // 分裂出来的节点
	y := x.children[i]    // z将是y的兄弟节点
	z.isLeaf = y.isLeaf
	d := t.t

	// y后半部分关键字分给z
	// y和z各有d-1个关键字，y为keys[0..d-1),z为keys[d,2d-1),keys[d-1]被提到父节点中
	for j := d; j < 2*d-1; j++ {
		z.keys = append(z.keys, y.keys[j])
	}
	upKey := y.keys[d-1] // 将要提升的关键字
	y.keys = y.keys[0 : d-1]

	// 如果y不是叶子，将y后半部分的孩子节点也分给z,分t个
	if !y.isLeaf {
		for j := d; j < 2*d; j++ {
			z.children = append(z.children, y.children[j])
		}
		y.children = y.children[0:d]
	}

	// 将z插入到x.children中
	// y是x.children[i],那么z现在是x.children[i+1]
	x.children = append(x.children, nil)
	for j := len(x.children) - 1; j > i+1; j-- {
		// x有n个关键字，必然有n+1个子结点
		x.children[j] = x.children[j-1]
	}
	x.children[i+1] = z

	// 将提升上来的关键字插入到x.keys中
	// 分裂前y中所有关键字都比x.keys[i]小，分裂后提升上来的关键字也比x.keys[i]小，所以插入到x.keys[i]之前
	x.keys = append(x.keys, nil)
	for j := len(x.keys) - 1; j >= i+1; j-- {
		x.keys[j] = x.keys[j-1]
	}
	x.keys[i] = upKey

	t.diskWrite(y)
	t.diskWrite(z)
	t.diskWrite(x)
}

// insertNotFull: 将k插入到x中，x不满。
func (t *BTree) insertNotFull(x *BTreeNode, k Key) {
	i := len(x.keys) - 1
	if x.isLeaf {
		x.keys = append(x.keys, nil)
		// 从后向前遍历，找到第一个小于或等于k的位置，将k插入到该位置后
		for ; i >= 0 && k.CompareTo(x.keys[i]) < 0; i-- {
			x.keys[i+1] = x.keys[i]
		}
		x.keys[i+1] = k
		t.diskWrite(x)
	} else {
		// 从后向前遍历，找到第一个小于或等于k的位置
		for ; i >= 0 && k.CompareTo(x.keys[i]) < 0; i-- {
		}
		i++
		t.diskRead(x.children[i])
		if len(x.children[i].keys) == 2*t.t-1 { // 满节点
			t.splitChild(x, i)
			// 分裂后x.keys[i]已经被替换成提升上来的那个关键字
			if k.CompareTo(x.keys[i]) > 0 {
				i++
			}
		}
		t.insertNotFull(x.children[i], k) // 尾递归，可优化
	}
}

/************** 删除关键字 ******************/

func (t *BTree) Delete(root *BTreeNode, k Key) {
	if root == nil {
		root = t.root
	}
	firstBig, compResult := root.notLessThan(k)
	// 关键字在root中
	if firstBig >= 0 && compResult == 0 {
		// 情况1：如果关键字在root中，并且root是叶子节点，则从root中删除k
		if root.isLeaf {
			root.deleteKeyAt(firstBig)
			return
		}

		// 情况2：如果关键字在root中，但root不是叶子节点
		// y是关键字k之前的结点,即小于k的最大孩子
		y := root.children[firstBig]
		// z是关键字k之后的结点,即大于k的最小孩子
		z := root.children[firstBig+1]

		// 情况[2.a]：root中前于k的子结点y包含至少t个关键字
		if len(y.keys) >= t.t {
			// 找出k在以y为根的子树中的前驱pre
			pre := t.predecessor(y)
			// 用前驱取代k
			k = pre.keys[len(pre.keys)-1]
			root.keys[firstBig] = k
			// 递归地删除k
			t.Delete(y, k)
			return
		}

		// 情况[2.b]: y少于t个关键字，但z至少t个关键字
		if len(z.keys) >= t.t {
			// 查找后继
			next := t.successor(z)
			// 用后继替换k
			k = next.keys[0]
			root.keys[firstBig] = k
			// 递归删除k
			t.Delete(z, k)
			return
		}

		// 情况[2.c]：y和z都只有t-1个关键字，将k和z中所有关键字合并进y，使得x失去k和指向z的指针
		// 将k关键字合并进y
		y.keys = append(y.keys, k)
		root.deleteKeyAt(firstBig)
		// 将z合并到y
		t.merge(y, z)
		root.deleteChildAt(firstBig + 1)
		// 如果root是树的根结点并且没有关键字了，替换根节点
		if root == t.root && len(root.keys) == 0 {
			t.root = y
		}
		// 将k从y中递归删除
		t.Delete(y, k)
		return
	} else { // 关键字不在 root 中,但在 root 的某个子节点中 或者 不在树中
		// root是叶子结点,找到叶子结点都没有找到k，则k不在树中
		if root.isLeaf {
			return
		}
		t.adjustForTKeys(root, firstBig)

		// 递归删除
		t.Delete(root.children[firstBig], k)
	}
}

// notLessThan 返回this节点中第一个不小于k的关键字所在的下标，并且返回keys[index]和k比较的结果。
// 如果this节点中所有关键字都比k小，返回len(this.keys), -1
func (this *BTreeNode) notLessThan(k Key) (index int, compResult int) {
	for i, n := 0, len(this.keys); i < n; i++ {
		if comp := this.keys[i].CompareTo(k); comp >= 0 {
			return i, comp
		}
	}
	return len(this.keys), -1
}

// deleteKeyAt 删除下标为 index 的关键字
func (this *BTreeNode) deleteKeyAt(index int) {
	len := len(this.keys)
	if len <= 0 || index < 0 || index >= len {
		return
	}
	if index == 0 {
		this.keys = this.keys[1:]
		return
	}
	if index == len-1 {
		this.keys = this.keys[:len-1]
		return
	}
	this.keys = append(this.keys[:index], this.keys[index+1:]...)
}

// deleteChildAt 删除下标为 index 的孩子节点
func (this *BTreeNode) deleteChildAt(index int) {
	len := len(this.children)
	if len <= 0 || index < 0 || index >= len {
		return
	}
	if index == 0 {
		this.children = this.children[1:]
		return
	}
	if index == len-1 {
		this.children = this.children[:len-1]
		return
	}
	this.children = append(this.children[:index], this.children[index+1:]...)
}

// 将关键字插入到下标为i的位置
func (this *BTreeNode) insertKeyAt(k Key, i int) {
	this.keys = append(this.keys, nil)
	for j := len(this.keys) - 1; j > i; j-- {
		this.keys[j] = this.keys[j-1]
	}
	this.keys[i] = k
}

// 将孩子节点 child 插入到 this.children[i]
func (this *BTreeNode) insertChildAt(child *BTreeNode, i int) {
	this.children = append(this.children, nil)
	for j := len(this.children) - 1; j > i; j-- {
		this.children[j] = this.children[j-1]
	}
	this.children[i] = child
}

// 寻找以x为根的子树的前驱
func (t *BTree) predecessor(x *BTreeNode) *BTreeNode {
	for !x.isLeaf {
		x = x.children[len(x.keys)]
	}
	return x
}

// 查找以x为根的子树的后继
func (t *BTree) successor(x *BTreeNode) *BTreeNode {
	for !x.isLeaf {
		x = x.children[0]
	}
	return x
}

// 将z合并到y
func (t *BTree) merge(y, z *BTreeNode) {
	// 将z中所有关键字合并进y
	y.keys = append(y.keys, z.keys...)
	// 如果有孩子，孩子也要合并
	if !y.isLeaf {
		y.children = append(y.children, z.children...)
	}
}

// adjustForTKeys 调整 x.children[i] 使关键字个数至少为 t 个
func (t *BTree) adjustForTKeys(x *BTreeNode, i int) {
	// y是关键字k之前的结点,即小于k的最大孩子
	y := x.children[i]

	if len(y.keys) >= t.t {
		return
	}

	// 情况3：y只有t-1个关键字（不可能比t-1少，否则就不是b树了）

	// 情况3a：y 右侧的兄弟至少t个关键字，从它那借
	if i+1 < len(x.children) && len(x.children[i+1].keys) >= t.t {
		right := x.children[i+1]
		// 将需要上升的关键字从 right 中删除
		upKey := right.keys[0]
		right.keys = right.keys[1:]
		// 下降关键字x.keys[i]
		y.keys = append(y.keys, x.keys[i])
		// 上升right第一个关键字
		x.keys[i] = upKey
		// 如果right不是叶子节点，将孩子节点移动到 y.children 的最后
		if !right.isLeaf {
			y.children = append(y.children, right.children[0])
			right.deleteChildAt(0)
		}
		return
	}

	// 情况3a：y的相邻兄弟x->child[i-1]包含至少t个关键字
	if i-1 >= 0 && len(x.children[i-1].keys) >= t.t {
		left := x.children[i-1]
		// 将需要上升的关键字从 right 中删除
		upKey := left.keys[len(left.keys)-1]
		left.keys = left.keys[:len(left.keys)-1]
		// 下降关键字x.keys[i-1]
		y.insertKeyAt(x.keys[i-1], 0)
		// 上升left最后一个关键字
		x.keys[i-1] = upKey
		// 如果left不是叶子节点，将孩子节点移动到 y.children 的开始
		if !left.isLeaf {
			lastIndex := len(left.children) - 1
			y.insertChildAt(left.children[lastIndex], 0)
			left.deleteChildAt(lastIndex)
		}
		return
	}

	// 情况3b：相邻兄弟节点都没有足够的关键字，和其中一个合并
	// 如果有右兄弟
	if i+1 < len(x.children) {
		// 下降关键字
		y.keys = append(y.keys, x.keys[i])
		x.deleteKeyAt(i)
		// 合并
		t.merge(y, x.children[i+1])
		x.deleteChildAt(i + 1)
		// 如果x是根结点并且没有关键字了，删除根节点
		if x == t.root && len(x.keys) == 0 {
			t.root = y
		}
		return
	}
	// 如果有左兄弟
	if i-1 >= 0 {
		// 下降关键字
		y.insertKeyAt(x.keys[i-1], 0)
		x.deleteKeyAt(i - 1)
		// 合并
		t.merge(y, x.children[i-1])
		x.deleteChildAt(i - 1)
		// 如果x是根结点并且没有关键字了，删除根节点
		if x == t.root && len(x.keys) == 0 {
			t.root = y
		}
		return
	}
}

/************** 搜索 ****************/

// 在以x为根节点的子树中搜索
// @x 若x为nil,从根开始搜索
// @return y.keys[index]即要搜索的结果
func (t *BTree) Search(x *BTreeNode, key Key) (y *BTreeNode, index int) {
	if x == nil {
		x = t.root
	}
	i := 0
	for i < len(x.keys) && key.CompareTo(x.keys[i]) > 0 {
		i++
	}
	if i < len(x.keys) && key.CompareTo(x.keys[i]) == 0 {
		return x, i
	}
	if x.isLeaf {
		return nil, -1
	}
	nextNode := x.children[i] // 如果i==n，x是有children[n]孩子的
	t.diskRead(nextNode)
	return t.Search(nextNode, key)
}



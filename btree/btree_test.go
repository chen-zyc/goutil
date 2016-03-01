package btree

import (
	"bytes"
	"fmt"
	"testing"
)

/************** 实现util.TreeNodePrinter接口，打印树到字符串 ******************/

func (node *BTreeNode) String() string {
	s := new(bytes.Buffer)
	n := len(node.keys)
	for i := 0; i < n-1; i++ {
		s.WriteString(fmt.Sprintf("%v, ", node.keys[i]))
	}
	s.WriteString(fmt.Sprintf("%v", node.keys[n-1]))
	return s.String()
}

func (n *BTreeNode) Children() []interface{} {
	leng := len(n.children)
	chil := make([]interface{}, leng)
	for i := 0; i < leng; i++ {
		chil[i] = n.children[i]
	}
	return chil
}

func TestSearch(t *testing.T) {
	node, index := bTree.Search(nil, &key{"C"})
	if node == nil || index < 0 {
		t.Error("没有找到关键字C")
	}
	//	t.Logf("C是第%d个关键字", index)

	node, index = bTree.Search(nil, &key{"H"})
	if node != nil || index >= 0 {
		t.Errorf("不应该在%d位置上找到关键字H", index)
	}
}

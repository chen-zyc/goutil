package printer

import (
	"fmt"
	"strconv"
	"testing"
)

type treeNode struct {
	data     string
	children []interface{}
}

func newTreeNode(data string) *treeNode {
	return &treeNode{
		data:     data,
		children: make([]interface{}, 0, 4),
	}
}

func (n treeNode) String() string {
	return n.data
}

func (n treeNode) Children() []interface{} {
	return n.children
}

func buildTestTree() (root *treeNode) {
	root = newTreeNode("root")
	for i := 0; i < 5; i++ {
		c := newTreeNode("child" + strconv.Itoa(i))
		root.children = append(root.children, c)

		for j := 0; j < 3; j++ {
			cc := newTreeNode(fmt.Sprintf("child%d-%d", i, j))
			c.children = append(c.children, cc)

			for k := 0; k < 4; k++ {
				ccc := newTreeNode(fmt.Sprintf("child%d-%d-%d", i, j, k))
				cc.children = append(cc.children, ccc)
			}
		}
	}

	return root
}

func TestTreePrinter(t *testing.T) {
	root := buildTestTree()

	t.Log("the whole tree is:\n", TreeString(root, "\n"))

	for _, i := range []int{-1, 0, 1, 3} {
		t.Logf("tree[depth:%d]:\n%s", i, TreeStringDepth(root, "\n", i))
	}
}

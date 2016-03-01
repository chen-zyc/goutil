package printer

import (
	"bytes"
)

// 树节点打印接口
type TreeNodePrinter interface {
	// 显示文本
	String() string
	// 子节点
	Children() []interface{}
}

func PrintTree(n TreeNodePrinter) []string {
	dataThisNode := n.String()
	var treeSlice []string
	treeSlice = append(treeSlice, dataThisNode)

	children := n.Children()
	for i, c := range children {
		var child TreeNodePrinter
		var ok bool
		if child, ok = c.(TreeNodePrinter); !ok {
			continue
		}
		childSlice := PrintTree(child)

		// 是否还有兄弟结点
		hasNextChild := false
		if i < len(children)-1 {
			hasNextChild = true
		}

		prefix := "|-- "
		// 先输出孩子结点的数据
		if hasNextChild {
			treeSlice = append(treeSlice, prefix+childSlice[0])
		} else {
			treeSlice = append(treeSlice, "`-- "+childSlice[0])
		}

		// 如果孩子结点还有孩子结点
		if len(childSlice) > 1 {
			// 修正前缀
			if hasNextChild {
				prefix = "|   "
			} else {
				prefix = "    "
			}
			// 输出孩子的孩子结点
			for i := 1; i < len(childSlice); i++ {
				treeSlice = append(treeSlice, prefix+childSlice[i])
			}
		}
	}

	return treeSlice
}

func TreeString(n TreeNodePrinter, lineBreak string) string {
	treeSlice := PrintTree(n)
	var s bytes.Buffer
	for _, v := range treeSlice {
		s.WriteString(v)
		s.WriteString(lineBreak)
	}
	return s.String()
}

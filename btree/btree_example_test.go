package btree

import (
	"bytes"
	"fmt"
	"github.com/zhangyuchen0411/goutil/printer"
)

type key struct {
	k string
}

func (k key) CompareTo(other Key) int {
	return bytes.Compare([]byte(k.k), []byte(other.(*key).k))
}

func (k key) String() string {
	return k.k
}

var bTree *BTree

func init() {
	bTree = NewBTree(3, nil, nil)
	bTree.Insert(&key{"A"})
	bTree.Insert(&key{"C"})
	bTree.Insert(&key{"G"})
	bTree.Insert(&key{"J"})
	bTree.Insert(&key{"K"})
	bTree.Insert(&key{"D"})
	bTree.Insert(&key{"E"})
	bTree.Insert(&key{"M"})
	bTree.Insert(&key{"N"})
	bTree.Insert(&key{"O"})
	bTree.Insert(&key{"P"})
	bTree.Insert(&key{"R"})
	bTree.Insert(&key{"S"})
	bTree.Insert(&key{"X"})
	bTree.Insert(&key{"Y"})
	bTree.Insert(&key{"Z"})
	bTree.Insert(&key{"T"})
	bTree.Insert(&key{"U"})
	bTree.Insert(&key{"V"})
}

// 测试插入，选择的例子见《算法导论》P285

func ExampleInsert() {
	str := printer.TreeString(bTree.root, "\n")
	fmt.Println(str)
	// Output:
	// G, M, P, X
	// |-- A, C, D, E
	// |-- J, K
	// |-- N, O
	// |-- R, S, T, U, V
	// `-- Y, Z
}

func ExampleInsertB() {
	bTree.Insert(&key{"B"})
	fmt.Println(printer.TreeString(bTree.root, "\n"))
	// Output:
	// G, M, P, X
	// |-- A, B, C, D, E
	// |-- J, K
	// |-- N, O
	// |-- R, S, T, U, V
	// `-- Y, Z
}

func ExampleInsertQ() {
	bTree.Insert(&key{"Q"})
	fmt.Println(printer.TreeString(bTree.root, "\n"))
	// Output:
	// G, M, P, T, X
	// |-- A, B, C, D, E
	// |-- J, K
	// |-- N, O
	// |-- Q, R, S
	// |-- U, V
	// `-- Y, Z
}

func ExampleInsertL() {
	bTree.Insert(&key{"L"})
	fmt.Println(printer.TreeString(bTree.root, "\n"))
	// Output:
	// P
	// |-- G, M
	// |   |-- A, B, C, D, E
	// |   |-- J, K, L
	// |   `-- N, O
	// `-- T, X
	//     |-- Q, R, S
	//     |-- U, V
	//     `-- Y, Z
}

func ExampleInsertF() {
	bTree.Insert(&key{"F"})
	fmt.Println(printer.TreeString(bTree.root, "\n"))
	// Output:
	// P
	// |-- C, G, M
	// |   |-- A, B
	// |   |-- D, E, F
	// |   |-- J, K, L
	// |   `-- N, O
	// `-- T, X
	//     |-- Q, R, S
	//     |-- U, V
	//     `-- Y, Z
}

func ExampleDeleteF() {
	bTree.Delete(nil, &key{"F"})
	fmt.Println(printer.TreeString(bTree.root, "\n"))
	// Output:
	// P
	// |-- C, G, M
	// |   |-- A, B
	// |   |-- D, E
	// |   |-- J, K, L
	// |   `-- N, O
	// `-- T, X
	//     |-- Q, R, S
	//     |-- U, V
	//     `-- Y, Z
}

func ExampleDeleteM() {
	bTree.Delete(nil, &key{"M"})
	fmt.Println(printer.TreeString(bTree.root, "\n"))
	// Output:
	// P
	// |-- C, G, L
	// |   |-- A, B
	// |   |-- D, E
	// |   |-- J, K
	// |   `-- N, O
	// `-- T, X
	//     |-- Q, R, S
	//     |-- U, V
	//     `-- Y, Z
}

// 测试 情况2.b
//func ExampleDeleteG() {
//	btree.Delete(nil, &key{"G"})
//	fmt.Println(util.TreeString(btree.root, "\n"))
//	// Output:
//	// P
//	// |-- C, J, M
//	// |   |-- A, B
//	// |   |-- D, E
//	// |   |-- K, L
//	// |   `-- N, O
//	// `-- T, X
//	//     |-- Q, R, S
//	//     |-- U, V
//	//     `-- Y, Z
//}

// 情况 2.c
func ExampleDeleteG() {
	bTree.Delete(nil, &key{"G"})
	fmt.Println(printer.TreeString(bTree.root, "\n"))
	// Output:
	// P
	// |-- C, L
	// |   |-- A, B
	// |   |-- D, E, J, K
	// |   `-- N, O
	// `-- T, X
	//     |-- Q, R, S
	//     |-- U, V
	//     `-- Y, Z
}

// 情况 3.b
func ExampleDeleteD() {
	bTree.Delete(nil, &key{"D"})
	fmt.Println(printer.TreeString(bTree.root, "\n"))
	// Output:
	// C, L, P, T, X
	// |-- A, B
	// |-- E, J, K
	// |-- N, O
	// |-- Q, R, S
	// |-- U, V
	// `-- Y, Z
}

//// 情况 3.b
//func ExampleDeleteQ() {
//	btree.Delete(nil, &key{"Q"})
//	fmt.Println(util.TreeString(btree.root, "\n"))
//	// Output:
//	// C, L, P, T, X
//	// |-- A, B
//	// |-- E, J, K
//	// |-- N, O
//	// |-- R, S
//	// |-- U, V
//	// `-- Y, Z
//}

// 情况 3.a
func ExampleDeleteB() {
	bTree.Delete(nil, &key{"B"})
	fmt.Println(printer.TreeString(bTree.root, "\n"))
	// Output:
	// E, L, P, T, X
	// |-- A, C
	// |-- J, K
	// |-- N, O
	// |-- Q, R, S
	// |-- U, V
	// `-- Y, Z
}

//// 情况 3.a
//func ExampleDeleteU() {
//	btree.Delete(nil, &key{"U"})
//	fmt.Println(util.TreeString(btree.root, "\n"))
//	// Output:
//	// C, L, P, S, X
//	// |-- A, B
//	// |-- E, J, K
//	// |-- N, O
//	// |-- Q, R
//	// |-- T, V
//	// `-- Y, Z
//}

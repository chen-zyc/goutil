package sorter

import (
	"fmt"
)

func ExampleSorterUse() {
	data := []int{3, 9, 1, 0, 6}
	label := []string{"A", "B", "C", "D", "E"}

	Sorter{
		LenFunc:  func() int { return len(data) },
		LessFunc: func(i, j int) bool { return data[i] < data[j] },
		SwapFunc: func(i, j int) {
			data[i], data[j] = data[j], data[i]
			label[i], label[j] = label[j], label[i]
		},
	}.Sort()

	for i, d := range data {
		fmt.Println(label[i], "=>", d)
	}
	// Output:
	// D => 0
	// C => 1
	// A => 3
	// E => 6
	// B => 9
}

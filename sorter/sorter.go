package sorter

import "sort"

type Sorter struct {
	LenFunc  func() int
	LessFunc func(i, j int) bool
	SwapFunc func(i, j int)
}

func (s Sorter) Sort() {
	sort.Sort(s)
}

func (s Sorter) Len() int {
	return s.LenFunc()
}

func (s Sorter) Less(i, j int) bool {
	return s.LessFunc(i, j)
}

func (s Sorter) Swap(i, j int) {
	s.SwapFunc(i, j)
}


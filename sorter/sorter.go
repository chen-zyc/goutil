package sorter

import "sort"

type Sorter struct {
	Len  func() int
	Less func(i, j int) bool
	Swap func(i, j int)
}

func (s Sorter) Sort() {
	sort.Sort(s)
}

func (s Sorter) Len() int {
	return s.Len()
}

func (s Sorter) Less(i, j int) bool {
	return s.Less(i, j)
}

func (s Sorter) Swap(i, j int) {
	s.Swap(i, j)
}

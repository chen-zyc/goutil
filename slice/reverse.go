package slice

type Swapper interface {
	Len() int
	Swap(i, j int)
}

func Reverse(swapper Swapper) {
	for i, j := 0, swapper.Len()-1; i < j; i, j = i+1, j-1 {
		swapper.Swap(i, j)
	}
}

type IntSwapper []int

func (self IntSwapper) Len() int { return len(self) }

func (self IntSwapper) Swap(i, j int) { self[i], self[j] = self[j], self[i] }

func ReverseInt(slice []int) { Reverse(IntSwapper(slice)) }

type StringSwapper []string

func (self StringSwapper) Len() int { return len(self) }

func (self StringSwapper) Swap(i, j int) { self[i], self[j] = self[j], self[i] }

func ReverseString(slice []string) { Reverse(StringSwapper(slice)) }

type Float64Swapper []float64

func (self Float64Swapper) Len() int { return len(self) }

func (self Float64Swapper) Swap(i, j int) { self[i], self[j] = self[j], self[i] }

func ReverseFloat64(slice []float64) { Reverse(Float64Swapper(slice)) }

package slice

func ShiftLeftInt(slice []int, n int) {
	ReverseInt(slice[:n])
	ReverseInt(slice[n:])
	ReverseInt(slice)
}

func ShiftLeftString(slice []string, n int) {
	ReverseString(slice[:n])
	ReverseString(slice[n:])
	ReverseString(slice)
}

func ShiftLeftFloat64(slice []float64, n int) {
	ReverseFloat64(slice[:n])
	ReverseFloat64(slice[n:])
	ReverseFloat64(slice)
}

func ShiftRightInt(slice []int, n int) {
	ReverseInt(slice)
	ReverseInt(slice[:n])
	ReverseInt(slice[n:])
}

func ShiftRightString(slice []string, n int) {
	ReverseString(slice)
	ReverseString(slice[:n])
	ReverseString(slice[n:])
}

func ShiftRightFloat64(slice []float64, n int) {
	ReverseFloat64(slice)
	ReverseFloat64(slice[:n])
	ReverseFloat64(slice[n:])
}

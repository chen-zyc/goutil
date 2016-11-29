package binary

func IsPowerOf2(x int64) bool {
	return x > 0 && (x&(x-1) == 0)
}

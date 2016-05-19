package lcs

// LCS 返回X与Y最长公共子序列
func LCS(X, Y []rune) []rune { // 去掉辅助数组b的版本
	if len(X) == 0 || len(Y) == 0 {
		return []rune{}
	}
	// c[i,j]用于记录x[:i+1]和x[:y+1]的最长公共子序列的长度,c[0][0]为0,表示两个空串的lcs为0
	c := make([][]int, len(X)+1)
	for i := 0; i < len(c); i++ {
		c[i] = make([]int, len(Y)+1)
	}

	for i, x := range X {
		for j, y := range Y {
			if x == y { // 找到一个公共元素
				c[i+1][j+1] = c[i][j] + 1
			} else if c[i][j+1] >= c[i+1][j] { // 上面那个比左面那个的lcs长,选择上面那个
				c[i+1][j+1] = c[i][j+1]
			} else {
				c[i+1][j+1] = c[i+1][j]
			}
		}
	}

	i := len(c) - 1
	j := len(c[0]) - 1
	seqIndex := c[i][j] - 1
	seq := make([]rune, seqIndex+1)
	var (
		upLeft int
		left   int
		up     int
		max    int
	)

	for {
		if i == 0 || j == 0 {
			break
		}
		upLeft, left, up = c[i-1][j-1], c[i][j-1], c[i-1][j]
		max = upLeft
		if up > upLeft {
			max = up
		}
		if left > upLeft {
			max = left
		}
		if max == upLeft {
			seq[seqIndex] = X[i-1]
			seqIndex--
			i--
			j--
		} else if max == up {
			i--
		} else {
			j--
		}
	}
	return seq
}

// 这个是帮助理解算法的
func _LCS(X, Y []rune) []rune {
	// c[i,j]用于记录x[:i+1]和x[:y+1]的最长公共子序列的长度,c[0][0]为0,表示两个空串的lcs为0
	c := make([][]int, len(X)+1)
	for i := 0; i < len(c); i++ {
		c[i] = make([]int, len(Y)+1)
	}
	// b[i,j]用于记录c[i,j]时下一步应该回溯的方向
	b := make([][]string, len(X)+1)
	for i := 0; i < len(b); i++ {
		b[i] = make([]string, len(Y)+1)
	}

	for i, x := range X {
		for j, y := range Y {
			if x == y { // 找到一个公共元素
				c[i+1][j+1] = c[i][j] + 1
				b[i+1][j+1] = "↖" // 下一个应该回溯到左上
			} else if c[i][j+1] >= c[i+1][j] { // 上面那个比左面那个的lcs长,选择上面那个
				c[i+1][j+1] = c[i][j+1]
				b[i+1][j+1] = "↑"
			} else {
				c[i+1][j+1] = c[i+1][j]
				b[i+1][j+1] = "←"
			}
		}
	}

	//for _, row := range b {
	//	for _, col := range row {
	//		fmt.Print(col, " ")
	//	}
	//	fmt.Println()
	//}

	i := len(c) - 1
	j := len(c[0]) - 1
	seqIndex := c[i][j] - 1
	seq := make([]rune, seqIndex+1)

	for {
		if i == 0 || j == 0 {
			break
		}
		switch b[i][j] {
		case "↖":
			seq[seqIndex] = X[i-1]
			seqIndex--
			i--
			j--
		case "↑":
			i--
		default:
			j--
		}
	}
	return seq
}

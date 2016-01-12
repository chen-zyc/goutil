# goutil

## sorter

有时候对一个数组进行排序时需要相应的改变另一个数组元素的顺序，但这两个数组又不在同一个结构体中。

比如有一下两个数组：

```go
	data := []int{3, 9, 1, 0, 6}
	label := []string{"A", "B", "C", "D", "E"}
```

对data排序时也希望改变label的顺序，此时可以这样：

```go
	sorter.Sorter{
		LenFunc:  func() int { return len(data) },
		LessFunc: func(i, j int) bool { return data[i] < data[j] },
		SwapFunc: func(i, j int) {
			data[i], data[j] = data[j], data[i]
			label[i], label[j] = label[j], label[i]
		},
	}.Sort()
```


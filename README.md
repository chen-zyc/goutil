# goutil

## sorter

有时候对一个数组进行排序时需要相应的改变另一个数组元素的顺序，但这两个数组又不在同一个结构体中。

比如有以下两个数组：

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

## reg.FindStringSubmatchMap

假如正则表达式是包含命名分组的（比如 `abc_(?P<n1>\d{4})`）, `FindStringSubmatchMap` 可以返回一个map，key就是正则中的命名，value就是匹配到的值。

```go
	reg, err := regexp.Compile(`abc_(?P<n1>\d{4}), def_(?P<n2>\d{4}), ghi_(?P<n3>\d{4})`)
	if err != nil {
		t.Error(err)
		return
	}
	text := "abc_1234, def_1235, ghi_1236"
	match := FindStringSubmatchMap(reg, text)
```

此时match的值是这样的：

```go
		"n1": "1234",
		"n2": "1235",
		"n3": "1236",
```

## slice

### Reverse

反转切片的值.

比如： `slice := []int{3, 9, 0, -3, -5}`

反转： `ReverseInt(slice)`

结果： `{-5, -3, 0, 9, 3}`

### Shift

左移/右移切片。

比如： `slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}`

反转： `ShiftLeftInt(slice, 3)`

结果： `{4, 5, 6, 7, 8, 9, 1, 2, 3}`


## printer

打印相关的包。

### TreeString()

打印一棵树，树结点实现了 `TreeNodePrinter` 接口，效果如下：

```
G, M, P, X
|-- A, C, D, E
|-- J, K
|-- N, O
|-- R, S, T, U, V
`-- Y, Z
```

### PrintTree()

返回打印结果的字符串切片，每一个切片的元素表示树的一行。


## btree

b树的实现，详情参考 <http://blog.csdn.net/u013790019/article/details/47456607>



package binary

// IsPowerOf2 判断x是否是2的次幂
func IsPowerOf2(x int64) bool {
	return x > 0 && (x&(x-1) == 0)
}

// BitCount1 通过查找最后bit为是否为1来统计1的个数
func BitCount1(x int64) int {
	c := 0
	for ; x > 0; x >>= 1 {
		c += int(x & 1)
	}
	return c
}

// BitCount 通过清除最右边的1来统计1的个数
func BitCount(x int64) int {
	c := 0
	for ; x > 0; c++ {
		x &= x - 1 // x & (x-1) 会清除最右边的1
	}
	return c
}

// BitCount32 通过相邻bit相加来统计1的个数
func BitCount32(x int32) int {
	x = x&0x55555555 + (x>>1)&0x55555555
	x = x&0x33333333 + (x>>2)&0x33333333
	x = x&0x0f0f0f0f + (x>>4)&0x0f0f0f0f
	x = x&0x00ff00ff + (x>>8)&0x00ff00ff
	x = x&0x0000ffff + (x>>16)&0x0000ffff
	return int(x)
}

// BitCount64 通过相邻bit相加来统计1的个数
func BitCount64(x int64) int {
	x = x&0x5555555555555555 + (x>>1)&0x5555555555555555
	x = x&0x3333333333333333 + (x>>2)&0x3333333333333333
	x = x&0x0f0f0f0f0f0f0f0f + (x>>4)&0x0f0f0f0f0f0f0f0f
	x = x&0x00ff00ff00ff00ff + (x>>8)&0x00ff00ff00ff00ff
	x = x&0x0000ffff0000ffff + (x>>16)&0x0000ffff0000ffff
	x = x&0x00000000ffffffff + (x>>32)&0x00000000ffffffff
	return int(x)
}

package bitmap

import (
	"bytes"
	"fmt"
)

type BitMap32 struct {
	a []uint32
}

func NewBitMap32(size int) *BitMap32 {
	return &BitMap32{
		a: make([]uint32, size),
	}
}

func (bm32 *BitMap32) Set(i uint32) {
	// i >> 5 => i / 32, 落在哪个槽里
	// i & 0x1F => i % 32
	bm32.a[i>>5] |= 1 << (i & 0x1F)
}

func (bm32 *BitMap32) Clear(i uint32) {
	bm32.a[i>>5] &= ^(1 << (i & 0x1F))
}

func (bm32 *BitMap32) In(i uint32) bool {
	return bm32.a[i>>5]&(1<<(i&0x1F)) > 0
}

func (bm32 *BitMap32) String() string {
	buf := bytes.NewBuffer(make([]byte, 0, len(bm32.a)*33))
	for i, x := range bm32.a {
		if i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(fmt.Sprintf("%032b", x))
	}
	return buf.String()
}

type BitMap64 struct {
	a []uint64
}

func NewBitMap64(size int) *BitMap64 {
	return &BitMap64{
		a: make([]uint64, size),
	}
}

func (bm64 *BitMap64) Set(i uint64) {
	// i >> 6 => i / 64, 落在哪个槽里
	// i & 0x3F => i % 64
	bm64.a[i>>6] |= 1 << (i & 0x3F)
}

func (bm64 *BitMap64) Clear(i uint64) {
	bm64.a[i>>6] &= ^(1 << (i & 0x3F))
}

func (bm64 *BitMap64) In(i uint64) bool {
	return bm64.a[i>>6]&(1<<(i&0x3F)) > 0
}

func (bm64 *BitMap64) String() string {
	buf := bytes.NewBuffer(make([]byte, 0, len(bm64.a)*65))
	for i, x := range bm64.a {
		if i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(fmt.Sprintf("%064b", x))
	}
	return buf.String()
}

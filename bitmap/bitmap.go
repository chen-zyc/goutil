package bitmap

import (
	"bytes"
	"fmt"
	"strconv"
)

const intSize = 32 << (^uint(0) >> 63)
const shift = 5 + (^uint(0) >> 63)
const mask = intSize - 1

func init() {
	println(intSize, shift, mask)
}

type BitMap struct {
	a    []uint
	size uint
}

func NewBitMap(size int) *BitMap {
	return &BitMap{
		a:    make([]uint, size),
		size: uint(size),
	}
}

func (bm *BitMap) Set(i uint) {
	bm.a[(i>>shift)%bm.size] |= 1 << (i & mask)
}

func (bm *BitMap) Clear(i uint) {
	bm.a[(i>>shift)%bm.size] &= ^(1 << (i & mask))
}

func (bm *BitMap) In(i uint) bool {
	return bm.a[(i>>shift)%bm.size]&(1<<(i&mask)) > 0
}

func (bm *BitMap) String() string {
	format := "%0" + strconv.Itoa(intSize) + "b"
	buf := bytes.NewBuffer(make([]byte, 0, len(bm.a)*(intSize+1)))
	for i, x := range bm.a {
		if i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(fmt.Sprintf(format, x))
	}
	return buf.String()
}

type BitMap32 struct {
	a    []uint32
	size uint32
}

func NewBitMap32(size int) *BitMap32 {
	return &BitMap32{
		a:    make([]uint32, size),
		size: uint32(size),
	}
}

func (bm32 *BitMap32) Set(i uint32) {
	// i >> 5 => i / 32, 落在哪个槽里
	// i & 0x1F => i % 32
	bm32.a[(i>>5)%bm32.size] |= 1 << (i & 0x1F)
}

func (bm32 *BitMap32) Clear(i uint32) {
	bm32.a[(i>>5)%bm32.size] &= ^(1 << (i & 0x1F))
}

func (bm32 *BitMap32) In(i uint32) bool {
	return bm32.a[(i>>5)%bm32.size]&(1<<(i&0x1F)) > 0
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
	a    []uint64
	size uint64
}

func NewBitMap64(size int) *BitMap64 {
	return &BitMap64{
		a:    make([]uint64, size),
		size: uint64(size),
	}
}

func (bm64 *BitMap64) Set(i uint64) {
	// i >> 6 => i / 64, 落在哪个槽里
	// i & 0x3F => i % 64
	bm64.a[(i>>6)%bm64.size] |= 1 << (i & 0x3F)
}

func (bm64 *BitMap64) Clear(i uint64) {
	bm64.a[(i>>6)%bm64.size] &= ^(1 << (i & 0x3F))
}

func (bm64 *BitMap64) In(i uint64) bool {
	return bm64.a[(i>>6)%bm64.size]&(1<<(i&0x3F)) > 0
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

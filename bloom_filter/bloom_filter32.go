package bloom_filter

import (
	"github.com/zhangyuchen0411/goutil/bitmap"
)

type HashFunc32 func(key string) uint32

type BloomFilter32 struct {
	m         *bitmap.BitMap32
	hashFuncs []HashFunc32
}

func NewBloomFilter32(bits int, hashFuncs []HashFunc32) *BloomFilter32 {
	bits /= 32
	if bits < 1 {
		bits = 1
	}
	return &BloomFilter32{
		m:         bitmap.NewBitMap32(bits),
		hashFuncs: hashFuncs,
	}
}

func (bf32 *BloomFilter32) Set(key string) {
	var hash uint32
	for _, f := range bf32.hashFuncs {
		hash = f(key)
		bf32.m.Set(hash)
	}
}

func (bf32 *BloomFilter32) In(key string) bool {
	for _, f := range bf32.hashFuncs {
		hash := f(key)
		if !bf32.m.In(hash) {
			return false
		}
	}
	return true
}

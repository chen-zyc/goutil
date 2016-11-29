package bloom_filter

import "github.com/zhangyuchen0411/goutil/bitmap"

type HashFunc64 func(key string) uint64

type BloomFilter64 struct {
	m         *bitmap.BitMap64
	hashFuncs []HashFunc64
}

func NewBloomFilter64(bits int, hashFuncs []HashFunc64) *BloomFilter64 {
	bits /= 64
	if bits < 1 {
		bits = 1
	}
	return &BloomFilter64{
		m:         bitmap.NewBitMap64(bits),
		hashFuncs: hashFuncs,
	}
}

func (bf64 *BloomFilter64) Set(key string) {
	var hash uint64
	for _, f := range bf64.hashFuncs {
		hash = f(key)
		bf64.m.Set(hash)
	}
}

func (bf64 *BloomFilter64) In(key string) bool {
	for _, f := range bf64.hashFuncs {
		hash := f(key)
		if !bf64.m.In(hash) {
			return false
		}
	}
	return true
}

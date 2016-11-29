package bloom_filter

import (
	"fmt"
	"github.com/zhangyuchen0411/goutil/hash"
	"math/rand"
	"testing"
	"time"
)

var testKeys []string
var hashFunc64 []HashFunc64
var hashFunc32 []HashFunc32

func init() {
	rand.Seed(time.Now().Unix())
	testKeys = make([]string, 100)
	for i := 0; i < len(testKeys); i++ {
		testKeys[i] = fmt.Sprintf("test-key-%d", rand.Int())
	}
	hashFunc64 = []HashFunc64{
		hash.APHash,
		hash.BKDRHash,
		hash.DEKHash,
		hash.DJBHash,
		hash.ELFHash,
		hash.FNVHash,
		hash.JSHash,
		hash.MYSQLHash,
		hash.PJWHash,
		hash.RSHash,
		hash.SDBMHash,
	}
	hashFunc32 = make([]HashFunc32, len(hashFunc64))
	for i, f := range hashFunc64 {
		hashFunc32[i] = func(key string) uint32 {
			return uint32(f(key) % 0x1F)
		}
	}
}

func TestBloomFilter32(t *testing.T) {
	f := NewBloomFilter32(128, hashFunc32)
	for _, key := range testKeys {
		f.Set(key)
	}
	for _, key := range testKeys {
		if !f.In(key) {
			t.Fatalf("%s not in bloom filter", key)
		}
	}
}

func TestBloomFilter64(t *testing.T) {
	f := NewBloomFilter64(128, hashFunc64)
	for _, key := range testKeys {
		f.Set(key)
	}
	for _, key := range testKeys {
		if !f.In(key) {
			t.Fatalf("%s not in bloom filter", key)
		}
	}
}

func BenchmarkBloomFilter32(b *testing.B) {
	f := NewBloomFilter32(128, hashFunc32)
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("test-key-%d", i)
		f.Set(key)
		f.In(key)
	}
}

func BenchmarkBloomFilter64(b *testing.B) {
	f := NewBloomFilter64(128, hashFunc64)
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("test-key-%d", i)
		f.Set(key)
		f.In(key)
	}
}

package bitmap

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestBitMap32(t *testing.T) {
	n := 1
	randomNums := make([]uint32, 10)
	for i := 0; i < len(randomNums); i++ {
		randomNums[i] = uint32(rand.Intn(n * 32))
	}
	m := NewBitMap32(n)
	for _, n := range randomNums {
		m.Set(n)
	}

	//t.Logf("\n%s", m.String())
	for _, n := range randomNums {
		if !m.In(n) {
			t.Fatalf("%d not in bit map", n)
		}
	}
}

func TestBitMap64(t *testing.T) {
	n := 1
	randomNums := make([]uint64, 10)
	for i := 0; i < len(randomNums); i++ {
		randomNums[i] = uint64(rand.Intn(n * 64))
	}
	m := NewBitMap64(n)
	for _, n := range randomNums {
		m.Set(n)
	}

	//t.Logf("\n%s", m.String())
	for _, n := range randomNums {
		if !m.In(n) {
			t.Fatalf("%d not in bit map", n)
		}
	}
}

func TestBitMap(t *testing.T) {
	n := 2
	randomNums := make([]uint, 10)
	for i := 0; i < len(randomNums); i++ {
		randomNums[i] = uint(rand.Intn(n * intSize))
	}
	m := NewBitMap(n)
	for _, n := range randomNums {
		m.Set(n)
	}

	//t.Logf("\n%s", m.String())
	for _, n := range randomNums {
		if !m.In(n) {
			t.Fatalf("%d not in bit map", n)
		}
	}
}

func BenchmarkBitMap32(b *testing.B) {
	m := NewBitMap32(20)
	for i := 0; i < b.N; i++ {
		m.Set(uint32(i))
	}
}

func BenchmarkBitMap64(b *testing.B) {
	m := NewBitMap64(10)
	for i := 0; i < b.N; i++ {
		m.Set(uint64(i))
	}
}

func BenchmarkBitMap(b *testing.B) {
	m := NewBitMap(10)
	for i := 0; i < b.N; i++ {
		m.Set(uint(i))
	}
}

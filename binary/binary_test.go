package binary

import "testing"

func TestIsPowerOf2(t *testing.T) {
	cases := map[int64]bool{
		0:       false,
		1:       true,
		2:       true,
		3:       false,
		4:       true,
		2 << 50: true,
		-1:      false,
		-2:      false,
	}
	for v, want := range cases {
		if v2 := IsPowerOf2(v); v2 != want {
			t.Fatalf("%d, want %t, got %t", v, want, v2)
		}
	}
}

func TestBitCount1(t *testing.T) {
	cases := map[int64]int{
		0x0:               0,
		0x1:               1,
		0xF:               4,
		0x5:               2,
		0xA:               2,
		0xFFFFFFFF:        32,
		0xFFFFFFFFFFFFFFF: 60,
	}
	for x, want := range cases {
		if act := BitCount1(x); act != want {
			t.Fatalf("%b, want %d, got %d", x, want, act)
		}
	}
}

func BenchmarkBitCount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitCount1(0xFFFFFFFF)
	}
}

func TestBitCount(t *testing.T) {
	cases := map[int64]int{
		0x0:               0,
		0x1:               1,
		0xF:               4,
		0x5:               2,
		0xA:               2,
		0xFFFFFFFF:        32,
		0xFFFFFFFFFFFFFFF: 60,
	}
	for x, want := range cases {
		if act := BitCount(x); act != want {
			t.Fatalf("%b, want %d, got %d", x, want, act)
		}
	}
}

func BenchmarkBitCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitCount(0xFFFFFFFF)
	}
}

func TestBitCount32(t *testing.T) {
	cases := map[int32]int{
		0x0:       0,
		0x1:       1,
		0xF:       4,
		0x5:       2,
		0xA:       2,
		0xFFFFFFF: 28,
	}
	for x, want := range cases {
		if act := BitCount32(x); act != want {
			t.Fatalf("%b, want %d, got %d", x, want, act)
		}
	}
}

func TestBitCount64(t *testing.T) {
	cases := map[int64]int{
		0x0:               0,
		0x1:               1,
		0xF:               4,
		0x5:               2,
		0xA:               2,
		0xFFFFFFF:         28,
		0xFFFFFFFFFFFFFFF: 60,
	}
	for x, want := range cases {
		if act := BitCount64(x); act != want {
			t.Fatalf("%b, want %d, got %d", x, want, act)
		}
	}
}

func BenchmarkBitCount32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitCount32(0xFFFFFFF)
	}
}

func BenchmarkBitCount64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitCount64(0xFFFFFFF)
	}
}

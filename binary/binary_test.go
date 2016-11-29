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

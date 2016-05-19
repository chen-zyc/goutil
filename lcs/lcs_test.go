package lcs

import "testing"

func TestLCS(t *testing.T) {
	tests := []struct {
		X    string
		Y    string
		want string
	}{
		{"ABCBDAB", "BDCABA", "BCBA"},
		{"10010101", "010110110", "101010"},
	}

	for _, te := range tests {
		result := LCS([]rune(te.X), []rune(te.Y))
		if string(result) != te.want {
			t.Errorf("LCS(%s, %s) ret %s, want %s", te.X, te.Y, string(result), te.want)
		}
	}
}

func BenchmarkLCS(b *testing.B) {
	x := []rune("ABCBDAB")
	y := []rune("BDCABA")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LCS(x, y)
	}
}

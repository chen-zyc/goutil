package hash

import "testing"

var testKey = "test-key"

func BenchmarkRSHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RSHash(testKey)
	}
}

func BenchmarkJSHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JSHash(testKey)
	}
}

func BenchmarkPJWHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PJWHash(testKey)
	}
}

func BenchmarkELFHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ELFHash(testKey)
	}
}

func BenchmarkBKDRHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BKDRHash(testKey)
	}
}

func BenchmarkSDBMHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SDBMHash(testKey)
	}
}

func BenchmarkDJBHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DJBHash(testKey)
	}
}

func BenchmarkDEKHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DEKHash(testKey)
	}
}

func BenchmarkAPHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		APHash(testKey)
	}
}

func BenchmarkFNVHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FNVHash(testKey)
	}
}

func BenchmarkMYSQLHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MYSQLHash(testKey)
	}
}

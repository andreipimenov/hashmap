package hashmap

import (
	"testing"
)

// BenchmarkFunc16 - test benchmarks for hash func with block size 16
func BenchmarkFunc16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hashFunc(16, Key(i))
	}
}

// BenchmarkFunc64 - test benchmarks for hash func with block size 64
func BenchmarkFunc64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hashFunc(64, Key(i))
	}
}

// BenchmarkFunc128 - test benchmarks for hash func with block size 128
func BenchmarkFunc128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hashFunc(128, Key(i))
	}
}

// BenchmarkFunc1024 - test benchmarks for hash func with block size 1024
func BenchmarkFunc1024(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hashFunc(1024, Key(i))
	}
}

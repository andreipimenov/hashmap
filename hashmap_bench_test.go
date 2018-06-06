package hashmap

import (
	"testing"
)

//BenchmarkSetMap - test benchmarks for operation Set on map type
func BenchmarkSetMap(b *testing.B) {
	m := map[Key]interface{}{}
	for i := 0; i < b.N; i++ {
		m[Key(i)] = i
	}
}

//BenchmarkGetMap - test benchmarks for operation Get on map type
func BenchmarkGetMap(b *testing.B) {
	m := map[Key]interface{}{}
	for i := 0; i < b.N; i++ {
		_ = m[Key(i)]
	}
}

//BenchmarkUnsetMap - test benchmarks for operation Unset on map type
func BenchmarkUnsetMap(b *testing.B) {
	m := map[Key]interface{}{}
	for i := 0; i < b.N; i++ {
		delete(m, Key(i))
	}
}

// BenchmarkCountMap - test benchmarks for operation Count on map type
func BenchmarkCountMap(b *testing.B) {
	m := map[Key]interface{}{}
	for i := 0; i < b.N; i++ {
		_ = len(m)
	}
}

// BenchmarkCount16 - test benchmarks for operation Count on hash map with block size 16
func BenchmarkCount16(b *testing.B) {
	hm, _ := NewHashMap(16)
	for i := 0; i < b.N; i++ {
		hm.Count()
	}
}

// BenchmarkCount64 - test benchmarks for operation Count on hash map with block size 64
func BenchmarkCount64(b *testing.B) {
	hm, _ := NewHashMap(64)
	for i := 0; i < b.N; i++ {
		hm.Count()
	}
}

// BenchmarkCount128 - test benchmarks for operation Count on hash map with block size 128
func BenchmarkCount128(b *testing.B) {
	hm, _ := NewHashMap(128)
	for i := 0; i < b.N; i++ {
		hm.Count()
	}
}

// BenchmarkCount1024 - test benchmarks for operation Count on hash map with block size 1024
func BenchmarkCount1024(b *testing.B) {
	hm, _ := NewHashMap(1024)
	for i := 0; i < b.N; i++ {
		hm.Count()
	}
}

// BenchmarkSet16 - test benchmarks for operation Set on hash map with block size 16 ignoring errors
func BenchmarkSet16(b *testing.B) {
	hm, _ := NewHashMap(16)
	for i := 0; i < b.N; i++ {
		hm.Set(Key(i), i)
	}
}

// BenchmarkSet64 - test benchmarks for operation Set on hash map with block size 64 ignoring errors
func BenchmarkSet64(b *testing.B) {
	hm, _ := NewHashMap(64)
	for i := 0; i < b.N; i++ {
		hm.Set(Key(i), i)
	}
}

// BenchmarkSet128 - test benchmarks for operation Set on hash map with block size 128 ignoring errors
func BenchmarkSet128(b *testing.B) {
	hm, _ := NewHashMap(128)
	for i := 0; i < b.N; i++ {
		hm.Set(Key(i), i)
	}
}

// BenchmarkSet1024 - test benchmarks for operation Set on hash map with block size 1024 ignoring errors
func BenchmarkSet1024(b *testing.B) {
	hm, _ := NewHashMap(1024)
	for i := 0; i < b.N; i++ {
		hm.Set(Key(i), i)
	}
}

// BenchmarkGet16 - test benchmarks for operation Get on hash map with block size 16 ignoring errors
func BenchmarkGet16(b *testing.B) {
	hm, _ := NewHashMap(16)
	for i := 0; i < b.N; i++ {
		hm.Get(Key(i))
	}
}

// BenchmarkGet64 - test benchmarks for operation Get on hash map with block size 64 ignoring errors
func BenchmarkGet64(b *testing.B) {
	hm, _ := NewHashMap(64)
	for i := 0; i < b.N; i++ {
		hm.Get(Key(i))
	}
}

// BenchmarkGet128 - test benchmarks for operation Get on hash map with block size 128 ignoring errors
func BenchmarkGet128(b *testing.B) {
	hm, _ := NewHashMap(128)
	for i := 0; i < b.N; i++ {
		hm.Get(Key(i))
	}
}

// BenchmarkGet1024 - test benchmarks for operation Get on hash map with block size 1024 ignoring errors
func BenchmarkGet1024(b *testing.B) {
	hm, _ := NewHashMap(1024)
	for i := 0; i < b.N; i++ {
		hm.Get(Key(i))
	}
}

// BenchmarkUnset16 - test benchmarks for operation Unset on hash map with block size 16 ignoring errors
func BenchmarkUnset16(b *testing.B) {
	hm, _ := NewHashMap(16)
	for i := 0; i < b.N; i++ {
		hm.Unset(Key(i))
	}
}

// BenchmarkUnset64 - test benchmarks for operation Unset on hash map with block size 64 ignoring errors
func BenchmarkUnset64(b *testing.B) {
	hm, _ := NewHashMap(64)
	for i := 0; i < b.N; i++ {
		hm.Unset(Key(i))
	}
}

// BenchmarkUnset128 - test benchmarks for operation Unset on hash map with block size 128 ignoring errors
func BenchmarkUnset128(b *testing.B) {
	hm, _ := NewHashMap(128)
	for i := 0; i < b.N; i++ {
		hm.Unset(Key(i))
	}
}

// BenchmarkUnset1024 - test benchmarks for operation Unset on hash map with block size 1024 ignoring errors
func BenchmarkUnset1024(b *testing.B) {
	hm, _ := NewHashMap(1024)
	for i := 0; i < b.N; i++ {
		hm.Unset(Key(i))
	}
}

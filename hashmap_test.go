package hashmap

import (
	"testing"
)

// Number of operations for testing Set, Get, Unset, Count
const testN = 1000000

// testSetN - helper for testing operation Set N times ignoring errors
func testSetN(n int, hm HashMaper) {
	for i := 0; i < n; i++ {
		hm.Set(Key(i), i)
	}
}

// testGetN - helper for testing operation Get N times ignoring errors
func testGetN(n int, hm HashMaper) {
	for i := 0; i < n; i++ {
		hm.Get(Key(i))
	}
}

// testUnsetN - helper for testing operation Unset N times ignoring errors
func testUnsetN(n int, hm HashMaper) {
	for i := 0; i < n; i++ {
		hm.Unset(Key(i))
	}
}

// testCountN - helper for testing operation Count N times ignoring errors
func testCountN(n int, hm HashMaper) {
	for i := 0; i < n; i++ {
		hm.Count()
	}
}

// testMapSetN - helper for testing operation Set N times on map type ignoring errors
func testMapSetN(n int, m map[Key]interface{}) {
	for i := 0; i < n; i++ {
		m[Key(i)] = i
	}
}

// testMapGetN - helper for testing operation Get N times on map type ignoring errors
func testMapGetN(n int, m map[Key]interface{}) {
	for i := 0; i < n; i++ {
		_ = m[Key(i)]
	}
}

// testMapUnsetN - helper for testing operation Unset N times on map type ignoring errors
func testMapUnsetN(n int, m map[Key]interface{}) {
	for i := 0; i < n; i++ {
		delete(m, Key(i))
	}
}

// testMapCountN - helper for testing operation Count N times on map type ignoring errors
func testMapCountN(n int, m map[Key]interface{}) {
	for i := 0; i < n; i++ {
		_ = len(m)
	}
}

// TestNewHashMap testing create new hash map by test table (check some different initial values)
func TestNewHashMap(t *testing.T) {
	tests := []struct {
		BlockSize     int
		ExpectedError bool
	}{
		{128, false},
		{0, true},
		{-100, true},
	}
	for _, test := range tests {
		_, err := NewHashMap(test.BlockSize)
		if (err != nil && !test.ExpectedError) || (err == nil && test.ExpectedError) {
			t.Errorf("Expected error: %t, received: %v", test.ExpectedError, err)
		}
	}
}

// TestSet testing set and get entries (check some different key->value entries being set or not and its returned values)
func TestSetGet(t *testing.T) {
	tests := []struct {
		Key           Key
		Value         interface{}
		BeingSet      bool
		ExpectedError bool
	}{
		{"key1", "value", true, false},
		{"key2", 123, false, true},
	}
	hm, _ := NewHashMap(128)
	for _, test := range tests {
		if test.BeingSet {
			err := hm.Set(test.Key, test.Value)
			if err != nil {
				t.Errorf("Expected error: %t, received: %v", test.ExpectedError, err)
			}
		}
		value, err := hm.Get(test.Key)
		if err != nil && test.BeingSet {
			t.Errorf("Expected value: %v, received error: %v", test.Value, err)
		}
		if test.BeingSet && test.Value != value {
			t.Errorf("Expected value: %v, received value: %v", test.Value, value)
		}
	}
}

// TestMapSet - test operation Set N times on map type
func TestMapSet(t *testing.T) {
	m := map[Key]interface{}{}
	testMapSetN(testN, m)
}

// TestMapGet - test operation Get N times on map type
func TestMapGet(t *testing.T) {
	m := map[Key]interface{}{}
	testMapGetN(testN, m)
}

// TestMapUnset - test operation Unset N times on map type
func TestMapUnset(t *testing.T) {
	m := map[Key]interface{}{}
	testMapUnsetN(testN, m)
}

// TestMapCount - test operation Count N times on map type
func TestMapCount(t *testing.T) {
	m := map[Key]interface{}{}
	testMapCountN(testN, m)
}

// TestSet16 - test operation Set N times on hash map with block size = 16
func TestSet16(t *testing.T) {
	hm, _ := NewHashMap(16)
	testSetN(testN, hm)
}

// TestSet64 - test operation Set N times on hash map with block size = 64
func TestSet64(t *testing.T) {
	hm, _ := NewHashMap(64)
	testSetN(testN, hm)
}

// TestSet128 - test operation Set N times on hash map with block size = 128
func TestSet128(t *testing.T) {
	hm, _ := NewHashMap(128)
	testSetN(testN, hm)
}

// TestSet1024 - test operation Set N times on hash map with block size = 1024
func TestSet1024(t *testing.T) {
	hm, _ := NewHashMap(1024)
	testSetN(testN, hm)
}

// TestGet16 - test operation Get N times on hash map with block size = 16
func TestGet16(t *testing.T) {
	hm, _ := NewHashMap(16)
	testGetN(testN, hm)
}

// TestGet64 - test operation Get N times on hash map with block size = 64
func TestGet64(t *testing.T) {
	hm, _ := NewHashMap(64)
	testGetN(testN, hm)
}

// TestGet128 - test operation Get N times on hash map with block size = 128
func TestGet128(t *testing.T) {
	hm, _ := NewHashMap(128)
	testGetN(testN, hm)
}

// TestGet1024 - test operation Get N times on hash map with block size = 1024
func TestGet1024(t *testing.T) {
	hm, _ := NewHashMap(1024)
	testGetN(testN, hm)
}

// TestUnset16 - test operation Unset N times on hash map with block size = 16
func TestUnset16(t *testing.T) {
	hm, _ := NewHashMap(16)
	testUnsetN(testN, hm)
}

// TestUnset64 - test operation Unset N times on hash map with block size = 64
func TestUnset64(t *testing.T) {
	hm, _ := NewHashMap(64)
	testUnsetN(testN, hm)
}

// TestUnset128 - test operation Unset N times on hash map with block size = 128
func TestUnset128(t *testing.T) {
	hm, _ := NewHashMap(128)
	testUnsetN(testN, hm)
}

// TestUnset1024 - test operation Unset N times on hash map with block size = 1024
func TestUnset1024(t *testing.T) {
	hm, _ := NewHashMap(1024)
	testUnsetN(testN, hm)
}

// TestCount16 - test operation Count N times on hash map with block size = 16
func TestCount16(t *testing.T) {
	hm, _ := NewHashMap(16)
	testCountN(testN, hm)
}

// TestCount64 - test operation Count N times on hash map with block size = 64
func TestCount64(t *testing.T) {
	hm, _ := NewHashMap(64)
	testCountN(testN, hm)
}

// TestCount128 - test operation Count N times on hash map with block size = 128
func TestCount128(t *testing.T) {
	hm, _ := NewHashMap(128)
	testCountN(testN, hm)
}

// TestCount1024 - test operation Count N times on hash map with block size = 1024
func TestCount1024(t *testing.T) {
	hm, _ := NewHashMap(1024)
	testCountN(testN, hm)
}

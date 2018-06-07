package hashmap

import (
	"testing"
)

// testCollisionsN logs count of collisions for N operations Set on hash map
func testCollisionN(blockSize int, t *testing.T) {
	collisions := 0
	hm, _ := NewHashMap(blockSize)
	for i := 0; i < testN; i++ {
		index := hashFunc(blockSize, Key(i))
		if hm.Buckets[index] != nil {
			collisions++
		}
		hm.Set(Key(i), i)
	}
	t.Logf("%-25d%-25d%-25d", blockSize, hm.BlockSize, collisions)
}

// TestHashFuncCollisions testing count of collisions for N operations on hash map with different initial block sizes
func TestHashFuncCollisions(t *testing.T) {
	t.Logf("%-25s%-25s%-25s\n", "Initial blockSize", "Resulted blockSize", "Collisions")
	testCollisionN(16, t)
	testCollisionN(64, t)
	testCollisionN(128, t)
	testCollisionN(1024, t)
}

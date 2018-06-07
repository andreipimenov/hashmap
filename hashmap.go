package hashmap

import (
	"errors"
	"fmt"
)

// loadFactor = count of entries / block size
// loadFactor using to check if hash map needs to be reconstructed
const (
	loadFactor = 0.75
)

// Key is of type string
type Key string

// HashMaper interface describes hash map methods:
// set pair of key->value, get value by key, unset pair by key, get count of key->value pairs
type HashMaper interface {
	Set(key Key, value interface{}) error
	Get(key Key) (value interface{}, err error)
	Unset(key Key) error
	Count() int
}

// Entry contains key, value and pointer to the next entry in the chain of entries
type Entry struct {
	Key       Key
	Value     interface{}
	NextEntry *Entry
}

// HashFunc type for function which creates hash and returns key's index
type HashFunc func(blockSize int, key Key) int

// HashMap implements hash map data structure with defined block size, hash function and list of buckets with count of entries
type HashMap struct {
	BlockSize int
	HashFunc  HashFunc
	Buckets   []*Entry
	Entries   int
}

// hashFunc is default hash function
func hashFunc(blockSize int, key Key) int {
	h := 0
	for i := 0; i < len(key); i++ {
		h = 31*h + int(key[i])
	}
	return h % blockSize
}

// NewHashMap creates new hash map and returns pointer to it
func NewHashMap(blockSize int, fn ...HashFunc) (*HashMap, error) {
	if blockSize <= 0 {
		return nil, errors.New("blockSize must be more than 0")
	}
	f := hashFunc
	if len(fn) > 0 {
		f = fn[0]
	}
	return &HashMap{
		BlockSize: blockSize,
		HashFunc:  f,
		Buckets:   make([]*Entry, blockSize),
		Entries:   0,
	}, nil
}

// Set sets key->value entry: insert or update
func (h *HashMap) Set(key Key, value interface{}) error {
	h.reconstruct()
	i := h.HashFunc(h.BlockSize, key)
	entry := h.Buckets[i]
	if entry == nil {
		h.Buckets[i] = &Entry{
			Key:   key,
			Value: value,
		}
		h.Entries++
		return nil
	}
	for entry != nil {
		if entry.Key == key {
			entry.Value = value
			return nil
		}
		if entry.NextEntry == nil {
			entry.NextEntry = &Entry{
				Key:   key,
				Value: value,
			}
			h.Entries++
			return nil
		}
		entry = entry.NextEntry
	}
	return nil
}

// Get returns value by key or error otherwise
func (h *HashMap) Get(key Key) (value interface{}, err error) {
	i := h.HashFunc(h.BlockSize, key)
	entry := h.Buckets[i]
	for entry != nil {
		if entry.Key == key {
			return entry.Value, nil
		}
		entry = entry.NextEntry
	}
	return nil, fmt.Errorf("key %s not found", key)
}

// Unset removes entry by key and returns nil or error otherwise
func (h *HashMap) Unset(key Key) error {
	i := h.HashFunc(h.BlockSize, key)
	entry := h.Buckets[i]
	prevEntry := entry
	for entry != nil {
		if entry.Key == key {
			if prevEntry == entry {
				h.Buckets[i] = entry.NextEntry
			} else {
				prevEntry.NextEntry = entry.NextEntry
			}
			entry = nil
			h.Entries--
			return nil
		}
		prevEntry = entry
		entry = entry.NextEntry
	}
	return fmt.Errorf("key %s not found", key)
}

// Count returns count of entries
func (h *HashMap) Count() int {
	return h.Entries
}

// reconstruct checks current load factor and makes reconstruction of buckets list if needed
func (h *HashMap) reconstruct() {
	if float64(h.Entries)/float64(h.BlockSize) > loadFactor {
		newBlockSize := h.BlockSize * 2
		newBuckets := make([]*Entry, newBlockSize)
		for _, bucket := range h.Buckets {
			if bucket != nil {
				i := h.HashFunc(newBlockSize, bucket.Key)
				newBuckets[i] = bucket
			}
		}
		h.BlockSize = newBlockSize
		h.Buckets = newBuckets
	}
}

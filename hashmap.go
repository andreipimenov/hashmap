package hashmap

import (
	"errors"
	"fmt"
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

// HashMap implements hash map data structure with defined block size, hash function and list of entries
type HashMap struct {
	BlockSize    int
	HashFunc     HashFunc
	Entries      []*Entry
	EntriesCount int
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
func NewHashMap(blockSize int, fn ...HashFunc) (HashMaper, error) {
	if blockSize <= 0 {
		return nil, errors.New("blockSize must be more than 0")
	}
	f := hashFunc
	if len(fn) > 0 {
		f = fn[0]
	}
	return &HashMap{
		BlockSize:    blockSize,
		HashFunc:     f,
		Entries:      make([]*Entry, blockSize),
		EntriesCount: 0,
	}, nil
}

// Set sets key->value entry: insert or update
func (h *HashMap) Set(key Key, value interface{}) error {
	i := h.HashFunc(h.BlockSize, key)
	if i > len(h.Entries)-1 {
		return fmt.Errorf("index %d is out of range", i)
	}
	entry := h.Entries[i]
	if entry == nil {
		h.Entries[i] = &Entry{
			Key:   key,
			Value: value,
		}
		h.EntriesCount++
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
			h.EntriesCount++
			return nil
		}
		entry = entry.NextEntry
	}
	return nil
}

// Get returns value by key or error otherwise
func (h *HashMap) Get(key Key) (value interface{}, err error) {
	i := h.HashFunc(h.BlockSize, key)
	if i > len(h.Entries)-1 {
		return nil, fmt.Errorf("index %d is out of range", i)
	}
	entry := h.Entries[i]
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
	if i > len(h.Entries)-1 {
		return fmt.Errorf("index %d is out of range", i)
	}
	entry := h.Entries[i]
	prevEntry := entry
	for entry != nil {
		if entry.Key == key {
			if prevEntry == entry {
				h.Entries[i] = entry.NextEntry
			} else {
				prevEntry.NextEntry = entry.NextEntry
			}
			entry = nil
			h.EntriesCount--
			return nil
		}
		prevEntry = entry
		entry = entry.NextEntry
	}
	return fmt.Errorf("key %s not found", key)
}

// Count returns count of entries
func (h *HashMap) Count() int {
	return h.EntriesCount
}

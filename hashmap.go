package hashmap

// Key is of dynamic type
type Key interface{}

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
	BlockSize int
	HashFunc  HashFunc
	Entries   []*Entry
}

// NewHashMap creates new hash map and returns pointer to it
func NewHashMap(blockSize int, fn ...HashFunc) HashMaper {
	hashFunc := HashFunc(nil)
	if len(fn) > 0 {
		hashFunc = fn[0]
	}
	return &HashMap{
		BlockSize: blockSize,
		HashFunc:  hashFunc,
		Entries:   make([]*Entry, blockSize),
	}
}

// Set sets key->value entry
func (h *HashMap) Set(key Key, value interface{}) error {
	return nil
}

// Get returns value by key or error otherwise
func (h *HashMap) Get(key Key) (value interface{}, err error) {
	return nil, nil
}

// Unset removes entry by key and returns nil or error otherwise
func (h *HashMap) Unset(key Key) error {
	return nil
}

// Count returns number of entries
func (h *HashMap) Count() int {
	return 0
}

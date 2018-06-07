# Hash map
Implementation of hash map data structure in Golang with separate chains as method of resolving collisions.

### Example

```
package main

import (
	"fmt"

	"github.com/andreipimenov/hashmap"
)

func main() {
	//Create hash map
	hm, _ := hashmap.NewHashMap(128)

	//Set entries
	hm.Set("Hello", "World")
	hm.Set("Names", []string{"John", "Jane", "Melissa"})
	hm.Set("Ages", map[string]float64{"John": 33, "Jane": 19})

	//Get entries
	v1, _ := hm.Get("Hello")
	v2, _ := hm.Get("Names")
	v3, _ := hm.Get("Ages")
	fmt.Printf("%v->%v\n%v->%v\n%v->%v\n", "Hello", v1, "Names", v2, "Ages", v3)

	//Get count of entries
	fmt.Printf("Count of entries is %d\n", hm.Count())

	//Unset entry
	hm.Unset("Hello")

	//Try to get removed entry
	_, err := hm.Get("Hello")
	if err != nil {
		fmt.Printf("Trying to get removed entry with key Hello. Returned error: %v\n", err)
	}

	//Check count of entries
	fmt.Printf("Count of entries now is %d\n", hm.Count())
}
```

### Collisions

Hash map automatically grows if load factor exceeded.
I use load factor = 0.75. Load factor = count of entries / block size

```
=== RUN   TestHashFuncCollisions
--- PASS: TestHashFuncCollisions (1.75s)
        hashfunc_test.go:23: Initial blockSize        Resulted blockSize       Collisions
        hashfunc_test.go:18: 16                       2097152                  999984
        hashfunc_test.go:18: 64                       2097152                  999936
        hashfunc_test.go:18: 128                      2097152                  999872
        hashfunc_test.go:18: 1024                     2097152                  462570
```

Conclusion: less count of collisions on hash maps with bigger initial block size. 

### Benchmarks

#### Compare time spent for 1000000 operations SET on native map versus hash map.
```
C:\Gocode\src\github.com\andreipimenov\hashmap>go test -v -count=1 -cpu=1 -run=TestMapSet
=== RUN   TestMapSet
--- PASS: TestMapSet (0.99s)
PASS
ok      github.com/andreipimenov/hashmap        1.214s

C:\Gocode\src\github.com\andreipimenov\hashmap>go test -v -count=1 -cpu=1 -run=TestSet16
=== RUN   TestSet16
--- PASS: TestSet16 (0.44s)
PASS
ok      github.com/andreipimenov/hashmap        0.527s
```
Conclusion: hash map up to 2 times faster than native map on operation SET.

#### Test benchmarks

```
go test -v -count=1 -cpu=1 -bench=. -run=^Benchmark -benchmem
```
```
pkg: github.com/andreipimenov/hashmap
BenchmarkFunc16         50000000                28.5 ns/op             0 B/op          0 allocs/op
BenchmarkFunc64         50000000                29.1 ns/op             0 B/op          0 allocs/op
BenchmarkFunc128        50000000                28.8 ns/op             0 B/op          0 allocs/op
BenchmarkFunc1024       50000000                29.1 ns/op             0 B/op          0 allocs/op
BenchmarkSetMap          2000000               509 ns/op              97 B/op          2 allocs/op
BenchmarkGetMap         100000000               11.3 ns/op             0 B/op          0 allocs/op
BenchmarkUnsetMap       100000000               10.8 ns/op             0 B/op          0 allocs/op
BenchmarkCountMap       2000000000               0.43 ns/op            0 B/op          0 allocs/op
BenchmarkCount16        2000000000               0.43 ns/op            0 B/op          0 allocs/op
BenchmarkCount64        2000000000               0.43 ns/op            0 B/op          0 allocs/op
BenchmarkCount128       2000000000               0.43 ns/op            0 B/op          0 allocs/op
BenchmarkCount1024      2000000000               0.43 ns/op            0 B/op          0 allocs/op
BenchmarkSet16           5000000               217 ns/op              33 B/op          2 allocs/op
BenchmarkSet64           5000000               203 ns/op              33 B/op          2 allocs/op
BenchmarkSet128          5000000               218 ns/op              33 B/op          2 allocs/op
BenchmarkSet1024         5000000               207 ns/op              33 B/op          2 allocs/op
BenchmarkGet16           5000000               376 ns/op              67 B/op          4 allocs/op
BenchmarkGet64           5000000               371 ns/op              67 B/op          4 allocs/op
BenchmarkGet128          5000000               371 ns/op              68 B/op          4 allocs/op
BenchmarkGet1024         5000000               374 ns/op              68 B/op          4 allocs/op
BenchmarkUnset16         3000000               374 ns/op              67 B/op          4 allocs/op
BenchmarkUnset64         5000000               370 ns/op              67 B/op          4 allocs/op
BenchmarkUnset128        5000000               371 ns/op              67 B/op          4 allocs/op
BenchmarkUnset1024       5000000               372 ns/op              68 B/op          4 allocs/op
```
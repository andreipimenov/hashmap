# Hash map

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

### Benchmarks
```
go test -v -count=1 -cpu=1 -bench=. -run=^Benchmark -benchmem
```
```
pkg: github.com/andreipimenov/hashmap
BenchmarkFunc16         50000000                28.4 ns/op             0 B/op          0 allocs/op
BenchmarkFunc64         50000000                28.6 ns/op             0 B/op          0 allocs/op
BenchmarkFunc128        50000000                28.9 ns/op             0 B/op          0 allocs/op
BenchmarkFunc1024       50000000                28.7 ns/op             0 B/op          0 allocs/op
BenchmarkSetMap          1000000              1026 ns/op             178 B/op          2 allocs/op
BenchmarkGetMap         100000000               11.0 ns/op             0 B/op          0 allocs/op
BenchmarkUnsetMap       100000000               10.5 ns/op             0 B/op          0 allocs/op
BenchmarkCountMap       2000000000               0.43 ns/op            0 B/op          0 allocs/op
BenchmarkCount16        1000000000               2.57 ns/op            0 B/op          0 allocs/op
BenchmarkCount64        1000000000               2.56 ns/op            0 B/op          0 allocs/op
BenchmarkCount128       1000000000               2.60 ns/op            0 B/op          0 allocs/op
BenchmarkCount1024      1000000000               2.60 ns/op            0 B/op          0 allocs/op
BenchmarkSet16          10000000               156 ns/op              24 B/op          2 allocs/op
BenchmarkSet64           5000000               203 ns/op              33 B/op          2 allocs/op
BenchmarkSet128          5000000               210 ns/op              33 B/op          2 allocs/op
BenchmarkSet1024        10000000               180 ns/op              24 B/op          2 allocs/op
BenchmarkGet16           5000000               381 ns/op              67 B/op          4 allocs/op
BenchmarkGet64           3000000               369 ns/op              67 B/op          4 allocs/op
BenchmarkGet128          3000000               381 ns/op              67 B/op          4 allocs/op
BenchmarkGet1024         5000000               372 ns/op              68 B/op          4 allocs/op
BenchmarkUnset16         5000000               374 ns/op              67 B/op          4 allocs/op
BenchmarkUnset64         5000000               380 ns/op              67 B/op          4 allocs/op
BenchmarkUnset128        5000000               378 ns/op              67 B/op          4 allocs/op
BenchmarkUnset1024       5000000               378 ns/op              68 B/op          4 allocs/op
```
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

package main

import "fmt"

func main() {
	// a map is a dynamic collection
	// relates a value to a key
	var m map[string]int
	// nil
	fmt.Println(m)
	m = map[string]int{"foo": 1, "bar": 2}
	// does not guarantee an order
	fmt.Println(m)
	m["foo"] = 3
	fmt.Println(m["foo"])
	delete(m, "foo")
	m["baz"] = 12
	fmt.Println(m)
	// even if the key does not exist, it will return a zero value
	fmt.Println(m["foofoo"])
	// comma ok syntax
	v, ok := m["foofoo"]
	fmt.Println(v, ok)
	// maps are also copied by reference
	// maps do also have a clone function like slices
	m2 := m
	fmt.Println(m2)
}

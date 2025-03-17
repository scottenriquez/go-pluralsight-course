package main

import "fmt"

// pointers are primarily used to share memory
// use copies whenever possible

func main() {
	// value semantics
	a := 42
	b := a
	a = 27
	fmt.Println(a, b)
	// pointers
	c := 34
	// address operator
	d := &c
	c = 100
	// dereference operator
	fmt.Println(c, *d)
	// built-in new function creates a pointer to anonymous variable
	// and returns a pointer to it
	e := new(int)
	fmt.Println(e)
}

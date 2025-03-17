package main

import (
	"fmt"
	"slices"
)

func main() {
	// a slice is a subset of an array
	// slices are reference types and do not container their own data
	// changes made to the array are reflected both there and in the slice
	var s []int
	// value is actually nil
	fmt.Println(s)
	s = []int{1, 2, 3}
	s[0] = 99
	fmt.Println(s)
	// same as arrays above
	// slices support append
	// the append method does not mutate the slice
	s = append(s, 5, 6, 7)
	fmt.Println(s)
	// can also delete a range from the slice
	// up to but not including the second index
	s = slices.Delete(s, 1, 3)
	fmt.Println(s)
	// unlike arrays, slices are copied by reference
	s2 := s
	s[0] = 42
	// both reflect the changes
	// can use slices.Clone() to copy by values
	fmt.Println(s, s2)
	// slices are not comparable because they are reference types
	// use slices.Equal() to compare slices by value
	fmt.Println(slices.Equal(s, s2))
}

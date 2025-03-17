package main

import "fmt"

func main() {
	// arrays are fixed length
	// zero-based indexing
	var arr [3]int
	// all 'zero' values
	// [0, 0, 0]
	fmt.Println(arr)
	var initializedArr = [3]int{1, 2, 3}
	fmt.Println(initializedArr)
	// arrays are mutable
	initializedArr[0] = 5
	fmt.Println(initializedArr)
	// can also get length
	fmt.Println(len(initializedArr))
	// arrays are copied by value
	arr2 := initializedArr
	arr3 := initializedArr
	initializedArr[0] = 12
	fmt.Println(initializedArr)
	// does not reflect the change to index 0
	fmt.Println(arr2)
	// arrays are comparable
	fmt.Println(arr2 == initializedArr)
	fmt.Println(arr2 == arr3)
}

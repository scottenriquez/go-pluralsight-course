package main

import "fmt"

func main() {
	// fixed-size collection but not limited to a single data type
	// comprised of field name, type, and value
	// cannot add fields dynamically
	// declares an anonymous struct
	// leads to lengthy signature definitions
	var s struct {
		firstName string
		lastName  string
		score     int
	}
	// a struct is a value type
	// defaults to zero values
	fmt.Println(s)
	type student struct {
		firstName string
		lastName  string
		score     int
	}
	arthurDent := student{
		firstName: "Arthur",
		lastName:  "Dent",
		score:     92,
	}
	fmt.Println(arthurDent)
	// structs copied by value
	otherStudent := arthurDent
	arthurDent.score = 88
	fmt.Println(arthurDent, otherStudent)
}

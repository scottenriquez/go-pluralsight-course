package main

import "fmt"

func main() {
	var myName string
	// uses the 'zero' value which is an empty string
	fmt.Println(myName)
	// string becomes the inferred type
	// we do not need to add the string keyword
	var anotherName string = "Scottie"
	fmt.Println(anotherName)
	// short declaration syntax
	aThirdName := "Christopher"
	fmt.Println(aThirdName)
}

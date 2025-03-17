package main

import (
	"fmt"
)

func main() {
	// supports escape characters
	interpretedString := "Hello, World!\n"
	fmt.Println(interpretedString)
	rawString := `Hello, World!\n`
	fmt.Println(rawString)
	rawStringWithNewLine := `Hello, World..!
	
...!
	`
	fmt.Println(rawStringWithNewLine)
}

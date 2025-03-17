package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Dent, Arthur"
	var score uint = 87
	// can also declare on the same line
	// name, score := "Dent, Arthur", 87
	fmt.Println("Student Scores")
	fmt.Println(strings.Repeat("-", 14))
	fmt.Println(name, score)
}

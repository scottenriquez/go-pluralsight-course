package main

import (
	"fmt"
	"strings"
)

func main() {
	type score struct {
		name  string
		score int
	}
	scores := []score{
		// score is optional
		score{"Dent, Arthur", 87},
		{"MacMillan, Tricia", 93},
		{"Prefect, Ford", 98},
	}
	fmt.Println("Select score to print (1-3): ")
	var option string
	fmt.Scanln(&option)
	var index int
	switch option {
	case "1":
		index = 0
	case "2":
		index = 1
	case "3":
		index = 2
	default:
		fmt.Println("Invalid option", option)
		fmt.Println("Defaulting to 1")
		index = 0
	}
	fmt.Println("Student Scores")
	fmt.Println(strings.Repeat("-", 14))
	fmt.Println(scores[index])
}

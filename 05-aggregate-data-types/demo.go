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
	fmt.Println("Student Scores")
	fmt.Println(strings.Repeat("-", 14))
	fmt.Println(scores[0])
	fmt.Println(scores[1])
	fmt.Println(scores[2])
}

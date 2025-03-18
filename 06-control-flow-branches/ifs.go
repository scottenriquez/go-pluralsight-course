package main

import "fmt"

func main() {
	conditionA := false
	conditionB := false
	if conditionA {
		fmt.Println("Condition A is true")
	} else if conditionB {
		fmt.Println("Condition B is true")
	} else {
		fmt.Println("Both conditions are false")
	}
}

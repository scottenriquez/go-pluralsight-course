package main

import "fmt"

func main() {
	integer := 10
	fmt.Println(integer)
	unsignedInteger := uint(integer)
	fmt.Println(unsignedInteger)
	// supports 32 and 64 bits
	float := 10.0
	fmt.Println(float)
	// supports 64 and 128
	complexNumeric := complex(1, 2)
	fmt.Println(complexNumeric)
}

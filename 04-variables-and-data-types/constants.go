package main

func main() {
	// implicitly-typed constant
	const fixedNumber = 42
	// explicitly-typed constant
	const typedFixedNumber uint64 = 42
	// constant group
	const (
		isEnabled         = true
		pi        float64 = 3.14
	)
	// constant expression
	// must be able to evaluate at compile time (e.g., no functions)
	const expression = 2 + 5
}

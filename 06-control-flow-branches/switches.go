package main

func main() {
	index := 3
	switch index {
	case 1:
		println("First")
	case 2, 3:
		println("Second")
	case 2 * 2:
		println("Third")
	default:
		println("default")
	}
}

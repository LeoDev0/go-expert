package main

func main() {
	a := 10
	b := 2

	// Go lang does not have else if and ternary operator
	if a < b {
		println("a < b")
	} else {
		println("a >= b")
	}

	switch a {
		case 1:
			println("a")
		case 2:
			println("b")
		case 3:
			println("c")
		default:
			println("d")
	}
}
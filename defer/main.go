package main

import "fmt"

func main() {
	defer fmt.Println("Primeira linha") // Defer is executed after the function ends
	fmt.Println("Segunda linha")
	fmt.Println("Terceira linha")

	// output:
	// - Segunda linha
	// - Terceira linha
	// - Primeira linha
}
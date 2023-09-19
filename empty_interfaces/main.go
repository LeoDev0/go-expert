package main

import "fmt"

// interfaces vazias são usadas para representar qualquer tipo de dado, é o tipo mais genérico do Go
func main() {
	var x interface{} = 10

	showType(x)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variável é %T e o valor é %v \n", t, t)
}
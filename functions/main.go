package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum(5, 10)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(valor)

	fmt.Println(sum2(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)) // 55

	// função anônima
	total := func() int {
		return sum2(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 20, -10)
	}()

	fmt.Println(total) // 65
}

func sum(a, b int) (int, error) {
	if a + b >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}

	return a + b, nil
}

// Função com parâmetros variádicos
func sum2(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}
	
	return total
}


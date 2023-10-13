package main

import (
	"fmt"

	"curso-go/math"

	"github.com/google/uuid"
)

func main() {
	s := math.Soma(10, 20)

	fmt.Printf("Resultado: %v \n", s)
	fmt.Println(uuid.New())
}
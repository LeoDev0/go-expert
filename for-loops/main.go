package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"um", "dois", "três"}
	for k, v := range numeros {
		println(k, v)
	}
	
	// "do-while" equivalent in Go
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	// infinite loop ("while true" equivalent in Go)
	for {
		println("Hello, World!")
	}
}
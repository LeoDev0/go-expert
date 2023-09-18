package main

func main() {
	// Memória -> Endereço -> Valor

	// Ponteiro é representado por um *
	// Variável que armazena o endereço de memória de outra variável
	// & é usado para obter o endereço de memória de uma variável
	// * é usado para obter o valor da variável através do endereço de memória
	// * é usado para definir um tipo de ponteiro

	var variavel int = 10
	var ponteiro *int = &variavel

	println(variavel) // 10
	println(ponteiro) // 0xc0000160b0

	// Alterando o valor da variável através do ponteiro
	*ponteiro = 20
	variavel2 := &variavel

	println(variavel) // 20
	println(variavel2) // 0xc0000160b0
	println(*variavel2) // 20
	println(ponteiro) // 0xc0000160b0
}
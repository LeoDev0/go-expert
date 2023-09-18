package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero 		 int
	Cidade 		 string
	Estado 		 string
}

type Cliente struct {
	Nome 	string
	Idade int
	Ativo bool
	Endereco
	// End   Endereco
}

// interface
type Pessoa interface {
	Desativar()
}

// método
// esse método auto implementa a interface Pessoa pela assinatura do método - Go não usa palavras reservadas para implementar interfaces
func (c *Cliente) Desativar() { // struct Cliente associada ao método acompanha ponteiro para alterar o valor
	c.Ativo = false
	fmt.Printf("cliente %s foi desativado \n", c.Nome)
}	

func main() {
	wesley := Cliente{
		Nome: "Wesley",
		Idade: 33,
		Ativo: true,
	}

	wesley.Desativar()
	wesley.Cidade = "São Paulo"

	fmt.Println(wesley.Nome)
	fmt.Println(wesley.Ativo)
	fmt.Println(wesley.Cidade)
}
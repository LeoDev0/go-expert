package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Address struct {
	Cep 			 	string `json:"cep"`
	Logradouro 	string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro 			string `json:"bairro"`
	Localidade 	string `json:"localidade"`
	UF 					string `json:"uf"`
	IBGE 				string `json:"ibge"`
	GIA 				string `json:"gia"`
	DDD 				string `json:"ddd"`
	SIAFI 			string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		address, err := GetViaCep(cep)
		if err != "" {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(address)
		}
	}
}

func GetViaCep(cep string) (Address, string) {
	var address Address
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	res, err := http.Get(url)

	if err != nil {
		return address, fmt.Sprintf("Requisição falhou: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return address, fmt.Sprintf("Cep %s não encontrado", cep)
	}

	dataBytes, err := io.ReadAll(res.Body)

	err = json.Unmarshal(dataBytes, &address)
	if err != nil {
		return address, err.Error()
	}

	return address, ""
}

// EXAMPLE: go run cli.go 95010010 01001000
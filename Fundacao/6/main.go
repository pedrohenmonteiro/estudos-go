package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome    string
	Idade   int
	Ativo   bool
	Address Endereco
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}

func main() {
	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}

	wesley.Cidade = "Londrina"
	wesley.Address.Cidade = "Londrina"

	wesley.Desativar()

	fmt.Printf("O nome do cliente É %s", wesley.Nome)
}

package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

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

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}

	wesley.Cidade = "Londrina"
	wesley.Address.Cidade = "Londrina"

	minhaEmpresa := Empresa{}
	Desativacao(minhaEmpresa)

	fmt.Printf("O nome do cliente É %s", wesley.Nome)
}

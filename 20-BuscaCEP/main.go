package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Cep struct {
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
}

func buscaCep(num int) string {
	viaCepLink := fmt.Sprintf("https://viacep.com.br/ws/%d/json/", num)

	req, err := http.Get(viaCepLink)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	var cep Cep
	json.Unmarshal(res, &cep)

	return cep.Localidade
}

func main() {

	println(buscaCep(86010150))

}

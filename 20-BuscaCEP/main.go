package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// type Cep struct {
// 	Bairro     string `json:"bairro"`
// 	Localidade string `json:"localidade"`
// }

// func buscaCep(num int) string {
// 	viaCepLink := fmt.Sprintf("https://viacep.com.br/ws/%d/json/", num)

// 	req, err := http.Get(viaCepLink)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer req.Body.Close()

// 	res, err := io.ReadAll(req.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	var cep Cep
// 	json.Unmarshal(res, &cep)

// 	return cep.Localidade
// }

type CepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {

	for _, url := range os.Args[1:] {
		req, err := http.Get("https://viacep.com.br/ws/" + url + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição: %v\n", err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler a resposta: %v\n", err)
		}
		var cep CepResponse
		err = json.Unmarshal(res, &cep)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao converter a resposta em JSON: %v\n", err)
		}

		file, err := os.Create("arquivo.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("CEP: %s\nLogradouro: %s\nComplemento: %s\nBairro: %s\nLocalidade: %s\nUF: %s\nIBGE: %s\nGIA: %s\nDDD: %s\n ", cep.Cep, cep.Logradouro, cep.Complemento, cep.Bairro, cep.Localidade, cep.Uf, cep.Ibge, cep.Gia, cep.Ddd))
		if err != nil {
			panic(err)
		}

	}

}

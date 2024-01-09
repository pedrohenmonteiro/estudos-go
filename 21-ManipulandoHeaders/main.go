package main

import (
	"encoding/json"
	"io"
	"net/http"
)

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
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8080", nil)
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep")

	req, err := http.Get("https://viacep.com.br/ws/" + cepParam + "/json/")
	if cepParam == "" || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := io.ReadAll(req.Body)

	var cepResponse CepResponse

	json.Unmarshal(res, &cepResponse)

	marshal, err := json.Marshal(cepResponse)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(marshal))
}

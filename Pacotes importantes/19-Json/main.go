package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"-"`
	Saldo  int `json:"saldo"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}

	//Ao usar Marshal, você consegue guardar o valor do json em uma váriavel.
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	println(string(res))

	// Ao usar NewEncoder, você consegue serializar o json direto na tela.
	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}

	// Ao usar Unmarshal, você consegue deserializar o json em uma struct.
	jsonPuro := []byte(`{"n":2,"s":200}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		panic(err)
	}
	println(contaX.Saldo)

}

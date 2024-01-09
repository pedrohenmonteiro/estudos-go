package main

func main() {

	// Memória -> Endereço -> Valor

	// variavel -> ponteiro que tem um endereço na memoria -> valor
	a := 10

	var ponteiro *int = &a
	*ponteiro = 20

	b := &a
	*b = 30

	println(a)
}

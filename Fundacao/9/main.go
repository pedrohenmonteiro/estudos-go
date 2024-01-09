// package main

// func soma(a, b int) int {
// 	a = 50
// 	return a + b
// }

// func main() {
// 	minhaVar1 := 10
// 	minhaVar2 := 20

// 	println(soma(minhaVar1, minhaVar2))
// 	println(minhaVar1) // na funçao soma, minhaVar1 é apenas uma cópia, portanto o valor nao é alterado
// }

package main

func soma(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20

	println(soma(&minhaVar1, &minhaVar2))
	println(minhaVar1) //na funçao soma, minhaVar1 acessa o endereçamento da memoria e altera o valor, sem fazer uma copia
}

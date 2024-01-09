package main

import "fmt"

func main() {

	// func() {
	// 	//rodar web server
	// }

	total := func() int {
		return sum(34, 43, 34, 13, 6, 3, 12, 435, 32, 3) * 2
	}()

	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}

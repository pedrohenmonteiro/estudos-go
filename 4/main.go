package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum(50, 10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(valor)
}

func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("O valor nao deve ultrapassar 50")
	}
	return a + b, nil
}

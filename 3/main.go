package main

import "fmt"

type ID int

const a ID = 1

func main() {

	salarios := map[string]int{
		"Wesley": 1000,
		"João":   2000,
		"Maria":  3000,
	}

	delete(salarios, "Wesley")
	salarios["Wes"] = 5000

	fmt.Println(salarios["Wes"])

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}
}

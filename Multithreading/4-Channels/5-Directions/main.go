package main

import "fmt"

// receive only
func recebe(nome string, hello chan<- string) {
	hello <- nome
}

// send only
func ler(data <-chan string) {
	fmt.Println(<-data)
}

func main() {
	hello := make(chan string)
	go recebe("Hello", hello)

	ler(hello)
}

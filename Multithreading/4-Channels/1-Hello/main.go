package main

import "fmt"

// Thread 1
func main() {
	canal := make(chan string) // Canal está vazio

	//Thread 2
	go func() {
		canal <- "Olá mundo" // Canal está cheio
		//canal <- "Olá mundo 2" //  Não dá pra colocar outro valor no canal enquanto ele não esvaziar
	}()

	// Thread 1
	msg := <-canal // Canal esvazia
	fmt.Println(msg)
}

package main

import "fmt"

// Thread 1
func main() {
	ch := make(chan int) //canal vazio
	go publish(ch)       //for para encher o canal. o proximo looping só executa depois que o canal está vazio
	reader(ch)           //for para verificar o valor do canal quando ele encher, ler o valor e esvaziar,

}

func reader(ch chan int) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)

	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

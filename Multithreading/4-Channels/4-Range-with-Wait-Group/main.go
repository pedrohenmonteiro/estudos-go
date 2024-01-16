package main

import (
	"fmt"
	"sync"
)

// Thread 1
func main() {
	ch := make(chan int) //canal vazio
	wg := sync.WaitGroup{}
	wg.Add(10)
	go publish(ch)     //for para encher o canal. o proximo looping só executa depois que o canal está vazio
	go reader(ch, &wg) //for para verificar o valor do canal quando ele encher, ler o valor e esvaziar,
	wg.Wait()

}

func reader(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
		wg.Done()
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

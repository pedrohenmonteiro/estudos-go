package main

// Thread 1
func main() {
	forever := make(chan bool) // canal está vazio

	// forever <- true  // gera deadlock pois o canal deve ficar cheio em uma go routine

	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		forever <- true
	}()

	<-forever //esperando canal ficar cheio para esvaziar
}

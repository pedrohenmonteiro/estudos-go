package main

import "fmt"

func main() {
	// req, err := http.Get("https://api.chucknorris.io/jokes/random")
	// if err != nil {
	// 	panic(err)
	// }
	// defer req.Body.Close()
	// res, err := io.ReadAll(req.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// println(string(res))

	defer fmt.Println("Primeira linha")
	fmt.Println("Segunda linha")
	fmt.Println("Terceira linha")
}

package main

func main() {
	var minhaVar interface{} = "Pedro Monteiro"

	println(minhaVar.(string))
	res, ok := minhaVar.(int)
	println(res)
	println(ok)
}

package main

func main() {
	a := 1
	b := 2

	if a < b {
		println(true)
	} else {
		println(false)
	}

	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	default:
		println("ovu")
	}
}

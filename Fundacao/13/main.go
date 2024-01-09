package main

type MyNumber int

type Number interface {
	~int | float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	return a == b
}

func main() {

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	m2 := map[string]float64{
		"a": 1.1,
		"b": 2.2,
		"c": 3.0,
	}

	m3 := map[string]MyNumber{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10))
	println(Compara(10, 1))

}

package main

type Conta struct {
	saldo int
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func (c *Conta) depositar(valor int) {
	c.saldo += valor
	println(c.saldo)
}

func main() {
	conta := Conta{
		saldo: 10,
	}

	conta.depositar(45)
	println(conta.saldo)

}

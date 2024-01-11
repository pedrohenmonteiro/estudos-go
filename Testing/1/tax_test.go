package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %.1f but got %.1f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expected {
			t.Errorf("Expected %.1f but got %.1f", item.expected, result)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.00)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.00)
	}
}

// go test -coverprofile=coverage.out  para ver a cobertura de testes

// go tool cover -html=coverage.out  para vizualizar a cobertura de testes em html

// go test -bench=.          para rodar o benchmark
// go test -bench=. -run=^#   para rodar o benchmark sem rodar testes
// go test -bench=. -run=^# -count=10 para rodar o benchmark 10 vezes
// go test -bench=. -run=^# -count=10 -benchtime=3s roda o benchmark por 3 segundos
// go test -bench=. -run=^# -benchmem   para verificar a alocaçao de memoria

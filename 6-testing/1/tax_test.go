package tax

import "testing"

// por padrão para lidar com testes no go, os metodos devem começar com Test
// go test . -coverprofile=coverage.out
// go tool cover -html=coverage
// go test -bench=.
// go test -bench=. -run=^# -> roda todos os testes que começam com #, no caso nenhum

func TestCalculateTax(t *testing.T) {
	t.Run("Tax is 5% of the amount", func(t *testing.T) {
		amount := 100.0
		expected := 5.0

		result := CalculateTax(amount)

		if result != expected {
			t.Errorf("Expected %.2f, got %.2f", expected, result)
		}
	})

	t.Run("Tax is 10% of the amount", func(t *testing.T) {
		amount := 1000.0
		expected := 100.0

		result := CalculateTax(amount)

		if result != expected {
			t.Errorf("Expected %.2f, got %.2f", expected, result)
		}
	})
}

func TestCalculateTaxInBatch(t *testing.T) {
	type calcText struct {
		amount, expected float64
	}

	// testa varios casos de uma vez
	tests := []calcText{
		{100.0, 5.0},
		{1000.0, 100.0},
		{0, 0},
	}

	for _, test := range tests {
		result := CalculateTax(test.amount)
		if result != test.expected {
			t.Errorf("Expected %.2f, got %.2f", test.expected, result)
		}
	}
}

// bechmark testa a performance do codigo
func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(1000)
	}
}

// Benchmark que usa o time.Sleep para simular um delay
func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(1000)
	}
}

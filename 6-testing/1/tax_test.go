package tax

import "testing"

// por padrão para lidar com testes no go, os metodos devem começar com Test

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

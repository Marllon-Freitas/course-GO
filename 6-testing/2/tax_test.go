// tests with testify
package tax

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	t.Run("Tax is 5% of the amount", func(t *testing.T) {
		amount := 100.0
		expected := 5.0

		result, err := CalculateTax(amount)
		assert.Nil(t, err)
		assert.Equal(t, expected, result)

		amount = 0.0
		expected = 0.0
		result, err = CalculateTax(amount)
		assert.Nil(t, err)
		assert.Equal(t, expected, result)

		amount = -100.0
		result, err = CalculateTax(amount)
		assert.Error(t, err, "amount cannot be negative")
		assert.Equal(t, 0.0, result)
	})
}

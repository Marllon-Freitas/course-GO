// testes com testify
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

func TestCalculateTaxAndSave(t *testing.T) {
	t.Run("Tax is 5% of the amount", func(t *testing.T) {
		amount := 100.0
		expected := 5.0

		repo := &TaxRepositoryMock{}
		repo.On("Save", expected).Return(nil).Once()

		err := CalculateTaxAndSave(amount, repo)
		assert.Nil(t, err)

		repo.AssertExpectations(t)
	})

	t.Run("Tax is 0 for negative amount", func(t *testing.T) {
		amount := -10.0
		expected := 0.0

		repo := &TaxRepositoryMock{}
		repo.On("Save", expected).Return(nil).Once()

		err := CalculateTaxAndSave(amount, repo)
		assert.Nil(t, err)

		repo.AssertExpectations(t)
	})

	t.Run("Tax is 10% for amount between 1000 and 20000", func(t *testing.T) {
		amount := 1500.0
		expected := 150.0

		repo := &TaxRepositoryMock{}
		repo.On("Save", expected).Return(nil).Once()

		err := CalculateTaxAndSave(amount, repo)
		assert.Nil(t, err)

		repo.AssertExpectations(t)
	})

	t.Run("Tax is 20% for amount 20000 or more", func(t *testing.T) {
		amount := 25000.0
		expected := 5000.0

		repo := &TaxRepositoryMock{}
		repo.On("Save", expected).Return(nil).Once()

		err := CalculateTaxAndSave(amount, repo)
		assert.Nil(t, err)

		repo.AssertExpectations(t)
	})
}

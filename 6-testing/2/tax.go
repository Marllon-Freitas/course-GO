// testes com testify
package tax

import "errors"

type Repository interface {
	Save(float64) error
}

func CalculateTax(amount float64) (float64, error) {
	if amount < 0 {
		return 0.0, errors.New("amount cannot be negative")
	}
	if amount >= 1000 && amount < 20000 {
		return amount * 0.1, nil
	}
	if amount >= 20000 {
		return amount * 0.2, nil
	}
	return amount * 0.05, nil
}

func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := CalculateTax2(amount)
	return repository.Save(tax)
}

func CalculateTax2(amount float64) float64 {
	if amount < 0 {
		return 0.0
	}
	if amount >= 1000 && amount < 20000 {
		return amount * 0.1
	}
	if amount >= 20000 {
		return amount * 0.2
	}
	return amount * 0.05
}

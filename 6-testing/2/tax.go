// tests with testify
package tax

import "errors"

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

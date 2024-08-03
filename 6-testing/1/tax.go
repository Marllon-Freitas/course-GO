package tax

import "time"

func CalculateTax(amount float64) float64 {
	if amount < 0 {
		return 0
	}
	if amount >= 1000 && amount < 20000 {
		return amount * 0.1
	}
	if amount >= 20000 {
		return amount * 0.2
	}
	return amount * 0.05
}

func CalculateTax2(amount float64) float64 {
	time.Sleep(time.Millisecond)
	if amount == 0 {
		return 0
	}
	if amount >= 1000 {
		return amount * 0.1
	}
	return amount * 0.05
}

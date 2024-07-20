// Aula 23 - Condicionais
package main

// no go não tem else if, só if e else

func main() {
	a := 10
	b := 20
	c := 30

	if a > b {
		println("a é maior que b")
	} else {
		println("b é maior que a")
	}

	if a > b && c > a {
		println("a é maior que b e c é maior que a")
	}

	// switch case
	switch a {
	case 10:
		println("a é igual a 10")
	case 20:
		println("a é igual a 20")
	default:
		println("a não é igual a 10 nem a 20")
	}

}

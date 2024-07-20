// Aula 06 - Slices:
package main

import "fmt"

func main() {
	// colchetes vazios indicam que é um slice
	// seguidos pelo tipo de dado que o slice irá armazenar
	// e por fim os valores que o slice irá armazenar
	s:= []int{10, 20, 30, 40, 50}

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// o valor [:0] indica que queremos apagar todos os elementos do slice
	// a partir do índice 0 inclusive o índice 0
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	// o valor [:2] indica que queremos apagar todos os elementos do slice
	// a partir do índice 2 inclusive o índice 2
	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[:2])

	// o valor [1:2] indica que queremos apagar todos os elementos do slice
	// do indice 2 adiante, inclusive o índice 2
	// mas ignramos tambem o indice 0, já que ele não esta no range de 1 a 2
	fmt.Printf("len=%d cap=%d %v\n", len(s[1:2]), cap(s[1:2]), s[1:2])

	// O valor [3:] indica que queremos apagar todos os elementos do slice
	// antes do índice 3
	fmt.Printf("len=%d cap=%d %v\n", len(s[3:]), cap(s[3:]), s[3:])

	// append adiciona um ou mais elementos ao slice
	// o primeiro argumento é o slice que queremos adicionar os elementos
	// e o segundo argumento é o elemento ou elementos que queremos adicionar
	s = append(s, 60) // o append dobrou a capacidade do slice
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

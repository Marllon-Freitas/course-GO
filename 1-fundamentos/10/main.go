// Aula 10 - Closure:
package main

import (
	"fmt"
)

func main() {
	// Closure é uma função que referencia variáveis de fora do seu corpo
	// Closure é uma função que "lembra" do ambiente em que foi criada
	// Closure é uma função que "captura" variáveis de fora do seu corpo
	// Funções anônimas são closures

	// Exemplo de closure
	x := 10
	closure := func() {
		fmt.Println(x)
	}
	closure()

	// Exemplo de closure com parâmetros
	closure2 := func(x int) {
		fmt.Println(x)
	}
	closure2(20)

	// Exemplo de closure com retorno
	closure3 := func() int {
		return sum(2, 10) * 2 // sum é uma função declarada no arquivo main.go
	}() // os parênteses no final da função fazem com que ela seja executada
	
	fmt.Println(closure3)
}

func sum(a, b int) int {
	return a + b
}

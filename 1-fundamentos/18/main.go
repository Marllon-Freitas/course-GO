// Aula 18 - Interfaces vazias, ou genericas
package main

import "fmt"

// Interface vazia, ou generica, é uma interface que não possui métodos
// Serve para armazenar qualquer tipo de valor
// É o tipo de dado mais genérico em Go

func main() {
	var x interface{} = 10
	var y interface{} = "Hello World"
	var z interface{} = true

	showType(x)
	showType(y)
	showType(z)
}

func showType(i interface{}) {
	fmt.Printf("Type: %T\n", i)
}

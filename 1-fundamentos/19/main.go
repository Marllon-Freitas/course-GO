// Aula 19 - Type assertion
package main

import "fmt"

// Type assertion é uma operação que permite converter um valor de uma interface para um tipo de dado específico
// Sintaxe: valor.(tipo)
// Se o valor armazenado na interface for do tipo especificado, a operação é bem-sucedida
// Caso contrário, a operação falha e o programa é encerrado com um erro

func main() {
	var x interface{} = "10"

	fmt.Println(x)

	// Type assertion
	fmt.Println(x.(string))

	// Type assertion com erro
	// fmt.Println(x.(string))
	res, ok := x.(string)
	if !ok {
		fmt.Println("Erro ao converter para string")
	} else {
		fmt.Println(res)
	}

	// Eu preciso do res e do ok, se não o programa vai dar erro
	res2 := x.(int)
	fmt.Println(res2)
}

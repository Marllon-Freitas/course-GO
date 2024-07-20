// Aula 08 - Funções:
package main

import (
	"errors"
	"fmt"
)

func main() {
	// Funções são trechos de código que podem ser chamados
	// em qualquer parte do código
	// Funções são declaradas com a palavra reservada func
	// func nomeDaFunção(parametros) tipoDeRetorno {
	//		// código
	// }

	// Função sem parâmetros e sem retorno
	helloworld()

	// Função com parâmetros e sem retorno
	helloPerson("José")

	// Função sem parâmetros e com retorno
	fmt.Println(returnHelloWorld())

	// Função com parâmetros e com retorno
	fmt.Println(sum(10, 20))

	// Função que rertorna mais de um valor
	a, b := swap(10, 20)
	fmt.Println(a, b)

	// Função que retorna um valor e um erro
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

// Função sem parâmetros e sem retorno
func helloworld() {
	fmt.Println("Olá Mundo!")
}

// Função com parâmetros e sem retorno
func helloPerson(nome string) {
	fmt.Println("Olá", nome)
}

// Função sem parâmetros e com retorno
func returnHelloWorld() string {
	return "Olá Mundo!"
}

// Função com parâmetros e com retorno
// se os parâmetros forem do mesmo tipo, podemos omitir o tipo de cada um
// ou seja, sum(a, b int) é a mesma coisa que sum(a int, b int)
func sum(a, b int) int {
	return a + b
}

// Função que rertorna mais de um valor
func swap(a, b int) (int, int) {
	return b, a
}

// Função que retorna um valor e um erro
// como Go não tem exceções, é comum retornar um erro
// quando algo inesperado acontece
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("is not possible to divide by zero")
	}
	return a / b, nil
}

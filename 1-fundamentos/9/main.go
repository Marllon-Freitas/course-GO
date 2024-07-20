// Aula 09 - Funções variádicas:
package main

import (
	"fmt"
)

func main() {
	// Funções variádicas são funções que aceitam um número variável de argumentos
	// Para criar uma função variádica, basta adicionar "..." antes do tipo do último parâmetro
	fmt.Println(sum(1, 2, 3, 4))

	// Função variádica com parâmetros normais
	fmt.Println(sum2(10, 20, 30, 40, 50))

	// Função variádica com parâmetros de tipos diferentes
	fmt.Println(sum3("Olá", 1, 2, 3, 4, 5))

	// Função variádica com retorno de mais de um valor
	result, length := sum4(1, 2, 3, 4, 5)
	fmt.Println(result, length)

	// Função variádica com nenhum argumento
	fmt.Println(sum())
}

func sum(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

// Função variádica com parâmetros normais
func sum2(a int, numbers ...int) int {
	result := a
	for _, number := range numbers {
		result += number
	}
	return result
}

// Função variádica com parâmetros de tipos diferentes
func sum3(a string, numbers ...int) string {
	result := a
	for _, number := range numbers {
		result += fmt.Sprint(number)
	}
	return result
}

// Função variádica com retorno de mais de um valor
func sum4(numbers ...int) (int, int) {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result, len(numbers)
}

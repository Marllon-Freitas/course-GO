// Aula 07 - Maps:
package main

import "fmt"

func main() {
	// Maps são uma coleção de chaves e valores HashTable - Objetos no js

	// Declarando um map
	// var nome map[TipoDaChave]TipoDoValor
	salarios := map[string]float64{
		"José":  1000.0,
		"Maria": 2000.0,
		"Pedro": 3000.0,
	}

	fmt.Println(salarios)

	// Adicionando um novo elemento
	salarios["Rafael"] = 4000.0
	fmt.Println(salarios)

	// Deletando um elemento
	delete(salarios, "José")
	fmt.Println(salarios)

	// Iterando sobre um map
	for nome, salario := range salarios {
		fmt.Println(nome, salario)
	}

	// Iterando só sobre as chaves
	for nome := range salarios {
		fmt.Println(nome)
	}

	// Iterando só sobre os valores
	// _ é um underline, é uma forma de ignorar o valor
	for _, salario := range salarios {
		fmt.Println(salario)
	}

	// criando um map vazio
	funcionarios := make(map[string]float64)
	fmt.Println(funcionarios)
	// ou
	funcionarios2 := map[string]float64{}
	fmt.Println(funcionarios2)
}

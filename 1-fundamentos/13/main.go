// Aula 13 - Métodos em Structs:
package main

import "fmt"

// Composição de structs é a capacidade de uma struct ter outra struct como campo
// Aqui estamos criando uma variável do tipo Endereco dentro da struct Cliente
// type Cliente struct {
// 	Nome string
// 	Idade int
// 	Ativo bool
// 	Endereco Endereco
// }


type Endereco struct {
	Logradouro string
	Numero     int
}

// Aqui estamos fazendo a composição de structs
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) maiorDeIdade() bool {
	return c.Idade >= 18
}

func (c *Cliente) desativar() {
	c.Ativo = false
}

func main() {
	marllon := Cliente{
		Nome:  "Marllon",
		Idade: 22,
		Ativo: true,
		Endereco: Endereco{
			Logradouro: "Rua 1",
			Numero:     123,
		},
	}

	// Acessando os campos da struct
	fmt.Println("==============================")
	fmt.Println("| Informações do cliente     |")
	fmt.Println("==============================")
	fmt.Println("| Nome: ", marllon.Nome)
	fmt.Println("| Idade: ", marllon.Idade)
	fmt.Println("| Ativo: ", marllon.Ativo)
	fmt.Println("| Logradouro: ", marllon.Endereco.Logradouro) // ou println(marllon.Logradouro)
	fmt.Println("| Número: ", marllon.Endereco.Numero)     // ou println(marllon.Numero)
	fmt.Println("==============================")

	// Chamando o método
	fmt.Println(marllon.maiorDeIdade())

	// Chamando o método
	marllon.desativar()

	// mostrando o resultado
	fmt.Println(marllon.Ativo)
}

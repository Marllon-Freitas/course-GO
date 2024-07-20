// Aula 14 - Interfaces:
package main

import "fmt"

// Interfaces são tipos que definem um conjunto de métodos que uma struct deve
// implementarpara satisfazer a interface
// Uma interface é um tipo abstrato que é implementado por um tipo concreto

// por exemplo, nosso cliente possui metodos maiorDeIdade e desativar,
// criamos uma interface Pessoa que possui esses métodos
// logo, Cliente implementa a interface Pessoa

// qualquer struct que possua os métodos da interface Pessoa, também implementa
// a interface Pessoa, as funções nas structs e interfaces precisam ser indenticas

// as interfaces do go somente aceitam métodos, não aceitam propriedades

type Endereco struct {
	Logradouro string
	Numero     int
}

type Pessoa interface {
	maiorDeIdade() bool
	desativar()
}

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

// funções mais genéricas
func verificarMaiorDeIdade(p Pessoa) bool {
	return p.maiorDeIdade()
}

func desativarCliente(p Pessoa) {
	p.desativar()
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
	fmt.Println("| Número: ", marllon.Endereco.Numero)         // ou println(marllon.Numero)
	fmt.Println("==============================")

	// Chamando o método
	fmt.Println(marllon.maiorDeIdade())

	// Chamando o método
	marllon.desativar()

	// mostrando o resultado
	fmt.Println(marllon.Ativo)

	// Chamando a função mais genérica, já que marllon é do tipo Cliente
	// que implementa a interface Pessoa
	fmt.Println(verificarMaiorDeIdade(&marllon))
	desativarCliente(&marllon)
}

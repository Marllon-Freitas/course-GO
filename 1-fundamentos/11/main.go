// Aula 11 - Struct:
package main

// Structs são tipos de dados que permitem armazenar diferentes tipos de dados
// em um único tipo de dado

type Cliente struct {
	Nome string
	Idade int
	Ativo bool
}

func main() {
	marllon := Cliente{
		Nome: "Marllon",
		Idade: 22,
		Ativo: true,
	}

	// Acessando os campos da struct
	println(marllon.Nome)
	println(marllon.Idade)
	println(marllon.Ativo)

	// Alterando os campos da struct
	marllon.Nome = "Marllon Teste"
	println(marllon.Nome)

	// Declarando uma struct sem inicializar
	var cliente Cliente
	cliente.Nome = "Cliente"
	cliente.Idade = 30
	cliente.Ativo = false

	println(cliente.Nome)
	println(cliente.Idade)
	println(cliente.Ativo)

	// loop de structs
	clientes := []Cliente{
		{Nome: "Cliente 1", Idade: 20, Ativo: true},
		{Nome: "Cliente 2", Idade: 30, Ativo: false},
		{Nome: "Cliente 3", Idade: 40, Ativo: true},
	}

	for _, cliente := range clientes {
		println(cliente.Nome)
		println(cliente.Idade)
		println(cliente.Ativo)
	}
}

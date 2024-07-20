// Aula 12 - Composição de Structs:
package main

// Composição de structs é a capacidade de uma struct ter outra struct como campo
type Endereco struct {
	Logradouro string
	Numero     int
}

// Aqui estamos criando uma variável do tipo Endereco dentro da struct Cliente
// type Cliente struct {
// 	Nome string
// 	Idade int
// 	Ativo bool
// 	Endereco Endereco
// }

// Aqui estamos fazendo a composição de structs
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
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
	println(marllon.Nome)
	println(marllon.Idade)
	println(marllon.Ativo)
	println(marllon.Endereco.Logradouro) // ou println(marllon.Logradouro)
	println(marllon.Endereco.Numero)     // ou println(marllon.Numero)
}

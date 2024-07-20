// Aula 17 - ponteiros com structs:
package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func (c Cliente) andou() {
	c.Nome = "Teste"
	fmt.Printf("%s andou\n", c.Nome)
}

func (c *Cliente) andouComPonteiro() {
	c.Nome = "Teste"
	fmt.Printf("%s andou\n", c.Nome)
}

type Conta struct {
	Cliente
	Saldo int
}

func (c Conta) sacar(valor int) {
	c.Saldo -= valor
}

func (c *Conta) sacarComPonteiro(valor int) {
	c.Saldo -= valor
}

func main() {
	marllon := Cliente{
		Nome:  "Marllon",
		Idade: 22,
		Ativo: true,
	}

	contaMarllon := Conta{
		Cliente: marllon,
		Saldo:   100,
	}

	// Acessando os campos da struct
	fmt.Println("==============================")
	fmt.Println("| Informações do cliente     |")
	fmt.Println("==============================")
	fmt.Println("| Nome: ", marllon.Nome)
	fmt.Println("| Idade: ", marllon.Idade)
	fmt.Println("| Ativo: ", marllon.Ativo)
	fmt.Println("==============================")

	// Chamando o método
	marllon.andou()
	fmt.Printf("%s\n", marllon.Nome)

	// o resultado será:
	// Teste andou
	// Marllon

	// o que aconteceu foi que o método andou() alterou o valor do campo Nome
	// da struct Cliente, mas como o método não recebe um ponteiro, ele não altera
	// o valor do campo Nome da struct Cliente marllon
	// para resolver isso, precisamos passar um ponteiro para o método andou()
	// para que ele possa alterar o valor do campo Nome da struct Cliente marllon

	// Chamando o método
	marllon.andouComPonteiro()
	fmt.Printf("%s\n", marllon.Nome)

	// o resultado será:
	// Teste andou
	// Teste

	// o que aconteceu foi que o método andouComPonteiro() alterou o valor do campo Nome
	// da struct Cliente, e como o método recebe um ponteiro, ele altera
	// o valor do campo Nome da struct Cliente marllon

	// Acessando os campos da struct
	fmt.Println("==============================")
	fmt.Println("| Informações da conta       |")
	fmt.Println("==============================")
	fmt.Println("| Nome: ", contaMarllon.Nome)
	fmt.Println("| Idade: ", contaMarllon.Idade)
	fmt.Println("| Ativo: ", contaMarllon.Ativo)
	fmt.Println("| Saldo: ", contaMarllon.Saldo)
	fmt.Println("==============================")

	contaMarllon.sacar(50)
	fmt.Printf("Saldo após sacar 50 reais sem ponteiro: %d\n", contaMarllon.Saldo)

	// o resultado será:
	// Saldo após sacar 50 reais sem ponteiro: 100

	contaMarllon.sacarComPonteiro(50)
	fmt.Printf("Saldo após sacar 50 reais com ponteiro: %d\n", contaMarllon.Saldo)
}

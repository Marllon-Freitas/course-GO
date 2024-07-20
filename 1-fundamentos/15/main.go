// Aula 15 - Ponteiros:
package main

// Ponteiros são variáveis que armazenam o endereço de memória de outra variável
// Em Go, os ponteiros são representados pelo operador & (comercial) seguido da
// variável que queremos obter o endereço de memória
// Para acessar o valor de uma variável ponteiro, utilizamos o operador * (asterisco)
// seguido da variável ponteiro, o * significa desreferenciar, ou seja, acessar o valor
// da variável que o ponteiro aponta

// Exemplo de ponteiro

func main() {
	x := 10
	y := &x
	z := &y // ponteiro de ponteiro

	// de maneira mais explicita
	var c int = 10;
	var d *int = &c;

	// podemos alterar o valor de x através do ponteiro y
	println("x antes:", x) // 10
	println("z => ponteiro do ponteiro de y:", z)

	*y = 20

	println("x depois:", x) // 20
	println("y:", y) // endereço de x
	println("*y:", *y) // 20

	println("c antes:", c) // 10
	println("d:", d) // endereço de c
	println("d:", *d) // 10

	// podemos alterar o valor de c através do ponteiro d
	*d = 30

	println("c depois:", c) // 30
	println("d:", *d) // 30
}

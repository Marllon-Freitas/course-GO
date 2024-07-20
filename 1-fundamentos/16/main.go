// Aula 16 - Quando usar ponteiros:
package main

// Ponteiros são utilizados para passar referências de memória de uma variável
// para outra, evitando assim a cópia de valores entre as variáveis
// Ponteiros são utilizados para alterar o valor de uma variável fora do escopo
// de uma função ou método

// Exemplo de ponteiro

// nessa função estamos passando uma cópia do valor de a e b
// ou seja, a função soma não altera os valores originais de a e b
func soma(a, b int) int {
	a = a + 1
	b = b + 1
	return a + b
}

func somaComPonteiro(a, b *int) int {
	*a = *a + 1
	*b = *b + 1
	return *a + *b
}
 
func main() {
	a := 10
	b := 20

	// a função soma não altera os valores originais de a e b
	println("soma:", soma(a, b)) // 30
	println("a:", a) // 10
	println("b:", b) // 20

	// a função somaComPonteiro altera os valores originais de a e b
	println("somaComPonteiro:", somaComPonteiro(&a, &b)) // 32
	println("a:", a) // 11
	println("b:", b) // 21
}

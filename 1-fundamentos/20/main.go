// Aula 20 - Generics
package main

// Generics é um recurso que permite escrever funções, estruturas e interfaces
// que trabalham com tipos genéricos

// codigo que soma os valores inteiros de um map
func somaInt(m map[string]int) int {
	s := 0
	for _, v := range m {
		s += v
	}
	return s
}

type MyNumber int

// nos tambem podemos usar constraints para limitar os tipos que podem ser usados
type Number interface {
	~int | ~float64
}

// se eu quiser somar valores float64, eu preciso criar uma nova função
// mas com generics, eu posso criar uma única função que soma
// valores do tipo estabelicido, nesse caso int ou float64
func soma[T Number | float64](m map[string]T) T {
	var s T
	for _, v := range m {
		s += v
	}
	return s
}

// comperable
// o codigo abaixo recebe dois valores de qualquer tipo e compara se eles são iguais
// mas o go reclama, pois não é possivel comparar valores de tipos diferentes, e não
// temos como saber se o tipo que estamos recebendo é comparavel
// func compara[T any] (a, b T) bool {
// 	if a == b {
// 		return true
// 	}
// 	return false
// }
// para resolver isso, podemos usar o comparable, que compara dois tipos, que sejam
// comparaveis
func compara[T comparable] (a, b T) bool {
	return a == b
}

func main() {
	m := map[string]int{
		"um":   1,
		"dois": 2,
		"tres": 3,
	}
	println(somaInt(m))

	m2 := map[string]float64{
		"um":   1.1,
		"dois": 2.2,
		"tres": 3.3,
	}

	println(soma(m2))

	// mesmo meu tipo novo MyNumber sendo do tipo int, e meu Number aceitando int
	// o go reclama, o codigo abaixo não funciona

	// type MyNumber int
	// type Number interface {
	// 	int | float64
	// }

	// m3 := map[string]MyNumber{
	// 	"um":   1,
	// 	"dois": 2.2,
	// 	"tres": 3,
	// }

	// println(soma(m3))

	// mas se eu quiser criar uma exceção para o MyNumber, eu posso fazer isso
	// adicionando o ~ antes do tipo na minha interface Number
	// type MyNumber int
	// type Number interface {
	// 	~int | float64
	// }

	m3 := map[string]MyNumber{
		"um":   1,
		"dois": 2,
		"tres": 3,
	}

	println(soma(m3))

	println(compara(10, 10.0))
}

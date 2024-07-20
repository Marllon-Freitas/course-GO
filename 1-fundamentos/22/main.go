// Aula 22 - Loops
package main

// no go não tem while, só tem for

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	// for com apenas uma condição
	j := 0
	for j < 10 {
		println(j)
		j++
	}

	// for infinito
	// for {
	// 	println("Loop infinito")
	// }

	// for com range
	slice := []int{1, 2, 3, 4, 5}
	for i, v := range slice {
		println(i, v)
	}

	// for com range e ignorando o índice
	for _, v := range slice {
		println(v)
	}

	// for com range em map
	m := map[string]int{
		"um": 1,
		"dois": 2,
		"tres": 3,
	}
	for k, v := range m {
		println(k, v)
	}

	// for com range em string
	str := "Hello World"
	for i, v := range str {
		println(i, string(v))
	}
}

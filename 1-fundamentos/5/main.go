// Aula 05 - Percorrendo Arrays:
package main

import "fmt"

func main() {
	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3

	//pegar ultimo elemento do array
	println(myArray[len(myArray)-1])
	// pegar o primeiro elemento do array
	println(myArray[0])
	// pegar o tamanho do array
	println(len(myArray))

	// percorrer o array
	for i := 0; i < len(myArray); i++ {
		println(myArray[i])
	}

	// percorrer o array com range
	for i, value := range myArray {
		fmt.Printf("Index: %v, Value: %v\n", i, value)
	}
}

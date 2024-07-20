// Aula 04 - Importação de fmt e tipagem:
package main

import "fmt"

type ID int

var (
	f ID = 1
)

func main() {
	fmt.Printf("O tipo de f é: %v", f)
}

// Aula 21 - Pacotes e módulos
package main

import (
	math "curso-go/full-cycle/1-fundamentos/21/matematica"
	"fmt"

	"github.com/google/uuid"
)

// funções que iniciam com letra maiúscula são públicas
// funções que iniciam com letra minúscula são privadas
// isso também vale para variáveis, constantes, structs, etc
// dentro de uma struct, se o nome de um campo começar com letra maiúscula, ele é público
func main() {
	s := math.Soma(10, 20)
	uuid := uuid.New()
	fmt.Println(uuid)
	fmt.Println(s)
}

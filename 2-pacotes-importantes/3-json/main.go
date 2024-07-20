package main

import (
	"encoding/json"
	"os"
)

// manipulação de json em go

type Conta struct {
	// o recursp a seguir se chamam tags, são informações adicionais que você pode passar para o compilador
	// podemos omitir o campo no json com `json:"-"`
	Numero int // tambem é possivel usar `json:"n"` para mudar o nome do campo no json
	Saldo  int // tambem é possivel usar `json:"s"` para mudar o nome do campo no json
}

func main() {
	conta := Conta{
		Numero: 1,
		Saldo:  1000,
	}

	res, err := json.Marshal(conta) // ao usar marshal, você pega o valor do json para você usar
	if err != nil {
		panic(err)
	}

	println(string(res))

	// criar um  encoder
	err = json.NewEncoder(os.Stdout).Encode(conta) // encode pega o valor do json e joga para algum lugar
	if err != nil {
		panic(err)
	}

	// decodificar um json
	var contaDecodificada Conta
	jsonPure := []byte(`{"Numero":1,"Saldo":1000}`) // json puro
	// obs: para a assossiação de valores, o nome das variáveis devem ser iguais, o json deve possuir os mesmo campos da struct
	// o código a seguir é inválido
	// jsonPure := []byte(`{"numero":1,"saldo":1000}`)
	// resultado: 0 0
	err = json.Unmarshal(jsonPure, &contaDecodificada)
	if err != nil {
		panic(err)
	}

	println(contaDecodificada.Numero)
	println(contaDecodificada.Saldo)
}

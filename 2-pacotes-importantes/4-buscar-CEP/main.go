package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type CEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição %v\n", err)
			os.Exit(1)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta %v\n", err)
			os.Exit(1)
		}

		var data CEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao decodificar json %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("====================================\n")
		fmt.Printf("CEP: %s\n", data.Cep)
		fmt.Printf("Logradouro: %s\n", data.Logradouro)
		fmt.Printf("Complemento: %s\n", data.Complemento)
		fmt.Printf("Bairro: %s\n", data.Bairro)
		fmt.Printf("Localidade: %s\n", data.Localidade)
		fmt.Printf("UF: %s\n", data.Uf)
		fmt.Printf("IBGE: %s\n", data.Ibge)
		fmt.Printf("GIA: %s\n", data.Gia)
		fmt.Printf("DDD: %s\n", data.Ddd)
		fmt.Printf("SIAFI: %s\n", data.Siafi)
		fmt.Printf("====================================\n")

		file, err := os.Create("cep.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Logradouro: %s, Complemento: %s, Bairro: %s, Localidade: %s, UF: %s, IBGE: %s, GIA: %s, DDD: %s, SIAFI: %s\n", data.Cep, data.Logradouro, data.Complemento, data.Bairro, data.Localidade, data.Uf, data.Ibge, data.Gia, data.Ddd, data.Siafi))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao escrever no arquivo %v\n", err)
			os.Exit(1)
		}
	}
}

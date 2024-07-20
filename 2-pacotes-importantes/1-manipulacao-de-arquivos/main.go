package main

import (
	"bufio"
	"os"
)

func main() {
	// Abrindo um arquivo para escrita
	arquivo, err := os.Create("log.txt")
	if err != nil {
		panic(err)
	}

	// tamanho, err := arquivo.WriteString("Olá mundo!")
	tamanho, err := arquivo.Write([]byte("Olá mundo!"))
	if err != nil {
		panic(err)
	}

	// quantidade de bytes escritos
	println(tamanho)

	// fechando o arquivo
	arquivo.Close()

	// lendo o conteúdo do arquivo
	var conteudo []byte
	conteudo, err = os.ReadFile("log-leitura.txt")
	if err != nil {
		panic(err)
	}
	println(string(conteudo))

	// Ler um arquivo de pedaço em pedaço
	arquivoGrande, err := os.Open("log-grande.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivoGrande)
	buffer := make([]byte, 100) // meu buffer ler de 100 em 100 bytes
	for {
		tamanho, err := reader.Read(buffer)
		if err != nil {
			break
		}
		println(tamanho)
		println(string(buffer[:tamanho]))
	}

	// removendo um arquivo
	err = os.Remove("log.txt")
	if err != nil {
		panic(err)
	}
}

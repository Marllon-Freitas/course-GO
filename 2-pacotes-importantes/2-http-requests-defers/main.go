package main

import (
	"io"
	"net/http"
)

// defer: adia a execução de uma função até o final da função atual

func main () {
	req, err := http.Get("https://www.google.com.br")
	if err != nil {
		panic(err)
	}

	defer req.Body.Close() // fecha a conexão após a execução da função main

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	println(string(res))

}
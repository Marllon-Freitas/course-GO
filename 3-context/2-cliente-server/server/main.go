/*
	o servidor vai esperar 5 segundos para processar a requisição e retornar ao
	usuário que a requisição foi processada com sucesso.
	Se passar 5 segundos, o usuário vai receber um erro de rquest timeout.
*/

package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // pega o contexto da requisição
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")

	select {
	case <-time.After(5 * time.Second):
		log.Println("Processamento finalizado com sucesso") // imprime no console
		// imprime no browser
		w.Write([]byte("Processamento finalizado com sucesso"))
	case <-ctx.Done():
		err := ctx.Err()
		log.Println("Request cancelada pelo cliente:", err) // imprime no console
	}
}

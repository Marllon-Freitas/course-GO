package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.Handle("/blog", &blogHandler{})
	http.ListenAndServe(":8080", mux)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", homeHandler2)
	http.ListenAndServe(":8081", mux2)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá, mundo!"))
}

func homeHandler2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá, mundo! 2"))
}

type blogHandler struct{}
func (b *blogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá, blog!"))
}
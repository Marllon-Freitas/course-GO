package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/blog", handleBlog)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handleBlog(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá, blog!"))
}

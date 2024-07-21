package main

import (
	"text/template"
	"log"
	"net/http"
)

type Curso struct {
	Titulo       string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := t.Execute(w, Cursos{
			Curso{"Go", 40},
			Curso{"Python", 45},
			Curso{"Java", 60},
		})
		if err != nil {
			log.Fatal(err)
		}
	})
	http.ListenAndServe(":8282", nil)
}

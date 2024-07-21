package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

type Curso struct {
	Titulo       string
	CargaHoraria int
}

type Cursos []Curso

func toUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.New("content.html")
	t.Funcs(template.FuncMap{
		"toUpper": toUpper,
	})
	t = template.Must(t.ParseFiles(templates...))
	err := t.Execute(os.Stdout, Cursos{
		Curso{"Go", 40},
		Curso{"Python", 45},
		Curso{"Java", 60},
	})
	if err != nil {
		log.Fatal(err)
	}
}

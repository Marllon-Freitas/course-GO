package main

import (
	"text/template"
	"log"
	"os"
)

type Curso struct {
	Titulo       string
	CargaHoraria int
}

func main() {
	curso := Curso{"Go", 40}
	tmp := template.New("curso")
	tmp, err := tmp.Parse("Curso: {{.Titulo}}\nCarga Horária: {{.CargaHoraria}} horas")
	if err != nil {
		log.Fatal(err)
	}

	err = tmp.Execute(os.Stdout, curso)
	if err != nil {
		log.Fatal(err)
	}
}

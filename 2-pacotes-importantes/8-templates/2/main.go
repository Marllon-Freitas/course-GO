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
	t := template.Must(template.New("curso").Parse("Curso: {{.Titulo}}\nCarga Horária: {{.CargaHoraria}} horas"))
	err := t.Execute(os.Stdout, curso)
	if err != nil {
		log.Fatal(err)
	}
}

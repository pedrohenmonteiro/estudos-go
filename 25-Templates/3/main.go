package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	t := template.Must(template.New("template.html").ParseFiles("/home/pedro/curso-go/25-Templates/3/template.html"))
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 40},
		{"Python", 40},
	})
	if err != nil {
		panic(err)
	}
}

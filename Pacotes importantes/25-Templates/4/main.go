package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("/home/pedro/curso-go/25-Templates/4/template.html"))
		err := t.Execute(w, Cursos{
			{"Go", 40},
			{"Python", 40},
			{"Java", 40},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8081", nil)
}

package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{"Go", 40}
	//tmp := template.New("CursoTemplate")
	//tmp, _ = tmp.Parse("Curso: {{.Nome}} - Carga Horaria: {{.CargaHoraria}}")
	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} - Carga Horaria: {{.CargaHoraria}}"))
	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}

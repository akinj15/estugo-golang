package main

import (
	"os"
	"text/template"
)

type curso struct {
	nome         string
	cargahoraria int
}

func main() {
	curso := curso{"go", 40}
	t := template.Must(template.New("cursotemplate").Parse("curso: {{.nome}} - carga horaria: {{.cargahoraria}}"))
	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}

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
	tmp := template.New("cursotemplate")
	tmp, _ = tmp.Parse("curso: {{.nome}} - carga horaria: {{.cargahoraria}}")
	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}

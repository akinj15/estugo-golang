package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	Cargahoraria int
}
type Cursos []Curso

func main() {
	http.ListenAndServer(":8888", nil)
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Cursos{
		{"go", 40},
		{"java", 40},
		{"python", 40},
	})
	if err != nil {
		panic(err)
	}
}

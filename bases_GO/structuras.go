package main

import "fmt"

func main() {

	curso1 := Curso{
		nombre:    "GO",
		url:       "http://",
		habilidad: []string{"backend", "2"},
	}

	fmt.Println(curso1)
}

type Curso struct {
	nombre    string
	url       string
	habilidad []string
}

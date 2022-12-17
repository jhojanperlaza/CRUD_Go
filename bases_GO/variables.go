package main

import "fmt"

func main() {

	// definir una variable en formato largo
	var nombre string

	// ahora le podemos asignar un valor
	nombre = "jhojan"

	fmt.Println(nombre)

	// definir varias variables de un mismo tipo
	var a, b, c int = 1, 2, 3

	fmt.Println(a, b, c)

	// definir variables de varios tipos
	var (
		pi       float64
		booleano bool
		cadena   = "texto" //al asignarle automaticamente se vuelve de tipo string
		edad     = 25      // al usar el var no necesito los := para asignar valor
	)
	pi = 3.14
	booleano = false

	fmt.Println(pi, booleano, cadena, edad)

	// definir variable de forma ganster
	v1 := 123

	// ya cuando lo vayamos a modificar:

	v1 = 456

	fmt.Println(v1)
}

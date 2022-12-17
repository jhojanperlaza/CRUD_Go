package main

import "fmt"

func main() {

	var (
		nombre string
		edad   int
		pi     float64
	)

	fmt.Print("Ingrese su Nombre")
	fmt.Scanln(&nombre)

	fmt.Print("Ingrese su Edad")
	fmt.Scanln(&edad)

	fmt.Print("Ingrese el valor de pi")
	fmt.Scanln(&pi)

	fmt.Printf("Nombre: %s Edad: %d\n pi: %f", nombre, edad, pi)
}

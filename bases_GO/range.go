package main

import "fmt"

func main() {

	array1 := []int{1, 2, 3, 4, 5}

	//cuando no necesito el indice

	for _, num := range array1 {
		fmt.Println(num)
	}

	// si quiero los dos:

	for index, num := range array1 {
		fmt.Println(index, "==>", num)
	}

	// tambien puedo iterar maps

	colores := map[string]string{
		"yellow": "amarillo",
		"red":    "rojo",
		"blue":   "azul",
	}

	for k, v := range colores {
		fmt.Println(k, "==>", v)
	}
}

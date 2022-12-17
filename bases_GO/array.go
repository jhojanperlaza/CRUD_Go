package main

import "fmt"

func main() {

	// declarar un array vacio con tamaño 2 y tipo strings
	var array1 [2]string

	array1[0] = "hello"
	array1[1] = "world"

	fmt.Println(array1)

	// definir y declarar un array al mimso tiempo
	array2 := [3]int{1, 2, 3}

	fmt.Println(array2)

	//slicen permite guardar un array sin definir la cantidad de espacio

	var slicen1 []string

	slicen1 = append(slicen1, "Dato1")
	slicen1 = append(slicen1, "Dato2", "Dato3", "Dato4")

	fmt.Println(slicen1)

}

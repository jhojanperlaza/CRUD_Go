package main

import "fmt"

func main() {

	// son como los diccionarios
	map1 := map[string]int{
		"uno":  1,
		"dos":  2,
		"tres": 3,
	}

	//agregar
	map1["cuatro"] = 4
	map1["cinco"] = 5

	fmt.Println(map1)

	//eliminar
	delete(map1, "cuatro")

	fmt.Println(map1)

	//definir directamente

	map2 := make(map[int]string)

	map2[1] = "uno"
	map2[2] = "dos"
	map2[3] = "tres"

	fmt.Println(map2)

}

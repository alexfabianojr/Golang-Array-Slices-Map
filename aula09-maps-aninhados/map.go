package main

import "fmt"

func main() {

	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela": 12.1,
			"Guga":     67.3,
		},
		"A": {
			"Alex": 999.9,
		},
		"M": {
			"Maria":  43.12,
			"Marcos": 31231.1,
			"Mateus": 907.3,
		},
	}

	fmt.Println(funcsPorLetra)

	delete(funcsPorLetra, "A")

	fmt.Println(funcsPorLetra)

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}

package main

import "fmt"

func main() {

	//arrays sao homogeneos (mesmo tipo de dados) e estatica (tamanho fixo)

	var notas [3]float64 //[0 0 0] - nao inicializado

	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1 //array out of bounds eh erro de compilacao

	fmt.Println(notas)

	total := 0.0

	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))

	fmt.Printf("Media: %.2f", media)
}

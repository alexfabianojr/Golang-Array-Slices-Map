package main

import "fmt"

func main() {

	aprovados := make(map[int]string)

	aprovados[123] = "Maria"
	aprovados[456] = "Pedro"
	aprovados[789] = "Joao"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Println(nome, cpf)
	}

	fmt.Println(aprovados[123])

	delete(aprovados, 456)

	fmt.Println(aprovados)
}

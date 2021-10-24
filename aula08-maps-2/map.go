package main

import "fmt"

func main() {

	funcESalarios := map[string]float64{
		"jose":     2312.43,
		"Gabriela": 6554342.21,
		"Antonio":  9987.32,
	}

	funcESalarios["Maria"] = 999.9

	fmt.Println(funcESalarios)

	delete(funcESalarios, "Inexistente")

	fmt.Println(funcESalarios)

	for _, salario := range funcESalarios {
		fmt.Println(salario)
	}

	for chave := range funcESalarios {
		fmt.Println(chave)
	}

	for nome, salario := range funcESalarios {
		fmt.Println(nome, salario)
	}

}

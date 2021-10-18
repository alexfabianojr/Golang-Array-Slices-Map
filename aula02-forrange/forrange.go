package main

import "fmt"

func main() {

	numeros := [...]int{7, 2, 9, 4, 5} //compilador conta! array

	for i, numero := range numeros {
		fmt.Printf("%d ) %d \n", i, numero)
	}

	for _, numero := range numeros {
		fmt.Println(numero)
	}

}

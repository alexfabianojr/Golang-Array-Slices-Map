package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} //array
	s1 := []int{1, 2, 3}  //slice

	fmt.Println(a1, s1)

	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	//Slice nao eh um array! O slice define um pedaco de um array

	a2 := [5]int{1, 2, 3, 4, 5}

	s2 := a2[1:3] //indo do indice um ao indice 3 sem incluir o 3

	//slice nao cria um array mas eh uma referencia de memoria a um pedaco de um array

	fmt.Println(s2)

	//Novo slice apontando para o mesmo slice
	s3 := a2[:2]

	fmt.Println(s3)

	s4 := s2[:1]

	fmt.Println(s4)
}

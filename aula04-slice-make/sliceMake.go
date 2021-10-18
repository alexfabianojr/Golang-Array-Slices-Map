package main

import "fmt"

func main() {

	s := make([]int, 10) //tipo, tamanho

	s[9] = 12

	fmt.Println(s[9])
	fmt.Println(s[2])

	s = make([]int, 10, 20) // tipo, tamanho inicial, capacidade maxima

	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)

	fmt.Println(s, len(s), cap(s))
}

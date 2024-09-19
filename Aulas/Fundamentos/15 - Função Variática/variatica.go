package main

import "fmt"

func soma(numeros ...int) (soma int) {
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func main() {
	fmt.Println(soma(1, 2, 3, 4))
}

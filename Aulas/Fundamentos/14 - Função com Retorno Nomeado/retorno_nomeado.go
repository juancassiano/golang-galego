package main

import "fmt"

func calculosMatematico(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	fmt.Println(calculosMatematico(5, 5))
}

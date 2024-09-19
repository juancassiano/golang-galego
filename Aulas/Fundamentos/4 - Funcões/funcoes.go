package main

import "fmt"

func somar(n1, n2 int) int {
	return n1 + n2
}

func calculos(n1, n2 int) (int8, int8) {
	soma := int8(n1 + n2)
	subtracao := int8(n1 - n2)
	return soma, subtracao
}

func main() {
	soma := somar(5, 8)
	fmt.Println(soma)
	resultado1, resultado2 := calculos(5, 8)
	resultado3, _ := calculos(5, 8)
	fmt.Println(resultado1, resultado2, resultado3)
}

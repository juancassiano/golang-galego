package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1"
	variavel2 := "Variável 2"
	variavel3, variavel4 := "Variável 3", "Variável 4"
	var (
		variavel5 string = "Variável 5"
		variavel6 string = "Variável 6"
	)

	const constante1 string = "Constante 1"
	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(variavel3)
	fmt.Println(variavel4)
	fmt.Println(variavel5)
	fmt.Println(variavel6)
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5)
	fmt.Println(variavel6)
}

package main

import "fmt"

func main() {

	func() {
		fmt.Println("Olá")
	}()
	retorno := func(texto string) string {
		return texto
	}("Mundo")
	fmt.Println(retorno)
}

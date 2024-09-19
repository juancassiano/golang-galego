package main

import "fmt"

func main() {
	numero := 10
	if numero > 10 {
		fmt.Println("Maior que 10")
	} else {
		fmt.Println("Menor ou igual a 10")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Maior que zero")
	} else if numero < -10 {
		fmt.Println("Menor que -10")
	} else {
		fmt.Println("Entre 0 e")
	}
}

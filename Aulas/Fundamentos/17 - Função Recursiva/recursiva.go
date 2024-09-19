package main

import "fmt"

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-1) + fibonacci(posicao-2)
}

func main() {
	// fibonacci := fibonacci(10)
	for i := 0; i < 15; i++ {
		fmt.Println(fibonacci(i))
	}
	// fmt.Println(fibonacci)

}

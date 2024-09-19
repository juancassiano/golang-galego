package main

import (
	"errors"
	"fmt"
)

func main() {
	//Inteiros
	var x int8 = 1
	var y int16 = 1
	var z int32 = 1
	var w int64 = 1
	var a uint8 = 1
	fmt.Println(x, y, z, w, a)

	//Reais
	var f1 float32 = 1.0
	var f2 float64 = 1.0
	fmt.Println(f1, f2)

	//String
	var str string = "Hello, World!"
	fmt.Println(str)

	//Booleano
	var boolean1 bool = true
	fmt.Println(boolean1)

	//Erro
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}

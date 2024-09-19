package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	auxiliar.Escrever()
	err := checkmail.ValidateFormat("juan@hotmail.com")
	fmt.Println(err)
}

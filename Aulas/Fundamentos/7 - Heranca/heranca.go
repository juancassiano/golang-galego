package main

import "fmt"

type pessoa struct {
	nome      string
	idade     uint8
	altura    uint8
	sobrenome string
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"João", 20, 178, "Silva"}
	fmt.Println(p1)
	e1 := estudante{p1, "java", "casa"}
	fmt.Println(e1)
}

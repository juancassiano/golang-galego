package main

import "fmt"

func funcao1() {
	fmt.Println("Executando função 1")
}

func funcao2() {
	fmt.Println("Executando função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Media calculada, o resultado será retornado")

	fmt.Println("Entrando na função verifica aluno aprovado")
	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main() {
	//Defer = adiar
	defer funcao1()
	funcao2()
	fmt.Println(alunoEstaAprovado(8, 5))
}

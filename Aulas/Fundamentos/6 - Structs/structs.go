package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	user := usuario{
		nome:  "Fulano",
		idade: 22,
	}
	fmt.Println(user.nome)
	fmt.Println(user.idade)
	fmt.Println(user)

	user2 := usuario{idade: 33}
	fmt.Println(user2)

	end := endereco{
		logradouro: "Rua dos Bobos",
		numero:     0,
	}
	user3 := usuario{
		nome:     "Beltrano",
		idade:    44,
		endereco: end,
	}
	fmt.Println(user3)
}

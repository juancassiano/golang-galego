package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario := usuario{"Jõao", 22}
	fmt.Println(usuario)
	usuario.salvar()
	fmt.Println(usuario.maiorDeIdade())
	usuario.fazerAniversario()
	fmt.Println(usuario.idade)
}

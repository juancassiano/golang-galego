package main

import (
	"fmt"
	"introducao/enderecos"
)

func main() {
	tipoDeEndereco := enderecos.TipoDeEndereco("Avenida Paulista")
	fmt.Println(tipoDeEndereco)
}

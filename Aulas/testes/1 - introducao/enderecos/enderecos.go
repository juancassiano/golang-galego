package enderecos

import "strings"

// TipoDeEndereco veriica se um endereço tem um tipo válido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	primeiraPalavraDoEndereco := strings.Split(strings.ToLower(endereco), " ")[0]

	enderecoTemTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemTipoValido = true
		}
	}
	if enderecoTemTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}
	return "Tipo Inválido"
}

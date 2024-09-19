package enderecos_test

import (
	. "introducao/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	cenariosDeTeste :=
		[]cenarioDeTeste{
			{"Rua ABC", "Rua"},
			{"Avenida ABC", "Avenida"},
			{"Estrada ABC", "Estrada"},
			{"Rodovia ABC", "Rodovia"},
			{"Praça ABC", "Tipo Inválido"},
			{"RUA ABC", "Rua"},
			{"aVENIDA ABC", "Avenida"},
			{"", "Tipo Inválido"},
		}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				tipoDeEnderecoRecebido, cenario.retornoEsperado)
		}
	}
}

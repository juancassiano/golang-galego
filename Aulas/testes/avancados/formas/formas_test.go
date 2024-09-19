package formas_test

import (
	"math"
	"teste-avancado/formas"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		ret := formas.Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaRecebida != areaEsperada {
			t.Errorf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Círculo", func(t *testing.T) {
		circ := formas.Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()

		if areaRecebida != areaEsperada {
			t.Errorf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})
}

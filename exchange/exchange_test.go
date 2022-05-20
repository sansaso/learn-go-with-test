package exchange

import (
	"testing"
)

func TestCarteira(t *testing.T) {
	carteira := Carteira{}

	carteira.Depositar(Bitcoin(10))

	resultado := carteira.Saldo()

	esperado := Bitcoin(10)

	if resultado != esperado {
		t.Errorf("resultado %d, esperado %d", resultado, esperado)
	}
}

package formas

import (
	"testing"
)

func TestPerimetro(t *testing.T) {
	retangulo := Retangulo{10.0, 10.0}
	resultado := Perimetro(retangulo)
	esperado := 40.0

	if resultado != esperado {
		t.Errorf("resultado %.2f esperado %.2f", resultado, esperado)
	}
}

func TestArea(t *testing.T) {

	testesArea := []struct {
		nome     string
		forma    Forma
		esperado float64
	}{
		{nome: "Retangulo", forma: Retangulo{Largura: 12, Altura: 6}, esperado: 72.0},
		{nome: "Circulo", forma: Circulo{Raio: 10}, esperado: 314.1592653589793},
		{nome: "Triangulo", forma: Triangulo{Base: 12, Altura: 6}, esperado: 36.0},
	}

	for _, tt := range testesArea {
		t.Run(tt.nome, func(t *testing.T) {
			resultado := tt.forma.Area()
			if resultado != tt.esperado {
				t.Errorf("resultado %.2f, esperado %.2f", resultado, tt.esperado)
			}
		})
	}
}

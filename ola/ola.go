package main

import (
	"fmt"
)

const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlafrances = "Bonjour, "

func prefixoSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case "Frances":
		prefixo = prefixoOlafrances
	case "Espanhol":
		prefixo = prefixoOlaEspanhol
	default:
		prefixo = prefixoOlaPortugues
	}
	return
}
func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "mundo"
	}
	return prefixoSaudacao(idioma) + nome
}

func main() {
	fmt.Println(Ola("mundo", ""))
}

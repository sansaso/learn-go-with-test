package main

import (
	"errors"
)

type Dicionario map[string]string

var (
	ErrNaoEncontrado      = errors.New("não foi possível encontrar a palavra que você procura")
	ErrPalavraExistente   = errors.New("não é possível adicionar a palavra pois ela já existe")
	ErrPalavraInexistente = errors.New("não foi possível atualizar a palavra pois ela não existe")
)

type ErrDicionario string

func (e ErrDicionario) Error() string {
	return string(e)
}

func (d Dicionario) Busca(palavra string) (string, error) {

	definicao, existe := d[palavra]

	if !existe {
		return "", ErrNaoEncontrado
	}

	return definicao, nil
}

func (d Dicionario) Adiciona(palavra, definicao string) error {
	_, err := d.Busca(palavra)

	switch err {
	case ErrNaoEncontrado:
		d[palavra] = definicao
	case nil:
		return ErrPalavraExistente
	default:
		return err
	}

	return nil
}

func (d Dicionario) Atualiza(palavra, novaDefinicao string) error {
	_, err := d.Busca(palavra)

	switch err {
	case ErrNaoEncontrado:
		return ErrPalavraInexistente
	case nil:
		d[palavra] = novaDefinicao
	default:
		return err
	}

	return nil
}

func (d Dicionario) Deleta(palavra string) {
	delete(d, palavra)
}

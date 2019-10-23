package model

import "errors"

//ErrValorDeImovelInvalido : é um erro que é apresentado quando é a atribuindo a um imóvel um valor muito baixo
var ErrValorDeImovelInvalido = errors.New("O valor informado não é valido para um imóvel")

//ErrValorDeImovelMuitoAlto : é um erro que é apresentado quando é a atribuindo a um imóvel um valor muito alto
var ErrValorDeImovelMuitoAlto = errors.New("O valor informado é muito alto")

//Imovel : é uma struct que armazena dados de um imóvel
type Imovel struct {
	X     int    `json:"coordendaX"`
	Y     int    `json:"coordendaY"`
	Nome  string `json:"nome"`
	valor int
}

//SetValor : define o valor do imóvel
func (i *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor < 50000 {
		err = ErrValorDeImovelInvalido
		return
	} else if valor > 10000000 {
		err = ErrValorDeImovelMuitoAlto
		return
	}
	i.valor = valor
	return
}

//GetValor : retorna o valor do imóvel
func (i *Imovel) GetValor() int {
	return i.valor
}

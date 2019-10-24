package main

import (
	"fmt"

	"./model"
)

func main() {

	jojo := model.Ave{}
	jojo.Nome = "Jojo da Silva"
	queroAcordarComUmCacarejo(jojo)
	queroOuvirUmPatoNoLago(jojo)
}

func queroAcordarComUmCacarejo(g model.Galinha) {
	fmt.Println(g.Cacareja())
}

func queroOuvirUmPatoNoLago(p model.Pato) {
	fmt.Println(p.Quack())
}

package main

import "fmt"

//Imovel : é uma struct que armazena dados de um imóvel
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {
	casa := Imovel{}
	fmt.Printf("A casa é: %+v\r\n", casa)

	apartamento := Imovel{17, 56, "Meu AP", 760000}
	fmt.Printf("A casa é: %+v\r\n", apartamento)

	chacara := Imovel{
		Y:     85,
		Nome:  "Chacara",
		X:     22,
		valor: 55,
	}
	fmt.Printf("A casa é: %+v\r\n", chacara)

	casa.Nome = "Lar Doce Lar"
	casa.valor = 450000
	casa.X = 18
	casa.Y = 31
	fmt.Printf("A casa é: %+v\r\n", casa)

	//ponteiros
	casaDois := new(Imovel)
	fmt.Printf("A casa é: %p - %+v\r\n", &casaDois, casaDois)

	chacaraDois := Imovel{17, 28, "Chacara Linda", 280000}
	fmt.Printf("A casa é: %p - %+v\r\n", &chacaraDois, chacaraDois)
	mudaImovel(&chacaraDois)
	fmt.Printf("A casa é: %p - %+v\r\n", &chacaraDois, chacaraDois)
}

func mudaImovel(imovel *Imovel) {
	imovel.X = imovel.X + 10
	imovel.Y = imovel.Y - 5
}

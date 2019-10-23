package main

import (
	"encoding/json"
	"fmt"

	"../struts_avancado/model"
)

var cache map[string]model.Imovel

func main() {
	cache = make(map[string]model.Imovel, 0)

	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	if err := casa.SetValor(1100000000); err != nil {
		fmt.Println("[main] - Houve um erro na atribuição de valor da casa: ", err.Error())
		if err == model.ErrValorDeImovelMuitoAlto {
			fmt.Println("Corretor, por favor verifique a sua avaliação")
		}
		return
	}
	cache["Casa Amarela"] = casa

	fmt.Println("Há uma casa amarela no cache?")
	imovel, achou := cache["Casa Amarela"]
	if achou {
		fmt.Printf("Sim, e o que achei foi: %+v\r\n", imovel)
	}

	apto := model.Imovel{}
	apto.Nome = "Apto Azul"
	apto.X = 19
	apto.Y = 26
	apto.SetValor(70000)
	cache[apto.Nome] = apto

	fmt.Println("Quantos itens há no cache? ", len(cache))

	for chave, imovel := range cache {
		fmt.Printf("Chave[%s] = %+v\r\n", chave, imovel)
	}

	_, achou = cache[casa.Nome]
	if achou {
		delete(cache, casa.Nome)
	}
	fmt.Println("Quantos itens há no cache? ", len(cache))

	objJSON, err := json.Marshal(casa)
	if err != nil {
		fmt.Println("[main] - Houve um erro ao gerar objeto JSON: ", err.Error())
		return
	}
	fmt.Println("A casa em JSON: ", string(objJSON))
}

package main

import (
	"fmt"

	"../pacotes/operadora"
	"../pacotes/prefixo"
)

//NomeDoUsuario : nome do usuário
var NomeDoUsuario = "Jeff"

func main() {
	fmt.Printf("Nome do usuário: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Prefixo da Capital: %s\r\n", operadora.NomeOperadora)
	fmt.Printf("Prefixo da Operadora: %s\r\n", operadora.PrefixoDaCapitalOperadora)
	fmt.Printf("Valor de teste: %s\r\n", prefixo.TesteComPrefixo)
}

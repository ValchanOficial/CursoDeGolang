package main

import "fmt"

var (
	//Endereco : é uma variável pública
	Endereco string
	//Telefone : é uma variável pública
	Telefone            string
	quantidade, estoque int
	comprou             bool
	valor               float64
	palavras            rune
)

func main() {
	teste := "Valor de teste"
	fmt.Printf("Endereço: %s\r\n", Endereco)
	fmt.Printf("Quantidade: %d\r\n", quantidade)
	fmt.Printf("Comprou: %v\r\n", comprou)
	fmt.Printf("Palavras: %v\r\n", palavras)
	fmt.Printf("Teste: %v\r\n", teste)
}

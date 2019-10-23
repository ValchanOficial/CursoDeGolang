package main

import (
	"fmt"

	"./matematica"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 2, 7)
	fmt.Printf("O resultado da multiplicação foi: %d\r\n", resultado)
	resultado = matematica.Soma(3, 3)
	fmt.Printf("O resultado da soma foi: %d\r\n", resultado)
	resultado = matematica.Calculo(matematica.Divisor, 12, 3)
	fmt.Printf("O resultado da divisão foi: %d\r\n", resultado)
	resultado, resto := matematica.DivisorComResto(13, 5)
	fmt.Printf("O resultado da divisão foi: %d, e resto da divisão foi: %d\r\n", resultado, resto)
}

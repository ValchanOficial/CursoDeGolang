package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	numero := 3
	fmt.Print("O número ", numero, " se escreve assim: ")
	switch numero {
	case 1:
		fmt.Println("um.")
	case 2:
		fmt.Println("dois.")
	case 3:
		fmt.Println("três.")
	}

	fmt.Println("Você é da família do Unix?")
	switch runtime.GOOS {
	case "darwin":
		fallthrough //segue para o próximo
	case "freebsd":
		fallthrough
	case "linux":
		fmt.Println("Sim!!")
	default:
		fmt.Println("Deixa essa pergunta para lá...")
		fmt.Println(runtime.GOOS)
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("OK, você pode dormir até mais tarde.")
	default:
		fmt.Println("Vamos lá que é dia de trabalho.")
		fmt.Println(time.Now().Weekday())
	}

	numero = 7
	fmt.Println("Este número cabe num dígito?")
	switch {
	case numero < 10:
		fmt.Println("Sim!")
	case numero >= 10 && numero < 100:
		fmt.Println("Serve dois dígitos?")
	case numero >= 100:
		fmt.Println("Não dá, o número é muito grande!")
	}
}

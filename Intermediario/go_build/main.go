package main

import "fmt"

func main() {
	fmt.Println("Para fazer o build, execute os comandos:")
	fmt.Println("1 - go build main.go") //faz o build do projeto
	//go build main.go
	//go build -o meuapp
	//GOOS=windows GOARCH=386 go build -o meuappwindows.exe
	//GOOS=linux GOARCH=arm go build -o meuappraspberry -v
	//-v //mostra pacotes que estão sendo incluídos
	fmt.Println("2 - ./main.exe") //executa o projeto // ./meuapp
}

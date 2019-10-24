package main

import (
	"fmt"
	"net/http"

	"./manipulador"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá Mundo!")
	})
	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)

	fmt.Println("Escutando na porta: 8181")
	http.ListenAndServe(":8181", nil)
}

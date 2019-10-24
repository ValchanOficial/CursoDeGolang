package main

import (
	"fmt"
	"net/http"

	"./manipulador"
	"./repo"
)

func init() {
	fmt.Println("Vamos começar a subir o servidor..")
}

func main() {

	//Conexão MySQL
	err := repo.AbreConexaoComBancoDeDadosSQL()
	if err != nil {
		fmt.Println("Parando a carga do servidor. Erro ao abrir o banco de dados: ", err.Error())
		return
	}

	//Conexão MongoDB
	err = repo.AbreSessaoComMongo()
	if err != nil {
		fmt.Println("Parando a carga do servidor. Erro ao abrir sessão com o banco de dados: ", err.Error())
		return
	}

	//go get github.com/jmoiron/sqlx

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá Mundo!")
	})
	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	http.HandleFunc("/local/", manipulador.Local)
	http.HandleFunc("/insert/", manipulador.Insert)
	http.HandleFunc("/select/", manipulador.SelectMongo)
	http.HandleFunc("/insertMongo/", manipulador.InsereMongo)

	fmt.Println("Escutando na porta: 8181")
	http.ListenAndServe(":8181", nil)
}

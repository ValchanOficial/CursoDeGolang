package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	resposta, err := cliente.Get("https://www.google.com.br")
	if err != nil {
		fmt.Println("[main] - Erro ao abrir página do Google Brasil - Erro: ", err.Error())
		return
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] - Erro ao ler o conteúdo do body - Erro: ", err.Error())
			return
		}
		fmt.Println(string(corpo))
	}

	//Autenticação

	request, err := http.NewRequest("GET", "https://www.google.com.br", nil)
	if err != nil {
		fmt.Println("[main] - Erro ao criar um request para o Google Brasil - Erro: ", err.Error())
		return
	}
	request.SetBasicAuth("teste", "teste") //autenticação

	resposta, err = cliente.Do(request)
	if err != nil {
		fmt.Println("[main] - Erro ao abrir página do Google Brasil - Erro: ", err.Error())
		return
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] - Erro ao ler o conteúdo do body - Erro: ", err.Error())
			return
		}
		fmt.Println("------------------------------------------")
		fmt.Println(string(corpo))
	}

}

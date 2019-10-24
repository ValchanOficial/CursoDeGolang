package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"./model"
)

func main() {

	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	//https://jsonplaceholder.typicode.com/posts/1
	request, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		fmt.Println("[main] - Erro ao fazer um request para o jsonplaceholder - Erro: ", err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	resposta, err := cliente.Do(request)
	if err != nil {
		fmt.Println("[main] - Erro ao chamar o serviço do jsonplaceholder - Erro: ", err.Error())
		return
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] - Erro ao ler o conteúdo do body - Erro: ", err.Error())
			return
		}

		post := model.BlogPost{}
		err = json.Unmarshal(corpo, &post)
		if err != nil {
			fmt.Println("[main] - Erro ao converter o retorno JSON do body - Erro: ", err.Error())
			return
		}
		fmt.Printf("Post é: %+v\r\n", post)
	}
}

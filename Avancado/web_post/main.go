package main

import (
	"bytes"
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

	usuario := model.Usuario{}
	usuario.ID = 1
	usuario.Nome = "Valchan"

	conteudoEnviar, err := json.Marshal(usuario)
	if err != nil {
		fmt.Println("[main] - Erro ao gerar o JSON do objeto usuario - Erro: ", err.Error())
		return
	}

	//https://requestbin.com/
	request, err := http.NewRequest("POST", "https://enjcl5vjx2ya.x.pipedream.net", bytes.NewBuffer(conteudoEnviar))
	if err != nil {
		fmt.Println("[main] - Erro ao criar um request para o request bin - Erro: ", err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	resposta, err := cliente.Do(request)
	if err != nil {
		fmt.Println("[main] - Erro ao chamar o serviço do request bin - Erro: ", err.Error())
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
}

package main

import "fmt"

func main() {
	meses := 6
	situacao := true
	cidade := "Teste"

	//operadores
	//< > != == <= >= && ||
	if meses <= 6 {
		fmt.Println("Esse credor deve a pouco tempo")
	}

	if !situacao {
		fmt.Println("Ele está adimplente")
	} else {
		fmt.Println("Ele está devendo")
	}

	if cidade == "Teste" {
		fmt.Println("O cliente vive no estado Teste")
	}

	if descricao, status := haQuantoTempoEstaDevendo(meses); status {
		fmt.Println("Qual a situação do cliente? ", descricao)
	} else {
		fmt.Println("Qual a situação do cliente? ", descricao)
	}

	fmt.Println("Obrigado por nos consultar.")

}

func haQuantoTempoEstaDevendo(meses int) (descricao string, status bool) {
	if meses > 0 {
		status = true
		descricao = "O cliente esta devendo."
		return
	}
	descricao = "O cliente esta com o carnê em dia."
	return
}

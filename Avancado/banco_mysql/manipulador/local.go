package manipulador

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"../model"
	"../repo"
)

//Local : é manipulador da requisição de rota /local
func Local(w http.ResponseWriter, r *http.Request) {
	local := model.Local{}
	codigoTelefone, err := strconv.Atoi(r.URL.Path[7:])
	if err != nil {
		http.Error(w, "Não foi enviado um número válido. Verifique.", http.StatusBadRequest)
		fmt.Println("[Local] - Erro ao converter o número enviado: ", err.Error())

	}

	//SELECT
	query := "SELECT country, city, telcode FROM place WHERE telcode = ?"
	linha, err := repo.DB.Queryx(query, codigoTelefone)
	if err != nil {
		http.Error(w, "Não foi possível pesquisar esse número.", http.StatusInternalServerError)
		fmt.Println("[Local] - Não foi possível executar a query: ", query, " - Erro: ", err.Error())
		return
	}

	for linha.Next() {
		err = linha.StructScan(&local)
		if err != nil {
			http.Error(w, "Não foi possível pesquisar esse número.", http.StatusInternalServerError)
			fmt.Println("[Local] - Não foi possível fazer o biding dos dados do banco na struct - Erro: ", err.Error())
			return
		}
	}

	if err := ModeloLocal.ExecuteTemplate(w, "local.html", local); err != nil {
		http.Error(w, "Houve um erro na renderização da página.", http.StatusInternalServerError)
		fmt.Println("[Local] - Erro na execução do modelo: ", err.Error())
	}
}

//Insert : inserindo dado
func Insert(w http.ResponseWriter, r *http.Request) {
	localInsert := model.Local{"Brasil", sql.NullString{"Porto Alegre", true}, 5551}
	query := "INSERT INTO db.place (country, city, telcode) VALUES (?,?,?)"
	resultado, err := repo.DB.Exec(query, localInsert.Pais, localInsert.Cidade, localInsert.CodigoTelefonico)
	if err != nil {
		fmt.Println("[Local] Erro na inclusao do local: ", query, " - Erro: ", err.Error())
	}
	linhasAfetadas, err := resultado.RowsAffected()
	if err != nil {
		fmt.Println("[Local] Erro ao pegar o numero de linhas afetadas pelo comando: ", query, " - Erro: ", err.Error())
	}
	fmt.Fprintln(w, "Sucesso! ", linhasAfetadas, " linha(s) afetada(s)")
}

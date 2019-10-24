package repo

import (
	"github.com/jmoiron/sqlx"

	/*
		github.com/go-sql-driver/mysql não é usado diretamente pela aplicação
	*/
	_ "github.com/go-sql-driver/mysql"
)

//DB : armazena a conexão com o banco de dados
var DB *sqlx.DB

//AbreConexaoComBancoDeDadosSQL  : abre coneão com o banco de dados SQL
func AbreConexaoComBancoDeDadosSQL() (err error) {
	err = nil
	DB, err = sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/db")
	if err != nil {
		return
	}

	err = DB.Ping() //verifica a conexão
	if err != nil {
		return
	}

	return
}

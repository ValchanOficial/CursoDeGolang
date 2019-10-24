package repo

import (
	"../model"
	"gopkg.in/mgo.v2/bson"
)

//ObtemLocal : retorna um local do MongoDB
func ObtemLocal(codigoTelefone int) (local model.Local, err error) {
	sessao := SessaoMongo.Copy()
	defer sessao.Close()
	colecao := sessao.DB("place").C("local")
	err = colecao.Find(bson.M{"telcode": codigoTelefone}).One(&local)
	return
}

//SalvaLog : registra a consulta ao local
func SalvaLog(reg model.RegistroLog) (err error) {
	sessao := SessaoMongo.Copy()
	defer sessao.Close()
	colecao := sessao.DB("place").C("logVisitas")
	err = colecao.Insert(reg)
	return
}

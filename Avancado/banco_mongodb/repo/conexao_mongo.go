package repo

import "gopkg.in/mgo.v2"

//SessaoMongo : armazena a sessão de conexão com o Mongo
var SessaoMongo *mgo.Session

//AbreSessaoComMongo : abre coneão com o banco de dados MongoDB
func AbreSessaoComMongo() (err error) {

	err = nil
	SessaoMongo, err = mgo.Dial("mongodb://localhost:27017/db")
	return
}

package prefixo

import "strconv"

//Capital : representa o número do prefixo de telefone da capital de um estado/ provincia
var Capital = 11

var teste = "teste"

//TesteComPrefixo : testeComPrefixo
var TesteComPrefixo = teste + " " + strconv.Itoa(Capital)

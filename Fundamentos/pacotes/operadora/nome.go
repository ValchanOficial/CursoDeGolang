package operadora

import (
	"strconv"

	"../prefixo"
)

//NomeOperadora : representa o nome da operadora
var NomeOperadora = "XPTO Telecom"

//PrefixoDaCapitalOperadora : prefixoDaCapitalOperadora
var PrefixoDaCapitalOperadora = strconv.Itoa(prefixo.Capital) + " " + NomeOperadora

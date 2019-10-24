package manipulador

import "text/template"

//ModeloOla : armazena o modelo de página ola
var ModeloOla = template.Must(template.ParseFiles("html/ola.html"))

//ModeloLocal : armazena o modelo de página local
var ModeloLocal = template.Must(template.ParseFiles("html/local.html"))

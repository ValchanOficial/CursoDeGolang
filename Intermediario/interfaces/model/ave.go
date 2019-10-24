package model

//Galinha : representa uma ave do tipo Galinha
type Galinha interface {
	Cacareja() string
}

//Pato : representa uma ave do tipo Pato
type Pato interface {
	Quack() string
}

//Ave : representa um animal
type Ave struct {
	Nome string
}

//Cacareja : representa um som emitido por uma galinha
func (a Ave) Cacareja() string {
	return "Cócóricó.."
}

//Quack : representa um som emitido por um pato
func (a Ave) Quack() string {
	return "Quack, quack.."
}

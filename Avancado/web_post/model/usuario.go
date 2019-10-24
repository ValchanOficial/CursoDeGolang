package model

//Usuario : representa um usuário do sistema
type Usuario struct {
	ID int `json:"id"`
	Nome string `json:"nome"`
}
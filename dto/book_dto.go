package dto

type BookDTO struct {
	Id       int64  `json:"id"`
	Nome     string `json:"nome"`
	Detalhes string `json:"detalhes"`
}

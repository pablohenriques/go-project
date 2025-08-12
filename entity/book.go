package entity

import (
	"time"

	"github.com/pablohenriques/go-project/dto"
)

type Book struct {
	Id       int64
	Nome     string
	Detalhes string
	Data     time.Time
}

func (b Book) ToResponseDTO() dto.BookResponseDTO {
	return dto.BookResponseDTO{
		Id:   b.Id,
		Nome: b.Nome,
		Data: b.Data,
	}
}

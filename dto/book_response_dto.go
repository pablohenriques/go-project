package dto

import "time"

type BookResponseDTO struct {
	Id   int64
	Nome string
	Data time.Time
}

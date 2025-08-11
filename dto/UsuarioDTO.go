package dto

type UsuarioDTO struct {
	Id    int64  `json:"id" validate:"required,min=3"`
	Nome  string `json:"nome" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

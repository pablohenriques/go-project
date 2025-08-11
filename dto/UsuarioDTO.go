package dto

type UsuarioDTO struct {
	Id    int64  `json:"id" validate:"required"`
	Nome  string `json:"nome" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

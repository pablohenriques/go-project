package handler

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func CustomErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error - Minha Mensagem Interna"

	var e *fiber.Error

	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
	}

	errorResponse := ErrorResponse{
		Code:    code,
		Message: message,
	}

	log.Printf("Erro: %v - Code: %d", err, code)
	return c.Status(code).JSON(errorResponse)
}

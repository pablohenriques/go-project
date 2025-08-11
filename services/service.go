package services

import (
	"errors"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pablohenriques/go-project/dto"
)

func Home(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Olá Mundo!",
	})
}

func Welcome(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Welcome",
	})
}

func Book(c *fiber.Ctx) error {
	param := c.Params("id")
	message := "book"

	if param != "" {
		log.Println("Parâmetro:")
		message = param
	}

	return c.JSON(fiber.Map{
		"message": message,
	})
}

func CreateNewUser(c *fiber.Ctx) error {
	var userDTO dto.UsuarioDTO

	if err := c.BodyParser(&userDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "JSON inválido"})
	}

	log.Printf("new user: %v", userDTO)

	responseDTO := dto.ResponseDTO{
		Id:    userDTO.Id,
		Email: userDTO.Email,
	}

	return c.Status(fiber.StatusCreated).JSON(responseDTO)
}

func Erro500(c *fiber.Ctx) error {

	if time.Now().Unix()%2 == 0 {
		return errors.New("Não foi possível conectar")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "active"})
}

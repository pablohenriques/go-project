package clients

import (
	"log"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/pablohenriques/go-project/dto"
)

func GetExternalTasks(c *fiber.Ctx) error {

	var todoDTO dto.TodoDTO
	var errorDTO map[string]interface{}

	client := resty.New().SetBaseURL("https://jsonplaceholder.typicode.com").SetTimeout(10 * time.Second)

	resp, err := client.R().SetResult(&todoDTO).SetError(&errorDTO).Get("/todos/1")

	if err != nil {
		log.Printf("Erro ao chamar API externa: %v", err)
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": "Erro comunicação API Externa"})
	}

	if resp.IsError() {
		log.Printf("API retornou erro. Status=%s, Corpo=%v", resp.Status(), errorDTO)
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   "Erro comunicação API Externa",
			"details": errorDTO,
		})
	}

	log.Printf("Dados convertidos com sucesso: %+v", todoDTO)
	return c.JSON(todoDTO)
}

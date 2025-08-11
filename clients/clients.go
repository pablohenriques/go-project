package clients

import (
	"errors"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/pablohenriques/go-project/dto"
	"github.com/sony/gobreaker"
)

var (
	client *resty.Client
	cb     *gobreaker.CircuitBreaker
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

func CircuitBreakerHandler() {
	var settings gobreaker.Settings
	settings.Name = "HTTP-GET-Httpbin"
	settings.MaxRequests = 1
	settings.Interval = 0
	settings.Timeout = 5 * time.Second

	settings.ReadyToTrip = func(counts gobreaker.Counts) bool {
		return counts.ConsecutiveFailures > 3
	}

	settings.OnStateChange = func(name string, from gobreaker.State, to gobreaker.State) {
		log.Printf("CircuitBreaker '%s' mudou de estado: %s -> %s\n", name, from, to)
	}

	cb = gobreaker.NewCircuitBreaker(settings)
}

func GetCallExternalAPICircuit(c *fiber.Ctx) error {
	var todoDTO dto.TodoDTO
	var errorDTO map[string]interface{}

	client := resty.New().SetBaseURL("https://httpbin.org/status").SetTimeout(10 * time.Second)

	resp, err := client.R().SetResult(&todoDTO).SetError(&errorDTO).Get("/200")

	if err != nil {
		log.Printf("Erro ao chamar API externa: %v", err)

		if errors.Is(err, gobreaker.ErrOpenState) {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"error":  "Circuit Breaker Aberto",
				"messge": "Serviço fora do ar",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "A operação falhou.",
			"details": err.Error(),
		})
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

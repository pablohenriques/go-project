package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/pablohenriques/go-project/clients"
	"github.com/pablohenriques/go-project/services"
)

func SetupRoutes(app *fiber.App) {
	app.Use(logger.New())
	v1 := app.Group("/v1")
	v1.Post("/create-book", services.CreateBook)
	v1.Get("/get-one-book", services.GetOneBook)
	v1.Get("/get-all-book", services.GetAllBook)
	v1.Put("/update-book/:id", services.UpdateBook)
	v1.Delete("/delete-book", services.DeleteBook)

	app.Get("/", services.Home)
	app.Get("/erro", services.Erro500)
	app.Get("/request-client", clients.GetExternalTasks)
	app.Get("/request-circuit", services.CallCircuitBreaker)
}

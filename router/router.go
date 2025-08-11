package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/pablohenriques/go-project/services"
)

func SetupRoutes(app *fiber.App) {
	app.Use(logger.New())
	app.Get("/", services.Home)
	app.Get("/new", services.Welcome)
	app.Post("/insert", services.CreateNewUser)
	app.Get("/erro", services.Erro500)

	// api := app.Group("/api")
	v1 := app.Group("/v1")

	v1.Get("/book", services.Book)

	v1.Get("/book/:id", services.Book)

}

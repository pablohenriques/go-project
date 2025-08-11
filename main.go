package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	handler "github.com/pablohenriques/go-project/handler"
	routes "github.com/pablohenriques/go-project/router"
)

func main() {
	app := fiber.New(fiber.Config{ErrorHandler: handler.CustomErrorHandler})
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}

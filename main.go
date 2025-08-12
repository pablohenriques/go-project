package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	routes "github.com/pablohenriques/go-project/controller"
	handler "github.com/pablohenriques/go-project/handler"
)

func main() {
	app := fiber.New(fiber.Config{ErrorHandler: handler.CustomErrorHandler})
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}

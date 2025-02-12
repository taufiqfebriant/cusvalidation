package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/taufiqfebriant/cusvalidation/routes"
	"github.com/taufiqfebriant/cusvalidation/validator"
)

func main() {
	app := fiber.New()

	validator.InitValidator()
	routes.SetupRoutes(app)

	app.Listen(":3000")
}

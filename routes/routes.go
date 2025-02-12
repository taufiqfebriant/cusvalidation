package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/taufiqfebriant/cusvalidation/api/v1/users"
)

func SetupRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")

	users.UserRoutes(v1)
}

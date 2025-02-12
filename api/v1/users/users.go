package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/taufiqfebriant/cusvalidation/api/v1/users/post"
)

func UserRoutes(app fiber.Router) {
	app.Post("/users", post.CreateUserHandler)
}

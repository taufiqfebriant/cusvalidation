package post

import (
	"github.com/gofiber/fiber/v2"
	"github.com/taufiqfebriant/cusvalidation/validator"
)

func CreateUserHandler(c *fiber.Ctx) error {
	u := new(ReqBody)

	if err := c.BodyParser(u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if err := validator.ValidateStruct(u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User successfully created",
	})
}

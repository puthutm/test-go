package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type SchoolTypeRequest struct {
	Name string `json:"name" validate:"required,max=255"`
}

func ValidateSchoolType(c *fiber.Ctx) error {
	var req SchoolTypeRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}

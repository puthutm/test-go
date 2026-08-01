package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type EmployeeStatusRequest struct {
	Code string `json:"code" validate:"required,max=10"`
	Name string `json:"name" validate:"required"`
}

func ValidateEmployeeStatus(c *fiber.Ctx) error {
	var req EmployeeStatusRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}

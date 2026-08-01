package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type FacultyRequest struct {
	Name        string  `json:"name" validate:"required,max=225"`
	Shortname   string  `json:"shortname" validate:"required,max=225"`
	ChairName   string  `json:"chair_name" validate:"required,max=225"`
	Description *string `json:"description"`
}

func ValidateFaculty(c *fiber.Ctx) error {
	var req FacultyRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}

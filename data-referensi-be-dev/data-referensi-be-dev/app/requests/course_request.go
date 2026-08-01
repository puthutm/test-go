package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type CourseRequest struct {
	Code        string  `json:"code" validate:"required,max=50"`
	Name        string  `json:"name" validate:"required,max=225"`
	Description *string `json:"description"`
}

func ValidateCourse(c *fiber.Ctx) error {
	var req CourseRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}

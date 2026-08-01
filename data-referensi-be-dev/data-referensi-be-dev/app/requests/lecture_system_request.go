package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type LectureSystemRequest struct {
	Code          string `json:"code" validate:"required,max=255"`
	LectureSystem string `json:"lecture_system" validate:"required,max=255"`
	Description   string `json:"description" `
}

func ValidateLectureSystem(c *fiber.Ctx) error {
	var req LectureSystemRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}

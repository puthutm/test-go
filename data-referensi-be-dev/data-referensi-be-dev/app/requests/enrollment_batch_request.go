package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type EnrollmentBatchRequest struct {
	Code        string `json:"code" validate:"required,max=255"`
	Batch       string `json:"batch" validate:"required,max=255"`
	Description string `json:"description" `
}

func ValidateEnrollmentBatch(c *fiber.Ctx) error {
	var req EnrollmentBatchRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}

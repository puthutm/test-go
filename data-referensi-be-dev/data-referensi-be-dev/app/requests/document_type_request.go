package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type DocumentTypeRequest struct {
	Name string `json:"name" validate:"required,max=255"`
	Mime string `json:"mimes" validate:"required,max=255"`
	Size int    `json:"size" validate:"required,numeric"`
}

func ValidateDocumentType(c *fiber.Ctx) error {
	var req DocumentTypeRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}

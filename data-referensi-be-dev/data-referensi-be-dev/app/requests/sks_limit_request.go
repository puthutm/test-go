package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type SksLimitRequest struct {
	IPMinimal string  `json:"ip_minimal" validate:"required,max=225"`
	IPMaximal *string `json:"ip_maximal" validate:"required,max=225"`
	SKSLimit  *string `json:"sks_limit"`
}

func ValidateSksLimit(c *fiber.Ctx) error {
	var req SksLimitRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}

package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type EthnicRequest struct {
	Name           string `json:"name" validate:"required,max=255"`
	RegionOfOrigin string `json:"region_of_origin" validate:"required,max=255"`
}

func ValidateEthnic(c *fiber.Ctx) error {
	var req EthnicRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}

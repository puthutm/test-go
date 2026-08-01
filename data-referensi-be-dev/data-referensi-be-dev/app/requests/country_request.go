package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type CountryRequest struct {
	Name         string `json:"name" validate:"required,max=255"`
	PhoneCode    string `json:"phone_code" validate:"required,max=10"`
	IconFlagPath string `json:"icon_flag_path" validate:"omitempty,max=255"`
}

func ValidateCountry(c *fiber.Ctx) error {
	var req CountryRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}

package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type CityRequest struct {
	ProvinceId string `json:"province_id" validate:"required,max=36"`
	Name       string `json:"name" validate:"required,max=255"`
	Code       string `json:"code" validate:"required,max=10"`
}

func ValidateCity(c *fiber.Ctx) error {
	var req CityRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}

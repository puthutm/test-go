package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type DistrictRequest struct {
	CityId string `json:"city_id" validate:"required,max=36"`
	Name   string `json:"name" validate:"required,max=255"`
	Code   string `json:"code" validate:"required,max=10"`
}

func ValidateDistrict(c *fiber.Ctx) error {
	var req DistrictRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}

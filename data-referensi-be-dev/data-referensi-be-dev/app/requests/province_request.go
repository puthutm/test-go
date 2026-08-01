package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type ProvinceRequest struct {
	CountryId  string `json:"country_id" validate:"required,max=36"`
	Name       string `json:"name" validate:"required,max=255"`
	Code       string `json:"code" validate:"required,max=5"`
	RegionCode string `json:"region_code" validate:"omitempty,max=255"`
}

func ValidateProvince(c *fiber.Ctx) error {
	var req ProvinceRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}

package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type SchoolRequest struct {
	Npsn          string `json:"npsn" validate:"required,max=255"`
	Name          string `json:"name" validate:"required,max=255"`
	EducationForm string `json:"education_form" validate:"required,max=255"`
	Status        string `json:"status" validate:"required,max=50"`
	PrivinceId    string `json:"province_id" validate:"required,max=36"`
	CityId        string `json:"city_id" validate:"required,max=36"`
	DistrictId    string `json:"district_id" validate:"required,max=36"`
}

func ValidateSchool(c *fiber.Ctx) error {
	var req SchoolRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}

package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type CollegeRequest struct {
	Name                 string  `json:"name" validate:"required,max=255"`
	ProvinceId           string  `json:"province_id" validate:"required,max=36"`
	CityId               string  `json:"city_id" validate:"required,max=36"`
	Type                 string  `json:"type" validate:"required,max=255"`
	Accreditation        string  `json:"accreditation" validate:"omitempty,max=50"`
	ShortName            string  `json:"short_name" validate:"omitempty,max=255"`
	NumberOfStudyProgram int     `json:"number_of_study_program"`
	LowerLimitTuitionFee float64 `json:"lower_limit_tuition_fee"`
	UpperLimitTuitionFee float64 `json:"upper_limit_tuition_fee"`
}

func ValidateCollege(c *fiber.Ctx) error {
	var req CollegeRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}

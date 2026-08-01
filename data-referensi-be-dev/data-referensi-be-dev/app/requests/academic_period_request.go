package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type AcademicPeriodRequest struct {
	Code                   string `json:"code" validate:"required,max=50"`
	AcademicYearId         string `json:"academic_year_id" validate:"required,max=36"`
	SemesterId             string `json:"semester_id" validate:"required,max=36"`
	Fullname               string `json:"fullname" validate:"required,max=255"`
	Shortname              string `json:"shortname" validate:"required,max=150"`
	StartDateOfCollege     string `json:"start_date_of_college" validate:"required"`
	EndDateOfCollege       string `json:"end_date_of_college" validate:"required"`
	StartDateOfUts         string `json:"start_date_of_uts" validate:"omitempty"`
	EndDateOfUts           string `json:"end_date_of_uts" validate:"omitempty"`
	StartDateOfUas         string `json:"start_date_of_uas" validate:"omitempty"`
	EndDateOfUas           string `json:"end_date_of_uas" validate:"omitempty"`
	NumberOfLectureMeeting string `json:"number_of_lecture_meeting" validate:"required,numeric"`
	IsActive               bool   `json:"is_active"`
}

func ValidateAcademicPeriod(c *fiber.Ctx) error {
	var req AcademicPeriodRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}

package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type AlmamaterSizeRequest struct {
	Code       string `json:"code" validate:"required,max=50"`
	Size       string `json:"size" validate:"required,max=255"`
	ChestSize  string `json:"chest_size" validate:"required,max=255"`
	ArmLength  string `json:"arm_length" validate:"required,max=255"`
	BodyLength string `json:"body_length" validate:"required,max=255"`
}

func ValidateAlmamaterSize(c *fiber.Ctx) error {
	var req AlmamaterSizeRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}

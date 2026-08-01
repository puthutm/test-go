package handlers

import (
	"data-referensi/helpers"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ModelValidate(c *fiber.Ctx, req interface{}) (map[string]string, error) {
	validate := validator.New()

	if err := c.BodyParser(req); err != nil {
		return nil, fmt.Errorf("invalid request body: %v", err)
	}

	if err := validate.Struct(req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make(map[string]string)

		for _, fieldError := range validationErrors {
			fieldName := helpers.ConvertCCToSC(fieldError.Field())
			errorMessages[fieldName] = fmt.Sprintf("Field '%s' failed validation: %s", fieldName, fieldError.Tag())
		}

		return errorMessages, fmt.Errorf("validation failed")
	}

	return nil, nil
}

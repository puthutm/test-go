package handlers

import (
	"data-referensi/helpers"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ErrorQuery(c *fiber.Ctx, err error) error {
	if strings.Contains(err.Error(), "duplicate key row") {
		return SendFailed(c, fiber.StatusBadRequest, nil, helpers.GenerateRM("exist"))
	} else if strings.Contains(err.Error(), "Update data Failed") {
		return SendFailed(c, fiber.StatusNotFound, nil, err.Error())
	} else if strings.Contains(err.Error(), "Delete data Failed") {
		return SendFailed(c, fiber.StatusNotFound, nil, err.Error())
	} else if strings.Contains(err.Error(), "Restore data Failed") {
		return SendFailed(c, fiber.StatusNotFound, nil, err.Error())
	}
	return SendFailed(c, fiber.StatusInternalServerError, nil, err.Error())
}

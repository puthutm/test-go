package controllers

import (
	"data-referensi/app/models"
	"data-referensi/app/requests"
	"data-referensi/handlers"
	"data-referensi/helpers"

	"github.com/gofiber/fiber/v2"
)

func GetSemesterNumbers(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	semesterNumbers, total, err := models.GetSemesterNumbers(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(semesterNumbers, page, pageSize, len(semesterNumbers), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func SearchSemesterNumbers(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	semesterNumbers, err := models.SearchSemesterNumbers(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	return handlers.SendSuccess(c, fiber.StatusOK, semesterNumbers, helpers.GenerateRM("get", true))
}

func ExportSemesterNumbers(c *fiber.Ctx) error {
	return handlers.ControllerExport(c, "SemesterNumbers", models.ExportSemesterNumbers)
}

func GetSemesterNumber(c *fiber.Ctx) error {
	id := c.Params("id")
	semesterNumber, err := models.GetSemesterNumber(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, semesterNumber, helpers.GenerateRM("get", true))
}

func CreateSemesterNumber(c *fiber.Ctx) error {
	var req requests.SemesterNumberRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	id, err := helpers.EnsureUUID(&models.MstSemester{})
	if err != nil {
		return err
	}

	err = models.CreateSemesterNumber(id, req.SemesterNumber)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	semesterNumber, err := models.GetSemesterNumber(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusCreated, semesterNumber, helpers.GenerateRM("insert", true))
}

func ImportSemesterNumbers(c *fiber.Ctx) error {
	return handlers.ControllerImport(c, models.ImportSemesterNumbers)
}

func UpdateSemesterNumber(c *fiber.Ctx) error {
	id := c.Params("id")

	var req requests.SemesterNumberRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	err := models.UpdateSemesterNumber(id, req.SemesterNumber)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	semesterNumber, err := models.GetSemesterNumber(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, semesterNumber, helpers.GenerateRM("update", true))
}

func DeleteSemesterNumber(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.DeleteSemesterNumber(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("delete", true))
}

func GetTrashSemesterNumbers(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	semesterNumbers, total, err := models.GetTrashSemesterNumbers(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(semesterNumbers, page, pageSize, len(semesterNumbers), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func RestoreSemesterNumber(c *fiber.Ctx) error {
	id := c.Params("id")
	err := models.RestoreSemesterNumber(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("restore", true))
}

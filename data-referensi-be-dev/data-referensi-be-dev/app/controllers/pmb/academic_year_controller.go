package controllers

import (
	"data-referensi/app/models"
	"data-referensi/app/requests"
	"data-referensi/handlers"
	"data-referensi/helpers"

	"github.com/gofiber/fiber/v2"
)

func GetAcademicYears(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetAcademicYears(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func SearchAcademicYears(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, err := models.SearchAcademicYears(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_results, helpers.GenerateRM("get", true))
}

func ExportAcademicYears(c *fiber.Ctx) error {
	return handlers.ControllerExport(c, "AcademicYears", models.ExportAcademicYears)
}

func GetAcademicYear(c *fiber.Ctx) error {
	id := c.Params("id")
	query_result, err := models.GetAcademicYear(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("get", true))
}

func CreateAcademicYear(c *fiber.Ctx) error {
	var req requests.AcademicYearRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	id, err := helpers.EnsureUUID(&models.MstAcademicYear{})
	if err != nil {
		return err
	}

	err = models.CreateAcademicYear(id, req.Year, req.Name)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetAcademicYear(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusCreated, query_result, helpers.GenerateRM("insert", true))
}

func ImportAcademicYears(c *fiber.Ctx) error {
	return handlers.ControllerImport(c, models.ImportAcademicYears)
}

func UpdateAcademicYear(c *fiber.Ctx) error {
	id := c.Params("id")

	var req requests.AcademicYearRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	err := models.UpdateAcademicYear(id, req.Year, req.Name)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetAcademicYear(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("update", true))
}

func DeleteAcademicYear(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.DeleteAcademicYear(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("delete", true))
}

func GetTrashAcademicYears(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetTrashAcademicYears(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func RestoreAcademicYear(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.RestoreAcademicYear(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("restore", true))
}

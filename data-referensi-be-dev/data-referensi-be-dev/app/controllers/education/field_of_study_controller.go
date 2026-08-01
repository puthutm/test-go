package controllers

import (
	"data-referensi/app/models"
	"data-referensi/app/requests"
	"data-referensi/handlers"
	"data-referensi/helpers"

	"github.com/gofiber/fiber/v2"
)

func GetFieldOfStudies(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetFieldOfStudies(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func SearchFieldOfStudies(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, err := models.SearchFieldOfStudies(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_results, helpers.GenerateRM("get", true))
}

func ExportFieldOfStudies(c *fiber.Ctx) error {
	return handlers.ControllerExport(c, "FieldOfStudies", models.ExportFieldOfStudies)
}

func GetFieldOfStudy(c *fiber.Ctx) error {
	id := c.Params("id")
	query_result, err := models.GetFieldOfStudy(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("get", true))
}

func CreateFieldOfStudy(c *fiber.Ctx) error {
	var req requests.FieldOfStudyRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	id, err := helpers.EnsureUUID(&models.MstFieldOfStudy{})
	if err != nil {
		return err
	}

	err = models.CreateFieldOfStudy(id, req.Name, req.Description)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetFieldOfStudy(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusCreated, query_result, helpers.GenerateRM("insert", true))
}

func ImportFieldOfStudies(c *fiber.Ctx) error {
	return handlers.ControllerImport(c, models.ImportFieldOfStudies)
}

func UpdateFieldOfStudy(c *fiber.Ctx) error {
	id := c.Params("id")

	var req requests.FieldOfStudyRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	err := models.UpdateFieldOfStudy(id, req.Name, req.Description)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetFieldOfStudy(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("update", true))
}

func DeleteFieldOfStudy(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.DeleteFieldOfStudy(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("delete", true))
}

func GetTrashFieldOfStudies(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetTrashFieldOfStudies(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func RestoreFieldOfStudy(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.RestoreFieldOfStudy(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("restore", true))
}

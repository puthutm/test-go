package controllers

import (
	"data-referensi/app/models"
	"data-referensi/app/requests"
	"data-referensi/handlers"
	"data-referensi/helpers"

	"github.com/gofiber/fiber/v2"
)

func GetFunctionalPositions(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetFunctionalPositions(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func SearchFunctionalPositions(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, err := models.SearchFunctionalPositions(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_results, helpers.GenerateRM("get", true))
}

func ExportFunctionalPositions(c *fiber.Ctx) error {
	return handlers.ControllerExport(c, "FunctionalPositions", models.ExportFunctionalPositions)
}

func GetFunctionalPosition(c *fiber.Ctx) error {
	id := c.Params("id")
	query_result, err := models.GetFunctionalPosition(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("get", true))
}

func CreateFunctionalPosition(c *fiber.Ctx) error {
	var req requests.FunctionalPositionRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	id, err := helpers.EnsureUUID(&models.MstFunctionalPosition{})
	if err != nil {
		return err
	}

	err = models.CreateFunctionalPosition(id, req.Name, req.Description)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetFunctionalPosition(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusCreated, query_result, helpers.GenerateRM("insert", true))
}

func ImportFunctionalPositions(c *fiber.Ctx) error {
	return handlers.ControllerImport(c, models.ImportFunctionalPositions)
}

func UpdateFunctionalPosition(c *fiber.Ctx) error {
	id := c.Params("id")

	var req requests.FunctionalPositionRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	err := models.UpdateFunctionalPosition(id, req.Name, req.Description)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetFunctionalPosition(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("update", true))
}

func DeleteFunctionalPosition(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.DeleteFunctionalPosition(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("delete", true))
}

func GetTrashFunctionalPositions(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetTrashFunctionalPositions(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func RestoreFunctionalPosition(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.RestoreFunctionalPosition(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("restore", true))
}

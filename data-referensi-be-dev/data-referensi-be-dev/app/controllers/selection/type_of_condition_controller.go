package controllers

import (
	"data-referensi/app/models"
	"data-referensi/app/requests"
	"data-referensi/handlers"
	"data-referensi/helpers"

	"github.com/gofiber/fiber/v2"
)

func GetTypeOfConditions(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetTypeOfConditions(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func SearchTypeOfConditions(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, err := models.SearchTypeOfConditions(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_results, helpers.GenerateRM("get", true))
}

func ExportTypeOfConditions(c *fiber.Ctx) error {
	return handlers.ControllerExport(c, "TypeOfConditions", models.ExportTypeOfConditions)
}

func GetTypeOfCondition(c *fiber.Ctx) error {
	id := c.Params("id")
	query_result, err := models.GetTypeOfCondition(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("get", true))
}

func CreateTypeOfCondition(c *fiber.Ctx) error {
	var req requests.TypeOfConditionRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	id, err := helpers.EnsureUUID(&models.MstTypeOfCondition{})
	if err != nil {
		return err
	}

	err = models.CreateTypeOfCondition(id, req.Code, req.TypeOfCondition, req.Note)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetTypeOfCondition(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusCreated, query_result, helpers.GenerateRM("insert", true))
}

func ImportTypeOfConditions(c *fiber.Ctx) error {
	return handlers.ControllerImport(c, models.ImportTypeOfConditions)
}

func UpdateTypeOfCondition(c *fiber.Ctx) error {
	id := c.Params("id")

	var req requests.TypeOfConditionRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	err := models.UpdateTypeOfCondition(id, req.Code, req.TypeOfCondition, req.Note)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetTypeOfCondition(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("update", true))
}

func DeleteTypeOfCondition(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.DeleteTypeOfCondition(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("delete", true))
}

func GetTrashTypeOfConditions(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetTrashTypeOfConditions(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func RestoreTypeOfCondition(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.RestoreTypeOfCondition(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("restore", true))
}

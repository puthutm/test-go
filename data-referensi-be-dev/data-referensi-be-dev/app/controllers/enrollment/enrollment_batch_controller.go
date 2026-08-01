package controllers

import (
	"data-referensi/app/models"
	"data-referensi/app/requests"
	"data-referensi/handlers"
	"data-referensi/helpers"

	"github.com/gofiber/fiber/v2"
)

func GetEnrollmentBatchs(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetEnrollmentBatchs(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func SearchEnrollmentBatchs(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, err := models.SearchEnrollmentBatchs(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_results, helpers.GenerateRM("get", true))
}

func ExportEnrollmentBatchs(c *fiber.Ctx) error {
	return handlers.ControllerExport(c, "EnrollmentBatchs", models.ExportEnrollmentBatchs)
}

func GetEnrollmentBatch(c *fiber.Ctx) error {
	id := c.Params("id")
	query_result, err := models.GetEnrollmentBatch(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("get", true))
}

func CreateEnrollmentBatch(c *fiber.Ctx) error {
	var req requests.EnrollmentBatchRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	id, err := helpers.EnsureUUID(&models.MstEnrollmentBatch{})
	if err != nil {
		return err
	}

	err = models.CreateEnrollmentBatch(id, req.Code, req.Batch, req.Description)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetEnrollmentBatch(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusCreated, query_result, helpers.GenerateRM("insert", true))
}

func ImportEnrollmentBatchs(c *fiber.Ctx) error {
	return handlers.ControllerImport(c, models.ImportEnrollmentBatchs)
}

func UpdateEnrollmentBatch(c *fiber.Ctx) error {
	id := c.Params("id")

	var req requests.EnrollmentBatchRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	err := models.UpdateEnrollmentBatch(id, req.Code, req.Batch, req.Description)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetEnrollmentBatch(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("update", true))
}

func DeleteEnrollmentBatch(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.DeleteEnrollmentBatch(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("delete", true))
}

func GetTrashEnrollmentBatchs(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetTrashEnrollmentBatchs(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func RestoreEnrollmentBatch(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.RestoreEnrollmentBatch(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("restore", true))
}

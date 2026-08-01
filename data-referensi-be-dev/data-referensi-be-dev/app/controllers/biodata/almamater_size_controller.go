package controllers

import (
	"data-referensi/app/models"
	"data-referensi/app/requests"
	"data-referensi/handlers"
	"data-referensi/helpers"

	"github.com/gofiber/fiber/v2"
)

func GetAlmamaterSizes(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetAlmamaterSizes(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func SearchAlmamaterSizes(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, err := models.SearchAlmamaterSizes(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_results, helpers.GenerateRM("get", true))
}

func ExportAlmamaterSizes(c *fiber.Ctx) error {
	return handlers.ControllerExport(c, "AlmamaterSizes", models.ExportAlmamaterSizes)
}

func GetAlmamaterSize(c *fiber.Ctx) error {
	id := c.Params("id")
	query_result, err := models.GetAlmamaterSize(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("get", true))
}

func CreateAlmamaterSize(c *fiber.Ctx) error {
	var req requests.AlmamaterSizeRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	id, err := helpers.EnsureUUID(&models.MstAlmamaterSize{})
	if err != nil {
		return err
	}

	err = models.CreateAlmamaterSize(id, req.Code, req.Size, req.ChestSize, req.ArmLength, req.BodyLength)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetAlmamaterSize(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusCreated, query_result, helpers.GenerateRM("insert", true))
}

func ImportAlmamaterSizes(c *fiber.Ctx) error {
	return handlers.ControllerImport(c, models.ImportAlmamaterSizes)
}

func UpdateAlmamaterSize(c *fiber.Ctx) error {
	id := c.Params("id")

	var req requests.AlmamaterSizeRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	err := models.UpdateAlmamaterSize(id, req.Code, req.Size, req.ChestSize, req.ArmLength, req.BodyLength)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetAlmamaterSize(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("update", true))
}

func DeleteAlmamaterSize(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.DeleteAlmamaterSize(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("delete", true))
}

func GetTrashAlmamaterSizes(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetTrashAlmamaterSizes(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func RestoreAlmamaterSize(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.RestoreAlmamaterSize(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("restore", true))
}

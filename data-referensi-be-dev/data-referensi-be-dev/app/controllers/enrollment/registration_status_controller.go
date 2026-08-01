package controllers

import (
	"data-referensi/app/models"
	"data-referensi/app/requests"
	"data-referensi/handlers"
	"data-referensi/helpers"

	"github.com/gofiber/fiber/v2"
)

func GetRegistrationStatuses(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetRegistrationStatuses(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func SearchRegistrationStatuses(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, err := models.SearchRegistrationStatuses(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_results, helpers.GenerateRM("get", true))
}

func ExportRegistrationStatuses(c *fiber.Ctx) error {
	return handlers.ControllerExport(c, "RegistrationStatuses", models.ExportRegistrationStatuses)
}

func GetRegistrationStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	query_result, err := models.GetRegistrationStatus(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("get", true))
}

func CreateRegistrationStatus(c *fiber.Ctx) error {
	var req requests.RegistrationStatusRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	id, err := helpers.EnsureUUID(&models.MstRegistrationStatus{})
	if err != nil {
		return err
	}

	err = models.CreateRegistrationStatus(id, req.Name, req.IsDefault, req.AcceptedMarker)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetRegistrationStatus(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusCreated, query_result, helpers.GenerateRM("insert", true))
}

func ImportRegistrationStatuses(c *fiber.Ctx) error {
	return handlers.ControllerImport(c, models.ImportRegistrationStatuses)
}

func UpdateRegistrationStatus(c *fiber.Ctx) error {
	id := c.Params("id")

	var req requests.RegistrationStatusRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	err := models.UpdateRegistrationStatus(id, req.Name, req.IsDefault, req.AcceptedMarker)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetRegistrationStatus(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("update", true))
}

func DeleteRegistrationStatus(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.DeleteRegistrationStatus(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("delete", true))
}

func GetTrashRegistrationStatuses(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetTrashRegistrationStatuses(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func RestoreRegistrationStatus(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.RestoreRegistrationStatus(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("restore", true))
}

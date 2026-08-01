package controllers

import (
	"data-referensi/app/models"
	"data-referensi/app/requests"
	"data-referensi/handlers"
	"data-referensi/helpers"

	"github.com/gofiber/fiber/v2"
)

func GetApprovals(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetApprovals(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func SearchApprovals(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, err := models.SearchApprovals(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_results, helpers.GenerateRM("get", true))
}

func ExportApprovals(c *fiber.Ctx) error {
	return handlers.ControllerExport(c, "Approvals", models.ExportApprovals)
}

func GetApproval(c *fiber.Ctx) error {
	id := c.Params("id")
	query_result, err := models.GetApproval(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("get", true))
}

func CreateApproval(c *fiber.Ctx) error {
	var req requests.ApprovalRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	id, err := helpers.EnsureUUID(&models.MstApproval{})
	if err != nil {
		return err
	}

	err = models.CreateApproval(id, req.Name, *req.Description)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetApproval(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusCreated, query_result, helpers.GenerateRM("insert", true))
}

func ImportApprovals(c *fiber.Ctx) error {
	return handlers.ControllerImport(c, models.ImportApprovals)
}

func UpdateApproval(c *fiber.Ctx) error {
	id := c.Params("id")

	var req requests.ApprovalRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	err := models.UpdateApproval(id, req.Name, *req.Description)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetApproval(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("update", true))
}

func DeleteApproval(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.DeleteApproval(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("delete", true))
}

func GetTrashApprovals(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetTrashApprovals(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func RestoreApproval(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.RestoreApproval(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("restore", true))
}

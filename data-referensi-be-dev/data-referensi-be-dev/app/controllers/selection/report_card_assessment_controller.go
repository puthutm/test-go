package controllers

import (
	"data-referensi/app/models"
	"data-referensi/app/requests"
	"data-referensi/handlers"
	"data-referensi/helpers"

	"github.com/gofiber/fiber/v2"
)

func GetReportCardAssessments(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetReportCardAssessments(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func SearchReportCardAssessments(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, err := models.SearchReportCardAssessments(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_results, helpers.GenerateRM("get", true))
}

func ExportReportCardAssessments(c *fiber.Ctx) error {
	return handlers.ControllerExport(c, "ReportCardAssessments", models.ExportReportCardAssessments)
}

func GetReportCardAssessment(c *fiber.Ctx) error {
	id := c.Params("id")
	query_result, err := models.GetReportCardAssessment(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("get", true))
}

func CreateReportCardAssessment(c *fiber.Ctx) error {
	var req requests.ReportCardAssessmentRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	id, err := helpers.EnsureUUID(&models.MstReportCardAssessment{})
	if err != nil {
		return err
	}

	err = models.CreateReportCardAssessment(id, req.Code, req.Name, req.Value, req.SubjectId, req.Note)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetReportCardAssessment(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusCreated, query_result, helpers.GenerateRM("insert", true))
}

func ImportReportCardAssessments(c *fiber.Ctx) error {
	return handlers.ControllerImport(c, models.ImportReportCardAssessments)
}

func UpdateReportCardAssessment(c *fiber.Ctx) error {
	id := c.Params("id")

	var req requests.ReportCardAssessmentRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	err := models.UpdateReportCardAssessment(id, req.Code, req.Name, req.Value, req.SubjectId, req.Note)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetReportCardAssessment(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("update", true))
}

func DeleteReportCardAssessment(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.DeleteReportCardAssessment(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("delete", true))
}

func GetTrashReportCardAssessments(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetTrashReportCardAssessments(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func RestoreReportCardAssessment(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.RestoreReportCardAssessment(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("restore", true))
}

package controllers

import (
	"data-referensi/app/models"
	"data-referensi/app/requests"
	"data-referensi/handlers"
	"data-referensi/helpers"

	"github.com/gofiber/fiber/v2"
)

func GetStudentCertificates(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetStudentCertificates(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func SearchStudentCertificates(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, err := models.SearchStudentCertificates(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_results, helpers.GenerateRM("get", true))
}

func ExportStudentCertificates(c *fiber.Ctx) error {
	return handlers.ControllerExport(c, "StudentCertificates", models.ExportStudentCertificates)
}

func GetStudentCertificate(c *fiber.Ctx) error {
	id := c.Params("id")
	query_result, err := models.GetStudentCertificate(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("get", true))
}

func CreateStudentCertificate(c *fiber.Ctx) error {
	var req requests.StudentCertificateRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	id, err := helpers.EnsureUUID(&models.MstStudentCertificate{})
	if err != nil {
		return err
	}

	description := ""
	if req.Description != nil {
		description = *req.Description
	}

	purposes := ""
	if req.Purposes != nil {
		purposes = *req.Purposes
	}

	err = models.CreateStudentCertificate(id, req.Name, description, purposes)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetStudentCertificate(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusCreated, query_result, helpers.GenerateRM("insert", true))
}

func ImportStudentCertificates(c *fiber.Ctx) error {
	return handlers.ControllerImport(c, models.ImportStudentCertificates)
}

func UpdateStudentCertificate(c *fiber.Ctx) error {
	id := c.Params("id")

	var req requests.StudentCertificateRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	description := ""
	if req.Description != nil {
		description = *req.Description
	}

	purposes := ""
	if req.Purposes != nil {
		purposes = *req.Purposes
	}

	err := models.UpdateStudentCertificate(id, req.Name, description, purposes)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetStudentCertificate(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("update", true))
}

func DeleteStudentCertificate(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.DeleteStudentCertificate(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("delete", true))
}

func GetTrashStudentCertificates(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetTrashStudentCertificates(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func RestoreStudentCertificate(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.RestoreStudentCertificate(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("restore", true))
}

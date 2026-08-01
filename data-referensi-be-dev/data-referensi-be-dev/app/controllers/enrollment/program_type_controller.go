package controllers

import (
	"data-referensi/app/models"
	"data-referensi/app/requests"
	"data-referensi/handlers"
	"data-referensi/helpers"

	"github.com/gofiber/fiber/v2"
)

func GetProgramTypes(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetProgramTypes(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func SearchProgramTypes(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, err := models.SearchProgramTypes(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_results, helpers.GenerateRM("get", true))
}

func ExportProgramTypes(c *fiber.Ctx) error {
	return handlers.ControllerExport(c, "ProgramTypes", models.ExportProgramTypes)
}

func GetProgramType(c *fiber.Ctx) error {
	id := c.Params("id")
	query_result, err := models.GetProgramType(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("get", true))
}

func CreateProgramType(c *fiber.Ctx) error {
	var req requests.ProgramTypeRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	id, err := helpers.EnsureUUID(&models.MstProgramType{})
	if err != nil {
		return err
	}

	err = models.CreateProgramType(id, req.Code, req.ProgramType, req.Description, req.IsIPC)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetProgramType(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusCreated, query_result, helpers.GenerateRM("insert", true))
}

func ImportProgramTypes(c *fiber.Ctx) error {
	return handlers.ControllerImport(c, models.ImportProgramTypes)
}

func UpdateProgramType(c *fiber.Ctx) error {
	id := c.Params("id")

	var req requests.ProgramTypeRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	err := models.UpdateProgramType(id, req.Code, req.ProgramType, req.Description, req.IsIPC)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetProgramType(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("update", true))
}

func DeleteProgramType(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.DeleteProgramType(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("delete", true))
}

func GetTrashProgramTypes(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetTrashProgramTypes(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func RestoreProgramType(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.RestoreProgramType(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("restore", true))
}

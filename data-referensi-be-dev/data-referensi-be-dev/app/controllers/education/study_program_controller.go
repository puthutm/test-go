package controllers

import (
	"data-referensi/app/models"
	"data-referensi/app/requests"
	"data-referensi/handlers"
	"data-referensi/helpers"

	"github.com/gofiber/fiber/v2"
)

func GetStudyPrograms(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)
	queryType := c.Query("type")

	query_results, total, err := models.GetStudyPrograms(filter, sortBy, sortDirection, page, pageSize, queryType)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func SearchStudyPrograms(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	queryType := c.Query("type")

	query_results, err := models.SearchStudyPrograms(filter, sortBy, sortDirection, page, pageSize, queryType)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_results, helpers.GenerateRM("get", true))
}

func ExportStudyPrograms(c *fiber.Ctx) error {
	return handlers.ControllerExport(c, "StudyPrograms", models.ExportStudyPrograms)
}

func GetStudyProgram(c *fiber.Ctx) error {
	id := c.Params("id")
	query_result, err := models.GetStudyProgram(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("get", true))
}

func CreateStudyProgram(c *fiber.Ctx) error {
	var req requests.StudyProgramRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	id, err := helpers.EnsureUUID(&models.MstStudyProgram{})
	if err != nil {
		return err
	}

	err = models.CreateStudyProgram(id, req.Code, req.Name, req.Type)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetStudyProgram(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusCreated, query_result, helpers.GenerateRM("insert", true))
}

func ImportStudyPrograms(c *fiber.Ctx) error {
	return handlers.ControllerImport(c, models.ImportStudyPrograms)
}

func UpdateStudyProgram(c *fiber.Ctx) error {
	id := c.Params("id")

	var req requests.StudyProgramRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	err := models.UpdateStudyProgram(id, req.Code, req.Name, req.Type)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetStudyProgram(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("update", true))
}

func DeleteStudyProgram(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.DeleteStudyProgram(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("delete", true))
}

func GetTrashStudyPrograms(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetTrashStudyPrograms(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func RestoreStudyProgram(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.RestoreStudyProgram(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("restore", true))
}

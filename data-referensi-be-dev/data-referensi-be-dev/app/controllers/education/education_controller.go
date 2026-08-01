package controllers

import (
	"data-referensi/app/models"
	"data-referensi/app/requests"
	"data-referensi/handlers"
	"data-referensi/helpers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetEducations(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetEducations(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func SearchEducations(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, err := models.SearchEducations(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_results, helpers.GenerateRM("get", true))
}

func ExportEducations(c *fiber.Ctx) error {
	return handlers.ControllerExport(c, "Educations", models.ExportEducations)
}

func GetEducationByEducationalLevelId(c *fiber.Ctx) error {
	educational_level_id := c.Params("educational_level_id")
	log.Print(educational_level_id)
	query_results, err := models.GetEducationByEducationalLevelId(educational_level_id)
	if err != nil {
		return handlers.SendSuccess(c, fiber.StatusBadRequest, nil, err.Error())
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_results, helpers.GenerateRM("get", true))
}

func GetEducation(c *fiber.Ctx) error {
	id := c.Params("id")
	query_result, err := models.GetEducation(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("get", true))
}

func CreateEducation(c *fiber.Ctx) error {
	var req requests.EducationRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	/* Check Existing ID */
	id, err := helpers.EnsureUUID(&models.MstEducation{})
	if err != nil {
		return err
	}

	err = models.CreateEducation(id, req.EducationalLevelId, req.StudyProgramId, req.Name)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetEducation(id)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	return handlers.SendSuccess(c, fiber.StatusCreated, query_result, helpers.GenerateRM("insert", true))
}

func ImportEducations(c *fiber.Ctx) error {
	return handlers.ControllerImport(c, models.ImportEducations)
}

func UpdateEducation(c *fiber.Ctx) error {
	id := c.Params("id")

	var req requests.EducationRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	err := models.UpdateEducation(id, req.EducationalLevelId, req.StudyProgramId, req.Name)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetEducation(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusCreated, query_result, helpers.GenerateRM("update", true))
}

func DeleteEducation(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.DeleteEducation(id)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusInternalServerError, nil, err.Error())
	}

	return handlers.SendSuccess(c, fiber.StatusCreated, nil, helpers.GenerateRM("delete", true))
}

func GetTrashEducations(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetTrashEducations(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func RestoreEducation(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.RestoreEducation(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusCreated, nil, helpers.GenerateRM("restore", true))
}

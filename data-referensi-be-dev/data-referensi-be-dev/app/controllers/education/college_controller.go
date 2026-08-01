package controllers

import (
	"data-referensi/app/models"
	"data-referensi/app/requests"
	"data-referensi/handlers"
	"data-referensi/helpers"

	"github.com/gofiber/fiber/v2"
)

func GetColleges(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetColleges(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func SearchColleges(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, err := models.SearchColleges(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_results, helpers.GenerateRM("get", true))
}

func ExportColleges(c *fiber.Ctx) error {
	return handlers.ControllerExport(c, "Colleges", models.ExportColleges)
}

func GetCollege(c *fiber.Ctx) error {
	id := c.Params("id")
	query_result, err := models.GetCollege(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("get", true))
}

func CreateCollege(c *fiber.Ctx) error {
	var req requests.CollegeRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	id, err := helpers.EnsureUUID(&models.MstCollege{})
	if err != nil {
		return err
	}

	err = models.CreateCollege(id, req.Name, req.ProvinceId, req.CityId, req.Type, req.Accreditation, req.ShortName, req.NumberOfStudyProgram, req.LowerLimitTuitionFee, req.UpperLimitTuitionFee)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetCollege(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusCreated, query_result, helpers.GenerateRM("insert", true))
}

func ImportColleges(c *fiber.Ctx) error {
	return handlers.ControllerImport(c, models.ImportColleges)
}

func UpdateCollege(c *fiber.Ctx) error {
	id := c.Params("id")

	var req requests.CollegeRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	err := models.UpdateCollege(id, req.Name, req.ProvinceId, req.CityId, req.Type, req.Accreditation, req.ShortName, req.NumberOfStudyProgram, req.LowerLimitTuitionFee, req.UpperLimitTuitionFee)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetCollege(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("update", true))
}

func DeleteCollege(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.DeleteCollege(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("delete", true))
}

func GetTrashColleges(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetTrashColleges(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func RestoreCollege(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.RestoreCollege(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("restore", true))
}

package controllers

import (
	"data-referensi/app/models"
	"data-referensi/app/requests"
	"data-referensi/handlers"
	"data-referensi/helpers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetGrades(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetGrades(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func SearchGrades(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, err := models.SearchGrades(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_results, helpers.GenerateRM("get", true))
}

func ExportGrades(c *fiber.Ctx) error {
	return handlers.ControllerExport(c, "Grades", models.ExportGrades)
}

func GetGrade(c *fiber.Ctx) error {
	id := c.Params("id")
	query_result, err := models.GetGrade(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("get", true))
}

func CreateGrade(c *fiber.Ctx) error {
	var req requests.GradeRequest

	log.Print(req)

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	id, err := helpers.EnsureUUID(&models.MstGrade{})
	if err != nil {
		return err
	}

	// lowerLimitStr := req.LowerLimit
	// lowerLimit, err := strconv.ParseFloat(lowerLimitStr, 64)
	// if err != nil {
	// 	fmt.Println("Error konversi lower_limit:", err)
	// }

	// upperLimitStr := req.UpperLimit
	// upperLimit, err := strconv.ParseFloat(upperLimitStr, 64)
	// if err != nil {
	// 	fmt.Println("Error konversi upper_limit:", err)
	// }

	err = models.CreateGrade(id, req.Code, req.Name, req.LowerLimit, req.UpperLimit, req.Description)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetGrade(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusCreated, query_result, helpers.GenerateRM("insert", true))
}

func ImportGrades(c *fiber.Ctx) error {
	return handlers.ControllerImport(c, models.ImportGrades)
}

func UpdateGrade(c *fiber.Ctx) error {
	id := c.Params("id")

	var req requests.GradeRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	// lowerLimitStr := req.LowerLimit
	// lowerLimit, err := strconv.ParseFloat(lowerLimitStr, 64)
	// if err != nil {
	// 	fmt.Println("Error konversi lower_limit:", err)
	// }

	// upperLimitStr := req.UpperLimit
	// upperLimit, err := strconv.ParseFloat(upperLimitStr, 64)
	// if err != nil {
	// 	fmt.Println("Error konversi upper_limit:", err)
	// }

	err := models.UpdateGrade(id, req.Code, req.Name, req.LowerLimit, req.UpperLimit, req.Description)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetGrade(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("update", true))
}

func DeleteGrade(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.DeleteGrade(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("delete", true))
}

func GetTrashGrades(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetTrashGrades(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func RestoreGrade(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.RestoreGrade(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("restore", true))
}

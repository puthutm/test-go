package controllers

import (
	"data-referensi/app/models"
	"data-referensi/app/requests"
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAcademicPeriods(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetAcademicPeriods(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func SearchAcademicPeriods(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, err := models.SearchAcademicPeriods(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_results, helpers.GenerateRM("get", true))
}

func ExportAcademicPeriods(c *fiber.Ctx) error {
	return handlers.ControllerExport(c, "AcademicPeriods", models.ExportAcademicPeriods)
}

func GetAcademicPeriod(c *fiber.Ctx) error {
	id := c.Params("id")
	query_result, err := models.GetAcademicPeriod(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("get", true))
}

func GetAcademicPeriodDetailWithSession(c *fiber.Ctx) error {
	id := c.Params("id")
	query_result, err := models.GetAcademicPeriodDetailWithSession(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("get", true))
}

func CreateAcademicPeriod(c *fiber.Ctx) error {
	var req requests.AcademicPeriodRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	id, err := helpers.EnsureUUID(&models.MstAcademicPeriod{})
	if err != nil {
		return err
	}

	semester, err := models.GetSemester(req.SemesterId)
	if err != nil {
		return err
	}

	StartDateOfCollege, err := time.Parse("2006-01-02", req.StartDateOfCollege)
	if err != nil {
		fmt.Println("Format tanggal salah:", err)
		return err
	}

	// Jadwal UTS
	UTSSession, err := strconv.Atoi(semester.UTSSession)
	if err != nil {
		fmt.Println("Format sesi UTS salah:", err)
		return err
	}

	StartDateOfUts := StartDateOfCollege.AddDate(0, 0, UTSSession*7)
	EndDateOfUts := StartDateOfUts.AddDate(0, 0, 7)

	// Jadwal UAS
	UASSession, err := strconv.Atoi(semester.UASSession)
	if err != nil {
		fmt.Println("Format sesi UTS salah:", err)
		return err
	}

	StartDateOfUas := StartDateOfCollege.AddDate(0, 0, UASSession*7)
	EndDateOfUas := StartDateOfUas.AddDate(0, 0, 7)

	req.StartDateOfUts = StartDateOfUts.Format("2006-01-02")
	req.EndDateOfUts = EndDateOfUts.Format("2006-01-02")
	req.StartDateOfUas = StartDateOfUas.Format("2006-01-02")
	req.EndDateOfUas = EndDateOfUas.Format("2006-01-02")
	req.IsActive = true

	err = models.CreateAcademicPeriod(id, req.Code, req.AcademicYearId, req.SemesterId, req.Fullname, req.Shortname, req.StartDateOfCollege, req.EndDateOfCollege, req.StartDateOfUts, req.EndDateOfUts, req.StartDateOfUas, req.EndDateOfUas, req.NumberOfLectureMeeting, req.IsActive)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetAcademicPeriod(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusCreated, query_result, helpers.GenerateRM("insert", true))
}

func ImportAcademicPeriods(c *fiber.Ctx) error {
	return handlers.ControllerImport(c, models.ImportAcademicPeriods)
}

func UpdateAcademicPeriod(c *fiber.Ctx) error {
	id := c.Params("id")

	var req requests.AcademicPeriodRequest

	if err := c.BodyParser(&req); err != nil {
		return handlers.SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	err := models.UpdateAcademicPeriod(id, req.Code, req.AcademicYearId, req.SemesterId, req.Fullname, req.Shortname, req.StartDateOfCollege, req.EndDateOfCollege, req.StartDateOfUts, req.EndDateOfUts, req.StartDateOfUas, req.EndDateOfUas, req.NumberOfLectureMeeting, req.IsActive)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	query_result, err := models.GetAcademicPeriod(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, query_result, helpers.GenerateRM("update", true))
}

func DeleteAcademicPeriod(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.DeleteAcademicPeriod(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("delete", true))
}

func GetTrashAcademicPeriods(c *fiber.Ctx) error {
	filter, sortBy, sortDirection, page, pageSize := handlers.ControllerQueryParams(c)

	query_results, total, err := models.GetTrashAcademicPeriods(filter, sortBy, sortDirection, page, pageSize)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("get", false))
	}

	results := handlers.ResponseWithMetadata(query_results, page, pageSize, len(query_results), total)

	return handlers.SendSuccess(c, fiber.StatusOK, results, helpers.GenerateRM("get", true))
}

func RestoreAcademicPeriod(c *fiber.Ctx) error {
	id := c.Params("id")
	err := models.RestoreAcademicPeriod(id)
	if err != nil {
		return handlers.ErrorQuery(c, err)
	}

	return handlers.SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("restore", true))
}

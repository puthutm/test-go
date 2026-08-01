package handlers

import (
	"data-referensi/helpers"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

/* Get Query Params */
func ControllerQueryParams(c *fiber.Ctx) (filter, sortBy, sortDirection string, page int, pageSize int64) {
	filter = c.Query("filter", "")
	sortBy = c.Query("sort_by", "name")
	sortDirection = c.Query("sort_direction", "asc")
	page = c.QueryInt("page", 1)
	pageSize = int64(c.QueryInt("page_size", 10))
	return
}

/* Controller Export */
func ControllerExport(c *fiber.Ctx, modelName string, exportFunc func(c *fiber.Ctx, filePath string) error) error {
	fileName := fmt.Sprintf("%s.xlsx", modelName)
	fileSaveAs := fmt.Sprintf("tmp/exports/%s", fileName)

	if err := exportFunc(c, fileSaveAs); err != nil {
		return SendFailed(c, fiber.StatusOK, nil, helpers.GenerateRM("export", false))
	}

	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	return c.SendFile(fileSaveAs, false)
}

/* Controller Import Handler */
func ControllerImport(c *fiber.Ctx, importFunc func(string) error) error {
	file, err := c.FormFile("file_import")
	if err != nil {
		return SendFailed(c, fiber.StatusBadRequest, nil, err.Error())
	}

	filePath := fmt.Sprintf("%s/%s", "./tmp/uploads", file.Filename)
	if err := c.SaveFile(file, filePath); err != nil {
		return SendFailed(c, fiber.StatusInternalServerError, nil, helpers.GenerateRM("save", false))
	}

	if err := importFunc(filePath); err != nil {
		if strings.Contains(err.Error(), "duplicate key row") {
			return SendFailed(c, fiber.StatusBadRequest, nil, helpers.GenerateRM("exist"))
		}
		return SendFailed(c, fiber.StatusInternalServerError, nil, err.Error())
	}

	if err := os.Remove(filePath); err != nil {
		log.Println("Error removing uploaded file:", err)
	}

	return SendSuccess(c, fiber.StatusOK, nil, helpers.GenerateRM("import", true))
}

package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstSemesterNumber struct {
	ID             string `json:"id"`
	SemesterNumber string `json:"semester_number"`
	CreatedAt      int64  `json:"created_at"`
	UpdatedAt      int64  `json:"updated_at"`
}

type MstSemesterNumberDetail struct {
	ID             string `json:"id"`
	SemesterNumber string `json:"semester_number"`
	CreatedAt      int64  `json:"created_at"`
	UpdatedAt      int64  `json:"updated_at"`
}

type MstSemesterNumberExport struct {
	ID             string `json:"id"`
	SemesterNumber string `json:"semester_number"`
}

type MstSemesterNumberSearch struct {
	ID             string `json:"id"`
	SemesterNumber string `json:"semester_number"`
}

type MstSemesterNumberRelation struct {
	ID             string `json:"id"`
	SemesterNumber string `json:"semester_number"`
}

/* Action */
func GetSemesterNumbers(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstSemesterNumber, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_semester_numbers_get", filter, sortBy, sortDirection, page, pageSize, &MstSemesterNumber{})
	if err != nil {
		return []MstSemesterNumber{}, 0, err
	}

	var modelResults []MstSemesterNumber
	for _, item := range results {
		level, ok := item.(*MstSemesterNumber)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchSemesterNumbers(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstSemesterNumberSearch, error) {
	var results []MstSemesterNumberSearch
	err := handlers.SPGet("sp_mst_semester_numbers_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstSemesterNumberSearch{}, err
	}
	return results, nil
}

func ExportSemesterNumbers(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstSemesterNumberExport
	err := handlers.SPGet("sp_mst_semester_numbers_get", "", "name", "asc", 1, CountSemesterNumbers(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":             result.ID,
			"SemesterNumber": result.SemesterNumber,
		}
	}

	headers := []string{
		"ID", "Semester Number",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetSemesterNumber(id string) (MstSemesterNumberDetail, error) {
	var result MstSemesterNumberDetail
	err := handlers.SPGetByID("sp_mst_semester_numbers_get_by_id", id, &result)
	if err != nil {
		return MstSemesterNumberDetail{}, err
	}
	return result, nil
}

func CreateSemesterNumber(id string, semester_number string) error {
	return QueryInsertSemesterNumber(id, semester_number)
}

func ImportSemesterNumbers(fileStatus string) error {
	headers := []string{
		"id", "semester_number",
	}

	return handlers.ModelImport(fileStatus, headers, func(row map[string]string) error {
		id := row["id"]
		semester_number := row["semester_number"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstSemesterNumber{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateSemesterNumber(id, semester_number); err != nil {
					return err
				}
			} else {
				if err := QueryInsertSemesterNumber(id, semester_number); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstSemesterNumber{})
			if err != nil {
				return err
			}
			if err := QueryInsertSemesterNumber(id, semester_number); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateSemesterNumber(id string, semester_number string) error {
	return QueryUpdateSemesterNumber(id, semester_number)
}

func DeleteSemesterNumber(id string) error {
	err := handlers.SPDelete("sp_mst_semester_numbers_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashSemesterNumbers(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstSemesterNumber, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_semester_numbers_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstSemesterNumber{})
	if err != nil {
		return []MstSemesterNumber{}, 0, err
	}

	var modelResults []MstSemesterNumber
	for _, item := range results {
		level, ok := item.(*MstSemesterNumber)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreSemesterNumber(id string) error {
	err := handlers.SPRestore("sp_mst_semester_numbers_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountSemesterNumbers(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstSemesterNumber{}, nullable)
}

/* Query */
func QueryInsertSemesterNumber(id string, semester_number string) error {
	query := `
		EXEC sp_mst_semester_numbers_insert
		@id = ?,
		@semester_number = ?,
		@created_at = ?,
		@created_by = ?
	`

	err := handlers.SPInsertSemesterNumber(query, id, semester_number)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateSemesterNumber(id string, semester_number string) error {
	query := `
		EXEC sp_mst_semester_numbers_update
		@id = ?,
		@semester_number = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, semester_number)
	if err != nil {
		return err
	}

	return nil
}

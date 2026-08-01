package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstEnrollmentBatch struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Batch       string `json:"batch"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

func (MstEnrollmentBatch) TableName() string {
	return "mst_enrollment_batchs"
}

type MstEnrollmentBatchDetail struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Batch       string `json:"batch"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstEnrollmentBatchExport struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Batch       string `json:"batch"`
	Description string `json:"description"`
}

type MstEnrollmentBatchSearch struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Batch       string `json:"batch"`
	Description string `json:"description"`
}

type MstEnrollmentBatchRelation struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Batch       string `json:"batch"`
	Description string `json:"description"`
}

/* Action */
func GetEnrollmentBatchs(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstEnrollmentBatch, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_enrollment_batchs_get", filter, sortBy, sortDirection, page, pageSize, &MstEnrollmentBatch{})
	if err != nil {
		return []MstEnrollmentBatch{}, 0, err
	}

	var modelResults []MstEnrollmentBatch
	for _, item := range results {
		level, ok := item.(*MstEnrollmentBatch)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchEnrollmentBatchs(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstEnrollmentBatchSearch, error) {
	var results []MstEnrollmentBatchSearch
	err := handlers.SPGet("sp_mst_enrollment_batchs_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstEnrollmentBatchSearch{}, err
	}
	return results, nil
}

func ExportEnrollmentBatchs(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstEnrollmentBatchExport
	err := handlers.SPGet("sp_mst_enrollment_batchs_get", "", "name", "asc", 1, CountEnrollmentBatchs(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":          result.ID,
			"Code":        result.Code,
			"Batch":       result.Batch,
			"Description": result.Description,
		}
	}

	headers := []string{
		"ID", "Code", "Batch", "Description",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetEnrollmentBatch(id string) (MstEnrollmentBatchDetail, error) {
	var result MstEnrollmentBatchDetail
	err := handlers.SPGetByID("sp_mst_enrollment_batchs_get_by_id", id, &result)
	if err != nil {
		return MstEnrollmentBatchDetail{}, err
	}
	return result, nil
}

func CreateEnrollmentBatch(id string, code string, batch string, description string) error {
	return QueryInsertEnrollmentBatch(id, code, batch, description)
}

func ImportEnrollmentBatchs(filePath string) error {
	headers := []string{
		"id", "code", "batch", "description",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		batch := row["batch"]
		description := row["description"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstEnrollmentBatch{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateEnrollmentBatch(id, code, batch, description); err != nil {
					return err
				}
			} else {
				if err := QueryInsertEnrollmentBatch(id, code, batch, description); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstEnrollmentBatch{})
			if err != nil {
				return err
			}
			if err := QueryInsertEnrollmentBatch(id, code, batch, description); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateEnrollmentBatch(id string, code string, batch string, description string) error {
	return QueryUpdateEnrollmentBatch(id, code, batch, description)
}

func DeleteEnrollmentBatch(id string) error {
	err := handlers.SPDelete("sp_mst_enrollment_batchs_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashEnrollmentBatchs(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstEnrollmentBatch, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_enrollment_batchs_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstEnrollmentBatch{})
	if err != nil {
		return []MstEnrollmentBatch{}, 0, err
	}

	var modelResults []MstEnrollmentBatch
	for _, item := range results {
		level, ok := item.(*MstEnrollmentBatch)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreEnrollmentBatch(id string) error {
	err := handlers.SPRestore("sp_mst_enrollment_batchs_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountEnrollmentBatchs(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstEnrollmentBatch{}, nullable)
}

/* Query */
func QueryInsertEnrollmentBatch(id string, code string, batch string, description string) error {
	query := `
		EXEC sp_mst_enrollment_batchs_insert
		@id = ?,
		@code = ?,
		@batch= ?,
		@description = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, code, batch, description)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateEnrollmentBatch(id string, code string, batch string, description string) error {
	query := `
		EXEC sp_mst_enrollment_batchs_update
		@id = ?,
		@code = ?,
		@batch= ?,
		@description = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, code, batch, description)
	if err != nil {
		return err
	}

	return nil
}

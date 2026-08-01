package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type MstValueScale struct {
	ID          string `json:"id"`
	Value       string `json:"value"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstValueScaleDetail struct {
	ID          string `json:"id"`
	Value       string `json:"value"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstValueScaleExport struct {
	ID          string `json:"id"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

type MstValueScaleSearch struct {
	ID          string `json:"id"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

type MstValueScaleRelation struct {
	ID          string `json:"id"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

/* Action */
func GetValueScales(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstValueScale, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_value_scales_get", filter, sortBy, sortDirection, page, pageSize, &MstValueScale{})
	if err != nil {
		return []MstValueScale{}, 0, err
	}

	var modelResults []MstValueScale
	for _, item := range results {
		level, ok := item.(*MstValueScale)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchValueScales(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstValueScaleSearch, error) {
	var results []MstValueScaleSearch
	err := handlers.SPGet("sp_mst_value_scales_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstValueScaleSearch{}, err
	}
	return results, nil
}

func ExportValueScales(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstValueScaleExport
	err := handlers.SPGet("sp_mst_value_scales_get", "", "value", "asc", 1, CountValueScales(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":          result.ID,
			"Value":       result.Value,
			"Description": result.Description,
		}
	}

	headers := []string{
		"ID", "Value", "Description",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetValueScale(id string) (MstValueScaleDetail, error) {
	var result MstValueScaleDetail
	err := handlers.SPGetByID("sp_mst_value_scales_get_by_id", id, &result)
	if err != nil {
		return MstValueScaleDetail{}, err
	}
	return result, nil
}

func CreateValueScale(id string, value string, description string) error {
	return QueryInsertValueScale(id, value, description)
}

func ImportValueScales(fileStatus string) error {
	headers := []string{
		"id", "value", "description",
	}

	return handlers.ModelImport(fileStatus, headers, func(row map[string]string) error {
		id := row["id"]
		value := row["value"]
		description := row["description"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstValueScale{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateValueScale(id, value, description); err != nil {
					return err
				}
			} else {
				if err := QueryInsertValueScale(id, value, description); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstValueScale{})
			if err != nil {
				return err
			}
			if err := QueryInsertValueScale(id, value, description); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateValueScale(id string, value string, description string) error {
	return QueryUpdateValueScale(id, value, description)
}

func DeleteValueScale(id string) error {
	err := handlers.SPDelete("sp_mst_value_scales_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashValueScales(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstValueScale, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_value_scales_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstValueScale{})
	if err != nil {
		return []MstValueScale{}, 0, err
	}

	var modelResults []MstValueScale
	for _, item := range results {
		level, ok := item.(*MstValueScale)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreValueScale(id string) error {
	err := handlers.SPRestore("sp_mst_value_scales_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountValueScales(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstValueScale{}, nullable)
}

/* Query */
func QueryInsertValueScale(id string, value string, description string) error {
	query := `
		EXEC sp_mst_value_scales_insert
		@id = ?,
		@value = ?, 
		@description = ?, 
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`
	log.Print(query)

	err := handlers.SPInsert(query, id, value, description)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateValueScale(id string, value string, description string) error {
	query := `
		EXEC sp_mst_value_scales_update
		@id = ?,
		@value = ?, 
		@description = ?, 
		@updated_at = ?,
		@updated_by = ?
	`
	err := handlers.SPUpdate(query, id, value, description)
	if err != nil {
		return err
	}

	return nil
}

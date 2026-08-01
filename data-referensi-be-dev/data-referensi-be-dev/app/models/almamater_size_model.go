package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstAlmamaterSize struct {
	ID         string `json:"id"`
	Code       string `json:"code"`
	Size       string `json:"size"`
	ChestSize  string `json:"chest_size"`
	ArmLength  string `json:"arm_length"`
	BodyLength string `json:"body_length"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}

type MstAlmamaterSizeDetail struct {
	ID         string `json:"id"`
	Code       string `json:"code"`
	Size       string `json:"size"`
	ChestSize  string `json:"chest_size"`
	ArmLength  string `json:"arm_length"`
	BodyLength string `json:"body_length"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}

type MstAlmamaterSizeExport struct {
	ID         string `json:"id"`
	Code       string `json:"code"`
	Size       string `json:"size"`
	ChestSize  string `json:"chest_size"`
	ArmLength  string `json:"arm_length"`
	BodyLength string `json:"body_length"`
}

type MstAlmamaterSizeSearch struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Size string `json:"size"`
}

type MstAlmamaterSizeRelation struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Size string `json:"size"`
}

/* Action */
func GetAlmamaterSizes(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstAlmamaterSize, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_almamater_sizes_get", filter, sortBy, sortDirection, page, pageSize, &MstAlmamaterSize{})
	if err != nil {
		return []MstAlmamaterSize{}, 0, err
	}

	var modelResults []MstAlmamaterSize
	for _, item := range results {
		level, ok := item.(*MstAlmamaterSize)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchAlmamaterSizes(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstAlmamaterSizeSearch, error) {
	var results []MstAlmamaterSizeSearch
	err := handlers.SPGet("sp_mst_almamater_sizes_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstAlmamaterSizeSearch{}, err
	}
	return results, nil
}

func ExportAlmamaterSizes(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstAlmamaterSizeExport
	err := handlers.SPGet("sp_mst_almamater_sizes_get", "", "name", "asc", 1, CountAlmamaterSizes(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":          result.ID,
			"Code":        result.Code,
			"Size":        result.Size,
			"Chest Size":  result.ChestSize,
			"Arm Length":  result.ArmLength,
			"Body Length": result.BodyLength,
		}
	}

	headers := []string{
		"ID",
		"Code",
		"Size",
		"Chest Size",
		"Arm Length",
		"Body Length",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetAlmamaterSize(id string) (MstAlmamaterSizeDetail, error) {
	var result MstAlmamaterSizeDetail
	err := handlers.SPGetByID("sp_mst_almamater_sizes_get_by_id", id, &result)
	if err != nil {
		return MstAlmamaterSizeDetail{}, err
	}
	return result, nil
}

func CreateAlmamaterSize(id string, code string, size string, chest_size string, arm_length string, body_length string) error {
	return QueryInsertAlmamaterSize(id, code, size, chest_size, arm_length, body_length)
}

func ImportAlmamaterSizes(filePath string) error {
	headers := []string{
		"id", "code", "size", "chest_size", "arm_length", "body_length",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		size := row["size"]
		chest_size := row["chest_size"]
		arm_length := row["arm_length"]
		body_length := row["body_length"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstAlmamaterSize{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateAlmamaterSize(id, code, size, chest_size, arm_length, body_length); err != nil {
					return err
				}
			} else {
				if err := QueryInsertAlmamaterSize(id, code, size, chest_size, arm_length, body_length); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstAlmamaterSize{})
			if err != nil {
				return err
			}
			if err := QueryInsertAlmamaterSize(id, code, size, chest_size, arm_length, body_length); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateAlmamaterSize(id string, code string, size string, chest_size string, arm_length string, body_length string) error {
	return QueryUpdateAlmamaterSize(id, code, size, chest_size, arm_length, body_length)
}

func DeleteAlmamaterSize(id string) error {
	err := handlers.SPDelete("sp_mst_almamater_sizes_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashAlmamaterSizes(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstAlmamaterSize, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_almamater_sizes_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstAlmamaterSize{})
	if err != nil {
		return []MstAlmamaterSize{}, 0, err
	}

	var modelResults []MstAlmamaterSize
	for _, item := range results {
		level, ok := item.(*MstAlmamaterSize)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreAlmamaterSize(id string) error {
	err := handlers.SPRestore("sp_mst_almamater_sizes_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountAlmamaterSizes(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstAlmamaterSize{}, nullable)
}

/* Query */
func QueryInsertAlmamaterSize(id string, code string, size string, chest_size string, arm_length string, body_length string) error {
	query := `
		EXEC sp_mst_almamater_sizes_insert
		@id = ?,
		@code = ?,
		@size = ?,
		@chest_size = ?,
		@arm_length = ?,
		@body_length = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, code, size, chest_size, arm_length, body_length)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateAlmamaterSize(id string, code string, size string, chest_size string, arm_length string, body_length string) error {
	query := `
		EXEC sp_mst_almamater_sizes_update
		@id = ?,
		@code = ?,
		@size = ?,
		@chest_size = ?,
		@arm_length = ?,
		@body_length = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, code, size, chest_size, arm_length, body_length)
	if err != nil {
		return err
	}

	return nil
}

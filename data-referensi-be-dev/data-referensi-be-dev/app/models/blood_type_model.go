package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstBloodType struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstBloodTypeDetail struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstBloodTypeExport struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type MstBloodTypeSearch struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}
type MstBloodTypeRelation struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

/* Action */
func GetBloodTypes(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstBloodType, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_blood_types_get", filter, sortBy, sortDirection, page, pageSize, &MstBloodType{})
	if err != nil {
		return []MstBloodType{}, 0, err
	}

	var modelResults []MstBloodType
	for _, item := range results {
		level, ok := item.(*MstBloodType)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchBloodTypes(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstBloodTypeSearch, error) {
	var results []MstBloodTypeSearch
	err := handlers.SPGet("sp_mst_blood_types_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstBloodTypeSearch{}, err
	}
	return results, nil
}

func ExportBloodTypes(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstBloodTypeExport
	err := handlers.SPGet("sp_mst_blood_types_get", "", "name", "asc", 1, CountBloodTypes(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":   result.ID,
			"Code": result.Code,
			"Name": result.Name,
		}
	}

	headers := []string{
		"ID",
		"Code",
		"Name",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetBloodType(id string) (MstBloodTypeDetail, error) {
	var result MstBloodTypeDetail
	err := handlers.SPGetByID("sp_mst_blood_types_get_by_id", id, &result)
	if err != nil {
		return MstBloodTypeDetail{}, err
	}
	return result, nil
}

func CreateBloodType(id string, code string, name string) error {
	return QueryInsertBloodType(id, code, name)
}

func ImportBloodTypes(filePath string) error {
	headers := []string{
		"id", "code", "name",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		name := row["name"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstBloodType{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateBloodType(id, code, name); err != nil {
					return err
				}
			} else {
				if err := QueryInsertBloodType(id, code, name); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstBloodType{})
			if err != nil {
				return err
			}
			if err := QueryInsertBloodType(id, code, name); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateBloodType(id string, code string, name string) error {
	return QueryUpdateBloodType(id, code, name)
}

func DeleteBloodType(id string) error {
	err := handlers.SPDelete("sp_mst_blood_types_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashBloodTypes(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstBloodType, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_blood_types_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstBloodType{})
	if err != nil {
		return []MstBloodType{}, 0, err
	}

	var modelResults []MstBloodType
	for _, item := range results {
		level, ok := item.(*MstBloodType)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreBloodType(id string) error {
	err := handlers.SPRestore("sp_mst_blood_types_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountBloodTypes(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstBloodType{}, nullable)
}

/* Query */
func QueryInsertBloodType(id string, code string, name string) error {
	query := `
		EXEC sp_mst_blood_types_insert
		@id = ?,
		@code = ?,
		@name = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, code, name)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateBloodType(id string, code string, name string) error {
	query := `
		EXEC sp_mst_blood_types_update
		@id = ?,
		@code = ?,
		@name = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, code, name)
	if err != nil {
		return err
	}

	return nil
}

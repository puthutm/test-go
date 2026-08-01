package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstSchoolType struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstSchoolTypeDetail struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstSchoolTypeExport struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type MstSchoolTypeSearch struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type MstSchoolTypeRelation struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

/* Action */
func GetSchoolTypes(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstSchoolType, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_school_types_get", filter, sortBy, sortDirection, page, pageSize, &MstSchoolType{})
	if err != nil {
		return []MstSchoolType{}, 0, err
	}

	var modelResults []MstSchoolType
	for _, item := range results {
		level, ok := item.(*MstSchoolType)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchSchoolTypes(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstSchoolTypeSearch, error) {
	var results []MstSchoolTypeSearch
	err := handlers.SPGet("sp_mst_school_types_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstSchoolTypeSearch{}, err
	}
	return results, nil
}

func ExportSchoolTypes(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstSchoolTypeExport
	err := handlers.SPGet("sp_mst_school_types_get", "", "name", "asc", 1, CountSchoolTypes(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":   result.ID,
			"Name": result.Name,
		}
	}

	headers := []string{
		"ID", "Name",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetSchoolType(id string) (MstSchoolTypeDetail, error) {
	var result MstSchoolTypeDetail
	err := handlers.SPGetByID("sp_mst_school_types_get_by_id", id, &result)
	if err != nil {
		return MstSchoolTypeDetail{}, err
	}
	return result, nil
}

func CreateSchoolType(id string, name string) error {
	return QueryInsertSchoolType(id, name)
}

func ImportSchoolTypes(fileStatus string) error {
	headers := []string{
		"id", "name",
	}

	return handlers.ModelImport(fileStatus, headers, func(row map[string]string) error {
		id := row["id"]
		name := row["name"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstSchoolType{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateSchoolType(id, name); err != nil {
					return err
				}
			} else {
				if err := QueryInsertSchoolType(id, name); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstSchoolType{})
			if err != nil {
				return err
			}
			if err := QueryInsertSchoolType(id, name); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateSchoolType(id string, name string) error {
	return QueryUpdateSchoolType(id, name)
}

func DeleteSchoolType(id string) error {
	err := handlers.SPDelete("sp_mst_school_types_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashSchoolTypes(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstSchoolType, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_school_types_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstSchoolType{})
	if err != nil {
		return []MstSchoolType{}, 0, err
	}

	var modelResults []MstSchoolType
	for _, item := range results {
		level, ok := item.(*MstSchoolType)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreSchoolType(id string) error {
	err := handlers.SPRestore("sp_mst_school_types_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountSchoolTypes(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstSchoolType{}, nullable)
}

/* Query */
func QueryInsertSchoolType(id string, name string) error {
	query := `
		EXEC sp_mst_school_types_insert
		@id = ?,
		@name = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, name)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateSchoolType(id string, name string) error {
	query := `
		EXEC sp_mst_school_types_update
		@id = ?,
		@name = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, name)
	if err != nil {
		return err
	}

	return nil
}

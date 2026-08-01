package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstEducationalLevel struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstEducationalLevelDetail struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstEducationalLevelExport struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MstEducationalLevelSearch struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type MstEducationalLevelRelation struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetEducationalLevels(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstEducationalLevel, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_educational_levels_get", filter, sortBy, sortDirection, page, pageSize, &MstEducationalLevel{})
	if err != nil {
		return []MstEducationalLevel{}, 0, err
	}
	var modelResults []MstEducationalLevel
	for _, item := range results {
		level, ok := item.(*MstEducationalLevel)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}
	return modelResults, total, nil
}

func SearchEducationalLevels(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstEducationalLevelSearch, error) {
	var results []MstEducationalLevelSearch
	err := handlers.SPGet("sp_mst_educational_levels_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstEducationalLevelSearch{}, err
	}
	return results, nil
}

func ExportEducationalLevels(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstEducationalLevelExport
	err := handlers.SPGet("sp_mst_educational_levels_get", "", "name", "asc", 1, CountEducationalLevels(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":          result.ID,
			"Code":        result.Code,
			"Name":        result.Name,
			"Description": result.Description,
		}
	}

	headers := []string{
		"ID",
		"Code",
		"Name",
		"Description",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetEducationalLevel(id string) (MstEducationalLevelDetail, error) {
	var result MstEducationalLevelDetail
	err := handlers.SPGetByID("sp_mst_educational_levels_get_by_id", id, &result)
	if err != nil {
		return MstEducationalLevelDetail{}, err
	}
	return result, nil
}

func CreateEducationalLevel(id string, code string, name string, description string) error {
	return QueryInsertEducationalLevel(id, code, name, description)
}

func ImportEducationalLevels(filePath string) error {
	headers := []string{
		"id", "code", "name", "description",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		name := row["name"]
		description := row["description"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstEducationalLevel{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateEducationalLevel(id, code, name, description); err != nil {
					return err
				}
			} else {
				if err := QueryInsertEducationalLevel(id, code, name, description); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstEducationalLevel{})
			if err != nil {
				return err
			}
			if err := QueryInsertEducationalLevel(id, code, name, description); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateEducationalLevel(id string, code string, name string, description string) error {
	return QueryUpdateEducationalLevel(id, code, name, description)
}

func DeleteEducationalLevel(id string) error {
	err := handlers.SPDelete("sp_mst_educational_levels_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashEducationalLevels(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstEducationalLevel, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_educational_levels_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstEducationalLevel{})
	if err != nil {
		return []MstEducationalLevel{}, 0, err
	}

	var modelResults []MstEducationalLevel
	for _, item := range results {
		level, ok := item.(*MstEducationalLevel)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreEducationalLevel(id string) error {
	err := handlers.SPRestore("sp_mst_educational_levels_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountEducationalLevels(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstEducationalLevel{}, nullable)
}

/* Query */
func QueryInsertEducationalLevel(id string, code string, name string, description string) error {
	query := `
		EXEC sp_mst_educational_levels_insert
		@id = ?,
		@code = ?,
		@name = ?,
		@description = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, code, name, description)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateEducationalLevel(id string, code string, name string, description string) error {
	query := `
		EXEC sp_mst_educational_levels_update
		@id = ?,
		@code = ?,
		@name = ?,
		@description = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, code, name, description)
	if err != nil {
		return err
	}

	return nil
}

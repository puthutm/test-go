package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstScience struct {
	ID          string `json:"id"`
	Name        string `json:"name" `
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstScienceDetail struct {
	ID          string `json:"id"`
	Name        string `json:"name" `
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstScienceExport struct {
	ID          string `json:"id"`
	Name        string `json:"name" `
	Description string `json:"description"`
}

type MstScienceSearch struct {
	ID          string `json:"id"`
	Name        string `json:"name" `
	Description string `json:"description"`
}

type MstScienceRelation struct {
	ID   string `json:"id"`
	Name string `json:"name" `
}

/* Action */
func GetSciences(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstScience, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_sciences_get", filter, sortBy, sortDirection, page, pageSize, &MstScience{})
	if err != nil {
		return []MstScience{}, 0, err
	}

	var modelResults []MstScience
	for _, item := range results {
		level, ok := item.(*MstScience)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchSciences(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstScienceSearch, error) {
	var results []MstScienceSearch
	err := handlers.SPGet("sp_mst_sciences_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstScienceSearch{}, err
	}
	return results, nil
}

func ExportSciences(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstScienceExport
	err := handlers.SPGet("sp_mst_sciences_get", "", "name", "asc", 1, CountSciences(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":          result.ID,
			"Name":        result.Name,
			"Description": result.Description,
		}
	}

	headers := []string{
		"ID", "Name", "Description",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetScience(id string) (MstScienceDetail, error) {
	var result MstScienceDetail
	err := handlers.SPGetByID("sp_mst_sciences_get_by_id", id, &result)
	if err != nil {
		return MstScienceDetail{}, err
	}
	return result, nil
}

func CreateScience(id string, name string, description string) error {
	return QueryInsertScience(id, name, description)
}

func ImportSciences(filePath string) error {
	headers := []string{
		"id", "name", "description",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		name := row["name"]
		description := row["description"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstScience{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateScience(id, name, description); err != nil {
					return err
				}
			} else {
				if err := QueryInsertScience(id, name, description); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstScience{})
			if err != nil {
				return err
			}
			if err := QueryInsertScience(id, name, description); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateScience(id string, name string, description string) error {
	return QueryUpdateScience(id, name, description)
}

func DeleteScience(id string) error {
	err := handlers.SPDelete("sp_mst_sciences_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashSciences(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstScience, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_sciences_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstScience{})
	if err != nil {
		return []MstScience{}, 0, err
	}

	var modelResults []MstScience
	for _, item := range results {
		level, ok := item.(*MstScience)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreScience(id string) error {
	err := handlers.SPRestore("sp_mst_sciences_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountSciences(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstScience{}, nullable)
}

/* Query */
func QueryInsertScience(id string, name string, description string) error {
	query := `
		EXEC sp_mst_sciences_insert
		@id = ?,
		@name = ?,
		@description = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, name, description)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateScience(id string, name string, description string) error {
	query := `
		EXEC sp_mst_sciences_update
		@id = ?,
		@name = ?,
		@description = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, name, description)
	if err != nil {
		return err
	}

	return nil
}

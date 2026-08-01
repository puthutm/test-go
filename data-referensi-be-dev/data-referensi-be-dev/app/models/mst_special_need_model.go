package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type MstSpecialNeed struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstSpecialNeedDetail struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstSpecialNeedExport struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MstSpecialNeedSearch struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MstSpecialNeedRelation struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

/* Action */
func GetSpecialNeeds(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstSpecialNeed, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_special_needs_get", filter, sortBy, sortDirection, page, pageSize, &MstSpecialNeed{})
	if err != nil {
		return []MstSpecialNeed{}, 0, err
	}

	var modelResults []MstSpecialNeed
	for _, item := range results {
		level, ok := item.(*MstSpecialNeed)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchSpecialNeeds(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstSpecialNeedSearch, error) {
	var results []MstSpecialNeedSearch
	err := handlers.SPGet("sp_mst_special_needs_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstSpecialNeedSearch{}, err
	}
	return results, nil
}

func ExportSpecialNeeds(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstSpecialNeedExport
	err := handlers.SPGet("sp_mst_special_needs_get", "", "name", "asc", 1, CountSpecialNeeds(), &results)
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

func GetSpecialNeed(id string) (MstSpecialNeedDetail, error) {
	var result MstSpecialNeedDetail
	err := handlers.SPGetByID("sp_mst_special_needs_get_by_id", id, &result)
	if err != nil {
		return MstSpecialNeedDetail{}, err
	}
	return result, nil
}

func CreateSpecialNeed(id string, name string, description string) error {
	return QueryInsertSpecialNeed(id, name, description)
}

func ImportSpecialNeeds(fileStatus string) error {
	headers := []string{
		"id", "name", "description",
	}

	return handlers.ModelImport(fileStatus, headers, func(row map[string]string) error {
		id := row["id"]
		name := row["name"]
		description := row["description"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstSpecialNeed{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateSpecialNeed(id, name, description); err != nil {
					return err
				}
			} else {
				if err := QueryInsertSpecialNeed(id, name, description); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstSpecialNeed{})
			if err != nil {
				return err
			}
			if err := QueryInsertSpecialNeed(id, name, description); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateSpecialNeed(id string, name string, description string) error {
	return QueryUpdateSpecialNeed(id, name, description)
}

func DeleteSpecialNeed(id string) error {
	err := handlers.SPDelete("sp_mst_special_needs_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashSpecialNeeds(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstSpecialNeed, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_special_needs_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstSpecialNeed{})
	if err != nil {
		return []MstSpecialNeed{}, 0, err
	}

	var modelResults []MstSpecialNeed
	for _, item := range results {
		level, ok := item.(*MstSpecialNeed)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreSpecialNeed(id string) error {
	err := handlers.SPRestore("sp_mst_special_needs_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountSpecialNeeds(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstSpecialNeed{}, nullable)
}

/* Query */
func QueryInsertSpecialNeed(id string, name string, description string) error {
	query := `
		EXEC sp_mst_special_needs_insert
		@id = ?,
		@name = ?, 
		@description = ?, 
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`
	log.Print(query)

	err := handlers.SPInsert(query, id, name, description)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateSpecialNeed(id string, name string, description string) error {
	query := `
		EXEC sp_mst_special_needs_update
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

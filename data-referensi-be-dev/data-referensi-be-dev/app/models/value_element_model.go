package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstValueElement struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstValueElementDetail struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstValueElementExport struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MstValueElementSearch struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MstValueElementRelation struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

/* Action */
func GetValueElements(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstValueElement, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_value_elements_get", filter, sortBy, sortDirection, page, pageSize, &MstValueElement{})
	if err != nil {
		return []MstValueElement{}, 0, err
	}

	var modelResults []MstValueElement
	for _, item := range results {
		level, ok := item.(*MstValueElement)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchValueElements(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstValueElementSearch, error) {
	var results []MstValueElementSearch
	err := handlers.SPGet("sp_mst_value_elements_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstValueElementSearch{}, err
	}
	return results, nil
}

func ExportValueElements(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstValueElementExport
	err := handlers.SPGet("sp_mst_value_elements_get", "", "name", "asc", 1, CountValueElements(), &results)
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

func GetValueElement(id string) (MstValueElementDetail, error) {
	var result MstValueElementDetail
	err := handlers.SPGetByID("sp_mst_value_elements_get_by_id", id, &result)
	if err != nil {
		return MstValueElementDetail{}, err
	}
	return result, nil
}

func CreateValueElement(id string, name string, description string) error {
	return QueryInsertValueElement(id, name, description)
}

func ImportValueElements(fileStatus string) error {
	headers := []string{
		"id", "name", "description",
	}

	return handlers.ModelImport(fileStatus, headers, func(row map[string]string) error {
		id := row["id"]
		name := row["name"]
		description := row["description"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstValueElement{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateValueElement(id, name, description); err != nil {
					return err
				}
			} else {
				if err := QueryInsertValueElement(id, name, description); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstValueElement{})
			if err != nil {
				return err
			}
			if err := QueryInsertValueElement(id, name, description); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateValueElement(id string, name string, description string) error {
	return QueryUpdateValueElement(id, name, description)
}

func DeleteValueElement(id string) error {
	err := handlers.SPDelete("sp_mst_value_elements_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashValueElements(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstValueElement, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_value_elements_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstValueElement{})
	if err != nil {
		return []MstValueElement{}, 0, err
	}

	var modelResults []MstValueElement
	for _, item := range results {
		level, ok := item.(*MstValueElement)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreValueElement(id string) error {
	err := handlers.SPRestore("sp_mst_value_elements_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountValueElements(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstValueElement{}, nullable)
}

/* Query */
func QueryInsertValueElement(id string, name string, description string) error {
	query := `
		EXEC sp_mst_value_elements_insert
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

func QueryUpdateValueElement(id string, name string, description string) error {
	query := `
		EXEC sp_mst_value_elements_update
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

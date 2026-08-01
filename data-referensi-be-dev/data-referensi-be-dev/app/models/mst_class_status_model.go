package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type MstClassStatus struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstClassStatusDetail struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstClassStatusExport struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MstClassStatusSearch struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MstClassStatusRelation struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

/* Action */
func GetClassStatuses(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstClassStatus, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_class_statuses_get", filter, sortBy, sortDirection, page, pageSize, &MstClassStatus{})
	if err != nil {
		return []MstClassStatus{}, 0, err
	}

	var modelResults []MstClassStatus
	for _, item := range results {
		level, ok := item.(*MstClassStatus)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchClassStatuses(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstClassStatusSearch, error) {
	var results []MstClassStatusSearch
	err := handlers.SPGet("sp_mst_class_statuses_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstClassStatusSearch{}, err
	}
	return results, nil
}

func ExportClassStatuses(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstClassStatusExport
	err := handlers.SPGet("sp_mst_class_statuses_get", "", "name", "asc", 1, CountClassStatuses(), &results)
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

func GetClassStatus(id string) (MstClassStatusDetail, error) {
	var result MstClassStatusDetail
	err := handlers.SPGetByID("sp_mst_class_statuses_get_by_id", id, &result)
	if err != nil {
		return MstClassStatusDetail{}, err
	}
	return result, nil
}

func CreateClassStatus(id string, name string, description string) error {
	return QueryInsertClassStatus(id, name, description)
}

func ImportClassStatuses(fileStatus string) error {
	headers := []string{
		"id", "name", "description",
	}

	return handlers.ModelImport(fileStatus, headers, func(row map[string]string) error {
		id := row["id"]
		name := row["name"]
		description := row["description"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstClassStatus{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateClassStatus(id, name, description); err != nil {
					return err
				}
			} else {
				if err := QueryInsertClassStatus(id, name, description); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstClassStatus{})
			if err != nil {
				return err
			}
			if err := QueryInsertClassStatus(id, name, description); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateClassStatus(id string, name string, description string) error {
	return QueryUpdateClassStatus(id, name, description)
}

func DeleteClassStatus(id string) error {
	err := handlers.SPDelete("sp_mst_class_statuses_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashClassStatuses(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstClassStatus, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_class_statuses_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstClassStatus{})
	if err != nil {
		return []MstClassStatus{}, 0, err
	}

	var modelResults []MstClassStatus
	for _, item := range results {
		level, ok := item.(*MstClassStatus)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreClassStatus(id string) error {
	err := handlers.SPRestore("sp_mst_class_statuses_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountClassStatuses(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstClassStatus{}, nullable)
}

/* Query */
func QueryInsertClassStatus(id string, name string, description string) error {
	query := `
		EXEC sp_mst_class_statuses_insert
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

func QueryUpdateClassStatus(id string, name string, description string) error {
	query := `
		EXEC sp_mst_class_statuses_update
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

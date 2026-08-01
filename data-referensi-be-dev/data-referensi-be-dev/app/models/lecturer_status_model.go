package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type MstLecturerStatus struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstLecturerStatusDetail struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstLecturerStatusExport struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MstLecturerStatusSearch struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MstLecturerStatusRelation struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

/* Action */
func GetLecturerStatuses(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstLecturerStatus, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_lecturer_statuses_get", filter, sortBy, sortDirection, page, pageSize, &MstLecturerStatus{})
	if err != nil {
		return []MstLecturerStatus{}, 0, err
	}

	var modelResults []MstLecturerStatus
	for _, item := range results {
		level, ok := item.(*MstLecturerStatus)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchLecturerStatuses(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstLecturerStatusSearch, error) {
	var results []MstLecturerStatusSearch
	err := handlers.SPGet("sp_mst_lecturer_statuses_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstLecturerStatusSearch{}, err
	}
	return results, nil
}

func ExportLecturerStatuses(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstLecturerStatusExport
	err := handlers.SPGet("sp_mst_lecturer_statuses_get", "", "name", "asc", 1, CountLecturerStatuses(), &results)
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

func GetLecturerStatus(id string) (MstLecturerStatusDetail, error) {
	var result MstLecturerStatusDetail
	err := handlers.SPGetByID("sp_mst_lecturer_statuses_get_by_id", id, &result)
	if err != nil {
		return MstLecturerStatusDetail{}, err
	}
	return result, nil
}

func CreateLecturerStatus(id string, name string, description string) error {
	return QueryInsertLecturerStatus(id, name, description)
}

func ImportLecturerStatuses(fileStatus string) error {
	headers := []string{
		"id", "name", "description",
	}

	return handlers.ModelImport(fileStatus, headers, func(row map[string]string) error {
		id := row["id"]
		name := row["name"]
		description := row["description"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstLecturerStatus{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateLecturerStatus(id, name, description); err != nil {
					return err
				}
			} else {
				if err := QueryInsertLecturerStatus(id, name, description); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstLecturerStatus{})
			if err != nil {
				return err
			}
			if err := QueryInsertLecturerStatus(id, name, description); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateLecturerStatus(id string, name string, description string) error {
	return QueryUpdateLecturerStatus(id, name, description)
}

func DeleteLecturerStatus(id string) error {
	err := handlers.SPDelete("sp_mst_lecturer_statuses_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashLecturerStatuses(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstLecturerStatus, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_lecturer_statuses_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstLecturerStatus{})
	if err != nil {
		return []MstLecturerStatus{}, 0, err
	}

	var modelResults []MstLecturerStatus
	for _, item := range results {
		level, ok := item.(*MstLecturerStatus)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreLecturerStatus(id string) error {
	err := handlers.SPRestore("sp_mst_lecturer_statuses_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountLecturerStatuses(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstLecturerStatus{}, nullable)
}

/* Query */
func QueryInsertLecturerStatus(id string, name string, description string) error {
	query := `
		EXEC sp_mst_lecturer_statuses_insert
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

func QueryUpdateLecturerStatus(id string, name string, description string) error {
	query := `
		EXEC sp_mst_lecturer_statuses_update
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

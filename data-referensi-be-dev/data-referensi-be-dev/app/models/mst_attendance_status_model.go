package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type MstAttendanceStatus struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstAttendanceStatusDetail struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstAttendanceStatusExport struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MstAttendanceStatusSearch struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MstAttendanceStatusRelation struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

/* Action */
func GetAttendanceStatuses(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstAttendanceStatus, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_attendance_statuses_get", filter, sortBy, sortDirection, page, pageSize, &MstAttendanceStatus{})
	if err != nil {
		return []MstAttendanceStatus{}, 0, err
	}

	var modelResults []MstAttendanceStatus
	for _, item := range results {
		level, ok := item.(*MstAttendanceStatus)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchAttendanceStatuses(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstAttendanceStatusSearch, error) {
	var results []MstAttendanceStatusSearch
	err := handlers.SPGet("sp_mst_attendance_statuses_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstAttendanceStatusSearch{}, err
	}
	return results, nil
}

func ExportAttendanceStatuses(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstAttendanceStatusExport
	err := handlers.SPGet("sp_mst_attendance_statuses_get", "", "name", "asc", 1, CountAttendanceStatuses(), &results)
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

func GetAttendanceStatus(id string) (MstAttendanceStatusDetail, error) {
	var result MstAttendanceStatusDetail
	err := handlers.SPGetByID("sp_mst_attendance_statuses_get_by_id", id, &result)
	if err != nil {
		return MstAttendanceStatusDetail{}, err
	}
	return result, nil
}

func CreateAttendanceStatus(id string, name string, description string) error {
	return QueryInsertAttendanceStatus(id, name, description)
}

func ImportAttendanceStatuses(fileStatus string) error {
	headers := []string{
		"id", "name", "description",
	}

	return handlers.ModelImport(fileStatus, headers, func(row map[string]string) error {
		id := row["id"]
		name := row["name"]
		description := row["description"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstAttendanceStatus{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateAttendanceStatus(id, name, description); err != nil {
					return err
				}
			} else {
				if err := QueryInsertAttendanceStatus(id, name, description); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstAttendanceStatus{})
			if err != nil {
				return err
			}
			if err := QueryInsertAttendanceStatus(id, name, description); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateAttendanceStatus(id string, name string, description string) error {
	return QueryUpdateAttendanceStatus(id, name, description)
}

func DeleteAttendanceStatus(id string) error {
	err := handlers.SPDelete("sp_mst_attendance_statuses_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashAttendanceStatuses(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstAttendanceStatus, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_attendance_statuses_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstAttendanceStatus{})
	if err != nil {
		return []MstAttendanceStatus{}, 0, err
	}

	var modelResults []MstAttendanceStatus
	for _, item := range results {
		level, ok := item.(*MstAttendanceStatus)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreAttendanceStatus(id string) error {
	err := handlers.SPRestore("sp_mst_attendance_statuses_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountAttendanceStatuses(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstAttendanceStatus{}, nullable)
}

/* Query */
func QueryInsertAttendanceStatus(id string, name string, description string) error {
	query := `
		EXEC sp_mst_attendance_statuses_insert
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

func QueryUpdateAttendanceStatus(id string, name string, description string) error {
	query := `
		EXEC sp_mst_attendance_statuses_update
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

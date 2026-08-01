package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MstStudentStatus struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsDefault string `json:"is_default"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstStudentStatusDetail struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsDefault string `json:"is_default"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstStudentStatusExport struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsDefault bool   `json:"is_default"`
}

type MstStudentStatusSearch struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsDefault bool   `json:"is_default"`
}

type MstStudentStatusRelation struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsDefault bool   `json:"is_default"`
}

/* Action */
func GetStudentStatuses(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstStudentStatus, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_student_statuses_get", filter, sortBy, sortDirection, page, pageSize, &MstStudentStatus{})
	if err != nil {
		return []MstStudentStatus{}, 0, err
	}

	var modelResults []MstStudentStatus
	for _, item := range results {
		level, ok := item.(*MstStudentStatus)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchStudentStatuses(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstStudentStatusSearch, error) {
	var results []MstStudentStatusSearch
	err := handlers.SPGet("sp_mst_student_statuses_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstStudentStatusSearch{}, err
	}
	return results, nil
}

func ExportStudentStatuses(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstStudentStatusExport
	err := handlers.SPGet("sp_mst_student_statuses_get", "", "name", "asc", 1, CountStudentStatuses(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":        result.ID,
			"Name":      result.Name,
			"IsDefault": result.IsDefault,
		}
	}

	headers := []string{
		"ID", "Name", "Is Default",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetStudentStatus(id string) (MstStudentStatusDetail, error) {
	var result MstStudentStatusDetail
	err := handlers.SPGetByID("sp_mst_student_statuses_get_by_id", id, &result)
	if err != nil {
		return MstStudentStatusDetail{}, err
	}
	return result, nil
}

func CreateStudentStatus(id string, name string, is_default bool) error {
	return QueryInsertStudentStatus(id, name, is_default)
}

func ImportStudentStatuses(filePath string) error {
	headers := []string{
		"id", "name", "is_default",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		name := row["name"]
		is_default, err := strconv.ParseBool(row["is_default"])
		if err != nil {
			return err
		}

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstStudentStatus{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateStudentStatus(id, name, is_default); err != nil {
					return err
				}
			} else {
				if err := QueryInsertStudentStatus(id, name, is_default); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstStudentStatus{})
			if err != nil {
				return err
			}
			if err := QueryInsertStudentStatus(id, name, is_default); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateStudentStatus(id string, name string, is_default bool) error {
	return QueryUpdateStudentStatus(id, name, is_default)
}

func DeleteStudentStatus(id string) error {
	err := handlers.SPDelete("sp_mst_student_statuses_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashStudentStatuses(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstStudentStatus, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_student_statuses_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstStudentStatus{})
	if err != nil {
		return []MstStudentStatus{}, 0, err
	}

	var modelResults []MstStudentStatus
	for _, item := range results {
		level, ok := item.(*MstStudentStatus)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreStudentStatus(id string) error {
	err := handlers.SPRestore("sp_mst_student_statuses_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountStudentStatuses(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstStudentStatus{}, nullable)
}

/* Query */
func QueryInsertStudentStatus(id string, name string, is_default bool) error {
	query := `
		EXEC sp_mst_student_statuses_insert
		@id = ?,
		@name = ?,
		@is_default = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, name, is_default)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateStudentStatus(id string, name string, is_default bool) error {
	query := `
		EXEC sp_mst_student_statuses_update
		@id = ?,
		@name = ?,
		@is_default = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, name, is_default)
	if err != nil {
		return err
	}

	return nil
}

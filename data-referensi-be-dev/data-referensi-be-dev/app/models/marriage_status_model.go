package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstMarriageStatus struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstMarriageStatusDetail struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstMarriageStatusExport struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type MstMarriageStatusSearch struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type MstMarriageStatusRelation struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

/* Action */
func GetMarriageStatuses(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstMarriageStatus, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_marriage_statuses_get", filter, sortBy, sortDirection, page, pageSize, &MstMarriageStatus{})
	if err != nil {
		return []MstMarriageStatus{}, 0, err
	}

	var modelResults []MstMarriageStatus
	for _, item := range results {
		level, ok := item.(*MstMarriageStatus)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchMarriageStatuses(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstMarriageStatusSearch, error) {
	var results []MstMarriageStatusSearch
	err := handlers.SPGet("sp_mst_marriage_statuses_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstMarriageStatusSearch{}, err
	}
	return results, nil
}

func ExportMarriageStatuses(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstMarriageStatusExport
	err := handlers.SPGet("sp_mst_marriage_statuses_get", "", "name", "asc", 1, CountMarriageStatuses(), &results)
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
		"ID",
		"Name",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetMarriageStatus(id string) (MstMarriageStatusDetail, error) {
	var result MstMarriageStatusDetail
	err := handlers.SPGetByID("sp_mst_marriage_statuses_get_by_id", id, &result)
	if err != nil {
		return MstMarriageStatusDetail{}, err
	}
	return result, nil
}

func CreateMarriageStatus(id string, name string) error {
	return QueryInsertMarriageStatus(id, name)
}

func ImportMarriageStatuses(filePath string) error {
	headers := []string{
		"id", "name",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		name := row["name"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstMarriageStatus{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateMarriageStatus(id, name); err != nil {
					return err
				}
			} else {
				if err := QueryInsertMarriageStatus(id, name); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstMarriageStatus{})
			if err != nil {
				return err
			}
			if err := QueryInsertMarriageStatus(id, name); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateMarriageStatus(id string, name string) error {
	return QueryUpdateMarriageStatus(id, name)
}

func DeleteMarriageStatus(id string) error {
	err := handlers.SPDelete("sp_mst_marriage_statuses_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashMarriageStatuses(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstMarriageStatus, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_marriage_statuses_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstMarriageStatus{})
	if err != nil {
		return []MstMarriageStatus{}, 0, err
	}

	var modelResults []MstMarriageStatus
	for _, item := range results {
		level, ok := item.(*MstMarriageStatus)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreMarriageStatus(id string) error {
	err := handlers.SPRestore("sp_mst_marriage_statuses_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountMarriageStatuses(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstMarriageStatus{}, nullable)
}

/* Query */
func QueryInsertMarriageStatus(id string, name string) error {
	query := `
		EXEC sp_mst_marriage_statuses_insert
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

func QueryUpdateMarriageStatus(id string, name string) error {
	query := `
		EXEC sp_mst_marriage_statuses_update
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

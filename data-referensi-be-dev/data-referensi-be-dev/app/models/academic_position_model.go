package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstAcademicPosition struct {
	ID          string `json:"id"`
	Name        string `json:"name" `
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstAcademicPositionDetail struct {
	ID          string `json:"id"`
	Name        string `json:"name" `
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstAcademicPositionExport struct {
	ID          string `json:"id"`
	Name        string `json:"name" `
	Description string `json:"description"`
}

type MstAcademicPositionSearch struct {
	ID          string `json:"id"`
	Name        string `json:"name" `
	Description string `json:"description"`
}

type MstAcademicPositionRelation struct {
	ID   string `json:"id"`
	Name string `json:"name" `
}

/* Action */
func GetAcademicPositions(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstAcademicPosition, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_academic_positions_get", filter, sortBy, sortDirection, page, pageSize, &MstAcademicPosition{})
	if err != nil {
		return []MstAcademicPosition{}, 0, err
	}

	var modelResults []MstAcademicPosition
	for _, item := range results {
		level, ok := item.(*MstAcademicPosition)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchAcademicPositions(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstAcademicPositionSearch, error) {
	var results []MstAcademicPositionSearch
	err := handlers.SPGet("sp_mst_academic_positions_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstAcademicPositionSearch{}, err
	}
	return results, nil
}

func ExportAcademicPositions(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstAcademicPositionExport
	err := handlers.SPGet("sp_mst_academic_positions_get", "", "name", "asc", 1, CountAcademicPositions(), &results)
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

func GetAcademicPosition(id string) (MstAcademicPositionDetail, error) {
	var result MstAcademicPositionDetail
	err := handlers.SPGetByID("sp_mst_academic_positions_get_by_id", id, &result)
	if err != nil {
		return MstAcademicPositionDetail{}, err
	}
	return result, nil
}

func CreateAcademicPosition(id string, name string, description string) error {
	return QueryInsertAcademicPosition(id, name, description)
}

func ImportAcademicPositions(filePath string) error {
	headers := []string{
		"id", "name", "description",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		name := row["name"]
		description := row["description"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstAcademicPosition{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateAcademicPosition(id, name, description); err != nil {
					return err
				}
			} else {
				if err := QueryInsertAcademicPosition(id, name, description); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstAcademicPosition{})
			if err != nil {
				return err
			}
			if err := QueryInsertAcademicPosition(id, name, description); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateAcademicPosition(id string, name string, description string) error {
	return QueryUpdateAcademicPosition(id, name, description)
}

func DeleteAcademicPosition(id string) error {
	err := handlers.SPDelete("sp_mst_academic_positions_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashAcademicPositions(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstAcademicPosition, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_academic_positions_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstAcademicPosition{})
	if err != nil {
		return []MstAcademicPosition{}, 0, err
	}

	var modelResults []MstAcademicPosition
	for _, item := range results {
		level, ok := item.(*MstAcademicPosition)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreAcademicPosition(id string) error {
	err := handlers.SPRestore("sp_mst_academic_positions_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountAcademicPositions(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstAcademicPosition{}, nullable)
}

/* Query */
func QueryInsertAcademicPosition(id string, name string, description string) error {
	query := `
		EXEC sp_mst_academic_positions_insert
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

func QueryUpdateAcademicPosition(id string, name string, description string) error {
	query := `
		EXEC sp_mst_academic_positions_update
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

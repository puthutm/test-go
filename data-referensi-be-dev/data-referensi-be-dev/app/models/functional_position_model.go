package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstFunctionalPosition struct {
	ID          string `json:"id"`
	Name        string `json:"name" `
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstFunctionalPositionDetail struct {
	ID          string `json:"id"`
	Name        string `json:"name" `
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstFunctionalPositionExport struct {
	ID          string `json:"id"`
	Name        string `json:"name" `
	Description string `json:"description"`
}

type MstFunctionalPositionSearch struct {
	ID          string `json:"id"`
	Name        string `json:"name" `
	Description string `json:"description"`
}

type MstFunctionalPositionRelation struct {
	ID   string `json:"id"`
	Name string `json:"name" `
}

/* Action */
func GetFunctionalPositions(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstFunctionalPosition, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_functional_positions_get", filter, sortBy, sortDirection, page, pageSize, &MstFunctionalPosition{})
	if err != nil {
		return []MstFunctionalPosition{}, 0, err
	}

	var modelResults []MstFunctionalPosition
	for _, item := range results {
		level, ok := item.(*MstFunctionalPosition)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchFunctionalPositions(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstFunctionalPositionSearch, error) {
	var results []MstFunctionalPositionSearch
	err := handlers.SPGet("sp_mst_functional_positions_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstFunctionalPositionSearch{}, err
	}
	return results, nil
}

func ExportFunctionalPositions(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstFunctionalPositionExport
	err := handlers.SPGet("sp_mst_functional_positions_get", "", "name", "asc", 1, CountFunctionalPositions(), &results)
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

func GetFunctionalPosition(id string) (MstFunctionalPositionDetail, error) {
	var result MstFunctionalPositionDetail
	err := handlers.SPGetByID("sp_mst_functional_positions_get_by_id", id, &result)
	if err != nil {
		return MstFunctionalPositionDetail{}, err
	}
	return result, nil
}

func CreateFunctionalPosition(id string, name string, description string) error {
	return QueryInsertFunctionalPosition(id, name, description)
}

func ImportFunctionalPositions(filePath string) error {
	headers := []string{
		"id", "name", "description",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		name := row["name"]
		description := row["description"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstFunctionalPosition{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateFunctionalPosition(id, name, description); err != nil {
					return err
				}
			} else {
				if err := QueryInsertFunctionalPosition(id, name, description); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstFunctionalPosition{})
			if err != nil {
				return err
			}
			if err := QueryInsertFunctionalPosition(id, name, description); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateFunctionalPosition(id string, name string, description string) error {
	return QueryUpdateFunctionalPosition(id, name, description)
}

func DeleteFunctionalPosition(id string) error {
	err := handlers.SPDelete("sp_mst_functional_positions_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashFunctionalPositions(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstFunctionalPosition, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_functional_positions_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstFunctionalPosition{})
	if err != nil {
		return []MstFunctionalPosition{}, 0, err
	}

	var modelResults []MstFunctionalPosition
	for _, item := range results {
		level, ok := item.(*MstFunctionalPosition)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreFunctionalPosition(id string) error {
	err := handlers.SPRestore("sp_mst_functional_positions_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountFunctionalPositions(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstFunctionalPosition{}, nullable)
}

/* Query */
func QueryInsertFunctionalPosition(id string, name string, description string) error {
	query := `
		EXEC sp_mst_functional_positions_insert
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

func QueryUpdateFunctionalPosition(id string, name string, description string) error {
	query := `
		EXEC sp_mst_functional_positions_update
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

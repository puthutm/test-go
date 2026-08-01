package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type MstFinalAssignmentStage struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstFinalAssignmentStageDetail struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstFinalAssignmentStageExport struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MstFinalAssignmentStageSearch struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MstFinalAssignmentStageRelation struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

/* Action */
func GetFinalAssignmentStages(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstFinalAssignmentStage, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_final_assignment_stages_get", filter, sortBy, sortDirection, page, pageSize, &MstFinalAssignmentStage{})
	if err != nil {
		return []MstFinalAssignmentStage{}, 0, err
	}

	var modelResults []MstFinalAssignmentStage
	for _, item := range results {
		level, ok := item.(*MstFinalAssignmentStage)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchFinalAssignmentStages(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstFinalAssignmentStageSearch, error) {
	var results []MstFinalAssignmentStageSearch
	err := handlers.SPGet("sp_mst_final_assignment_stages_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstFinalAssignmentStageSearch{}, err
	}
	return results, nil
}

func ExportFinalAssignmentStages(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstFinalAssignmentStageExport
	err := handlers.SPGet("sp_mst_final_assignment_stages_get", "", "name", "asc", 1, CountFinalAssignmentStages(), &results)
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

func GetFinalAssignmentStage(id string) (MstFinalAssignmentStageDetail, error) {
	var result MstFinalAssignmentStageDetail
	err := handlers.SPGetByID("sp_mst_final_assignment_stages_get_by_id", id, &result)
	if err != nil {
		return MstFinalAssignmentStageDetail{}, err
	}
	return result, nil
}

func CreateFinalAssignmentStage(id string, name string, description string) error {
	return QueryInsertFinalAssignmentStage(id, name, description)
}

func ImportFinalAssignmentStages(fileStatus string) error {
	headers := []string{
		"id", "name", "description",
	}

	return handlers.ModelImport(fileStatus, headers, func(row map[string]string) error {
		id := row["id"]
		name := row["name"]
		description := row["description"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstFinalAssignmentStage{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateFinalAssignmentStage(id, name, description); err != nil {
					return err
				}
			} else {
				if err := QueryInsertFinalAssignmentStage(id, name, description); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstFinalAssignmentStage{})
			if err != nil {
				return err
			}
			if err := QueryInsertFinalAssignmentStage(id, name, description); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateFinalAssignmentStage(id string, name string, description string) error {
	return QueryUpdateFinalAssignmentStage(id, name, description)
}

func DeleteFinalAssignmentStage(id string) error {
	err := handlers.SPDelete("sp_mst_final_assignment_stages_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashFinalAssignmentStages(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstFinalAssignmentStage, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_final_assignment_stages_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstFinalAssignmentStage{})
	if err != nil {
		return []MstFinalAssignmentStage{}, 0, err
	}

	var modelResults []MstFinalAssignmentStage
	for _, item := range results {
		level, ok := item.(*MstFinalAssignmentStage)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreFinalAssignmentStage(id string) error {
	err := handlers.SPRestore("sp_mst_final_assignment_stages_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountFinalAssignmentStages(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstFinalAssignmentStage{}, nullable)
}

/* Query */
func QueryInsertFinalAssignmentStage(id string, name string, description string) error {
	query := `
		EXEC sp_mst_final_assignment_stages_insert
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

func QueryUpdateFinalAssignmentStage(id string, name string, description string) error {
	query := `
		EXEC sp_mst_final_assignment_stages_update
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

package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type MstApproval struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstApprovalDetail struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstApprovalExport struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MstApprovalSearch struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MstApprovalRelation struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

/* Action */
func GetApprovals(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstApproval, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_approvals_get", filter, sortBy, sortDirection, page, pageSize, &MstApproval{})
	if err != nil {
		return []MstApproval{}, 0, err
	}

	var modelResults []MstApproval
	for _, item := range results {
		level, ok := item.(*MstApproval)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchApprovals(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstApprovalSearch, error) {
	var results []MstApprovalSearch
	err := handlers.SPGet("sp_mst_approvals_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstApprovalSearch{}, err
	}
	return results, nil
}

func ExportApprovals(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstApprovalExport
	err := handlers.SPGet("sp_mst_approvals_get", "", "name", "asc", 1, CountApprovals(), &results)
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

func GetApproval(id string) (MstApprovalDetail, error) {
	var result MstApprovalDetail
	err := handlers.SPGetByID("sp_mst_approvals_get_by_id", id, &result)
	if err != nil {
		return MstApprovalDetail{}, err
	}
	return result, nil
}

func CreateApproval(id string, name string, description string) error {
	return QueryInsertApproval(id, name, description)
}

func ImportApprovals(fileStatus string) error {
	headers := []string{
		"id", "name", "description",
	}

	return handlers.ModelImport(fileStatus, headers, func(row map[string]string) error {
		id := row["id"]
		name := row["name"]
		description := row["description"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstApproval{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateApproval(id, name, description); err != nil {
					return err
				}
			} else {
				if err := QueryInsertApproval(id, name, description); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstApproval{})
			if err != nil {
				return err
			}
			if err := QueryInsertApproval(id, name, description); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateApproval(id string, name string, description string) error {
	return QueryUpdateApproval(id, name, description)
}

func DeleteApproval(id string) error {
	err := handlers.SPDelete("sp_mst_approvals_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashApprovals(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstApproval, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_approvals_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstApproval{})
	if err != nil {
		return []MstApproval{}, 0, err
	}

	var modelResults []MstApproval
	for _, item := range results {
		level, ok := item.(*MstApproval)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreApproval(id string) error {
	err := handlers.SPRestore("sp_mst_approvals_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountApprovals(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstApproval{}, nullable)
}

/* Query */
func QueryInsertApproval(id string, name string, description string) error {
	query := `
		EXEC sp_mst_approvals_insert
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

func QueryUpdateApproval(id string, name string, description string) error {
	query := `
		EXEC sp_mst_approvals_update
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

package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstReferral struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstReferralDetail struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstReferralExport struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type MstReferralSearch struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type MstReferralRelation struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

/* Action */
func GetReferrals(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstReferral, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_referrals_get", filter, sortBy, sortDirection, page, pageSize, &MstReferral{})
	if err != nil {
		return []MstReferral{}, 0, err
	}

	var modelResults []MstReferral
	for _, item := range results {
		level, ok := item.(*MstReferral)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchReferrals(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstReferralSearch, error) {
	var results []MstReferralSearch
	err := handlers.SPGet("sp_mst_referrals_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstReferralSearch{}, err
	}
	return results, nil
}

func ExportReferrals(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstReferralExport
	err := handlers.SPGet("sp_mst_referrals_get", "", "name", "asc", 1, CountReferrals(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":          result.ID,
			"Code":        result.Code,
			"Name":        result.Name,
			"Description": result.Description,
			"Status":      result.Status,
		}
	}

	headers := []string{
		"ID", "Code", "Name", "Description", "Status",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetReferral(id string) (MstReferralDetail, error) {
	var result MstReferralDetail
	err := handlers.SPGetByID("sp_mst_referrals_get_by_id", id, &result)
	if err != nil {
		return MstReferralDetail{}, err
	}
	return result, nil
}

func CreateReferral(id string, code string, name string, description string, status string) error {
	return QueryInsertReferral(id, code, name, description, status)
}

func ImportReferrals(fileStatus string) error {
	headers := []string{
		"id", "code", "name", "description", "status",
	}

	return handlers.ModelImport(fileStatus, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		name := row["name"]
		description := row["description"]
		status := row["status"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstReferral{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateReferral(id, code, name, description, status); err != nil {
					return err
				}
			} else {
				if err := QueryInsertReferral(id, code, name, description, status); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstReferral{})
			if err != nil {
				return err
			}
			if err := QueryInsertReferral(id, code, name, description, status); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateReferral(id string, code string, name string, description string, status string) error {
	return QueryUpdateReferral(id, code, name, description, status)
}

func DeleteReferral(id string) error {
	err := handlers.SPDelete("sp_mst_referrals_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashReferrals(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstReferral, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_referrals_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstReferral{})
	if err != nil {
		return []MstReferral{}, 0, err
	}

	var modelResults []MstReferral
	for _, item := range results {
		level, ok := item.(*MstReferral)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreReferral(id string) error {
	err := handlers.SPRestore("sp_mst_referrals_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountReferrals(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstReferral{}, nullable)
}

/* Query */
func QueryInsertReferral(id string, code string, name string, description string, status string) error {
	query := `
		EXEC sp_mst_referrals_insert
		@id = ?,
		@code = ?,
		@name = ?,
		@description = ?,
		@status = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, code, name, description, status)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateReferral(id string, code string, name string, description string, status string) error {
	query := `
		EXEC sp_mst_referrals_update
		@id = ?,
		@code = ?,
		@name = ?,
		@description = ?,
		@status = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, code, name, description, status)
	if err != nil {
		return err
	}

	return nil
}

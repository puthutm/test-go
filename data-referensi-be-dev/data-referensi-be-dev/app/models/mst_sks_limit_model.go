package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstSksLimit struct {
	ID        string   `json:"id"`
	IPMinimal float64  `json:"ip_minimal"`
	IPMaximal *float64 `json:"ip_maximal"`
	SKSLimit  *int     `json:"sks_limit"`
	CreatedAt int64    `json:"created_at"`
	UpdatedAt int64    `json:"updated_at"`
}

type MstSksLimitDetail struct {
	ID        string   `json:"id"`
	IPMinimal float64  `json:"ip_minimal"`
	IPMaximal *float64 `json:"ip_maximal"`
	SKSLimit  *int     `json:"sks_limit"`
	CreatedAt int64    `json:"created_at"`
	UpdatedAt int64    `json:"updated_at"`
}

type MstSksLimitExport struct {
	ID        string `json:"id"`
	IPMinimal string `json:"ip_minimal"`
	IPMaximal string `json:"ip_maximal"`
	SKSLimit  string `json:"sks_limit"`
}

type MstSksLimitSearch struct {
	ID        string   `json:"id"`
	IPMinimal float64  `json:"ip_minimal"`
	IPMaximal *float64 `json:"ip_maximal"`
	SKSLimit  *int     `json:"sks_limit"`
}

type MstSksLimitRelation struct {
	ID        string   `json:"id"`
	IPMinimal float64  `json:"ip_minimal"`
	IPMaximal *float64 `json:"ip_maximal"`
	SKSLimit  *int     `json:"sks_limit"`
}

/* Action */
func GetSksLimits(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstSksLimit, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_sks_limits_get", filter, sortBy, sortDirection, page, pageSize, &MstSksLimit{})
	if err != nil {
		return []MstSksLimit{}, 0, err
	}

	var modelResults []MstSksLimit
	for _, item := range results {
		level, ok := item.(*MstSksLimit)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchSksLimits(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstSksLimitSearch, error) {
	var results []MstSksLimitSearch
	err := handlers.SPGet("sp_mst_sks_limits_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstSksLimitSearch{}, err
	}
	return results, nil
}

func ExportSksLimits(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstSksLimitExport
	err := handlers.SPGet("sp_mst_sks_limits_get", "", "name", "asc", 1, CountSksLimits(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":        result.ID,
			"IPMinimal": result.IPMinimal,
			"IPMaximal": result.IPMaximal,
			"SKSLimit":  result.SKSLimit,
		}
	}

	headers := []string{
		"ID", "IPMinimal", "IPMaximal", "SKSLimit",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetSksLimit(id string) (MstSksLimitDetail, error) {
	var result MstSksLimitDetail
	err := handlers.SPGetByID("sp_mst_sks_limits_get_by_id", id, &result)
	if err != nil {
		return MstSksLimitDetail{}, err
	}
	return result, nil
}

func CreateSksLimit(id string, ip_minimal string, ip_maximal string, sks_limit string) error {
	return QueryInsertSksLimit(id, ip_minimal, ip_maximal, sks_limit)
}

func ImportSksLimits(fileStatus string) error {
	headers := []string{
		"id", "ip_minimal", "ip_maximal", "sks_limit",
	}

	return handlers.ModelImport(fileStatus, headers, func(row map[string]string) error {
		id := row["id"]
		ip_minimal := row["ip_minimal"]
		ip_maximal := row["ip_maximal"]
		sks_limit := row["sks_limit"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstSksLimit{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateSksLimit(id, ip_minimal, ip_maximal, sks_limit); err != nil {
					return err
				}
			} else {
				if err := QueryInsertSksLimit(id, ip_minimal, ip_maximal, sks_limit); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstSksLimit{})
			if err != nil {
				return err
			}
			if err := QueryInsertSksLimit(id, ip_minimal, ip_maximal, sks_limit); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateSksLimit(id string, ip_minimal string, ip_maximal string, sks_limit string) error {
	return QueryUpdateSksLimit(id, ip_minimal, ip_maximal, sks_limit)
}

func DeleteSksLimit(id string) error {
	err := handlers.SPDelete("sp_mst_sks_limits_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashSksLimits(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstSksLimit, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_sks_limits_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstSksLimit{})
	if err != nil {
		return []MstSksLimit{}, 0, err
	}

	var modelResults []MstSksLimit
	for _, item := range results {
		level, ok := item.(*MstSksLimit)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreSksLimit(id string) error {
	err := handlers.SPRestore("sp_mst_sks_limits_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountSksLimits(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstSksLimit{}, nullable)
}

/* Query */
func QueryInsertSksLimit(id string, ip_minimal string, ip_maximal string, sks_limit string) error {
	query := `
		EXEC sp_mst_sks_limits_insert
		@id = ?,
		@ip_minimal = ?, 
		@ip_maximal = ?, 
		@sks_limit = ?, 
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`
	err := handlers.SPInsert(query, id, ip_minimal, ip_maximal, sks_limit)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateSksLimit(id string, ip_minimal string, ip_maximal string, sks_limit string) error {
	query := `
		EXEC sp_mst_sks_limits_update
		@id = ?,
		@ip_minimal = ?, 
		@ip_maximal = ?, 
		@sks_limit = ?, 
		@updated_at = ?,
		@updated_by = ?
	`
	err := handlers.SPUpdate(query, id, ip_minimal, ip_maximal, sks_limit)
	if err != nil {
		return err
	}

	return nil
}

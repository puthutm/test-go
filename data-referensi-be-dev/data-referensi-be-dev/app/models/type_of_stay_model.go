package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstTypeOfStay struct {
	ID        string `json:"id"`
	Code      string `json:"code" `
	Type      string `json:"type"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstTypeOfStayDetail struct {
	ID        string `json:"id"`
	Code      string `json:"code" `
	Type      string `json:"type"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstTypeOfStayExport struct {
	ID   string `json:"id"`
	Code string `json:"code" `
	Type string `json:"type"`
}

type MstTypeOfStaySearch struct {
	ID   string `json:"id"`
	Code string `json:"code" `
	Type string `json:"type"`
}

type MstTypeOfStayRelation struct {
	ID   string `json:"id"`
	Code string `json:"code" `
}

/* Action */
func GetTypeOfStays(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstTypeOfStay, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_type_of_stays_get", filter, sortBy, sortDirection, page, pageSize, &MstTypeOfStay{})
	if err != nil {
		return []MstTypeOfStay{}, 0, err
	}

	var modelResults []MstTypeOfStay
	for _, item := range results {
		level, ok := item.(*MstTypeOfStay)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchTypeOfStays(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstTypeOfStaySearch, error) {
	var results []MstTypeOfStaySearch
	err := handlers.SPGet("sp_mst_type_of_stays_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstTypeOfStaySearch{}, err
	}
	return results, nil
}

func ExportTypeOfStays(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstTypeOfStayExport
	err := handlers.SPGet("sp_mst_type_of_stays_get", "", "code", "asc", 1, CountTypeOfStays(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":   result.ID,
			"Code": result.Code,
			"Type": result.Type,
		}
	}

	headers := []string{
		"ID", "Code", "Type",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetTypeOfStay(id string) (MstTypeOfStayDetail, error) {
	var result MstTypeOfStayDetail
	err := handlers.SPGetByID("sp_mst_type_of_stays_get_by_id", id, &result)
	if err != nil {
		return MstTypeOfStayDetail{}, err
	}
	return result, nil
}

func CreateTypeOfStay(id string, code string, type_of_stay string) error {
	return QueryInsertTypeOfStay(id, code, type_of_stay)
}

func ImportTypeOfStays(filePath string) error {
	headers := []string{
		"id", "code", "type_of_stay",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		type_of_stay := row["type_of_stay"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstTypeOfStay{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateTypeOfStay(id, code, type_of_stay); err != nil {
					return err
				}
			} else {
				if err := QueryInsertTypeOfStay(id, code, type_of_stay); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstTypeOfStay{})
			if err != nil {
				return err
			}
			if err := QueryInsertTypeOfStay(id, code, type_of_stay); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateTypeOfStay(id string, code string, type_of_stay string) error {
	return QueryUpdateTypeOfStay(id, code, type_of_stay)
}

func DeleteTypeOfStay(id string) error {
	err := handlers.SPDelete("sp_mst_type_of_stays_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashTypeOfStays(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstTypeOfStay, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_type_of_stays_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstTypeOfStay{})
	if err != nil {
		return []MstTypeOfStay{}, 0, err
	}

	var modelResults []MstTypeOfStay
	for _, item := range results {
		level, ok := item.(*MstTypeOfStay)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreTypeOfStay(id string) error {
	err := handlers.SPRestore("sp_mst_type_of_stays_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountTypeOfStays(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstTypeOfStay{}, nullable)
}

/* Query */
func QueryInsertTypeOfStay(id string, code string, type_of_stay string) error {
	query := `
		EXEC sp_mst_type_of_stays_insert
		@id = ?,
		@code = ?,
		@type = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, code, type_of_stay)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateTypeOfStay(id string, code string, type_of_stay string) error {
	query := `
		EXEC sp_mst_type_of_stays_update
		@id = ?,
		@code = ?,
		@type = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, code, type_of_stay)
	if err != nil {
		return err
	}

	return nil
}

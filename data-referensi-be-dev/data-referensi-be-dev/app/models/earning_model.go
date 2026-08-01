package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstEarning struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Range     string `json:"range" gorm:"column:range"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstEarningDetail struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Range     string `json:"range"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstEarningExport struct {
	ID    string `json:"id"`
	Code  string `json:"code"`
	Name  string `json:"name"`
	Range string `json:"range"`
}

type MstEarningSearch struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type MstEarningRelation struct {
	ID    string `json:"id"`
	Code  string `json:"code"`
	Name  string `json:"name"`
	Range string `json:"range"`
}

/* Action */
func GetEarnings(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstEarning, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_earnings_get", filter, sortBy, sortDirection, page, pageSize, &MstEarning{})
	if err != nil {
		return []MstEarning{}, 0, err
	}

	var modelResults []MstEarning
	for _, item := range results {
		level, ok := item.(*MstEarning)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchEarnings(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstEarningSearch, error) {
	var results []MstEarningSearch
	err := handlers.SPGet("sp_mst_earnings_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstEarningSearch{}, err
	}
	return results, nil
}

func ExportEarnings(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstEarningExport
	err := handlers.SPGet("sp_mst_earnings_get", "", "name", "asc", 1, CountEarnings(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":    result.ID,
			"Code":  result.Code,
			"Name":  result.Name,
			"Range": result.Range,
		}
	}

	headers := []string{
		"ID",
		"Code",
		"Name",
		"Range",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetEarning(id string) (MstEarningDetail, error) {
	var result MstEarningDetail
	err := handlers.SPGetByID("sp_mst_earnings_get_by_id", id, &result)
	if err != nil {
		return MstEarningDetail{}, err
	}
	return result, nil
}

func CreateEarning(id string, code string, name string, range_earning string) error {
	return QueryInsertEarning(id, code, name, range_earning)
}

func ImportEarnings(filePath string) error {
	headers := []string{
		"id", "code", "name", "range",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		name := row["name"]
		range_earning := row["range"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstEarning{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateEarning(id, code, name, range_earning); err != nil {
					return err
				}
			} else {
				if err := QueryInsertEarning(id, code, name, range_earning); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstEarning{})
			if err != nil {
				return err
			}
			if err := QueryInsertEarning(id, code, name, range_earning); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateEarning(id string, code string, name string, range_earning string) error {
	return QueryUpdateEarning(id, code, name, range_earning)
}

func DeleteEarning(id string) error {
	err := handlers.SPDelete("sp_mst_earnings_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashEarnings(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstEarning, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_earnings_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstEarning{})
	if err != nil {
		return []MstEarning{}, 0, err
	}

	var modelResults []MstEarning
	for _, item := range results {
		level, ok := item.(*MstEarning)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreEarning(id string) error {
	err := handlers.SPRestore("sp_mst_earnings_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountEarnings(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstEarning{}, nullable)
}

/* Query */
func QueryInsertEarning(id string, code string, name string, range_earning string) error {
	query := `
		EXEC sp_mst_earnings_insert
		@id = ?,
		@code= ?,
		@name = ?,
		@range = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, code, name, range_earning)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateEarning(id string, code string, name string, range_earning string) error {
	query := `
		EXEC sp_mst_earnings_update
		@id = ?,
		@code= ?,
		@name = ?,
		@range = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, code, name, range_earning)
	if err != nil {
		return err
	}

	return nil
}

package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstBank struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstBankDetail struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstBankExport struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type MstBankSearch struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}
type MstBankRelation struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

/* Action */
func GetBanks(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstBank, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_banks_get", filter, sortBy, sortDirection, page, pageSize, &MstBank{})
	if err != nil {
		return []MstBank{}, 0, err
	}

	var modelResults []MstBank
	for _, item := range results {
		level, ok := item.(*MstBank)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchBanks(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstBankSearch, error) {
	var results []MstBankSearch
	err := handlers.SPGet("sp_mst_banks_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstBankSearch{}, err
	}
	return results, nil
}

func ExportBanks(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstBankExport
	err := handlers.SPGet("sp_mst_banks_get", "", "name", "asc", 1, CountBanks(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":   result.ID,
			"Code": result.Code,
			"Name": result.Name,
		}
	}

	headers := []string{
		"ID",
		"Code",
		"Name",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetBank(id string) (MstBankDetail, error) {
	var result MstBankDetail
	err := handlers.SPGetByID("sp_mst_banks_get_by_id", id, &result)
	if err != nil {
		return MstBankDetail{}, err
	}
	return result, nil
}

func CreateBank(id string, code string, name string) error {
	return QueryInsertBank(id, code, name)
}

func ImportBanks(filePath string) error {
	headers := []string{
		"id", "code", "name",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		name := row["name"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstBank{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateBank(id, code, name); err != nil {
					return err
				}
			} else {
				if err := QueryInsertBank(id, code, name); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstBank{})
			if err != nil {
				return err
			}
			if err := QueryInsertBank(id, code, name); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateBank(id string, code string, name string) error {
	return QueryUpdateBank(id, code, name)
}

func DeleteBank(id string) error {
	err := handlers.SPDelete("sp_mst_banks_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashBanks(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstBank, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_banks_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstBank{})
	if err != nil {
		return []MstBank{}, 0, err
	}

	var modelResults []MstBank
	for _, item := range results {
		level, ok := item.(*MstBank)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreBank(id string) error {
	err := handlers.SPRestore("sp_mst_banks_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountBanks(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstBank{}, nullable)
}

/* Query */
func QueryInsertBank(id string, code string, name string) error {
	query := `
		EXEC sp_mst_banks_insert
		@id = ?,
		@code = ?,
		@name = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, code, name)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateBank(id string, code string, name string) error {
	query := `
		EXEC sp_mst_banks_update
		@id = ?,
		@code = ?,
		@name = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, code, name)
	if err != nil {
		return err
	}

	return nil
}

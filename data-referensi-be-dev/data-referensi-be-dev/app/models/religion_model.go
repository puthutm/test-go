package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstReligion struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstReligionDetail struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstReligionExport struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type MstReligionSearch struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type MstReligionRelation struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

/* Action */
func GetReligions(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstReligion, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_religions_get", filter, sortBy, sortDirection, page, pageSize, &MstReligion{})
	if err != nil {
		return []MstReligion{}, 0, err
	}

	var modelResults []MstReligion
	for _, item := range results {
		level, ok := item.(*MstReligion)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchReligions(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstReligionSearch, error) {
	var results []MstReligionSearch
	err := handlers.SPGet("sp_mst_religions_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstReligionSearch{}, err
	}
	return results, nil
}

func ExportReligions(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstReligionExport
	err := handlers.SPGet("sp_mst_religions_get", "", "name", "asc", 1, CountReligions(), &results)
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

func GetReligion(id string) (MstReligionDetail, error) {
	var result MstReligionDetail
	err := handlers.SPGetByID("sp_mst_religions_get_by_id", id, &result)
	if err != nil {
		return MstReligionDetail{}, err
	}
	return result, nil
}

func CreateReligion(id string, code string, name string) error {
	return QueryInsertReligion(id, code, name)
}

func ImportReligions(filePath string) error {
	headers := []string{
		"id", "code", "name",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		name := row["name"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstReligion{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateReligion(id, code, name); err != nil {
					return err
				}
			} else {
				if err := QueryInsertReligion(id, code, name); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstReligion{})
			if err != nil {
				return err
			}
			if err := QueryInsertReligion(id, code, name); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateReligion(id string, code string, name string) error {
	return QueryUpdateReligion(id, code, name)
}

func DeleteReligion(id string) error {
	err := handlers.SPDelete("sp_mst_religions_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashReligions(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstReligion, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_religions_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstReligion{})
	if err != nil {
		return []MstReligion{}, 0, err
	}

	var modelResults []MstReligion
	for _, item := range results {
		level, ok := item.(*MstReligion)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreReligion(id string) error {
	err := handlers.SPRestore("sp_mst_religions_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountReligions(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstReligion{}, nullable)
}

/* Query */
func QueryInsertReligion(id string, code string, name string) error {
	query := `
		EXEC sp_mst_religions_insert
		@id = ?,
		@code= ?,
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

func QueryUpdateReligion(id string, code string, name string) error {
	query := `
		EXEC sp_mst_religions_update
		@id = ?,
		@code= ?,
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

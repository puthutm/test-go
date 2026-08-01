package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstTransportation struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstTransportationDetail struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstTransportationExport struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type MstTransportationSearch struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type MstTransportationRelation struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

/* Action */
func GetTransportations(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstTransportation, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_transportations_get", filter, sortBy, sortDirection, page, pageSize, &MstTransportation{})
	if err != nil {
		return []MstTransportation{}, 0, err
	}

	var modelResults []MstTransportation
	for _, item := range results {
		level, ok := item.(*MstTransportation)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchTransportations(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstTransportationSearch, error) {
	var results []MstTransportationSearch
	err := handlers.SPGet("sp_mst_transportations_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstTransportationSearch{}, err
	}
	return results, nil
}

func ExportTransportations(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstTransportationExport
	err := handlers.SPGet("sp_mst_transportations_get", "", "name", "asc", 1, CountTransportations(), &results)
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
		"ID", "Code", "Name",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetTransportation(id string) (MstTransportationDetail, error) {
	var result MstTransportationDetail
	err := handlers.SPGetByID("sp_mst_transportations_get_by_id", id, &result)
	if err != nil {
		return MstTransportationDetail{}, err
	}
	return result, nil
}

func CreateTransportation(id string, code string, name string) error {
	return QueryInsertTransportation(id, code, name)
}

func ImportTransportations(filePath string) error {
	headers := []string{
		"id", "code", "name",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		name := row["name"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstTransportation{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateTransportation(id, code, name); err != nil {
					return err
				}
			} else {
				if err := QueryInsertTransportation(id, code, name); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstTransportation{})
			if err != nil {
				return err
			}
			if err := QueryInsertTransportation(id, code, name); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateTransportation(id string, code string, name string) error {
	return QueryUpdateTransportation(id, code, name)
}

func DeleteTransportation(id string) error {
	err := handlers.SPDelete("sp_mst_transportations_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashTransportations(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstTransportation, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_transportations_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstTransportation{})
	if err != nil {
		return []MstTransportation{}, 0, err
	}

	var modelResults []MstTransportation
	for _, item := range results {
		level, ok := item.(*MstTransportation)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreTransportation(id string) error {
	err := handlers.SPRestore("sp_mst_transportations_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountTransportations(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstTransportation{}, nullable)
}

/* Query */
func QueryInsertTransportation(id string, code string, name string) error {
	query := `
		EXEC sp_mst_transportations_insert
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

func QueryUpdateTransportation(id string, code string, name string) error {
	query := `
		EXEC sp_mst_transportations_update
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

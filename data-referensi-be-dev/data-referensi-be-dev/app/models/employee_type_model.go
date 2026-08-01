package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstEmployeeType struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstEmployeeTypeDetail struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstEmployeeTypeExport struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type MstEmployeeTypeSearch struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type MstEmployeeTypeRelation struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

/* Action */
func GetEmployeeTypes(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstEmployeeType, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_employee_types_get", filter, sortBy, sortDirection, page, pageSize, &MstEmployeeType{})
	if err != nil {
		return []MstEmployeeType{}, 0, err
	}

	var modelResults []MstEmployeeType
	for _, item := range results {
		level, ok := item.(*MstEmployeeType)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchEmployeeTypes(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstEmployeeTypeSearch, error) {
	var results []MstEmployeeTypeSearch
	err := handlers.SPGet("sp_mst_employee_types_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstEmployeeTypeSearch{}, err
	}
	return results, nil
}

func ExportEmployeeTypes(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstEmployeeTypeExport
	err := handlers.SPGet("sp_mst_employee_types_get", "", "name", "asc", 1, CountEmployeeTypes(), &results)
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

func GetEmployeeType(id string) (MstEmployeeTypeDetail, error) {
	var result MstEmployeeTypeDetail
	err := handlers.SPGetByID("sp_mst_employee_types_get_by_id", id, &result)
	if err != nil {
		return MstEmployeeTypeDetail{}, err
	}
	return result, nil
}

func CreateEmployeeType(id string, code string, name string) error {
	return QueryInsertEmployeeType(id, code, name)
}

func ImportEmployeeTypes(filePath string) error {
	headers := []string{
		"id", "code", "name",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		name := row["name"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstEmployeeType{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateEmployeeType(id, code, name); err != nil {
					return err
				}
			} else {
				if err := QueryInsertEmployeeType(id, code, name); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstEmployeeType{})
			if err != nil {
				return err
			}
			if err := QueryInsertEmployeeType(id, code, name); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateEmployeeType(id string, code string, name string) error {
	return QueryUpdateEmployeeType(id, code, name)
}

func DeleteEmployeeType(id string) error {
	err := handlers.SPDelete("sp_mst_employee_types_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashEmployeeTypes(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstEmployeeType, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_employee_types_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstEmployeeType{})
	if err != nil {
		return []MstEmployeeType{}, 0, err
	}

	var modelResults []MstEmployeeType
	for _, item := range results {
		level, ok := item.(*MstEmployeeType)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreEmployeeType(id string) error {
	err := handlers.SPRestore("sp_mst_employee_types_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountEmployeeTypes(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstEmployeeType{}, nullable)
}

/* Query */
func QueryInsertEmployeeType(id string, code string, name string) error {
	query := `
		EXEC sp_mst_employee_types_insert
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

func QueryUpdateEmployeeType(id string, code string, name string) error {
	query := `
		EXEC sp_mst_employee_types_update
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

package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstEmployeeStatus struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstEmployeeStatusDetail struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstEmployeeStatusExport struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type MstEmployeeStatusSearch struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type MstEmployeeStatusRelation struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

/* Action */
func GetEmployeeStatuses(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstEmployeeStatus, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_employee_statuses_get", filter, sortBy, sortDirection, page, pageSize, &MstEmployeeStatus{})
	if err != nil {
		return []MstEmployeeStatus{}, 0, err
	}

	var modelResults []MstEmployeeStatus
	for _, item := range results {
		level, ok := item.(*MstEmployeeStatus)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchEmployeeStatuses(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstEmployeeStatusSearch, error) {
	var results []MstEmployeeStatusSearch
	err := handlers.SPGet("sp_mst_employee_statuses_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstEmployeeStatusSearch{}, err
	}
	return results, nil
}

func ExportEmployeeStatuses(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstEmployeeStatusExport
	err := handlers.SPGet("sp_mst_employee_statuses_get", "", "name", "asc", 1, CountEmployeeStatuses(), &results)
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

func GetEmployeeStatus(id string) (MstEmployeeStatusDetail, error) {
	var result MstEmployeeStatusDetail
	err := handlers.SPGetByID("sp_mst_employee_statuses_get_by_id", id, &result)
	if err != nil {
		return MstEmployeeStatusDetail{}, err
	}
	return result, nil
}

func CreateEmployeeStatus(id string, code string, name string) error {
	return QueryInsertEmployeeStatus(id, code, name)
}

func ImportEmployeeStatuses(filePath string) error {
	headers := []string{
		"id", "code", "name",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		name := row["name"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstEmployeeStatus{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateEmployeeStatus(id, code, name); err != nil {
					return err
				}
			} else {
				if err := QueryInsertEmployeeStatus(id, code, name); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstEmployeeStatus{})
			if err != nil {
				return err
			}
			if err := QueryInsertEmployeeStatus(id, code, name); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateEmployeeStatus(id string, code string, name string) error {
	return QueryUpdateEmployeeStatus(id, code, name)
}

func DeleteEmployeeStatus(id string) error {
	err := handlers.SPDelete("sp_mst_employee_statuses_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashEmployeeStatuses(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstEmployeeStatus, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_employee_statuses_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstEmployeeStatus{})
	if err != nil {
		return []MstEmployeeStatus{}, 0, err
	}

	var modelResults []MstEmployeeStatus
	for _, item := range results {
		level, ok := item.(*MstEmployeeStatus)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreEmployeeStatus(id string) error {
	err := handlers.SPRestore("sp_mst_employee_statuses_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountEmployeeStatuses(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstEmployeeStatus{}, nullable)
}

/* Query */
func QueryInsertEmployeeStatus(id string, code string, name string) error {
	query := `
		EXEC sp_mst_employee_statuses_insert
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

func QueryUpdateEmployeeStatus(id string, code string, name string) error {
	query := `
		EXEC sp_mst_employee_statuses_update
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

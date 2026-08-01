package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstTypeOfCondition struct {
	ID              string `json:"id"`
	Code            string `json:"code" `
	TypeOfCondition string `json:"type_of_condition"`
	Note            string `json:"note"`
	CreatedAt       int64  `json:"created_at"`
	UpdatedAt       int64  `json:"updated_at"`
}

type MstTypeOfConditionDetail struct {
	ID              string `json:"id"`
	Code            string `json:"code" `
	TypeOfCondition string `json:"type_of_condition"`
	Note            string `json:"note"`
	CreatedAt       int64  `json:"created_at"`
	UpdatedAt       int64  `json:"updated_at"`
}

type MstTypeOfConditionExport struct {
	ID              string `json:"id"`
	Code            string `json:"code" `
	TypeOfCondition string `json:"type_of_condition"`
	Note            string `json:"note"`
}

type MstTypeOfConditionSearch struct {
	ID              string `json:"id"`
	Code            string `json:"code" `
	TypeOfCondition string `json:"type_of_condition"`
	Note            string `json:"note"`
}

type MstTypeOfConditionRelation struct {
	ID              string `json:"id"`
	Code            string `json:"code" `
	TypeOfCondition string `json:"type_of_condition"`
	Note            string `json:"note"`
}

/* Action */
func GetTypeOfConditions(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstTypeOfCondition, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_type_of_conditions_get", filter, sortBy, sortDirection, page, pageSize, &MstTypeOfCondition{})
	if err != nil {
		return []MstTypeOfCondition{}, 0, err
	}

	var modelResults []MstTypeOfCondition
	for _, item := range results {
		level, ok := item.(*MstTypeOfCondition)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchTypeOfConditions(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstTypeOfConditionSearch, error) {
	var results []MstTypeOfConditionSearch
	err := handlers.SPGet("sp_mst_type_of_conditions_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstTypeOfConditionSearch{}, err
	}
	return results, nil
}

func ExportTypeOfConditions(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstTypeOfConditionExport
	err := handlers.SPGet("sp_mst_type_of_conditions_get", "", "code", "asc", 1, CountTypeOfConditions(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":              result.ID,
			"Code":            result.Code,
			"TypeOfCondition": result.TypeOfCondition,
			"Note":            result.Note,
		}
	}

	headers := []string{
		"ID", "Code", "Type Of Condition", "Note",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetTypeOfCondition(id string) (MstTypeOfConditionDetail, error) {
	var result MstTypeOfConditionDetail
	err := handlers.SPGetByID("sp_mst_type_of_conditions_get_by_id", id, &result)
	if err != nil {
		return MstTypeOfConditionDetail{}, err
	}
	return result, nil
}

func CreateTypeOfCondition(id string, code string, type_of_condition string, note string) error {
	return QueryInsertTypeOfCondition(id, code, type_of_condition, note)
}

func ImportTypeOfConditions(filePath string) error {
	headers := []string{
		"id", "code", "type_of_condition", "note",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		type_of_condition := row["type_of_condition"]
		note := row["note"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstTypeOfCondition{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateTypeOfCondition(id, code, type_of_condition, note); err != nil {
					return err
				}
			} else {
				if err := QueryInsertTypeOfCondition(id, code, type_of_condition, note); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstTypeOfCondition{})
			if err != nil {
				return err
			}
			if err := QueryInsertTypeOfCondition(id, code, type_of_condition, note); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateTypeOfCondition(id string, code string, type_of_condition string, note string) error {
	return QueryUpdateTypeOfCondition(id, code, type_of_condition, note)
}

func DeleteTypeOfCondition(id string) error {
	err := handlers.SPDelete("sp_mst_type_of_conditions_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashTypeOfConditions(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstTypeOfCondition, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_type_of_conditions_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstTypeOfCondition{})
	if err != nil {
		return []MstTypeOfCondition{}, 0, err
	}

	var modelResults []MstTypeOfCondition
	for _, item := range results {
		level, ok := item.(*MstTypeOfCondition)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreTypeOfCondition(id string) error {
	err := handlers.SPRestore("sp_mst_type_of_conditions_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountTypeOfConditions(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstTypeOfCondition{}, nullable)
}

/* Query */
func QueryInsertTypeOfCondition(id string, code string, type_of_condition string, note string) error {
	query := `
		EXEC sp_mst_type_of_conditions_insert
		@id = ?,
		@code = ?,
		@type_of_condition = ?,
		@note = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, code, type_of_condition, note)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateTypeOfCondition(id string, code string, type_of_condition string, note string) error {
	query := `
		EXEC sp_mst_type_of_conditions_update
		@id = ?,
		@code = ?,
		@type_of_condition = ?,
		@note = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, code, type_of_condition, note)
	if err != nil {
		return err
	}

	return nil
}

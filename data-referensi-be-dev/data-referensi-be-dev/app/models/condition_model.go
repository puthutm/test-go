package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MstCondition struct {
	ID        string `json:"id"`
	Code      string `json:"code" `
	Name      string `json:"name"`
	Point     string `json:"point"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstConditionDetail struct {
	ID        string `json:"id"`
	Code      string `json:"code" `
	Name      string `json:"name"`
	Point     string `json:"point"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstConditionExport struct {
	ID    string  `json:"id"`
	Code  string  `json:"code" `
	Name  string  `json:"name"`
	Point float64 `json:"point"`
}

type MstConditionSearch struct {
	ID    string  `json:"id"`
	Code  string  `json:"code" `
	Name  string  `json:"name"`
	Point float64 `json:"point"`
}

type MstConditionRelation struct {
	ID    string  `json:"id"`
	Code  string  `json:"code" `
	Name  string  `json:"name"`
	Point float64 `json:"point"`
}

/* Action */
func GetConditions(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstCondition, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_conditions_get", filter, sortBy, sortDirection, page, pageSize, &MstCondition{})
	if err != nil {
		return []MstCondition{}, 0, err
	}

	var modelResults []MstCondition
	for _, item := range results {
		level, ok := item.(*MstCondition)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchConditions(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstConditionSearch, error) {
	var results []MstConditionSearch
	err := handlers.SPGet("sp_mst_conditions_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstConditionSearch{}, err
	}
	return results, nil
}

func ExportConditions(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstConditionExport
	err := handlers.SPGet("sp_mst_conditions_get", "", "code", "asc", 1, CountConditions(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":    result.ID,
			"Code":  result.Code,
			"Name":  result.Name,
			"Point": result.Point,
		}
	}

	headers := []string{
		"ID", "Code", "Name", "Point",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetCondition(id string) (MstConditionDetail, error) {
	var result MstConditionDetail
	err := handlers.SPGetByID("sp_mst_conditions_get_by_id", id, &result)
	if err != nil {
		return MstConditionDetail{}, err
	}
	return result, nil
}

func CreateCondition(id string, code string, name string, point float64) error {
	return QueryInsertCondition(id, code, name, point)
}

func ImportConditions(filePath string) error {
	headers := []string{
		"id", "code", "name", "point",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		name := row["name"]
		point := row["point"]

		floatPoint, err := strconv.ParseFloat(point, 64)
		if err != nil {
			log.Fatalf("Gagal mengonversi string ke float64: %v", err)
		}

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstCondition{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateCondition(id, code, name, floatPoint); err != nil {
					return err
				}
			} else {
				if err := QueryInsertCondition(id, code, name, floatPoint); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstCondition{})
			if err != nil {
				return err
			}
			if err := QueryInsertCondition(id, code, name, floatPoint); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateCondition(id string, code string, name string, point float64) error {
	return QueryUpdateCondition(id, code, name, point)
}

func DeleteCondition(id string) error {
	err := handlers.SPDelete("sp_mst_conditions_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashConditions(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstCondition, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_conditions_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstCondition{})
	if err != nil {
		return []MstCondition{}, 0, err
	}

	var modelResults []MstCondition
	for _, item := range results {
		level, ok := item.(*MstCondition)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreCondition(id string) error {
	err := handlers.SPRestore("sp_mst_conditions_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountConditions(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstCondition{}, nullable)
}

/* Query */
func QueryInsertCondition(id string, code string, name string, point float64) error {
	query := `
		EXEC sp_mst_conditions_insert
		@id = ?,
		@code = ?,
		@name = ?,
		@point = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, code, name, point)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateCondition(id string, code string, name string, point float64) error {
	query := `
		EXEC sp_mst_conditions_update
		@id = ?,
		@code = ?,
		@name = ?,
		@point = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, code, name, point)
	if err != nil {
		return err
	}

	return nil
}

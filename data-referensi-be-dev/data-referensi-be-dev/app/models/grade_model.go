package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MstGrade struct {
	ID          string     `json:"id"`
	Code        string     `json:"code"`
	Name        string     `json:"name"`
	LowerLimit  *float64   `json:"lower_limit"`
	UpperLimit  *float64   `json:"upper_limit"`
	Description *string    `json:"description"`
	CreatedAt   *int64     `json:"created_at"`
	CreatedBy   *uuid.UUID `json:"created_by"`
	UpdatedAt   *int64     `json:"updated_at"`
	UpdatedBy   *uuid.UUID `json:"updated_by"`
	DeletedAt   *int64     `json:"deleted_at"`
	DeletedBy   *uuid.UUID `json:"deleted_by"`
}

type MstGradeDetail struct {
	ID          string     `json:"id"`
	Code        string     `json:"code"`
	Name        string     `json:"name"`
	LowerLimit  *float64   `json:"lower_limit"`
	UpperLimit  *float64   `json:"upper_limit"`
	Description *string    `json:"description"`
	CreatedAt   *int64     `json:"created_at"`
	CreatedBy   *uuid.UUID `json:"created_by"`
	UpdatedAt   *int64     `json:"updated_at"`
	UpdatedBy   *uuid.UUID `json:"updated_by"`
	DeletedAt   *int64     `json:"deleted_at"`
	DeletedBy   *uuid.UUID `json:"deleted_by"`
}

type MstGradeExport struct {
	ID          string   `json:"id"`
	Code        string   `json:"code"`
	Name        string   `json:"name"`
	LowerLimit  *float64 `json:"lower_limit"`
	UpperLimit  *float64 `json:"upper_limit"`
	Description *string  `json:"description"`
}

type MstGradeSearch struct {
	ID         string   `json:"id"`
	Code       string   `json:"code"`
	Name       string   `json:"name"`
	LowerLimit *float64 `json:"lower_limit"`
	UpperLimit *float64 `json:"upper_limit"`
}

type MstGradeRelation struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

/* Action */
func GetGrades(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstGrade, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_grades_get", filter, sortBy, sortDirection, page, pageSize, &MstGrade{})
	if err != nil {
		return []MstGrade{}, 0, err
	}

	var modelResults []MstGrade
	for _, item := range results {
		level, ok := item.(*MstGrade)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchGrades(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstGradeSearch, error) {
	var results []MstGradeSearch
	err := handlers.SPGet("sp_mst_grades_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstGradeSearch{}, err
	}
	return results, nil
}

func ExportGrades(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstGradeExport
	err := handlers.SPGet("sp_mst_grades_get", "", "name", "asc", 1, CountGrades(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":          result.ID,
			"Code":        result.Code,
			"Name":        result.Name,
			"LowerLimit":  result.LowerLimit,
			"UpperLimit":  result.UpperLimit,
			"Description": result.Description,
		}
	}

	headers := []string{
		"ID", "Code", "Name", "Lower Limit", "Upper Limit", "Description",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetGrade(id string) (MstGradeDetail, error) {
	var result MstGradeDetail
	err := handlers.SPGetByID("sp_mst_grades_get_by_id", id, &result)
	if err != nil {
		return MstGradeDetail{}, err
	}
	return result, nil
}

func CreateGrade(id string, code string, name string, lowerLimit *float64, upperLimit *float64, description *string) error {
	return QueryInsertGrade(id, code, name, lowerLimit, upperLimit, description)
}

func ImportGrades(fileStatus string) error {
	headers := []string{
		"id", "code", "name", "lower_limit", "upper_limit", "description",
	}

	return handlers.ModelImport(fileStatus, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		name := row["name"]
		lowerLimitStr := row["lower_limit"]
		lowerLimit, err1 := strconv.ParseFloat(lowerLimitStr, 64)

		if err1 != nil {
			fmt.Println("Error konversi lower_limit:", err1)
		}

		upperLimitStr := row["upper_limit"]
		upperLimit, err2 := strconv.ParseFloat(upperLimitStr, 64)
		if err2 != nil {
			fmt.Println("Error konversi upper_limit:", err2)
		}

		description := row["description"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstGrade{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateGrade(id, code, name, &lowerLimit, &upperLimit, &description); err != nil {
					return err
				}
			} else {
				if err := QueryInsertGrade(id, code, name, &lowerLimit, &upperLimit, &description); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstGrade{})
			if err != nil {
				return err
			}
			if err := QueryInsertGrade(id, code, name, &lowerLimit, &upperLimit, &description); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateGrade(id string, code string, name string, lowerLimit *float64, upperLimit *float64, description *string) error {
	return QueryUpdateGrade(id, code, name, lowerLimit, upperLimit, description)
}

func DeleteGrade(id string) error {
	err := handlers.SPDelete("sp_mst_grades_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashGrades(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstGrade, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_grades_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstGrade{})
	if err != nil {
		return []MstGrade{}, 0, err
	}

	var modelResults []MstGrade
	for _, item := range results {
		level, ok := item.(*MstGrade)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreGrade(id string) error {
	err := handlers.SPRestore("sp_mst_grades_restore_by_id", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountGrades(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstGrade{}, nullable)
}

/* Query */
func QueryInsertGrade(id string, code string, name string, lowerLimit *float64, upperlimit *float64, description *string) error {
	query := `
		EXEC sp_mst_grades_insert
		@id = ?,
		@code = ?, 
		@name = ?, 
		@lower_limit = ?, 
		@upper_limit = ?, 
		@description = ?, 
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, code, name, lowerLimit, upperlimit, description)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateGrade(id string, code string, name string, lowerLimit *float64, upperlimit *float64, description *string) error {
	query := `
		EXEC sp_mst_grades_update
		@id = ?,
		@code = ?, 
		@name = ?, 
		@lower_limit = ?, 
		@upper_limit = ?, 
		@description = ?, 
		@updated_at = ?,
		@updated_by = ?
	`
	err := handlers.SPUpdate(query, id, code, name, lowerLimit, upperlimit, description)
	if err != nil {
		return err
	}

	return nil
}

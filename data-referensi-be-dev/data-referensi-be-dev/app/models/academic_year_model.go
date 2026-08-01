package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type MstAcademicYear struct {
	ID        string  `json:"id"`
	Year      *string `json:"years" gorm:"column:years"`
	Name      string  `json:"name"`
	CreatedAt int64   `json:"created_at"`
	UpdatedAt int64   `json:"updated_at"`
}

type MstAcademicYearDetail struct {
	ID        string  `json:"id"`
	Year      *string `json:"years" gorm:"column:years"`
	Name      string  `json:"name"`
	CreatedAt int64   `json:"created_at"`
	UpdatedAt int64   `json:"updated_at"`
}

type MstAcademicYearExport struct {
	ID   string `json:"id"`
	Year string `json:"years"`
	Name string `json:"name"`
}

type MstAcademicYearSearch struct {
	ID   string `json:"id"`
	Year string `json:"years" gorm:"column:years"`
	Name string `json:"name"`
}

type MstAcademicYearRelation struct {
	ID   string `json:"id"`
	Year string `json:"years"`
	Name string `json:"name"`
}

/* Action */
func GetAcademicYears(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstAcademicYear, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_academic_years_get", filter, sortBy, sortDirection, page, pageSize, &MstAcademicYear{})
	if err != nil {
		return []MstAcademicYear{}, 0, err
	}

	var modelResults []MstAcademicYear
	for _, item := range results {
		level, ok := item.(*MstAcademicYear)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchAcademicYears(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstAcademicYearSearch, error) {
	var results []MstAcademicYearSearch
	err := handlers.SPGet("sp_mst_academic_years_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstAcademicYearSearch{}, err
	}
	return results, nil
}

func ExportAcademicYears(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstAcademicYearExport
	err := handlers.SPGet("sp_mst_academic_years_get", "", "name", "asc", 1, CountAcademicYears(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":   result.ID,
			"Year": result.Year,
			"Name": result.Name,
		}
	}

	headers := []string{
		"ID", "Year", "Name",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetAcademicYear(id string) (MstAcademicYearDetail, error) {
	var result MstAcademicYearDetail
	err := handlers.SPGetByID("sp_mst_academic_years_get_by_id", id, &result)
	if err != nil {
		return MstAcademicYearDetail{}, err
	}
	return result, nil
}

func CreateAcademicYear(id string, years string, name string) error {
	return QueryInsertAcademicYear(id, years, name)
}

func ImportAcademicYears(fileStatus string) error {
	headers := []string{
		"id", "years", "name",
	}

	return handlers.ModelImport(fileStatus, headers, func(row map[string]string) error {
		id := row["id"]
		years := row["years"]
		name := row["name"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstAcademicYear{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateAcademicYear(id, years, name); err != nil {
					return err
				}
			} else {
				if err := QueryInsertAcademicYear(id, years, name); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstAcademicYear{})
			if err != nil {
				return err
			}
			if err := QueryInsertAcademicYear(id, years, name); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateAcademicYear(id string, years string, name string) error {
	return QueryUpdateAcademicYear(id, years, name)
}

func DeleteAcademicYear(id string) error {
	err := handlers.SPDelete("sp_mst_academic_years_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashAcademicYears(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstAcademicYear, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_academic_years_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstAcademicYear{})
	if err != nil {
		return []MstAcademicYear{}, 0, err
	}

	var modelResults []MstAcademicYear
	for _, item := range results {
		level, ok := item.(*MstAcademicYear)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreAcademicYear(id string) error {
	err := handlers.SPRestore("sp_mst_academic_years_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountAcademicYears(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstAcademicYear{}, nullable)
}

/* Query */
func QueryInsertAcademicYear(id string, years string, name string) error {
	query := `
		EXEC sp_mst_academic_years_insert
		@id = ?,
		@years = ?,
		@name = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, years, name)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateAcademicYear(id string, years string, name string) error {
	log.Print(years)

	query := `
		EXEC sp_mst_academic_years_update
		@id = ?,
		@years = ?,
		@name = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, years, name)
	if err != nil {
		return err
	}

	return nil
}

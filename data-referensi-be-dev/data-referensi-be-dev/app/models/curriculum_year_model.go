package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

type MstCurriculumYear struct {
	ID                  string    `json:"id"`
	Years               string    `json:"years"`
	Starts              string    `json:"starts"`
	StartDate           time.Time `json:"start_date"`
	EndDate             time.Time `json:"end_date"`
	Description         *string   `json:"description"`
	AcademicPeriodeName *string   `json:"academic_periode_name"`
	CreatedAt           int64     `json:"created_at"`
	UpdatedAt           int64     `json:"updated_at"`
}

type MstCurriculumYearDetail struct {
	ID                  string    `json:"id"`
	Years               string    `json:"years"`
	Starts              string    `json:"starts"`
	StartDate           time.Time `json:"start_date"`
	EndDate             time.Time `json:"end_date"`
	Description         string    `json:"description"`
	AcademicPeriodeName *string   `json:"academic_periode_name"`
	CreatedAt           int64     `json:"created_at"`
	UpdatedAt           int64     `json:"updated_at"`
}

type MstCurriculumYearExport struct {
	ID          string `json:"id"`
	Years       string `json:"years"`
	Starts      string `json:"starts"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Description string `json:"description"`
}

type MstCurriculumYearSearch struct {
	ID          string    `json:"id"`
	Years       string    `json:"years"`
	Starts      string    `json:"starts"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Description string    `json:"description"`
}

type MstCurriculumYearRelation struct {
	ID          string    `json:"id"`
	Years       string    `json:"years"`
	Starts      string    `json:"starts"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Description string    `json:"description"`
}

/* Action */
func GetCurriculumYears(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstCurriculumYear, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_curriculum_years_get", filter, sortBy, sortDirection, page, pageSize, &MstCurriculumYear{})
	if err != nil {
		log.Print(results)
		return []MstCurriculumYear{}, 0, err
	}

	var modelResults []MstCurriculumYear
	for _, item := range results {
		level, ok := item.(*MstCurriculumYear)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchCurriculumYears(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstCurriculumYearSearch, error) {
	var results []MstCurriculumYearSearch
	err := handlers.SPGet("sp_mst_curriculum_years_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstCurriculumYearSearch{}, err
	}
	return results, nil
}

func ExportCurriculumYears(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstCurriculumYearExport
	err := handlers.SPGet("sp_mst_curriculum_years_get", "", "name", "asc", 1, CountCurriculumYears(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":          result.ID,
			"Years":       result.Years,
			"Starts":      result.Starts,
			"StartDate":   result.StartDate,
			"EndDate":     result.EndDate,
			"Description": result.Description,
		}
	}

	headers := []string{
		"ID", "Years", "Starts", "StartDate", "EndDate", "Description",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetCurriculumYear(id string) (MstCurriculumYearDetail, error) {
	var result MstCurriculumYearDetail
	err := handlers.SPGetByID("sp_mst_curriculum_years_get_by_id", id, &result)
	if err != nil {
		return MstCurriculumYearDetail{}, err
	}
	return result, nil
}

func CreateCurriculumYear(id string, years string, starts string, start_date string, end_date string, description string) error {
	return QueryInsertCurriculumYear(id, years, starts, start_date, end_date, description)
}

func ImportCurriculumYears(fileStatus string) error {
	headers := []string{
		"id", "years", "starts", "start_date", "end_date", "description",
	}

	return handlers.ModelImport(fileStatus, headers, func(row map[string]string) error {
		id := row["id"]
		years := row["years"]
		starts := row["starts"]
		start_date := row["start_date"]
		end_date := row["end_date"]
		description := row["description"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstCurriculumYear{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateCurriculumYear(id, years, starts, start_date, end_date, description); err != nil {
					return err
				}
			} else {
				if err := QueryInsertCurriculumYear(id, years, starts, start_date, end_date, description); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstCurriculumYear{})
			if err != nil {
				return err
			}
			if err := QueryInsertCurriculumYear(id, years, starts, start_date, end_date, description); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateCurriculumYear(id string, years string, starts string, start_date string, end_date string, description string) error {
	return QueryUpdateCurriculumYear(id, years, starts, start_date, end_date, description)
}

func DeleteCurriculumYear(id string) error {
	err := handlers.SPDelete("sp_mst_curriculum_years_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashCurriculumYears(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstCurriculumYear, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_curriculum_years_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstCurriculumYear{})
	if err != nil {
		return []MstCurriculumYear{}, 0, err
	}

	var modelResults []MstCurriculumYear
	for _, item := range results {
		level, ok := item.(*MstCurriculumYear)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreCurriculumYear(id string) error {
	err := handlers.SPRestore("sp_mst_curriculum_years_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountCurriculumYears(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstCurriculumYear{}, nullable)
}

/* Query */
func QueryInsertCurriculumYear(id string, years string, starts string, start_date string, end_date string, description string) error {
	new_start_date, _ := helpers.StringToDate(start_date)
	new_end_date, _ := helpers.StringToDate(end_date)

	query := `
		EXEC sp_mst_curriculum_years_insert
		@id = ?,
		@years = ?, 
		@starts = ?, 
		@start_date = ?, 
		@end_date = ?, 
		@description = ?, 
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, years, starts, new_start_date, new_end_date, description)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func QueryUpdateCurriculumYear(id string, years string, starts string, start_date string, end_date string, description string) error {
	new_start_date, _ := helpers.StringToDate(start_date)
	new_end_date, _ := helpers.StringToDate(end_date)

	query := `
		EXEC sp_mst_curriculum_years_update
		@id = ?,
		@years = ?, 
		@starts = ?, 
		@start_date = ?, 
		@end_date = ?, 
		@description = ?, 
		@updated_at = ?,
		@updated_by = ?
	`
	err := handlers.SPInsert(query, id, years, starts, new_start_date, new_end_date, description)
	if err != nil {
		return err
	}

	return nil
}

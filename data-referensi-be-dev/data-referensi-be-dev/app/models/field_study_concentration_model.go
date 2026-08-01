package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstFieldStudyConcentration struct {
	ID               string `json:"id"`
	Code             string `json:"code"`
	StudyProgramID   string `json:"study_program_id" gorm:"column:study_program_id"`
	StudyProgramName string `json:"study_program_name"`
	StudyProgramCode string `json:"study_program_code"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	CreatedAt        int64  `json:"created_at"`
	UpdatedAt        int64  `json:"updated_at"`
}

type MstFieldStudyConcentrationDetail struct {
	ID               string `json:"id"`
	Code             string `json:"code"`
	StudyProgramID   string `json:"study_program_id" gorm:"column:study_program_id"`
	StudyProgramName string `json:"study_program_name"`
	StudyProgramCode string `json:"study_program_code"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	CreatedAt        int64  `json:"created_at"`
	UpdatedAt        int64  `json:"updated_at"`
}

type MstFieldStudyConcentrationExport struct {
	ID               string `json:"id"`
	Code             string `json:"code"`
	StudyProgramID   string `json:"study_program_id" gorm:"column:study_program_id"`
	StudyProgramName string `json:"study_program_name"`
	StudyProgramCode string `json:"study_program_code"`
	Name             string `json:"name"`
	Description      string `json:"description"`
}

type MstFieldStudyConcentrationSearch struct {
	ID               string `json:"id"`
	Code             string `json:"code"`
	StudyProgramID   string `json:"study_program_id" gorm:"column:study_program_id"`
	StudyProgramName string `json:"study_program_name"`
	StudyProgramCode string `json:"study_program_code"`
	Name             string `json:"name"`
	Description      string `json:"description"`
}

type MstFieldStudyConcentrationRelation struct {
	ID               string `json:"id"`
	Code             string `json:"code"`
	StudyProgramID   string `json:"study_program_id" gorm:"column:study_program_id"`
	StudyProgramName string `json:"study_program_name"`
	StudyProgramCode string `json:"study_program_code"`
	Name             string `json:"name"`
	Description      string `json:"description"`
}

/* Action */
func GetFieldStudyConcentrations(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstFieldStudyConcentration, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_field_study_concentrations_get", filter, sortBy, sortDirection, page, pageSize, &MstFieldStudyConcentration{})
	if err != nil {
		return []MstFieldStudyConcentration{}, 0, err
	}

	var modelResults []MstFieldStudyConcentration
	for _, item := range results {
		level, ok := item.(*MstFieldStudyConcentration)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchFieldStudyConcentrations(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstFieldStudyConcentrationSearch, error) {
	var results []MstFieldStudyConcentrationSearch
	err := handlers.SPGet("sp_mst_field_study_concentrations_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstFieldStudyConcentrationSearch{}, err
	}
	return results, nil
}

func ExportFieldStudyConcentrations(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstFieldStudyConcentrationExport
	err := handlers.SPGet("sp_mst_field_study_concentrations_get", "", "name", "asc", 1, CountFieldStudyConcentrations(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":               result.ID,
			"Code":             result.Code,
			"StudyProgramID":   result.StudyProgramID,
			"StudyProgramName": result.StudyProgramName,
			"StudyProgramCode": result.StudyProgramCode,
			"Name":             result.Name,
			"Description":      result.Description,
		}
	}

	headers := []string{
		"ID", "Code", "Study Program ID", "Study Program Name", "Study Program Code", "Name", "Description",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetFieldStudyConcentration(id string) (MstFieldStudyConcentrationDetail, error) {
	var result MstFieldStudyConcentrationDetail
	err := handlers.SPGetByID("sp_mst_field_study_concentrations_get_by_id", id, &result)
	if err != nil {
		return MstFieldStudyConcentrationDetail{}, err
	}
	return result, nil
}

func CreateFieldStudyConcentration(id string, code string, study_program_id string, name string, description string) error {
	return QueryInsertFieldStudyConcentration(id, code, study_program_id, name, description)
}

func ImportFieldStudyConcentrations(fileStatus string) error {
	headers := []string{
		"id", "code", "study_program_id", "name", "description",
	}

	return handlers.ModelImport(fileStatus, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		study_program_id := row["study_program_id"]
		name := row["name"]
		description := row["description"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstFieldStudyConcentration{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateFieldStudyConcentration(id, code, study_program_id, name, description); err != nil {
					return err
				}
			} else {
				if err := QueryInsertFieldStudyConcentration(id, code, study_program_id, name, description); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstFieldStudyConcentration{})
			if err != nil {
				return err
			}
			if err := QueryInsertFieldStudyConcentration(id, code, study_program_id, name, description); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateFieldStudyConcentration(id string, code string, study_program_id string, name string, description string) error {
	return QueryUpdateFieldStudyConcentration(id, code, study_program_id, name, description)
}

func DeleteFieldStudyConcentration(id string) error {
	err := handlers.SPDelete("sp_mst_field_study_concentrations_delete_by_id", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashFieldStudyConcentrations(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstFieldStudyConcentration, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_field_study_concentrations_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstFieldStudyConcentration{})
	if err != nil {
		return []MstFieldStudyConcentration{}, 0, err
	}

	var modelResults []MstFieldStudyConcentration
	for _, item := range results {
		level, ok := item.(*MstFieldStudyConcentration)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreFieldStudyConcentration(id string) error {
	err := handlers.SPRestore("sp_mst_field_study_concentrations_restore_by_id", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountFieldStudyConcentrations(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstFieldStudyConcentration{}, nullable)
}

/* Query */
func QueryInsertFieldStudyConcentration(id string, code string, study_program_id string, name string, description string) error {
	query := `
		EXEC sp_mst_field_study_concentrations_insert
		@id = ?,
		@code = ?, 
		@study_program_id = ?, 
		@name = ?, 
		@description = ?, 
		@created_at = ?,
		@created_by = ?
	`

	err := handlers.SPInsertSemesterNumber(query, id, code, study_program_id, name, description)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateFieldStudyConcentration(id string, code string, study_program_id string, name string, description string) error {
	query := `
		EXEC sp_mst_field_study_concentrations_update
		@id = ?,
		@code = ?, 
		@study_program_id = ?, 
		@name = ?, 
		@description = ?, 
		@updated_at = ?,
		@updated_by = ?
	`
	err := handlers.SPUpdate(query, id, code, study_program_id, name, description)
	if err != nil {
		return err
	}

	return nil
}

package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstEducation struct {
	ID                   string `json:"id"`
	EducationalLevelName string `json:"educational_level_name"`
	StudyProgramName     string `json:"study_program_name"`
	Name                 string `json:"name"`
	CreatedAt            int64  `json:"created_at"`
	UpdatedAt            int64  `json:"updated_at"`
}

type MstEducationDetail struct {
	ID                   string `json:"id"`
	EducationalLevelId   string `json:"educational_level_id"`
	EducationalLevelName string `json:"educational_level_name"`
	StudyProgramId       string `json:"study_program_id"`
	StudyProgramName     string `json:"study_program_name"`
	Name                 string `json:"name"`
	CreatedAt            int64  `json:"created_at"`
	UpdatedAt            int64  `json:"updated_at"`
}

type MstEducationExport struct {
	ID                   string `json:"id"`
	EducationalLevelId   string `json:"educational_level_id"`
	EducationalLevelName string `json:"educational_level_name"`
	StudyProgramId       string `json:"study_program_id"`
	StudyProgramName     string `json:"study_program_name"`
	Name                 string `json:"name"`
}

type MstEducationSearch struct {
	ID                   string `json:"id"`
	EducationalLevelName string `json:"educational_level_name"`
	StudyProgramName     string `json:"study_program_name"`
	Name                 string `json:"name"`
}
type MstEducationRelation struct {
	ID                   string `json:"id"`
	EducationalLevelName string `json:"educational_level_name"`
	StudyProgramName     string `json:"study_program_name"`
	Name                 string `json:"name"`
}

/* Action */
func GetEducations(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstEducation, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_educations_get", filter, sortBy, sortDirection, page, pageSize, &MstEducation{})
	if err != nil {
		return []MstEducation{}, 0, err
	}

	var modelResults []MstEducation
	for _, item := range results {
		level, ok := item.(*MstEducation)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchEducations(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstEducationSearch, error) {
	var results []MstEducationSearch
	err := handlers.SPGet("sp_mst_educations_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstEducationSearch{}, err
	}
	return results, nil
}

func ExportEducations(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstEducationExport
	err := handlers.SPGet("sp_mst_educations_get", "", "name", "asc", 1, CountEducations(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":                     result.ID,
			"Educational Level ID":   result.EducationalLevelId,
			"Educational Level Name": result.EducationalLevelName,
			"Study Program ID":       result.StudyProgramId,
			"Study Program Name":     result.StudyProgramName,
			"Name":                   result.Name,
		}
	}

	headers := []string{
		"ID",
		"Educational Level ID",
		"Educational Level Name",
		"Study Program ID",
		"Study Program Name",
		"Name",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetEducation(id string) (MstEducationDetail, error) {
	var result MstEducationDetail
	err := handlers.SPGetByID("sp_mst_educations_get_by_id", id, &result)
	if err != nil {
		return MstEducationDetail{}, err
	}
	return result, nil
}

func GetEducationByEducationalLevelId(educational_level_id string) ([]MstEducationSearch, error) {
	var results []MstEducationSearch

	query := fmt.Sprintf(`
		EXEC sp_mst_educations_get_by_education_level_id
		@education_level_id = '%s'
	`, educational_level_id)

	err := handlers.SPGetByQuery(query, &results)
	if err != nil {
		return []MstEducationSearch{}, err
	}
	return results, nil
}

func CreateEducation(id string, educational_level_id string, study_program_id string, name string) error {
	return QueryInsertEducation(id, educational_level_id, study_program_id, name)
}

func ImportEducations(filePath string) error {
	headers := []string{
		"id",
		"educational_level_id",
		"educational_level_name",
		"study_program_id",
		"study_program_name",
		"name",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		educational_level_id := row["educational_level_id"]
		study_program_id := row["study_program_id"]
		name := row["name"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstEducation{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateEducation(id, educational_level_id, study_program_id, name); err != nil {
					return err
				}
			} else {
				if err := QueryInsertEducation(id, educational_level_id, study_program_id, name); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstEducation{})
			if err != nil {
				return err
			}
			if err := QueryInsertEducation(id, educational_level_id, study_program_id, name); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateEducation(id string, educational_level_id string, study_program_id string, name string) error {
	return QueryUpdateEducation(id, educational_level_id, study_program_id, name)
}

func DeleteEducation(id string) error {
	err := handlers.SPDelete("sp_mst_educations_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashEducations(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstEducation, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_educations_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstEducation{})
	if err != nil {
		return []MstEducation{}, 0, err
	}

	var modelResults []MstEducation
	for _, item := range results {
		level, ok := item.(*MstEducation)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreEducation(id string) error {
	err := handlers.SPRestore("sp_mst_educations_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountEducations(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstEducation{}, nullable)
}

/* Query */
func QueryInsertEducation(id string, educational_level_id string, study_program_id string, name string) error {
	query := `
		EXEC sp_mst_educations_insert
		@id = ?,
		@educational_level_id = ?,
		@study_program_id = ?,
		@name = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, educational_level_id, study_program_id, name)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateEducation(id string, educational_level_id string, study_program_id string, name string) error {
	query := `
		EXEC sp_mst_educations_update
		@id = ?,
		@educational_level_id = ?,
		@study_program_id = ?,
		@name = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, educational_level_id, study_program_id, name)
	if err != nil {
		return err
	}

	return nil
}

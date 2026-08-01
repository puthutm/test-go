package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstUnsiaStudyProgram struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstUnsiaStudyProgramDetail struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstUnsiaStudyProgramExport struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type MstUnsiaStudyProgramSearch struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type MstUnsiaStudyProgramRelation struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

/* Action */
func GetUnsiaStudyPrograms(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstUnsiaStudyProgram, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_unsia_study_programs_get", filter, sortBy, sortDirection, page, pageSize, &MstUnsiaStudyProgram{})
	if err != nil {
		return []MstUnsiaStudyProgram{}, 0, err
	}

	var modelResults []MstUnsiaStudyProgram
	for _, item := range results {
		level, ok := item.(*MstUnsiaStudyProgram)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchUnsiaStudyPrograms(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstUnsiaStudyProgramSearch, error) {
	var results []MstUnsiaStudyProgramSearch
	err := handlers.SPGet("sp_mst_unsia_study_programs_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstUnsiaStudyProgramSearch{}, err
	}
	return results, nil
}

func ExportUnsiaStudyPrograms(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstUnsiaStudyProgramExport
	err := handlers.SPGet("sp_mst_unsia_study_programs_get", "", "name", "asc", 1, CountUnsiaStudyPrograms(), &results)
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
		"ID",
		"Code",
		"Name",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetUnsiaStudyProgram(id string) (MstUnsiaStudyProgramDetail, error) {
	var result MstUnsiaStudyProgramDetail
	err := handlers.SPGetByID("sp_mst_unsia_study_programs_get_by_id", id, &result)
	if err != nil {
		return MstUnsiaStudyProgramDetail{}, err
	}
	return result, nil
}

func CreateUnsiaStudyProgram(id string, code string, name string) error {
	return QueryInsertUnsiaStudyProgram(id, code, name)
}

func ImportUnsiaStudyPrograms(filePath string) error {
	headers := []string{
		"id", "code", "name",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		name := row["name"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstUnsiaStudyProgram{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateUnsiaStudyProgram(id, code, name); err != nil {
					return err
				}
			} else {
				if err := QueryInsertUnsiaStudyProgram(id, code, name); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstUnsiaStudyProgram{})
			if err != nil {
				return err
			}
			if err := QueryInsertUnsiaStudyProgram(id, code, name); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateUnsiaStudyProgram(id string, code string, name string) error {
	return QueryUpdateUnsiaStudyProgram(id, code, name)
}

func DeleteUnsiaStudyProgram(id string) error {
	err := handlers.SPDelete("sp_mst_unsia_study_programs_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashUnsiaStudyPrograms(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstUnsiaStudyProgram, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_unsia_study_programs_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstUnsiaStudyProgram{})
	if err != nil {
		return []MstUnsiaStudyProgram{}, 0, err
	}

	var modelResults []MstUnsiaStudyProgram
	for _, item := range results {
		level, ok := item.(*MstUnsiaStudyProgram)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreUnsiaStudyProgram(id string) error {
	err := handlers.SPRestore("sp_mst_unsia_study_programs_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountUnsiaStudyPrograms(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstUnsiaStudyProgram{}, nullable)
}

/* Query */
func QueryInsertUnsiaStudyProgram(id string, code string, name string) error {
	query := `
		EXEC sp_mst_unsia_study_programs_insert
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

func QueryUpdateUnsiaStudyProgram(id string, code string, name string) error {
	query := `
		EXEC sp_mst_unsia_study_programs_update
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

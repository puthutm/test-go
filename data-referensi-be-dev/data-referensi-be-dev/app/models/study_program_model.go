package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type MstStudyProgram struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstStudyProgramDetail struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstStudyProgramExport struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type MstStudyProgramSearch struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Type string `json:"type"`
}
type MstStudyProgramRelation struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Type string `json:"type"`
}

/* Action */
func GetStudyPrograms(filter string, sortBy string, sortDirection string, page int, pageSize int64, queryType string) ([]MstStudyProgram, int64, error) {
	results, total, err := handlers.SPGetWithCountStudyProgram("sp_mst_study_programs_get", filter, sortBy, sortDirection, page, pageSize, queryType, &MstStudyProgram{})
	if err != nil {
		log.Print(err)
		return []MstStudyProgram{}, 0, err
	}

	var modelResults []MstStudyProgram
	for _, item := range results {
		level, ok := item.(*MstStudyProgram)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchStudyPrograms(filter string, sortBy string, sortDirection string, page int, pageSize int64, queryType string) ([]MstStudyProgramSearch, error) {
	var results []MstStudyProgramSearch
	err := handlers.SPGetStudyProgram("sp_mst_study_programs_get", filter, sortBy, sortDirection, page, pageSize, queryType, &results)
	if err != nil {
		return []MstStudyProgramSearch{}, err
	}
	return results, nil
}

func ExportStudyPrograms(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstStudyProgramExport
	err := handlers.SPGet("sp_mst_study_programs_get", "", "name", "asc", 1, CountStudyPrograms(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":   result.ID,
			"Code": result.Code,
			"Name": result.Name,
			"Type": result.Type,
		}
	}

	headers := []string{
		"ID",
		"Name",
		"Code",
		"Type",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetStudyProgram(id string) (MstStudyProgramDetail, error) {
	var result MstStudyProgramDetail
	err := handlers.SPGetByID("sp_mst_study_programs_get_by_id", id, &result)
	if err != nil {
		return MstStudyProgramDetail{}, err
	}
	return result, nil
}

func CreateStudyProgram(id string, code string, name string, typePS string) error {
	return QueryInsertStudyProgram(id, code, name, typePS)
}

func ImportStudyPrograms(filePath string) error {
	headers := []string{
		"id", "code", "name", "type",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		name := row["name"]
		typePS := row["type"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstStudyProgram{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateStudyProgram(id, code, name, typePS); err != nil {
					return err
				}
			} else {
				if err := QueryInsertStudyProgram(id, code, name, typePS); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstStudyProgram{})
			if err != nil {
				return err
			}
			if err := QueryInsertStudyProgram(id, code, name, typePS); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateStudyProgram(id string, code string, name string, typePS string) error {
	return QueryUpdateStudyProgram(id, code, name, typePS)
}

func DeleteStudyProgram(id string) error {
	err := handlers.SPDelete("sp_mst_study_programs_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashStudyPrograms(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstStudyProgram, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_study_programs_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstStudyProgram{})
	if err != nil {
		return []MstStudyProgram{}, 0, err
	}

	var modelResults []MstStudyProgram
	for _, item := range results {
		level, ok := item.(*MstStudyProgram)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreStudyProgram(id string) error {
	err := handlers.SPRestore("sp_mst_study_programs_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountStudyPrograms(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstStudyProgram{}, nullable)
}

/* Query */
func QueryInsertStudyProgram(id string, code string, name string, typePS string) error {
	query := `
		EXEC sp_mst_study_programs_insert
		@id = ?,
		@code = ?,
		@name = ?,
		@type = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, code, name, typePS)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateStudyProgram(id string, code string, name string, typePS string) error {
	query := `
		EXEC sp_mst_study_programs_update
		@id = ?,
		@code = ?,
		@name = ?,
		@type = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, code, name, typePS)
	if err != nil {
		return err
	}

	return nil
}

package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstFieldOfStudy struct {
	ID          string `json:"id"`
	Name        string `json:"name" `
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstFieldOfStudyDetail struct {
	ID          string `json:"id"`
	Name        string `json:"name" `
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstFieldOfStudyExport struct {
	ID          string `json:"id"`
	Name        string `json:"name" `
	Description string `json:"description"`
}

type MstFieldOfStudySearch struct {
	ID          string `json:"id"`
	Name        string `json:"name" `
	Description string `json:"description"`
}

type MstFieldOfStudyRelation struct {
	ID   string `json:"id"`
	Name string `json:"name" `
}

/* Action */
func GetFieldOfStudies(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstFieldOfStudy, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_field_of_studies_get", filter, sortBy, sortDirection, page, pageSize, &MstFieldOfStudy{})
	if err != nil {
		return []MstFieldOfStudy{}, 0, err
	}

	var modelResults []MstFieldOfStudy
	for _, item := range results {
		level, ok := item.(*MstFieldOfStudy)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchFieldOfStudies(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstFieldOfStudySearch, error) {
	var results []MstFieldOfStudySearch
	err := handlers.SPGet("sp_mst_field_of_studies_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstFieldOfStudySearch{}, err
	}
	return results, nil
}

func ExportFieldOfStudies(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstFieldOfStudyExport

	err := handlers.SPGet("sp_mst_field_of_studies_get", "", "name", "asc", 1, CountFieldOfStudies(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":          result.ID,
			"Name":        result.Name,
			"Description": result.Description,
		}
	}

	headers := []string{
		"ID", "Name", "Description",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetFieldOfStudy(id string) (MstFieldOfStudyDetail, error) {
	var result MstFieldOfStudyDetail
	err := handlers.SPGetByID("sp_mst_field_of_studies_get_by_id", id, &result)
	if err != nil {
		return MstFieldOfStudyDetail{}, err
	}
	return result, nil
}

func CreateFieldOfStudy(id string, name string, description string) error {
	return QueryInsertFieldOfStudy(id, name, description)
}

func ImportFieldOfStudies(filePath string) error {
	headers := []string{
		"id", "name", "description",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		name := row["name"]
		description := row["description"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstFieldOfStudy{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateFieldOfStudy(id, name, description); err != nil {
					return err
				}
			} else {
				if err := QueryInsertFieldOfStudy(id, name, description); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstFieldOfStudy{})
			if err != nil {
				return err
			}
			if err := QueryInsertFieldOfStudy(id, name, description); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateFieldOfStudy(id string, name string, description string) error {
	return QueryUpdateFieldOfStudy(id, name, description)
}

func DeleteFieldOfStudy(id string) error {
	err := handlers.SPDelete("sp_mst_field_of_studies_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashFieldOfStudies(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstFieldOfStudy, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_field_of_studies_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstFieldOfStudy{})
	if err != nil {
		return []MstFieldOfStudy{}, 0, err
	}

	var modelResults []MstFieldOfStudy
	for _, item := range results {
		level, ok := item.(*MstFieldOfStudy)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreFieldOfStudy(id string) error {
	err := handlers.SPRestore("sp_mst_field_of_studies_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountFieldOfStudies(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstFieldOfStudy{}, nullable)
}

/* Query */
func QueryInsertFieldOfStudy(id string, name string, description string) error {
	query := `
		EXEC sp_mst_field_of_studies_insert
		@id = ?,
		@name = ?,
		@description = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, name, description)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateFieldOfStudy(id string, name string, description string) error {
	query := `
		EXEC sp_mst_field_of_studies_update
		@id = ?,
		@name = ?,
		@description = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, name, description)
	if err != nil {
		return err
	}

	return nil
}

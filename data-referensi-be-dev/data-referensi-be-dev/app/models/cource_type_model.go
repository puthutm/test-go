package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstCourseType struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstCourseTypeDetail struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstCourseTypeExport struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MstCourseTypeSearch struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MstCourseTypeRelation struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

/* Action */
func GetCourseTypes(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstCourseType, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_course_types_get", filter, sortBy, sortDirection, page, pageSize, &MstCourseType{})
	if err != nil {
		return []MstCourseType{}, 0, err
	}

	var modelResults []MstCourseType
	for _, item := range results {
		level, ok := item.(*MstCourseType)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchCourseTypes(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstCourseTypeSearch, error) {
	var results []MstCourseTypeSearch
	err := handlers.SPGet("sp_mst_course_types_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstCourseTypeSearch{}, err
	}
	return results, nil
}

func ExportCourseTypes(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstCourseTypeExport
	err := handlers.SPGet("sp_mst_course_types_get", "", "name", "asc", 1, CountCourseTypes(), &results)
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

func GetCourseType(id string) (MstCourseTypeDetail, error) {
	var result MstCourseTypeDetail
	err := handlers.SPGetByID("sp_mst_course_types_get_by_id", id, &result)
	if err != nil {
		return MstCourseTypeDetail{}, err
	}
	return result, nil
}

func CreateCourseType(id string, name string, description string) error {
	return QueryInsertCourseType(id, name, description)
}

func ImportCourseTypes(fileStatus string) error {
	headers := []string{
		"id", "name", "description",
	}

	return handlers.ModelImport(fileStatus, headers, func(row map[string]string) error {
		id := row["id"]
		name := row["name"]
		description := row["description"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstCourseType{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateCourseType(id, name, description); err != nil {
					return err
				}
			} else {
				if err := QueryInsertCourseType(id, name, description); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstCourseType{})
			if err != nil {
				return err
			}
			if err := QueryInsertCourseType(id, name, description); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateCourseType(id string, name string, description string) error {
	return QueryUpdateCourseType(id, name, description)
}

func DeleteCourseType(id string) error {
	err := handlers.SPDelete("sp_mst_course_types_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashCourseTypes(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstCourseType, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_course_types_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstCourseType{})
	if err != nil {
		return []MstCourseType{}, 0, err
	}

	var modelResults []MstCourseType
	for _, item := range results {
		level, ok := item.(*MstCourseType)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreCourseType(id string) error {
	err := handlers.SPRestore("sp_mst_course_types_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountCourseTypes(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstCourseType{}, nullable)
}

/* Query */
func QueryInsertCourseType(id string, name string, description string) error {
	query := `
		EXEC sp_mst_course_types_insert
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

func QueryUpdateCourseType(id string, name string, description string) error {
	query := `
		EXEC sp_mst_course_types_update
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

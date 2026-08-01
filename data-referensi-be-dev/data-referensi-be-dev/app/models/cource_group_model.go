package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstCourseGroup struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstCourseGroupDetail struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstCourseGroupExport struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MstCourseGroupSearch struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MstCourseGroupRelation struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

/* Action */
func GetCourseGroups(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstCourseGroup, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_course_groups_get", filter, sortBy, sortDirection, page, pageSize, &MstCourseGroup{})
	if err != nil {
		return []MstCourseGroup{}, 0, err
	}

	var modelResults []MstCourseGroup
	for _, item := range results {
		level, ok := item.(*MstCourseGroup)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchCourseGroups(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstCourseGroupSearch, error) {
	var results []MstCourseGroupSearch
	err := handlers.SPGet("sp_mst_course_groups_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstCourseGroupSearch{}, err
	}
	return results, nil
}

func ExportCourseGroups(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstCourseGroupExport
	err := handlers.SPGet("sp_mst_course_groups_get", "", "name", "asc", 1, CountCourseGroups(), &results)
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

func GetCourseGroup(id string) (MstCourseGroupDetail, error) {
	var result MstCourseGroupDetail
	err := handlers.SPGetByID("sp_mst_course_groups_get_by_id", id, &result)
	if err != nil {
		return MstCourseGroupDetail{}, err
	}
	return result, nil
}

func CreateCourseGroup(id string, name string, description string) error {
	return QueryInsertCourseGroup(id, name, description)
}

func ImportCourseGroups(fileStatus string) error {
	headers := []string{
		"id", "name", "description",
	}

	return handlers.ModelImport(fileStatus, headers, func(row map[string]string) error {
		id := row["id"]
		name := row["name"]
		description := row["description"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstCourseGroup{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateCourseGroup(id, name, description); err != nil {
					return err
				}
			} else {
				if err := QueryInsertCourseGroup(id, name, description); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstCourseGroup{})
			if err != nil {
				return err
			}
			if err := QueryInsertCourseGroup(id, name, description); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateCourseGroup(id string, name string, description string) error {
	return QueryUpdateCourseGroup(id, name, description)
}

func DeleteCourseGroup(id string) error {
	err := handlers.SPDelete("sp_mst_course_groups_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashCourseGroups(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstCourseGroup, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_course_groups_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstCourseGroup{})
	if err != nil {
		return []MstCourseGroup{}, 0, err
	}

	var modelResults []MstCourseGroup
	for _, item := range results {
		level, ok := item.(*MstCourseGroup)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreCourseGroup(id string) error {
	err := handlers.SPRestore("sp_mst_course_groups_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountCourseGroups(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstCourseGroup{}, nullable)
}

/* Query */
func QueryInsertCourseGroup(id string, name string, description string) error {
	query := `
		EXEC sp_mst_course_groups_insert
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

func QueryUpdateCourseGroup(id string, name string, description string) error {
	query := `
		EXEC sp_mst_course_groups_update
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

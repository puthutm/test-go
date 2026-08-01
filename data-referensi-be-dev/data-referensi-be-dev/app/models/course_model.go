package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type MstCourse struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstCourseDetail struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstCourseExport struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MstCourseSearch struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MstCourseRelation struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

/* Action */
func GetCourses(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstCourse, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_courses_get", filter, sortBy, sortDirection, page, pageSize, &MstCourse{})
	if err != nil {
		return []MstCourse{}, 0, err
	}

	var modelResults []MstCourse
	for _, item := range results {
		level, ok := item.(*MstCourse)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchCourses(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstCourseSearch, error) {
	var results []MstCourseSearch
	err := handlers.SPGet("sp_mst_courses_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstCourseSearch{}, err
	}
	return results, nil
}

func ExportCourses(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstCourseExport
	err := handlers.SPGet("sp_mst_courses_get", "", "name", "asc", 1, CountCourses(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":          result.ID,
			"Code":        result.Code,
			"Name":        result.Name,
			"Description": result.Description,
		}
	}

	headers := []string{
		"ID", "Code", "Name", "Description",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetCourse(id string) (MstCourseDetail, error) {
	var result MstCourseDetail
	err := handlers.SPGetByID("sp_mst_courses_get_by_id", id, &result)
	if err != nil {
		return MstCourseDetail{}, err
	}
	return result, nil
}

func CreateCourse(id string, code string, name string, description string) error {
	return QueryInsertCourse(id, code, name, description)
}

func ImportCourses(fileStatus string) error {
	headers := []string{
		"id", "code", "name", "description",
	}

	return handlers.ModelImport(fileStatus, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		name := row["name"]
		description := row["description"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstCourse{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateCourse(id, code, name, description); err != nil {
					return err
				}
			} else {
				if err := QueryInsertCourse(id, code, name, description); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstCourse{})
			if err != nil {
				return err
			}
			if err := QueryInsertCourse(id, code, name, description); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateCourse(id string, code string, name string, description string) error {
	return QueryUpdateCourse(id, code, name, description)
}

func DeleteCourse(id string) error {
	err := handlers.SPDelete("sp_mst_courses_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashCourses(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstCourse, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_courses_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstCourse{})
	if err != nil {
		return []MstCourse{}, 0, err
	}

	var modelResults []MstCourse
	for _, item := range results {
		level, ok := item.(*MstCourse)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreCourse(id string) error {
	err := handlers.SPRestore("sp_mst_courses_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountCourses(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstCourse{}, nullable)
}

/* Query */
func QueryInsertCourse(id string, code string, name string, description string) error {
	query := `
		EXEC sp_mst_courses_insert
		@id = ?,
		@code = ?, 
		@name = ?, 
		@description = ?, 
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`
	log.Print(query)

	err := handlers.SPInsert(query, id, code, name, description)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateCourse(id string, code string, name string, description string) error {
	query := `
		EXEC sp_mst_courses_update
		@id = ?,
		@code = ?, 
		@name = ?, 
		@description = ?, 
		@updated_at = ?,
		@updated_by = ?
	`
	err := handlers.SPUpdate(query, id, code, name, description)
	if err != nil {
		return err
	}

	return nil
}

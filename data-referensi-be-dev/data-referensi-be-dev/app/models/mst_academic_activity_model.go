package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type MstAcademicActivity struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstAcademicActivityDetail struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstAcademicActivityExport struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MstAcademicActivitySearch struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MstAcademicActivityRelation struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

/* Action */
func GetAcademicActivities(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstAcademicActivity, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_academic_activities_get", filter, sortBy, sortDirection, page, pageSize, &MstAcademicActivity{})
	if err != nil {
		return []MstAcademicActivity{}, 0, err
	}

	var modelResults []MstAcademicActivity
	for _, item := range results {
		level, ok := item.(*MstAcademicActivity)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchAcademicActivities(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstAcademicActivitySearch, error) {
	var results []MstAcademicActivitySearch
	err := handlers.SPGet("sp_mst_academic_activities_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstAcademicActivitySearch{}, err
	}
	return results, nil
}

func ExportAcademicActivities(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstAcademicActivityExport
	err := handlers.SPGet("sp_mst_academic_activities_get", "", "name", "asc", 1, CountAcademicActivities(), &results)
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

func GetAcademicActivity(id string) (MstAcademicActivityDetail, error) {
	var result MstAcademicActivityDetail
	err := handlers.SPGetByID("sp_mst_academic_activities_get_by_id", id, &result)
	if err != nil {
		return MstAcademicActivityDetail{}, err
	}
	return result, nil
}

func CreateAcademicActivity(id string, name string, description string) error {
	return QueryInsertAcademicActivity(id, name, description)
}

func ImportAcademicActivities(fileStatus string) error {
	headers := []string{
		"id", "name", "description",
	}

	return handlers.ModelImport(fileStatus, headers, func(row map[string]string) error {
		id := row["id"]
		name := row["name"]
		description := row["description"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstAcademicActivity{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateAcademicActivity(id, name, description); err != nil {
					return err
				}
			} else {
				if err := QueryInsertAcademicActivity(id, name, description); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstAcademicActivity{})
			if err != nil {
				return err
			}
			if err := QueryInsertAcademicActivity(id, name, description); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateAcademicActivity(id string, name string, description string) error {
	return QueryUpdateAcademicActivity(id, name, description)
}

func DeleteAcademicActivity(id string) error {
	err := handlers.SPDelete("sp_mst_academic_activities_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashAcademicActivities(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstAcademicActivity, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_academic_activities_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstAcademicActivity{})
	if err != nil {
		return []MstAcademicActivity{}, 0, err
	}

	var modelResults []MstAcademicActivity
	for _, item := range results {
		level, ok := item.(*MstAcademicActivity)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreAcademicActivity(id string) error {
	err := handlers.SPRestore("sp_mst_academic_activities_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountAcademicActivities(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstAcademicActivity{}, nullable)
}

/* Query */
func QueryInsertAcademicActivity(id string, name string, description string) error {
	query := `
		EXEC sp_mst_academic_activities_insert
		@id = ?,
		@name = ?, 
		@description = ?, 
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`
	log.Print(query)

	err := handlers.SPInsert(query, id, name, description)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateAcademicActivity(id string, name string, description string) error {
	query := `
		EXEC sp_mst_academic_activities_update
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

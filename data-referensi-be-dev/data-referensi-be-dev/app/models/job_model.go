package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstJob struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstJobDetail struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstJobExport struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MstJobSearch struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MstJobRelation struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

/* Action */
func GetJobs(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstJob, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_jobs_get", filter, sortBy, sortDirection, page, pageSize, &MstJob{})
	if err != nil {
		return []MstJob{}, 0, err
	}

	var modelResults []MstJob
	for _, item := range results {
		level, ok := item.(*MstJob)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchJobs(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstJobSearch, error) {
	var results []MstJobSearch
	err := handlers.SPGet("sp_mst_jobs_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstJobSearch{}, err
	}
	return results, nil
}

func ExportJobs(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstJobExport
	err := handlers.SPGet("sp_mst_jobs_get", "", "name", "asc", 1, CountJobs(), &results)
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
		"ID",
		"Code",
		"Name",
		"Description",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetJob(id string) (MstJobDetail, error) {
	var result MstJobDetail
	err := handlers.SPGetByID("sp_mst_jobs_get_by_id", id, &result)
	if err != nil {
		return MstJobDetail{}, err
	}
	return result, nil
}

func CreateJob(id string, code string, name string, description string) error {
	return QueryInsertJob(id, code, name, description)
}

func ImportJobs(filePath string) error {
	headers := []string{
		"id", "code", "name", "description",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		name := row["name"]
		description := row["description"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstJob{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateJob(id, code, name, description); err != nil {
					return err
				}
			} else {
				if err := QueryInsertJob(id, code, name, description); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstJob{})
			if err != nil {
				return err
			}
			if err := QueryInsertJob(id, code, name, description); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateJob(id string, code string, name string, description string) error {
	return QueryUpdateJob(id, code, name, description)
}

func DeleteJob(id string) error {
	err := handlers.SPDelete("sp_mst_jobs_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashJobs(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstJob, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_jobs_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstJob{})
	if err != nil {
		return []MstJob{}, 0, err
	}

	var modelResults []MstJob
	for _, item := range results {
		level, ok := item.(*MstJob)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreJob(id string) error {
	err := handlers.SPRestore("sp_mst_jobs_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountJobs(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstJob{}, nullable)
}

/* Query */
func QueryInsertJob(id string, code string, name string, description string) error {
	query := `
		EXEC sp_mst_jobs_insert
		@id = ?,
		@code= ?,
		@name = ?,
		@description = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, code, name, description)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateJob(id string, code string, name string, description string) error {
	query := `
		EXEC sp_mst_jobs_update
		@id = ?,
		@code= ?,
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

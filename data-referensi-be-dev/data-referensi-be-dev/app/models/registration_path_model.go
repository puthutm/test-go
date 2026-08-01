package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstRegistrationPath struct {
	ID               string `json:"id"`
	Code             string `json:"code"`
	RegistrationPath string `json:"registration_path"`
	RegistrationType string `json:"registration_type"`
	Description      string `json:"description"`
	CreatedAt        int64  `json:"created_at"`
	UpdatedAt        int64  `json:"updated_at"`
}

type MstRegistrationPathDetail struct {
	ID               string `json:"id"`
	Code             string `json:"code"`
	RegistrationPath string `json:"registration_path"`
	RegistrationType string `json:"registration_type"`
	Description      string `json:"description"`
	CreatedAt        int64  `json:"created_at"`
	UpdatedAt        int64  `json:"updated_at"`
}

type MstRegistrationPathExport struct {
	ID               string `json:"id"`
	Code             string `json:"code"`
	RegistrationPath string `json:"registration_path"`
	RegistrationType string `json:"registration_type"`
	Description      string `json:"description"`
}

type MstRegistrationPathSearch struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type MstRegistrationPathRelation struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

/* Action */
func GetRegistrationPaths(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstRegistrationPath, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_registration_paths_get", filter, sortBy, sortDirection, page, pageSize, &MstRegistrationPath{})
	if err != nil {
		return []MstRegistrationPath{}, 0, err
	}

	var modelResults []MstRegistrationPath
	for _, item := range results {
		level, ok := item.(*MstRegistrationPath)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchRegistrationPaths(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstRegistrationPathSearch, error) {
	var results []MstRegistrationPathSearch
	err := handlers.SPGet("sp_mst_registration_paths_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstRegistrationPathSearch{}, err
	}
	return results, nil
}

func ExportRegistrationPaths(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstRegistrationPathExport
	err := handlers.SPGet("sp_mst_registration_paths_get", "", "name", "asc", 1, CountRegistrationPaths(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":                result.ID,
			"Code":              result.Code,
			"Registration Path": result.RegistrationPath,
			"Registration Type": result.RegistrationType,
			"Description":       result.Description,
		}
	}

	headers := []string{
		"ID", "Code", "Registration Path", "Registration Type", "Description",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetRegistrationPath(id string) (MstRegistrationPathDetail, error) {
	var result MstRegistrationPathDetail
	err := handlers.SPGetByID("sp_mst_registration_paths_get_by_id", id, &result)
	if err != nil {
		return MstRegistrationPathDetail{}, err
	}
	return result, nil
}

func CreateRegistrationPath(id string, code string, registration_path string, registration_type string, description string) error {
	return QueryInsertRegistrationPath(id, code, registration_path, registration_type, description)
}

func ImportRegistrationPaths(filePath string) error {
	headers := []string{
		"id", "code", "registration_path", "registration_type", "description",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		registration_path := row["registration_path"]
		registration_type := row["registration_type"]
		description := row["description"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstRegistrationPath{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateRegistrationPath(id, code, registration_path, registration_type, description); err != nil {
					return err
				}
			} else {
				if err := QueryInsertRegistrationPath(id, code, registration_path, registration_type, description); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstRegistrationPath{})
			if err != nil {
				return err
			}
			if err := QueryInsertRegistrationPath(id, code, registration_path, registration_type, description); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateRegistrationPath(id string, code string, registration_path string, registration_type string, description string) error {
	return QueryUpdateRegistrationPath(id, code, registration_path, registration_type, description)
}

func DeleteRegistrationPath(id string) error {
	err := handlers.SPDelete("sp_mst_registration_paths_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashRegistrationPaths(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstRegistrationPath, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_registration_paths_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstRegistrationPath{})
	if err != nil {
		return []MstRegistrationPath{}, 0, err
	}

	var modelResults []MstRegistrationPath
	for _, item := range results {
		level, ok := item.(*MstRegistrationPath)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreRegistrationPath(id string) error {
	err := handlers.SPRestore("sp_mst_registration_paths_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountRegistrationPaths(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstRegistrationPath{}, nullable)
}

/* Query */
func QueryInsertRegistrationPath(id string, code string, registration_path string, registration_type string, description string) error {
	query := `
		EXEC sp_mst_registration_paths_insert
		@id = ?,
		@code = ?,
		@registration_path= ?,
		@registration_type= ?,
		@description = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, code, registration_path, registration_type, description)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateRegistrationPath(id string, code string, registration_path string, registration_type string, description string) error {
	query := `
		EXEC sp_mst_registration_paths_update
		@id = ?,
		@code = ?,
		@registration_path= ?,
		@registration_type= ?,
		@description = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, code, registration_path, registration_type, description)
	if err != nil {
		return err
	}

	return nil
}

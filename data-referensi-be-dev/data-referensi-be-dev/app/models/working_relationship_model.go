package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstWorkingRelationship struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstWorkingRelationshipDetail struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstWorkingRelationshipExport struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type MstWorkingRelationshipSearch struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}
type MstWorkingRelationshipRelation struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

/* Action */
func GetWorkingRelationships(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstWorkingRelationship, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_working_relationships_get", filter, sortBy, sortDirection, page, pageSize, &MstWorkingRelationship{})
	if err != nil {
		return []MstWorkingRelationship{}, 0, err
	}

	var modelResults []MstWorkingRelationship
	for _, item := range results {
		level, ok := item.(*MstWorkingRelationship)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchWorkingRelationships(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstWorkingRelationshipSearch, error) {
	var results []MstWorkingRelationshipSearch
	err := handlers.SPGet("sp_mst_working_relationships_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstWorkingRelationshipSearch{}, err
	}
	return results, nil
}

func ExportWorkingRelationships(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstWorkingRelationshipExport
	err := handlers.SPGet("sp_mst_working_relationships_get", "", "name", "asc", 1, CountWorkingRelationships(), &results)
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

func GetWorkingRelationship(id string) (MstWorkingRelationshipDetail, error) {
	var result MstWorkingRelationshipDetail
	err := handlers.SPGetByID("sp_mst_working_relationships_get_by_id", id, &result)
	if err != nil {
		return MstWorkingRelationshipDetail{}, err
	}
	return result, nil
}

func CreateWorkingRelationship(id string, code string, name string) error {
	return QueryInsertWorkingRelationship(id, code, name)
}

func ImportWorkingRelationships(filePath string) error {
	headers := []string{
		"id", "code", "name",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		name := row["name"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstWorkingRelationship{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateWorkingRelationship(id, code, name); err != nil {
					return err
				}
			} else {
				if err := QueryInsertWorkingRelationship(id, code, name); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstWorkingRelationship{})
			if err != nil {
				return err
			}
			if err := QueryInsertWorkingRelationship(id, code, name); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateWorkingRelationship(id string, code string, name string) error {
	return QueryUpdateWorkingRelationship(id, code, name)
}

func DeleteWorkingRelationship(id string) error {
	err := handlers.SPDelete("sp_mst_working_relationships_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashWorkingRelationships(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstWorkingRelationship, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_working_relationships_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstWorkingRelationship{})
	if err != nil {
		return []MstWorkingRelationship{}, 0, err
	}

	var modelResults []MstWorkingRelationship
	for _, item := range results {
		level, ok := item.(*MstWorkingRelationship)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreWorkingRelationship(id string) error {
	err := handlers.SPRestore("sp_mst_working_relationships_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountWorkingRelationships(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstWorkingRelationship{}, nullable)
}

/* Query */
func QueryInsertWorkingRelationship(id string, code string, name string) error {
	query := `
		EXEC sp_mst_working_relationships_insert
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

func QueryUpdateWorkingRelationship(id string, code string, name string) error {
	query := `
		EXEC sp_mst_working_relationships_update
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

package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type MstPredicate struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstPredicateDetail struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstPredicateExport struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MstPredicateSearch struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MstPredicateRelation struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

/* Action */
func GetPredicates(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstPredicate, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_predicates_get", filter, sortBy, sortDirection, page, pageSize, &MstPredicate{})
	if err != nil {
		return []MstPredicate{}, 0, err
	}

	var modelResults []MstPredicate
	for _, item := range results {
		level, ok := item.(*MstPredicate)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchPredicates(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstPredicateSearch, error) {
	var results []MstPredicateSearch
	err := handlers.SPGet("sp_mst_predicates_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstPredicateSearch{}, err
	}
	return results, nil
}

func ExportPredicates(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstPredicateExport
	err := handlers.SPGet("sp_mst_predicates_get", "", "name", "asc", 1, CountPredicates(), &results)
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

func GetPredicate(id string) (MstPredicateDetail, error) {
	var result MstPredicateDetail
	err := handlers.SPGetByID("sp_mst_predicates_get_by_id", id, &result)
	if err != nil {
		return MstPredicateDetail{}, err
	}
	return result, nil
}

func CreatePredicate(id string, name string, description string) error {
	return QueryInsertPredicate(id, name, description)
}

func ImportPredicates(fileStatus string) error {
	headers := []string{
		"id", "name", "description",
	}

	return handlers.ModelImport(fileStatus, headers, func(row map[string]string) error {
		id := row["id"]
		name := row["name"]
		description := row["description"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstPredicate{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdatePredicate(id, name, description); err != nil {
					return err
				}
			} else {
				if err := QueryInsertPredicate(id, name, description); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstPredicate{})
			if err != nil {
				return err
			}
			if err := QueryInsertPredicate(id, name, description); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdatePredicate(id string, name string, description string) error {
	return QueryUpdatePredicate(id, name, description)
}

func DeletePredicate(id string) error {
	err := handlers.SPDelete("sp_mst_predicates_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashPredicates(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstPredicate, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_predicates_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstPredicate{})
	if err != nil {
		return []MstPredicate{}, 0, err
	}

	var modelResults []MstPredicate
	for _, item := range results {
		level, ok := item.(*MstPredicate)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestorePredicate(id string) error {
	err := handlers.SPRestore("sp_mst_predicates_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountPredicates(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstPredicate{}, nullable)
}

/* Query */
func QueryInsertPredicate(id string, name string, description string) error {
	query := `
		EXEC sp_mst_predicates_insert
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

func QueryUpdatePredicate(id string, name string, description string) error {
	query := `
		EXEC sp_mst_predicates_update
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

package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type MstFaculty struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Shortname   string `json:"shortname"`
	ChairName   string `json:"chair_name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstFacultyDetail struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Shortname   string `json:"shortname"`
	ChairName   string `json:"chair_name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstFacultyExport struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Shortname   string `json:"shortname"`
	ChairName   string `json:"chair_name"`
	Description string `json:"description"`
}

type MstFacultySearch struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Shortname   string `json:"shortname"`
	ChairName   string `json:"chair_name"`
	Description string `json:"description"`
}

type MstFacultyRelation struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Shortname   string `json:"shortname"`
	ChairName   string `json:"chair_name"`
	Description string `json:"description"`
}

/* Action */
func GetFaculties(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstFaculty, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_faculties_get", filter, sortBy, sortDirection, page, pageSize, &MstFaculty{})
	if err != nil {
		return []MstFaculty{}, 0, err
	}

	var modelResults []MstFaculty
	for _, item := range results {
		level, ok := item.(*MstFaculty)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchFaculties(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstFacultySearch, error) {
	var results []MstFacultySearch
	err := handlers.SPGet("sp_mst_faculties_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstFacultySearch{}, err
	}
	return results, nil
}

func ExportFaculties(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstFacultyExport
	err := handlers.SPGet("sp_mst_faculties_get", "", "name", "asc", 1, CountFaculties(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":          result.ID,
			"Name":        result.Name,
			"Shortname":   result.Shortname,
			"ChairName":   result.ChairName,
			"Description": result.Description,
		}
	}

	headers := []string{
		"ID", "Name", "Shortname", "ChairName", "Description",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetFaculty(id string) (MstFacultyDetail, error) {
	var result MstFacultyDetail
	err := handlers.SPGetByID("sp_mst_faculties_get_by_id", id, &result)
	if err != nil {
		return MstFacultyDetail{}, err
	}
	return result, nil
}

func CreateFaculty(id string, name string, shortname string, chair_name string, description string) error {
	return QueryInsertFaculty(id, name, shortname, chair_name, description)
}

func ImportFaculties(fileStatus string) error {
	headers := []string{
		"id", "name", "shortname", "chair_name", "description",
	}

	return handlers.ModelImport(fileStatus, headers, func(row map[string]string) error {
		id := row["id"]
		name := row["name"]
		shortname := row["shortname"]
		chair_name := row["chair_name"]
		description := row["description"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstFaculty{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateFaculty(id, name, shortname, chair_name, description); err != nil {
					return err
				}
			} else {
				if err := QueryInsertFaculty(id, name, shortname, chair_name, description); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstFaculty{})
			if err != nil {
				return err
			}
			if err := QueryInsertFaculty(id, name, shortname, chair_name, description); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateFaculty(id string, name string, shortname string, chair_name string, description string) error {
	return QueryUpdateFaculty(id, name, shortname, chair_name, description)
}

func DeleteFaculty(id string) error {
	err := handlers.SPDelete("sp_mst_faculties_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashFaculties(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstFaculty, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_faculties_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstFaculty{})
	if err != nil {
		return []MstFaculty{}, 0, err
	}

	var modelResults []MstFaculty
	for _, item := range results {
		level, ok := item.(*MstFaculty)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreFaculty(id string) error {
	err := handlers.SPRestore("sp_mst_faculties_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountFaculties(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstFaculty{}, nullable)
}

/* Query */
func QueryInsertFaculty(id string, name string, shortname string, chair_name string, description string) error {
	query := `
		EXEC sp_mst_faculties_insert
		@id = ?,
		@name = ?, 
		@shortname = ?, 
		@chair_name = ?, 
		@description = ?, 
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`
	log.Print(query)

	err := handlers.SPInsert(query, id, name, shortname, chair_name, description)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateFaculty(id string, name string, shortname string, chair_name string, description string) error {
	query := `
		EXEC sp_mst_faculties_update
		@id = ?,
		@name = ?, 
		@shortname = ?, 
		@chair_name = ?, 
		@description = ?, 
		@updated_at = ?,
		@updated_by = ?
	`
	err := handlers.SPUpdate(query, id, name, shortname, chair_name, description)
	if err != nil {
		return err
	}

	return nil
}

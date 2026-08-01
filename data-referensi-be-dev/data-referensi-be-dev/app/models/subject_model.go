package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MstSubject struct {
	ID           string `json:"id"`
	Code         string `json:"code" `
	Name         string `json:"name"`
	MinimumPoint string `json:"minimum_point"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
}

type MstSubjectDetail struct {
	ID           string `json:"id"`
	Code         string `json:"code" `
	Name         string `json:"name"`
	MinimumPoint string `json:"minimum_point"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
}

type MstSubjectExport struct {
	ID           string  `json:"id"`
	Code         string  `json:"code" `
	Name         string  `json:"name"`
	MinimumPoint float64 `json:"minimum_point"`
}

type MstSubjectSearch struct {
	ID           string  `json:"id"`
	Code         string  `json:"code" `
	Name         string  `json:"name"`
	MinimumPoint float64 `json:"minimum_point"`
}

type MstSubjectRelation struct {
	ID           string  `json:"id"`
	Code         string  `json:"code" `
	Name         string  `json:"name"`
	MinimumPoint float64 `json:"minimum_point"`
}

/* Action */
func GetSubjects(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstSubject, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_subjects_get", filter, sortBy, sortDirection, page, pageSize, &MstSubject{})
	if err != nil {
		return []MstSubject{}, 0, err
	}

	var modelResults []MstSubject
	for _, item := range results {
		level, ok := item.(*MstSubject)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchSubjects(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstSubjectSearch, error) {
	var results []MstSubjectSearch
	err := handlers.SPGet("sp_mst_subjects_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstSubjectSearch{}, err
	}
	return results, nil
}

func ExportSubjects(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstSubjectExport
	err := handlers.SPGet("sp_mst_subjects_get", "", "code", "asc", 1, CountSubjects(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":           result.ID,
			"Code":         result.Code,
			"Name":         result.Name,
			"MinimumPoint": result.MinimumPoint,
		}
	}

	headers := []string{
		"ID", "Code", "Name", "MinimumPoint",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetSubject(id string) (MstSubjectDetail, error) {
	var result MstSubjectDetail
	err := handlers.SPGetByID("sp_mst_subjects_get_by_id", id, &result)
	if err != nil {
		return MstSubjectDetail{}, err
	}
	return result, nil
}

func CreateSubject(id string, code string, name string, minimum_point float64) error {
	return QueryInsertSubject(id, code, name, minimum_point)
}

func ImportSubjects(filePath string) error {
	headers := []string{
		"id", "code", "name", "minimum_point",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		name := row["name"]
		minimum_point := row["minimum_point"]

		floatMinimumPoint, err := strconv.ParseFloat(minimum_point, 64)
		if err != nil {
			log.Fatalf("Gagal mengonversi string ke float64: %v", err)
		}

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstSubject{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateSubject(id, code, name, floatMinimumPoint); err != nil {
					return err
				}
			} else {
				if err := QueryInsertSubject(id, code, name, floatMinimumPoint); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstSubject{})
			if err != nil {
				return err
			}
			if err := QueryInsertSubject(id, code, name, floatMinimumPoint); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateSubject(id string, code string, name string, minimum_point float64) error {
	return QueryUpdateSubject(id, code, name, minimum_point)
}

func DeleteSubject(id string) error {
	err := handlers.SPDelete("sp_mst_subjects_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashSubjects(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstSubject, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_subjects_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstSubject{})
	if err != nil {
		return []MstSubject{}, 0, err
	}

	var modelResults []MstSubject
	for _, item := range results {
		level, ok := item.(*MstSubject)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreSubject(id string) error {
	err := handlers.SPRestore("sp_mst_subjects_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountSubjects(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstSubject{}, nullable)
}

/* Query */
func QueryInsertSubject(id string, code string, name string, minimum_point float64) error {
	query := `
		EXEC sp_mst_subjects_insert
		@id = ?,
		@code = ?,
		@name = ?,
		@minimum_point = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, code, name, minimum_point)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateSubject(id string, code string, name string, minimum_point float64) error {
	query := `
		EXEC sp_mst_subjects_update
		@id = ?,
		@code = ?,
		@name = ?,
		@minimum_point = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, code, name, minimum_point)
	if err != nil {
		return err
	}

	return nil
}

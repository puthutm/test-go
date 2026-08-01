package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstLectureSystem struct {
	ID            string `json:"id"`
	Code          string `json:"code"`
	LectureSystem string `json:"lecture_system"`
	Description   string `json:"description"`
	CreatedAt     int64  `json:"created_at"`
	UpdatedAt     int64  `json:"updated_at"`
}

type MstLectureSystemDetail struct {
	ID            string `json:"id"`
	Code          string `json:"code"`
	LectureSystem string `json:"lecture_system"`
	Description   string `json:"description"`
	CreatedAt     int64  `json:"created_at"`
	UpdatedAt     int64  `json:"updated_at"`
}

type MstLectureSystemExport struct {
	ID            string `json:"id"`
	Code          string `json:"code"`
	LectureSystem string `json:"lecture_system"`
	Description   string `json:"description"`
}

type MstLectureSystemSearch struct {
	ID            string `json:"id"`
	Code          string `json:"code"`
	LectureSystem string `json:"lecture_system"`
	Description   string `json:"description"`
}

type MstLectureSystemRelation struct {
	ID            string `json:"id"`
	Code          string `json:"code"`
	LectureSystem string `json:"lecture_system"`
	Description   string `json:"description"`
}

/* Action */
func GetLectureSystems(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstLectureSystem, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_lecture_systems_get", filter, sortBy, sortDirection, page, pageSize, &MstLectureSystem{})
	if err != nil {
		return []MstLectureSystem{}, 0, err
	}

	var modelResults []MstLectureSystem
	for _, item := range results {
		level, ok := item.(*MstLectureSystem)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchLectureSystems(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstLectureSystemSearch, error) {
	var results []MstLectureSystemSearch
	err := handlers.SPGet("sp_mst_lecture_systems_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstLectureSystemSearch{}, err
	}
	return results, nil
}

func ExportLectureSystems(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstLectureSystemExport
	err := handlers.SPGet("sp_mst_lecture_systems_get", "", "name", "asc", 1, CountLectureSystems(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":             result.ID,
			"Code":           result.Code,
			"Lecture System": result.LectureSystem,
			"Description":    result.Description,
		}
	}

	headers := []string{
		"ID", "Code", "Lecture System", "Description",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetLectureSystem(id string) (MstLectureSystemDetail, error) {
	var result MstLectureSystemDetail
	err := handlers.SPGetByID("sp_mst_lecture_systems_get_by_id", id, &result)
	if err != nil {
		return MstLectureSystemDetail{}, err
	}
	return result, nil
}

func CreateLectureSystem(id string, code string, lecture_system string, description string) error {
	return QueryInsertLectureSystem(id, code, lecture_system, description)
}

func ImportLectureSystems(filePath string) error {
	headers := []string{
		"id", "code", "lecture_system", "description",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		lecture_system := row["lecture_system"]
		description := row["description"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstLectureSystem{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateLectureSystem(id, code, lecture_system, description); err != nil {
					return err
				}
			} else {
				if err := QueryInsertLectureSystem(id, code, lecture_system, description); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstLectureSystem{})
			if err != nil {
				return err
			}
			if err := QueryInsertLectureSystem(id, code, lecture_system, description); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateLectureSystem(id string, code string, lecture_system string, description string) error {
	return QueryUpdateLectureSystem(id, code, lecture_system, description)
}

func DeleteLectureSystem(id string) error {
	err := handlers.SPDelete("sp_mst_lecture_systems_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashLectureSystems(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstLectureSystem, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_lecture_systems_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstLectureSystem{})
	if err != nil {
		return []MstLectureSystem{}, 0, err
	}

	var modelResults []MstLectureSystem
	for _, item := range results {
		level, ok := item.(*MstLectureSystem)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreLectureSystem(id string) error {
	err := handlers.SPRestore("sp_mst_lecture_systems_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountLectureSystems(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstLectureSystem{}, nullable)
}

/* Query */
func QueryInsertLectureSystem(id string, code string, lecture_system string, description string) error {
	query := `
		EXEC sp_mst_lecture_systems_insert
		@id = ?,
		@code = ?,
		@lecture_system= ?,
		@description = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, code, lecture_system, description)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateLectureSystem(id string, code string, lecture_system string, description string) error {
	query := `
		EXEC sp_mst_lecture_systems_update
		@id = ?,
		@code = ?,
		@lecture_system= ?,
		@description = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, code, lecture_system, description)
	if err != nil {
		return err
	}

	return nil
}

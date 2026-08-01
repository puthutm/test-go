package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MstProgramType struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	ProgramType string `json:"program_type"`
	Description string `json:"description"`
	IsIPC       bool   `json:"is_ipc"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstProgramTypeDetail struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	ProgramType string `json:"program_type"`
	Description string `json:"description"`
	IsIPC       bool   `json:"is_ipc"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstProgramTypeExport struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	ProgramType string `json:"program_type"`
	Description string `json:"description"`
	IsIPC       bool   `json:"is_ipc"`
}

type MstProgramTypeSearch struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	ProgramType string `json:"program_type"`
	Description string `json:"description"`
	IsIPC       bool   `json:"is_ipc"`
}

type MstProgramTypeRelation struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	ProgramType string `json:"program_type"`
	Description string `json:"description"`
	IsIPC       bool   `json:"is_ipc"`
}

/* Action */
func GetProgramTypes(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstProgramType, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_program_types_get", filter, sortBy, sortDirection, page, pageSize, &MstProgramType{})
	if err != nil {
		return []MstProgramType{}, 0, err
	}

	var modelResults []MstProgramType
	for _, item := range results {
		level, ok := item.(*MstProgramType)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchProgramTypes(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstProgramTypeSearch, error) {
	var results []MstProgramTypeSearch
	err := handlers.SPGet("sp_mst_program_types_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstProgramTypeSearch{}, err
	}
	return results, nil
}

func ExportProgramTypes(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstProgramTypeExport
	err := handlers.SPGet("sp_mst_program_types_get", "", "name", "asc", 1, CountProgramTypes(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":           result.ID,
			"Code":         result.Code,
			"Program Type": result.ProgramType,
			"Description":  result.Description,
			"Is IPC":       result.IsIPC,
		}
	}

	headers := []string{
		"ID", "Code", "Program Type", "Description", "Is IPC",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetProgramType(id string) (MstProgramTypeDetail, error) {
	var result MstProgramTypeDetail
	err := handlers.SPGetByID("sp_mst_program_types_get_by_id", id, &result)
	if err != nil {
		return MstProgramTypeDetail{}, err
	}
	return result, nil
}

func CreateProgramType(id string, code string, program_type string, description string, is_ipc bool) error {
	return QueryInsertProgramType(id, code, program_type, description, is_ipc)
}

func ImportProgramTypes(filePath string) error {
	headers := []string{
		"id", "code", "program_type", "description", "is_ipc",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		program_type := row["program_type"]
		description := row["description"]
		is_ipc, _ := strconv.ParseBool(row["is_ipc"])

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstProgramType{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateProgramType(id, code, program_type, description, is_ipc); err != nil {
					return err
				}
			} else {
				if err := QueryInsertProgramType(id, code, program_type, description, is_ipc); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstProgramType{})
			if err != nil {
				return err
			}
			if err := QueryInsertProgramType(id, code, program_type, description, is_ipc); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateProgramType(id string, code string, program_type string, description string, is_ipc bool) error {
	return QueryUpdateProgramType(id, code, program_type, description, is_ipc)
}

func DeleteProgramType(id string) error {
	err := handlers.SPDelete("sp_mst_program_types_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashProgramTypes(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstProgramType, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_program_types_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstProgramType{})
	if err != nil {
		return []MstProgramType{}, 0, err
	}

	var modelResults []MstProgramType
	for _, item := range results {
		level, ok := item.(*MstProgramType)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreProgramType(id string) error {
	err := handlers.SPRestore("sp_mst_program_types_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountProgramTypes(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstProgramType{}, nullable)
}

/* Query */
func QueryInsertProgramType(id string, code string, program_type string, description string, is_ipc bool) error {
	query := `
		EXEC sp_mst_program_types_insert
		@id = ?,
		@code = ?,
		@program_type= ?,
		@description = ?,
		@is_ipc = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, code, program_type, description, is_ipc)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateProgramType(id string, code string, program_type string, description string, is_ipc bool) error {
	query := `
		EXEC sp_mst_program_types_update
		@id = ?,
		@code = ?,
		@program_type= ?,
		@description = ?,
		@is_ipc = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, code, program_type, description, is_ipc)
	if err != nil {
		return err
	}

	return nil
}

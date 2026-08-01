package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstComposition struct {
	ID          string `json:"id"`
	Code        string `json:"code" `
	Composition string `json:"composition"`
	Note        string `json:"note"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstCompositionDetail struct {
	ID          string `json:"id"`
	Code        string `json:"code" `
	Composition string `json:"composition"`
	Note        string `json:"note"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstCompositionExport struct {
	ID          string `json:"id"`
	Code        string `json:"code" `
	Composition string `json:"composition"`
	Note        string `json:"note"`
}

type MstCompositionSearch struct {
	ID          string `json:"id"`
	Code        string `json:"code" `
	Composition string `json:"composition"`
	Note        string `json:"note"`
}

type MstCompositionRelation struct {
	ID          string `json:"id"`
	Code        string `json:"code" `
	Composition string `json:"composition"`
	Note        string `json:"note"`
}

/* Action */
func GetCompositions(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstComposition, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_compositions_get", filter, sortBy, sortDirection, page, pageSize, &MstComposition{})
	if err != nil {
		return []MstComposition{}, 0, err
	}

	var modelResults []MstComposition
	for _, item := range results {
		level, ok := item.(*MstComposition)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchCompositions(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstCompositionSearch, error) {
	var results []MstCompositionSearch
	err := handlers.SPGet("sp_mst_compositions_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstCompositionSearch{}, err
	}
	return results, nil
}

func ExportCompositions(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstCompositionExport
	err := handlers.SPGet("sp_mst_compositions_get", "", "code", "asc", 1, CountCompositions(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":          result.ID,
			"Code":        result.Code,
			"Composition": result.Composition,
			"Note":        result.Note,
		}
	}

	headers := []string{
		"ID", "Code", "Composition", "Note",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetComposition(id string) (MstCompositionDetail, error) {
	var result MstCompositionDetail
	err := handlers.SPGetByID("sp_mst_compositions_get_by_id", id, &result)
	if err != nil {
		return MstCompositionDetail{}, err
	}
	return result, nil
}

func CreateComposition(id string, code string, composition string, note string) error {
	return QueryInsertComposition(id, code, composition, note)
}

func ImportCompositions(filePath string) error {
	headers := []string{
		"id", "code", "composition", "note",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		composition := row["composition"]
		note := row["note"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstComposition{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateComposition(id, code, composition, note); err != nil {
					return err
				}
			} else {
				if err := QueryInsertComposition(id, code, composition, note); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstComposition{})
			if err != nil {
				return err
			}
			if err := QueryInsertComposition(id, code, composition, note); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateComposition(id string, code string, composition string, note string) error {
	return QueryUpdateComposition(id, code, composition, note)
}

func DeleteComposition(id string) error {
	err := handlers.SPDelete("sp_mst_compositions_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashCompositions(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstComposition, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_compositions_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstComposition{})
	if err != nil {
		return []MstComposition{}, 0, err
	}

	var modelResults []MstComposition
	for _, item := range results {
		level, ok := item.(*MstComposition)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreComposition(id string) error {
	err := handlers.SPRestore("sp_mst_compositions_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountCompositions(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstComposition{}, nullable)
}

/* Query */
func QueryInsertComposition(id string, code string, composition string, note string) error {
	query := `
		EXEC sp_mst_compositions_insert
		@id = ?,
		@code = ?,
		@composition = ?,
		@note = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, code, composition, note)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateComposition(id string, code string, composition string, note string) error {
	query := `
		EXEC sp_mst_compositions_update
		@id = ?,
		@code = ?,
		@composition = ?,
		@note = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, code, composition, note)
	if err != nil {
		return err
	}

	return nil
}

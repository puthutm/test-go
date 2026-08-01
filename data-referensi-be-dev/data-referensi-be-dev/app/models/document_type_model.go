package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MstDocumentType struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Mimes     string `json:"mimes"`
	Size      int    `json:"size"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstDocumentTypeDetail struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Mimes     string `json:"mimes"`
	Size      int    `json:"size"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstDocumentTypeExport struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Mimes string `json:"mimes"`
	Size  int    `json:"size"`
}

type MstDocumentTypeSearch struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Mimes string `json:"mimes"`
}

type MstDocumentTypeRelation struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Mimes string `json:"mimes"`
}

/* Action */
func GetDocumentTypes(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstDocumentType, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_document_types_get", filter, sortBy, sortDirection, page, pageSize, &MstDocumentType{})
	if err != nil {
		return []MstDocumentType{}, 0, err
	}

	var modelResults []MstDocumentType
	for _, item := range results {
		level, ok := item.(*MstDocumentType)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchDocumentTypes(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstDocumentTypeSearch, error) {
	var results []MstDocumentTypeSearch
	err := handlers.SPGet("sp_mst_document_types_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstDocumentTypeSearch{}, err
	}
	return results, nil
}

func ExportDocumentTypes(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstDocumentTypeExport
	err := handlers.SPGet("sp_mst_document_types_get", "", "name", "asc", 1, CountDocumentTypes(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":    result.ID,
			"Name":  result.Name,
			"Mimes": result.Mimes,
			"Size":  result.Size,
		}
	}

	headers := []string{
		"ID", "Name", "Mimes", "Size",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetDocumentType(id string) (MstDocumentTypeDetail, error) {
	var result MstDocumentTypeDetail
	err := handlers.SPGetByID("sp_mst_document_types_get_by_id", id, &result)
	if err != nil {
		return MstDocumentTypeDetail{}, err
	}
	return result, nil
}

func CreateDocumentType(id string, name string, mimes string, size int) error {
	return QueryInsertDocumentType(id, name, mimes, size)
}

func ImportDocumentTypes(filePath string) error {
	headers := []string{
		"id", "name", "mimes", "size",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		name := row["name"]
		mimes := row["mimes"]
		size, _ := strconv.Atoi(row["size"])

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstDocumentType{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateDocumentType(id, name, mimes, size); err != nil {
					return err
				}
			} else {
				if err := QueryInsertDocumentType(id, name, mimes, size); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstDocumentType{})
			if err != nil {
				return err
			}
			if err := QueryInsertDocumentType(id, name, mimes, size); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateDocumentType(id string, name string, mimes string, size int) error {
	return QueryUpdateDocumentType(id, name, mimes, size)
}

func DeleteDocumentType(id string) error {
	err := handlers.SPDelete("sp_mst_document_types_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashDocumentTypes(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstDocumentType, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_document_types_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstDocumentType{})
	if err != nil {
		return []MstDocumentType{}, 0, err
	}

	var modelResults []MstDocumentType
	for _, item := range results {
		level, ok := item.(*MstDocumentType)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreDocumentType(id string) error {
	err := handlers.SPRestore("sp_mst_document_types_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountDocumentTypes(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstDocumentType{}, nullable)
}

/* Query */
func QueryInsertDocumentType(id string, name string, mimes string, size int) error {
	query := `
		EXEC sp_mst_document_types_insert
		@id = ?,
		@name = ?,
		@mimes= ?,
		@size = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, name, mimes, size)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateDocumentType(id string, name string, mimes string, size int) error {
	query := `
		EXEC sp_mst_document_types_update
		@id = ?,
		@name = ?,
		@mimes= ?,
		@size = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, name, mimes, size)
	if err != nil {
		return err
	}

	return nil
}

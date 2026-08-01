package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstEthnic struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	RegionOfOrigin string `json:"region_of_origin"`
	CreatedAt      int64  `json:"created_at"`
	UpdatedAt      int64  `json:"updated_at"`
}

type MstEthnicDetail struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	RegionOfOrigin string `json:"region_of_origin"`
	CreatedAt      int64  `json:"created_at"`
	UpdatedAt      int64  `json:"updated_at"`
}

type MstEthnicExport struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	RegionOfOrigin string `json:"region_of_origin"`
}

type MstEthnicSearch struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	RegionOfOrigin string `json:"region_of_origin"`
}

type MstEthnicRelation struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	RegionOfOrigin string `json:"region_of_origin"`
}

/* Action */
func GetEthnics(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstEthnic, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_ethnics_get", filter, sortBy, sortDirection, page, pageSize, &MstEthnic{})
	if err != nil {
		return []MstEthnic{}, 0, err
	}

	var modelResults []MstEthnic
	for _, item := range results {
		level, ok := item.(*MstEthnic)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchEthnics(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstEthnicSearch, error) {
	var results []MstEthnicSearch
	err := handlers.SPGet("sp_mst_ethnics_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstEthnicSearch{}, err
	}
	return results, nil
}

func ExportEthnics(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstEthnicExport
	err := handlers.SPGet("sp_mst_ethnics_get", "", "name", "asc", 1, CountEthnics(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":               result.ID,
			"Name":             result.Name,
			"Region Of Origin": result.RegionOfOrigin,
		}
	}

	headers := []string{
		"ID",
		"Name",
		"Region Of Origin",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetEthnic(id string) (MstEthnicDetail, error) {
	var result MstEthnicDetail
	err := handlers.SPGetByID("sp_mst_ethnics_get_by_id", id, &result)
	if err != nil {
		return MstEthnicDetail{}, err
	}
	return result, nil
}

func CreateEthnic(id string, name string, region_of_origin string) error {
	return QueryInsertEthnic(id, name, region_of_origin)
}

func ImportEthnics(filePath string) error {
	headers := []string{
		"id", "name", "region_of_origin",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		name := row["name"]
		region_of_origin := row["region_of_origin"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstEthnic{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateEthnic(id, name, region_of_origin); err != nil {
					return err
				}
			} else {
				if err := QueryInsertEthnic(id, name, region_of_origin); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstEthnic{})
			if err != nil {
				return err
			}
			if err := QueryInsertEthnic(id, name, region_of_origin); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateEthnic(id string, name string, region_of_origin string) error {
	return QueryUpdateEthnic(id, name, region_of_origin)
}

func DeleteEthnic(id string) error {
	err := handlers.SPDelete("sp_mst_ethnics_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashEthnics(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstEthnic, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_ethnics_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstEthnic{})
	if err != nil {
		return []MstEthnic{}, 0, err
	}

	var modelResults []MstEthnic
	for _, item := range results {
		level, ok := item.(*MstEthnic)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreEthnic(id string) error {
	err := handlers.SPRestore("sp_mst_ethnics_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountEthnics(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstEthnic{}, nullable)
}

/* Query */
func QueryInsertEthnic(id string, name string, region_of_origin string) error {
	query := `
		EXEC sp_mst_ethnics_insert
		@id = ?,
		@name = ?,
		@region_of_origin = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, name, region_of_origin)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateEthnic(id string, name string, region_of_origin string) error {
	query := `
		EXEC sp_mst_ethnics_update
		@id = ?,
		@name = ?,
		@region_of_origin = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, name, region_of_origin)
	if err != nil {
		return err
	}

	return nil
}

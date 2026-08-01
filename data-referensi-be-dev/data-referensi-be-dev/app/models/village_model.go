package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstVillage struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Code         string `json:"code"`
	DistrictName string `json:"district_name"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
}

type MstVillageDetail struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Code         string `json:"code"`
	DistrictId   string `json:"district_id"`
	DistrictName string `json:"district_name"`
	DistrictCode string `json:"district_code"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
}

type MstVillageExport struct {
	ID           string `json:"id"`
	DistrictId   string `json:"district_id"`
	DistrictName string `json:"district_name"`
	Name         string `json:"name"`
	Code         string `json:"code"`
}

type MstVillageSearch struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	DistrictId string `json:"district_id"`
}

type MstVillageRelation struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

/* Action */
func GetVillages(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstVillage, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_villages_get", filter, sortBy, sortDirection, page, pageSize, &MstVillage{})
	if err != nil {
		return []MstVillage{}, 0, err
	}

	var modelResults []MstVillage
	for _, item := range results {
		level, ok := item.(*MstVillage)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchVillages(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstVillageSearch, error) {
	var results []MstVillageSearch
	err := handlers.SPGet("sp_mst_villages_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstVillageSearch{}, err
	}
	return results, nil
}

func ExportVillages(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstVillageExport
	err := handlers.SPGet("sp_mst_villages_get", "", "name", "asc", 1, CountVillages(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":            result.ID,
			"District ID":   result.DistrictId,
			"District Name": result.DistrictName,
			"Name":          result.Name,
			"Code":          result.Code,
		}
	}

	headers := []string{
		"ID",
		"District ID",
		"District Name",
		"Name",
		"Code",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetVillage(id string) (MstVillageDetail, error) {
	var result MstVillageDetail
	err := handlers.SPGetByID("sp_mst_villages_get_by_id", id, &result)
	if err != nil {
		return MstVillageDetail{}, err
	}
	return result, nil
}

func GetVillageByDistrictId(district_id string) ([]MstVillageSearch, error) {
	var results []MstVillageSearch

	query := fmt.Sprintf(`
		EXEC sp_mst_villages_get_by_district_id
		@district_id = '%s'
	`, district_id)

	err := handlers.SPGetByQuery(query, &results)
	if err != nil {
		return []MstVillageSearch{}, err
	}
	return results, nil
}

func CreateVillage(id string, district_id string, name string, code string) error {
	return QueryInsertVillage(id, district_id, name, code)
}

func ImportVillages(filePath string) error {
	headers := []string{
		"id", "district_id", "district_name", "name", "code",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		district_id := row["district_id"]
		name := row["name"]
		code := row["code"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstVillage{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateVillage(id, district_id, name, code); err != nil {
					return err
				}
			} else {
				if err := QueryInsertVillage(id, district_id, name, code); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstVillage{})
			if err != nil {
				return err
			}
			if err := QueryInsertVillage(id, district_id, name, code); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateVillage(id string, district_id string, name string, code string) error {
	return QueryUpdateVillage(id, district_id, name, code)
}

func DeleteVillage(id string) error {
	err := handlers.SPDelete("sp_mst_villages_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashVillages(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstVillage, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_villages_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstVillage{})
	if err != nil {
		return []MstVillage{}, 0, err
	}

	var modelResults []MstVillage
	for _, item := range results {
		level, ok := item.(*MstVillage)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreVillage(id string) error {
	err := handlers.SPRestore("sp_mst_villages_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountVillages(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstVillage{}, nullable)
}

/* Query */
func QueryInsertVillage(id string, district_id string, name string, code string) error {
	query := `
		EXEC sp_mst_villages_insert
		@id = ?,
		@district_id = ?,
		@name = ?,
		@code= ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, district_id, name, code)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateVillage(id string, district_id string, name string, code string) error {
	query := `
		EXEC sp_mst_villages_update
		@id = ?,
		@district_id = ?,
		@name = ?,
		@code= ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, district_id, name, code)
	if err != nil {
		return err
	}

	return nil
}

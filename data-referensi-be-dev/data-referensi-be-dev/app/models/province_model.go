package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstProvince struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	RegionCode  string `json:"region_code"`
	CountryName string `json:"country_name"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstProvinceDetail struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	RegionCode  string `json:"region_code"`
	CountryId   string `json:"country_id"`
	CountryName string `json:"country_name"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstProvinceExport struct {
	ID          string `json:"id"`
	CountryId   string `json:"country_id"`
	CountryName string `json:"country_name"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	RegionCode  string `json:"region_code"`
}

type MstProvinceSearch struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	CountryId string `json:"country_id"`
}

type MstProvinceRelation struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

/* Action */
func GetProvinces(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstProvince, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_provinces_get", filter, sortBy, sortDirection, page, pageSize, &MstProvince{})
	if err != nil {
		return []MstProvince{}, 0, err
	}

	var modelResults []MstProvince
	for _, item := range results {
		level, ok := item.(*MstProvince)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchProvinces(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstProvinceSearch, error) {
	var results []MstProvinceSearch
	err := handlers.SPGet("sp_mst_provinces_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstProvinceSearch{}, err
	}
	return results, nil
}

func ExportProvinces(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstProvinceExport
	err := handlers.SPGet("sp_mst_provinces_get", "", "name", "asc", 1, CountProvinces(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":           result.ID,
			"Country ID":   result.CountryId,
			"Country Name": result.CountryName,
			"Name":         result.Name,
			"Code":         result.Code,
			"Region Code":  result.RegionCode,
		}
	}

	headers := []string{
		"ID",
		"Country ID",
		"Country Name",
		"Name",
		"Code",
		"Region Code",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetProvince(id string) (MstProvinceDetail, error) {
	var result MstProvinceDetail
	err := handlers.SPGetByID("sp_mst_provinces_get_by_id", id, &result)
	if err != nil {
		return MstProvinceDetail{}, err
	}
	return result, nil
}

func GetProvinceByCountryId(country_id string) ([]MstProvinceSearch, error) {
	var results []MstProvinceSearch

	query := fmt.Sprintf(`
		EXEC sp_mst_provinces_get_by_country_id
		@country_id = '%s'
	`, country_id)

	err := handlers.SPGetByQuery(query, &results)
	if err != nil {
		return []MstProvinceSearch{}, err
	}
	return results, nil
}

func CreateProvince(id string, country_id string, name string, code string, region_code string) error {
	return QueryInsertProvince(id, country_id, name, code, region_code)
}

func ImportProvinces(filePath string) error {
	headers := []string{
		"id", "country_id", "country_name", "name", "code", "region_code",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		country_id := row["country_id"]
		name := row["name"]
		code := row["code"]
		region_code := row["region_code"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstProvince{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateProvince(id, country_id, name, code, region_code); err != nil {
					return err
				}
			} else {
				if err := QueryInsertProvince(id, country_id, name, code, region_code); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstProvince{})
			if err != nil {
				return err
			}
			if err := QueryInsertProvince(id, country_id, name, code, region_code); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateProvince(id string, country_id string, name string, code string, region_code string) error {
	return QueryUpdateProvince(id, country_id, name, code, region_code)
}

func DeleteProvince(id string) error {
	err := handlers.SPDelete("sp_mst_provinces_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashProvinces(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstProvince, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_provinces_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstProvince{})
	if err != nil {
		return []MstProvince{}, 0, err
	}

	var modelResults []MstProvince
	for _, item := range results {
		level, ok := item.(*MstProvince)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreProvince(id string) error {
	err := handlers.SPRestore("sp_mst_provinces_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountProvinces(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstProvince{}, nullable)
}

/* Query */
func QueryInsertProvince(id string, country_id string, name string, code string, region_code string) error {
	query := `
		EXEC sp_mst_provinces_insert
		@id = ?,
		@country_id = ?,
		@name = ?,
		@code= ?,
		@region_code = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, country_id, name, code, region_code)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateProvince(id string, country_id string, name string, code string, region_code string) error {
	query := `
		EXEC sp_mst_provinces_update
		@id = ?,
		@country_id = ?,
		@name = ?,
		@code= ?,
		@region_code = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, country_id, name, code, region_code)
	if err != nil {
		return err
	}

	return nil
}

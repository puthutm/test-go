package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstCity struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Code         string `json:"code"`
	ProvinceName string `json:"province_name"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
}

type MstCityDetail struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Code         string `json:"code"`
	ProvinceId   string `json:"province_id"`
	ProvinceName string `json:"province_name"`
	ProvinceCode string `json:"province_code"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
}

type MstCityExport struct {
	ID           string `json:"id"`
	ProvinceId   string `json:"province_id"`
	ProvinceName string `json:"province_name"`
	Name         string `json:"name"`
	Code         string `json:"code"`
}

type MstCitySearch struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	ProvinceId string `json:"province_id"`
}

type MstCityRelation struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

/* Action */
func GetCities(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstCity, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_cities_get", filter, sortBy, sortDirection, page, pageSize, &MstCity{})
	if err != nil {
		return []MstCity{}, 0, err
	}

	var modelResults []MstCity
	for _, item := range results {
		level, ok := item.(*MstCity)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchCities(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstCitySearch, error) {
	var results []MstCitySearch
	err := handlers.SPGet("sp_mst_cities_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstCitySearch{}, err
	}
	return results, nil
}

func ExportCities(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstCityExport
	err := handlers.SPGet("sp_mst_cities_get", "", "name", "asc", 1, CountCities(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":            result.ID,
			"Province ID":   result.ProvinceId,
			"Province Name": result.ProvinceName,
			"Name":          result.Name,
			"Code":          result.Code,
		}
	}

	headers := []string{
		"ID",
		"Province ID",
		"Province Name",
		"Name",
		"Code",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetCity(id string) (MstCityDetail, error) {
	var result MstCityDetail
	err := handlers.SPGetByID("sp_mst_cities_get_by_id", id, &result)
	if err != nil {
		return MstCityDetail{}, err
	}
	return result, nil
}

func GetCityByProvinceId(province_id string) ([]MstCitySearch, error) {
	var results []MstCitySearch

	query := fmt.Sprintf(`
		EXEC sp_mst_cities_get_by_province_id
		@province_id = '%s'
	`, province_id)

	err := handlers.SPGetByQuery(query, &results)
	if err != nil {
		return []MstCitySearch{}, err
	}
	return results, nil
}

func CreateCity(id string, province_id string, name string, code string) error {
	return QueryInsertCity(id, province_id, name, code)
}

func ImportCities(filePath string) error {
	headers := []string{
		"id", "province_id", "province_name", "name", "code",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		province_id := row["province_id"]
		name := row["name"]
		code := row["code"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstCity{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateCity(id, province_id, name, code); err != nil {
					return err
				}
			} else {
				if err := QueryInsertCity(id, province_id, name, code); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstCity{})
			if err != nil {
				return err
			}
			if err := QueryInsertCity(id, province_id, name, code); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateCity(id string, province_id string, name string, code string) error {
	return QueryUpdateCity(id, province_id, name, code)
}

func DeleteCity(id string) error {
	err := handlers.SPDelete("sp_mst_cities_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashCities(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstCity, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_cities_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstCity{})
	if err != nil {
		return []MstCity{}, 0, err
	}

	var modelResults []MstCity
	for _, item := range results {
		level, ok := item.(*MstCity)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreCity(id string) error {
	err := handlers.SPRestore("sp_mst_cities_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountCities(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstCity{}, nullable)
}

/* Query */
func QueryInsertCity(id string, province_id string, name string, code string) error {
	query := `
		EXEC sp_mst_cities_insert
		@id = ?,
		@province_id = ?,
		@name = ?,
		@code= ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, province_id, name, code)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateCity(id string, province_id string, name string, code string) error {
	query := `
		EXEC sp_mst_cities_update
		@id = ?,
		@province_id = ?,
		@name = ?,
		@code= ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, province_id, name, code)
	if err != nil {
		return err
	}

	return nil
}

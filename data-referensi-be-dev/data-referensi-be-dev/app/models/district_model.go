package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstDistrict struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	CityName  string `json:"city_name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstDistrictDetail struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	CityId    string `json:"city_id"`
	CityName  string `json:"city_name"`
	CityCode  string `json:"city_code"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstDistrictExport struct {
	ID       string `json:"id"`
	CityId   string `json:"city_id"`
	CityName string `json:"city_name"`
	Name     string `json:"name"`
	Code     string `json:"code"`
}

type MstDistrictSearch struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	CityId string `json:"city_id"`
}

type MstDistrictRelation struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

/* Action */
func GetDistricts(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstDistrict, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_districts_get", filter, sortBy, sortDirection, page, pageSize, &MstDistrict{})
	if err != nil {
		return []MstDistrict{}, 0, err
	}

	var modelResults []MstDistrict
	for _, item := range results {
		level, ok := item.(*MstDistrict)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchDistricts(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstDistrictSearch, error) {
	var results []MstDistrictSearch
	err := handlers.SPGet("sp_mst_districts_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstDistrictSearch{}, err
	}
	return results, nil
}

func ExportDistricts(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstDistrictExport
	err := handlers.SPGet("sp_mst_districts_get", "", "name", "asc", 1, CountDistricts(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":        result.ID,
			"City ID":   result.CityId,
			"City Name": result.CityName,
			"Name":      result.Name,
			"Code":      result.Code,
		}
	}

	headers := []string{
		"ID",
		"City ID",
		"City Name",
		"Name",
		"Code",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetDistrict(id string) (MstDistrictDetail, error) {
	var result MstDistrictDetail
	err := handlers.SPGetByID("sp_mst_districts_get_by_id", id, &result)
	if err != nil {
		return MstDistrictDetail{}, err
	}
	return result, nil
}

func GetDistrictByCityId(city_id string) ([]MstDistrictSearch, error) {
	var results []MstDistrictSearch

	query := fmt.Sprintf(`
		EXEC sp_mst_districts_get_by_city_id
		@city_id = '%s'
	`, city_id)

	err := handlers.SPGetByQuery(query, &results)
	if err != nil {
		return []MstDistrictSearch{}, err
	}
	return results, nil
}

func CreateDistrict(id string, city_id string, name string, code string) error {
	return QueryInsertDistrict(id, city_id, name, code)
}

func ImportDistricts(filePath string) error {
	headers := []string{
		"id", "city_id", "city_name", "name", "code",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		city_id := row["city_id"]
		name := row["name"]
		code := row["code"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstDistrict{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateDistrict(id, city_id, name, code); err != nil {
					return err
				}
			} else {
				if err := QueryInsertDistrict(id, city_id, name, code); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstDistrict{})
			if err != nil {
				return err
			}
			if err := QueryInsertDistrict(id, city_id, name, code); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateDistrict(id string, city_id string, name string, code string) error {
	return QueryUpdateDistrict(id, city_id, name, code)
}

func DeleteDistrict(id string) error {
	err := handlers.SPDelete("sp_mst_districts_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashDistricts(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstDistrict, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_districts_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstDistrict{})
	if err != nil {
		return []MstDistrict{}, 0, err
	}

	var modelResults []MstDistrict
	for _, item := range results {
		level, ok := item.(*MstDistrict)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreDistrict(id string) error {
	err := handlers.SPRestore("sp_mst_districts_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountDistricts(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstDistrict{}, nullable)
}

/* Query */
func QueryInsertDistrict(id string, city_id string, name string, code string) error {
	query := `
		EXEC sp_mst_districts_insert
		@id = ?,
		@city_id = ?,
		@name = ?,
		@code= ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, city_id, name, code)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateDistrict(id string, city_id string, name string, code string) error {
	query := `
		EXEC sp_mst_districts_update
		@id = ?,
		@city_id = ?,
		@name = ?,
		@code= ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, city_id, name, code)
	if err != nil {
		return err
	}

	return nil
}

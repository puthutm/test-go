package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstSchool struct {
	ID            string `json:"id"`
	Npsn          string `json:"npsn"`
	Name          string `json:"name"`
	EducationForm string `json:"education_form"`
	Status        string `json:"status"`
	ProvinceName  string `json:"province_name"`
	CityName      string `json:"city_name"`
	DistrictName  string `json:"district_name"`
	CreatedAt     int64  `json:"created_at"`
	UpdatedAt     int64  `json:"updated_at"`
}

type MstSchoolDetail struct {
	ID            string `json:"id"`
	Npsn          string `json:"npsn"`
	Name          string `json:"name"`
	EducationForm string `json:"education_form"`
	Status        string `json:"status"`
	ProvinceId    string `json:"province_id"`
	ProvinceCode  string `json:"province_code"`
	ProvinceName  string `json:"province_name"`
	CityId        string `json:"city_id"`
	CityCode      string `json:"city_code"`
	CityName      string `json:"city_name"`
	DistrictId    string `json:"district_id"`
	DistrictCode  string `json:"district_code"`
	DistrictName  string `json:"district_name"`
	CreatedAt     int64  `json:"created_at"`
	UpdatedAt     int64  `json:"updated_at"`
}

type MstSchoolExport struct {
	ID            string `json:"id"`
	Npsn          string `json:"npsn"`
	Name          string `json:"name"`
	EducationForm string `json:"education_form"`
	Status        string `json:"status"`
	ProvinceId    string `json:"province_id"`
	ProvinceName  string `json:"province_name"`
	CityId        string `json:"city_id"`
	CityName      string `json:"city_name"`
	DistrictId    string `json:"district_id"`
	DistrictName  string `json:"district_name"`
}

type MstSchoolSearch struct {
	ID   string `json:"id"`
	Npsn string `json:"npsn"`
	Name string `json:"name"`
}
type MstSchoolRelation struct {
	ID            string `json:"id"`
	Npsn          string `json:"npsn"`
	Name          string `json:"name"`
	EducationForm string `json:"education_form"`
	Status        string `json:"status"`
	ProvinceName  string `json:"province_name"`
	CityName      string `json:"city_name"`
	DistrictName  string `json:"district_name"`
}

/* Action */
func GetSchools(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstSchool, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_schools_get", filter, sortBy, sortDirection, page, pageSize, &MstSchool{})
	if err != nil {
		return []MstSchool{}, 0, err
	}

	var modelResults []MstSchool
	for _, item := range results {
		level, ok := item.(*MstSchool)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchSchools(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstSchoolSearch, error) {
	var results []MstSchoolSearch
	err := handlers.SPGet("sp_mst_schools_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstSchoolSearch{}, err
	}
	return results, nil
}

func ExportSchools(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstSchoolExport
	err := handlers.SPGet("sp_mst_schools_get", "", "name", "asc", 1, CountSchools(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":            result.ID,
			"NPSN":          result.Npsn,
			"Name":          result.Name,
			"School Form":   result.EducationForm,
			"Status":        result.Status,
			"Province ID":   result.ProvinceId,
			"Province Name": result.ProvinceName,
			"City ID":       result.CityId,
			"City Name":     result.CityName,
			"District ID":   result.DistrictId,
			"District Name": result.DistrictName,
		}
	}

	headers := []string{
		"ID",
		"NPSN",
		"Name",
		"School Form",
		"Status",
		"Province ID",
		"Province Name",
		"City ID",
		"City Name",
		"District ID",
		"District Name",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetSchool(id string) (MstSchoolDetail, error) {
	var result MstSchoolDetail
	err := handlers.SPGetByID("sp_mst_schools_get_by_id", id, &result)
	if err != nil {
		return MstSchoolDetail{}, err
	}
	return result, nil
}

func CreateSchool(id string, npsn string, name string, education_form string, status string, province_id string, city_id string, district_id string) error {
	return QueryInsertSchool(id, npsn, name, education_form, status, province_id, city_id, district_id)
}

func ImportSchools(filePath string) error {
	headers := []string{
		"id",
		"npsn",
		"name",
		"education_form",
		"status",
		"province_id",
		"province_name",
		"city_id",
		"city_name",
		"district_id",
		"district_name",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		npsn := row["npsn"]
		name := row["name"]
		education_form := row["education_form"]
		status := row["status"]
		province_id := row["province_id"]
		city_id := row["city_id"]
		district_id := row["district_id"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstSchool{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateSchool(id, npsn, name, education_form, status, province_id, city_id, district_id); err != nil {
					return err
				}
			} else {
				if err := QueryInsertSchool(id, npsn, name, education_form, status, province_id, city_id, district_id); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstSchool{})
			if err != nil {
				return err
			}
			if err := QueryInsertSchool(id, npsn, name, education_form, status, province_id, city_id, district_id); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateSchool(id string, npsn string, name string, education_form string, status string, province_id string, city_id string, district_id string) error {
	return QueryUpdateSchool(id, npsn, name, education_form, status, province_id, city_id, district_id)
}

func DeleteSchool(id string) error {
	err := handlers.SPDelete("sp_mst_schools_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashSchools(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstSchool, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_schools_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstSchool{})
	if err != nil {
		return []MstSchool{}, 0, err
	}

	var modelResults []MstSchool
	for _, item := range results {
		level, ok := item.(*MstSchool)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreSchool(id string) error {
	err := handlers.SPRestore("sp_mst_schools_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountSchools(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstSchool{}, nullable)
}

/* Query */
func QueryInsertSchool(id string, npsn string, name string, education_form string, status string, province_id string, city_id string, district_id string) error {
	query := `
		EXEC sp_mst_schools_insert
		@id = ?,
		@npsn = ?,
		@name = ?,
		@education_form = ?,
		@status = ?,
		@province_id = ?,
		@city_id = ?,
		@district_id = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, npsn, name, education_form, status, province_id, city_id, district_id)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateSchool(id string, npsn string, name string, education_form string, status string, province_id string, city_id string, district_id string) error {
	query := `
		EXEC sp_mst_schools_update
		@id = ?,
		@npsn = ?,
		@name = ?,
		@education_form = ?,
		@status = ?,
		@province_id = ?,
		@city_id = ?,
		@district_id = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, npsn, name, education_form, status, province_id, city_id, district_id)
	if err != nil {
		return err
	}

	return nil
}

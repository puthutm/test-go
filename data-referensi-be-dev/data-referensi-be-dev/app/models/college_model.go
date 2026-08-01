package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MstCollege struct {
	ID                   string  `json:"id"`
	Name                 string  `json:"name"`
	ProvinceName         string  `json:"province_name"`
	CityName             string  `json:"city_name"`
	Type                 string  `json:"type"`
	Accreditation        string  `json:"accreditation"`
	ShortName            string  `json:"short_name"`
	NumberOfStudyProgram int     `json:"number_of_study_program"`
	LowerLimitTuitionFee float64 `json:"lower_limit_tuition_fee"`
	UpperLimitTuitionFee float64 `json:"upper_limit_tuition_fee"`
	CreatedAt            int64   `json:"created_at"`
	UpdatedAt            int64   `json:"updated_at"`
}

type MstCollegeDetail struct {
	ID                   string  `json:"id"`
	Name                 string  `json:"name"`
	ProvinceId           string  `json:"province_id"`
	ProvinceName         string  `json:"province_name"`
	ProvinceCode         string  `json:"province_code"`
	CityId               string  `json:"city_id"`
	CityName             string  `json:"city_name"`
	CityCode             string  `json:"city_code"`
	Type                 string  `json:"type"`
	Accreditation        string  `json:"accreditation"`
	ShortName            string  `json:"short_name"`
	NumberOfStudyProgram int     `json:"number_of_study_program"`
	LowerLimitTuitionFee float64 `json:"lower_limit_tuition_fee"`
	UpperLimitTuitionFee float64 `json:"upper_limit_tuition_fee"`
	CreatedAt            int64   `json:"created_at"`
	UpdatedAt            int64   `json:"updated_at"`
}

type MstCollegeExport struct {
	ID                   string  `json:"id"`
	Name                 string  `json:"name"`
	ProvinceId           string  `json:"province_id"`
	ProvinceName         string  `json:"province_name"`
	CityId               string  `json:"city_id"`
	CityName             string  `json:"city_name"`
	Type                 string  `json:"type"`
	Accreditation        string  `json:"accreditation"`
	ShortName            string  `json:"short_name"`
	NumberOfStudyProgram int     `json:"number_of_study_program"`
	LowerLimitTuitionFee float64 `json:"lower_limit_tuition_fee"`
	UpperLimitTuitionFee float64 `json:"upper_limit_tuition_fee"`
}

type MstCollegeSearch struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
}

type MstCollegeRelation struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	ProvinceName  string `json:"province_name"`
	CityName      string `json:"city_name"`
	Type          string `json:"type"`
	Accreditation string `json:"accreditation"`
	ShortName     string `json:"short_name"`
}

/* Action */
func GetColleges(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstCollege, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_colleges_get", filter, sortBy, sortDirection, page, pageSize, &MstCollege{})
	if err != nil {
		return []MstCollege{}, 0, err
	}

	var modelResults []MstCollege
	for _, item := range results {
		level, ok := item.(*MstCollege)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchColleges(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstCollegeSearch, error) {
	var results []MstCollegeSearch
	err := handlers.SPGet("sp_mst_colleges_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstCollegeSearch{}, err
	}
	return results, nil
}

func ExportColleges(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstCollegeExport
	err := handlers.SPGet("sp_mst_colleges_get", "", "name", "asc", 1, CountColleges(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":                      result.ID,
			"Name":                    result.Name,
			"Province ID":             result.ProvinceId,
			"Province Name":           result.ProvinceName,
			"City ID":                 result.CityId,
			"City Name":               result.CityName,
			"Type":                    result.Type,
			"Accreditation":           result.Accreditation,
			"Short Name":              result.ShortName,
			"Number Of Study Program": result.NumberOfStudyProgram,
			"Lower Limit Tuition Fee": result.LowerLimitTuitionFee,
			"Upper Limit Tuition Fee": result.UpperLimitTuitionFee,
		}
	}

	headers := []string{
		"ID", "Name", "Province ID", "Province Name", "City ID", "City Name", "Type",
		"Accreditation", "Short Name", "Number Of Study Program", "Lower Limit Tuition Fee",
		"Upper Limit Tuition Fee",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetCollege(id string) (MstCollegeDetail, error) {
	var result MstCollegeDetail
	err := handlers.SPGetByID("sp_mst_colleges_get_by_id", id, &result)
	if err != nil {
		return MstCollegeDetail{}, err
	}
	return result, nil
}

func CreateCollege(id string, name string, provinceId string, cityId string, typeCollege string, accreditation string, shortName string, numberOfStudyProgram int, lowerLimitTuitionFee float64, upperLimitTuitionFee float64) error {
	return QueryInsertCollege(id, name, provinceId, cityId, typeCollege, accreditation, shortName, numberOfStudyProgram, lowerLimitTuitionFee, upperLimitTuitionFee)
}

func ImportColleges(filePath string) error {
	headers := []string{
		"id", "name", "province_id", "province_name", "city_id", "city_name", "type", "accreditation", "short_name",
		"number_of_study_program", "lower_limit_tuition_fee", "upper_limit_tuition_fee",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		name := row["name"]
		provinceId := row["province_id"]
		cityId := row["city_id"]
		typeCollege := row["type"]
		accreditation := row["accreditation"]
		shortName := row["short_name"]

		numberOfStudyProgram, _ := strconv.Atoi(row["number_of_study_program"])
		lowerLimitTuitionFee, _ := strconv.ParseFloat(row["lower_limit_tuition_fee"], 64)
		upperLimitTuitionFee, _ := strconv.ParseFloat(row["upper_limit_tuition_fee"], 64)

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstCollege{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateCollege(id, name, provinceId, cityId, typeCollege, accreditation, shortName, numberOfStudyProgram, lowerLimitTuitionFee, upperLimitTuitionFee); err != nil {
					return err
				}
			} else {
				if err := QueryInsertCollege(id, name, provinceId, cityId, typeCollege, accreditation, shortName, numberOfStudyProgram, lowerLimitTuitionFee, upperLimitTuitionFee); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstCollege{})
			if err != nil {
				return err
			}
			if err := QueryInsertCollege(id, name, provinceId, cityId, typeCollege, accreditation, shortName, numberOfStudyProgram, lowerLimitTuitionFee, upperLimitTuitionFee); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateCollege(id string, name string, provinceId string, cityId string, typeCollege string, accreditation string, shortName string, numberOfStudyProgram int, lowerLimitTuitionFee float64, upperLimitTuitionFee float64) error {
	return QueryUpdateCollege(id, name, provinceId, cityId, typeCollege, accreditation, shortName, numberOfStudyProgram, lowerLimitTuitionFee, upperLimitTuitionFee)
}

func DeleteCollege(id string) error {
	err := handlers.SPDelete("sp_mst_colleges_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashColleges(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstCollege, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_colleges_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstCollege{})
	if err != nil {
		return []MstCollege{}, 0, err
	}

	var modelResults []MstCollege
	for _, item := range results {
		level, ok := item.(*MstCollege)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreCollege(id string) error {
	err := handlers.SPRestore("sp_mst_colleges_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountColleges(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstCollege{}, nullable)
}

/* Query */
func QueryInsertCollege(id string, name string, provinceId string, cityId string, typeCollege string, accreditation string, shortName string, numberOfStudyProgram int, lowerLimitTuitionFee float64, upperLimitTuitionFee float64) error {
	query := `
		EXEC sp_mst_colleges_insert
		@id = ?,
		@name = ?,
		@province_id = ?,
		@city_id = ?,
		@type = ?,
		@accreditation = ?,
		@short_name = ?,
		@number_of_study_program = ?,
		@lower_limit_tuition_fee = ?,
		@upper_limit_tuition_fee = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, name, provinceId, cityId, typeCollege, accreditation, shortName, numberOfStudyProgram, lowerLimitTuitionFee, upperLimitTuitionFee)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateCollege(id string, name string, provinceId string, cityId string, typeCollege string, accreditation string, shortName string, numberOfStudyProgram int, lowerLimitTuitionFee float64, upperLimitTuitionFee float64) error {
	query := `
		EXEC sp_mst_colleges_update
		@id = ?,
		@name = ?,
		@province_id = ?,
		@city_id = ?,
		@type = ?,
		@accreditation = ?,
		@short_name = ?,
		@number_of_study_program = ?,
		@lower_limit_tuition_fee = ?,
		@upper_limit_tuition_fee = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, name, provinceId, cityId, typeCollege, accreditation, shortName, numberOfStudyProgram, lowerLimitTuitionFee, upperLimitTuitionFee)
	if err != nil {
		return err
	}

	return nil
}

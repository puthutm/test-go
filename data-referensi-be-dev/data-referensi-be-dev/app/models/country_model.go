package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type MstCountry struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	PhoneCode    string `json:"phone_code"`
	IconFlagPath string `json:"icon_flag_path"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
}

type MstCountryDetail struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	PhoneCode    string `json:"phone_code"`
	IconFlagPath string `json:"icon_flag_path"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
}

type MstCountryExport struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	PhoneCode    string `json:"phone_code"`
	IconFlagPath string `json:"icon_flag_path"`
}

type MstCountrySearch struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type MstCountryRelation struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	PhoneCode    string `json:"phone_code"`
	IconFlagPath string `json:"icon_flag_path"`
}

/* Action */
func GetCountries(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstCountry, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_countries_get", filter, sortBy, sortDirection, page, pageSize, &MstCountry{})
	if err != nil {
		log.Print(err)
		return []MstCountry{}, 0, err
	}
	var modelResults []MstCountry
	for _, item := range results {
		level, ok := item.(*MstCountry)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}
	return modelResults, total, nil
}

func SearchCountries(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstCountrySearch, error) {
	var results []MstCountrySearch
	err := handlers.SPGet("sp_mst_countries_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstCountrySearch{}, err
	}
	return results, nil
}

func ExportCountries(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstCountryExport
	err := handlers.SPGet("sp_mst_countries_get", "", "name", "asc", 1, CountCountries(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":             result.ID,
			"Name":           result.Name,
			"Phone Code":     result.PhoneCode,
			"Icon Flag Path": result.IconFlagPath,
		}
	}

	headers := []string{
		"ID", "Name", "Phone Code", "Icon Flag Path",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetCountry(id string) (MstCountryDetail, error) {
	var result MstCountryDetail
	err := handlers.SPGetByID("sp_mst_countries_get_by_id", id, &result)
	if err != nil {
		return MstCountryDetail{}, err
	}
	return result, nil
}

func CreateCountry(id string, name string, phone_code string, icon_flag_path string) error {
	return QueryInsertCountry(id, name, phone_code, icon_flag_path)
}

func ImportCountries(filePath string) error {
	headers := []string{
		"id", "name", "phone_code", "icon_flag_path",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		name := row["name"]
		phone_code := row["phone_code"]
		icon_flag_path := row["icon_flag_path"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstCountry{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateCountry(id, name, phone_code, icon_flag_path); err != nil {
					return err
				}
			} else {
				if err := QueryInsertCountry(id, name, phone_code, icon_flag_path); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstCountry{})
			if err != nil {
				return err
			}
			if err := QueryInsertCountry(id, name, phone_code, icon_flag_path); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateCountry(id string, name string, phone_code string, icon_flag_path string) error {
	return QueryUpdateCountry(id, name, phone_code, icon_flag_path)
}

func DeleteCountry(id string) error {
	err := handlers.SPDelete("sp_mst_countries_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashCountries(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstCountry, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_countries_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstCountry{})
	if err != nil {
		return []MstCountry{}, 0, err
	}
	var modelResults []MstCountry
	for _, item := range results {
		level, ok := item.(*MstCountry)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}
	return modelResults, total, nil
}

func RestoreCountry(id string) error {
	err := handlers.SPRestore("sp_mst_countries_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountCountries(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstCountry{}, nullable)
}

/* Query */
func QueryInsertCountry(id string, name string, phone_code string, icon_flag_path string) error {
	query := `
		EXEC sp_mst_countries_insert
		@id = ?,
		@name = ?,
		@phone_code= ?,
		@icon_flag_path = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, name, phone_code, icon_flag_path)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateCountry(id string, name string, phone_code string, icon_flag_path string) error {
	query := `
		EXEC sp_mst_countries_update
		@id = ?,
		@name = ?,
		@phone_code= ?,
		@icon_flag_path = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, name, phone_code, icon_flag_path)
	if err != nil {
		return err
	}

	return nil
}

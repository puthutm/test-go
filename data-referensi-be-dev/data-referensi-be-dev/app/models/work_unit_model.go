package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstWorkUnit struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	Acronym   string `json:"acronym"`
	Fullname  string `json:"fullname"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstWorkUnitDetail struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	Acronym   string `json:"acronym"`
	Fullname  string `json:"fullname"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstWorkUnitExport struct {
	ID       string `json:"id"`
	Code     string `json:"code"`
	Acronym  string `json:"acronym"`
	Fullname string `json:"fullname"`
}

type MstWorkUnitSearch struct {
	ID       string `json:"id"`
	Code     string `json:"code"`
	Acronym  string `json:"acronym"`
	Fullname string `json:"fullname"`
}

type MstWorkUnitRelation struct {
	ID       string `json:"id"`
	Code     string `json:"code"`
	Acronym  string `json:"acronym"`
	Fullname string `json:"fullname"`
}

/* Action */
func GetWorkUnits(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstWorkUnit, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_work_units_get", filter, sortBy, sortDirection, page, pageSize, &MstWorkUnit{})
	if err != nil {
		return []MstWorkUnit{}, 0, err
	}

	var modelResults []MstWorkUnit
	for _, item := range results {
		level, ok := item.(*MstWorkUnit)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchWorkUnits(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstWorkUnitSearch, error) {
	var results []MstWorkUnitSearch
	err := handlers.SPGet("sp_mst_work_units_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstWorkUnitSearch{}, err
	}
	return results, nil
}

func ExportWorkUnits(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstWorkUnitExport
	err := handlers.SPGet("sp_mst_work_units_get", "", "name", "asc", 1, CountWorkUnits(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":       result.ID,
			"Code":     result.Code,
			"Acronym":  result.Acronym,
			"Fullname": result.Fullname,
		}
	}

	headers := []string{
		"ID", "Code", "Acronym", "Fullname",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetWorkUnit(id string) (MstWorkUnitDetail, error) {
	var result MstWorkUnitDetail
	err := handlers.SPGetByID("sp_mst_work_units_get_by_id", id, &result)
	if err != nil {
		return MstWorkUnitDetail{}, err
	}
	return result, nil
}

func CreateWorkUnit(id string, code string, acronym string, fullname string) error {
	return QueryInsertWorkUnit(id, code, acronym, fullname)
}

func ImportWorkUnits(filePath string) error {
	headers := []string{
		"id", "code", "acronym", "fullname",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		acronym := row["acronym"]
		fullname := row["fullname"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstWorkUnit{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateWorkUnit(id, code, acronym, fullname); err != nil {
					return err
				}
			} else {
				if err := QueryInsertWorkUnit(id, code, acronym, fullname); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstWorkUnit{})
			if err != nil {
				return err
			}
			if err := QueryInsertWorkUnit(id, code, acronym, fullname); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateWorkUnit(id string, code string, acronym string, fullname string) error {
	return QueryUpdateWorkUnit(id, code, acronym, fullname)
}

func DeleteWorkUnit(id string) error {
	err := handlers.SPDelete("sp_mst_work_units_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashWorkUnits(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstWorkUnit, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_work_units_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstWorkUnit{})
	if err != nil {
		return []MstWorkUnit{}, 0, err
	}

	var modelResults []MstWorkUnit
	for _, item := range results {
		level, ok := item.(*MstWorkUnit)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreWorkUnit(id string) error {
	err := handlers.SPRestore("sp_mst_work_units_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountWorkUnits(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstWorkUnit{}, nullable)
}

/* Query */
func QueryInsertWorkUnit(id string, code string, acronym string, fullname string) error {
	query := `
		EXEC sp_mst_work_units_insert
		@id = ?,
		@code = ?,
		@acronym = ?,
		@fullname = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, code, acronym, fullname)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateWorkUnit(id string, code string, acronym string, fullname string) error {
	query := `
		EXEC sp_mst_work_units_update
		@id = ?,
		@code = ?,
		@acronym = ?,
		@fullname = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, code, acronym, fullname)
	if err != nil {
		return err
	}

	return nil
}

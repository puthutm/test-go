package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MstSemester struct {
	ID         string `json:"id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	UTSSession string `json:"uts_session" gorm:"uts_session"`
	UASSession string `json:"uas_session"  gorm:"uas_session"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}

type MstSemesterDetail struct {
	ID         string `json:"id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	UTSSession string `json:"uts_session" gorm:"uts_session"`
	UASSession string `json:"uas_session"  gorm:"uas_session"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}

type MstSemesterExport struct {
	ID         string `json:"id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	UTSSession string `json:"uts_session" gorm:"uts_session"`
	UASSession string `json:"uas_session"  gorm:"uas_session"`
}

type MstSemesterSearch struct {
	ID         string `json:"id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	UTSSession string `json:"uts_session" gorm:"uts_session"`
	UASSession string `json:"uas_session"  gorm:"uas_session"`
}

type MstSemesterRelation struct {
	ID         string `json:"id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	UTSSession string `json:"uts_session" gorm:"uts_session"`
	UASSession string `json:"uas_session"  gorm:"uas_session"`
}

/* Action */
func GetSemesters(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstSemester, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_semesters_get", filter, sortBy, sortDirection, page, pageSize, &MstSemester{})
	if err != nil {
		return []MstSemester{}, 0, err
	}

	var modelResults []MstSemester
	for _, item := range results {
		level, ok := item.(*MstSemester)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchSemesters(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstSemesterSearch, error) {
	var results []MstSemesterSearch
	err := handlers.SPGet("sp_mst_semesters_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstSemesterSearch{}, err
	}
	return results, nil
}

func ExportSemesters(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstSemesterExport
	err := handlers.SPGet("sp_mst_semesters_get", "", "name", "asc", 1, CountSemesters(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":         result.ID,
			"Code":       result.Code,
			"Name":       result.Name,
			"UTSSession": result.UTSSession,
			"UASSession": result.UASSession,
		}
	}

	headers := []string{
		"ID", "Code", "Name", "UTS Session", "UAS Session",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetSemester(id string) (MstSemesterDetail, error) {
	var result MstSemesterDetail
	err := handlers.SPGetByID("sp_mst_semesters_get_by_id", id, &result)
	if err != nil {
		return MstSemesterDetail{}, err
	}
	return result, nil
}

func CreateSemester(id string, code string, name string, UTSSession string, UASSession string) error {
	return QueryInsertSemester(id, code, name, UTSSession, UASSession)
}

func ImportSemesters(fileStatus string) error {
	headers := []string{
		"id", "code", "name", "uts_session", "uas_session",
	}

	return handlers.ModelImport(fileStatus, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		name := row["name"]
		uts_session := row["uts_session"]
		uas_session := row["uas_session"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstSemester{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateSemester(id, code, name, uts_session, uas_session); err != nil {
					return err
				}
			} else {
				if err := QueryInsertSemester(id, code, name, uts_session, uas_session); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstSemester{})
			if err != nil {
				return err
			}
			if err := QueryInsertSemester(id, code, name, uts_session, uas_session); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateSemester(id string, code string, name string, UTSSession string, UASSession string) error {
	return QueryUpdateSemester(id, code, name, UTSSession, UASSession)
}

func DeleteSemester(id string) error {
	err := handlers.SPDelete("sp_mst_semesters_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashSemesters(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstSemester, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_semesters_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstSemester{})
	if err != nil {
		return []MstSemester{}, 0, err
	}

	var modelResults []MstSemester
	for _, item := range results {
		level, ok := item.(*MstSemester)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreSemester(id string) error {
	err := handlers.SPRestore("sp_mst_semesters_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountSemesters(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstSemester{}, nullable)
}

/* Query */
func QueryInsertSemester(id string, code string, name string, UTSSession string, UASSession string) error {
	query := `
		EXEC sp_mst_semesters_insert
		@id = ?,
		@code = ?,
		@name = ?,
		@uts_session = ?,
		@uas_session = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, code, name, UTSSession, UASSession)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateSemester(id string, code string, name string, UTSSession string, UASSession string) error {
	query := `
		EXEC sp_mst_semesters_update
		@id = ?,
		@code = ?,
		@name = ?,
		@uts_session = ?,
		@uas_session = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, code, name, UTSSession, UASSession)
	if err != nil {
		return err
	}

	return nil
}

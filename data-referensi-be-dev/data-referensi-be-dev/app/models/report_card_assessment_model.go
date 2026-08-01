package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MstReportCardAssessment struct {
	ID          string `json:"id"`
	Code        string `json:"code" `
	Name        string `json:"name"`
	Value       string `json:"value"`
	SubjectId   string `json:"subject_id"`
	SubjectName string `json:"subject_name"`
	Note        string `json:"note"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstReportCardAssessmentDetail struct {
	ID          string `json:"id"`
	Code        string `json:"code" `
	Name        string `json:"name"`
	Value       string `json:"value"`
	SubjectId   string `json:"subject_id"`
	SubjectName string `json:"subject_name"`
	Note        string `json:"note"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstReportCardAssessmentExport struct {
	ID          string  `json:"id"`
	Code        string  `json:"code" `
	Name        string  `json:"name"`
	Value       float64 `json:"value"`
	SubjectId   string  `json:"subject_id"`
	SubjectName string  `json:"subject_name"`
	Note        string  `json:"note"`
}

type MstReportCardAssessmentSearch struct {
	ID          string  `json:"id"`
	Code        string  `json:"code" `
	Name        string  `json:"name"`
	Value       float64 `json:"value"`
	SubjectId   string  `json:"subject_id"`
	SubjectName string  `json:"subject_name"`
	Note        string  `json:"note"`
}

type MstReportCardAssessmentRelation struct {
	ID          string  `json:"id"`
	Code        string  `json:"code" `
	Name        string  `json:"name"`
	Value       float64 `json:"value"`
	SubjectId   string  `json:"subject_id"`
	SubjectName string  `json:"subject_name"`
	Note        string  `json:"note"`
}

/* Action */
func GetReportCardAssessments(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstReportCardAssessment, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_report_card_assessments_get", filter, sortBy, sortDirection, page, pageSize, &MstReportCardAssessment{})
	if err != nil {
		return []MstReportCardAssessment{}, 0, err
	}

	var modelResults []MstReportCardAssessment
	for _, item := range results {
		level, ok := item.(*MstReportCardAssessment)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchReportCardAssessments(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstReportCardAssessmentSearch, error) {
	var results []MstReportCardAssessmentSearch
	err := handlers.SPGet("sp_mst_report_card_assessments_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstReportCardAssessmentSearch{}, err
	}
	return results, nil
}

func ExportReportCardAssessments(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstReportCardAssessmentExport
	err := handlers.SPGet("sp_mst_report_card_assessments_get", "", "code", "asc", 1, CountReportCardAssessments(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":          result.ID,
			"Code":        result.Code,
			"Name":        result.Name,
			"Value":       result.Value,
			"SubjectId":   result.SubjectId,
			"SubjectName": result.SubjectName,
			"Note":        result.Note,
		}
	}

	headers := []string{
		"ID", "Code", "Name", "Value", "Subject ID", "Subject Name", "Note",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetReportCardAssessment(id string) (MstReportCardAssessmentDetail, error) {
	var result MstReportCardAssessmentDetail
	err := handlers.SPGetByID("sp_mst_report_card_assessments_get_by_id", id, &result)
	if err != nil {
		return MstReportCardAssessmentDetail{}, err
	}
	return result, nil
}

func CreateReportCardAssessment(id string, code string, name string, value float64, subject_id string, note string) error {
	return QueryInsertReportCardAssessment(id, code, name, value, subject_id, note)
}

func ImportReportCardAssessments(filePath string) error {
	headers := []string{
		"id", "code", "name", "value", "subject_id", "subject_name", "note",
	}

	return handlers.ModelImport(filePath, headers, func(row map[string]string) error {
		id := row["id"]
		code := row["code"]
		name := row["name"]
		value := row["value"]
		subject_id := row["subject_id"]
		note := row["note"]

		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Fatalf("Gagal mengonversi string ke float64: %v", err)
		}

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstReportCardAssessment{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateReportCardAssessment(id, code, name, floatValue, subject_id, note); err != nil {
					return err
				}
			} else {
				if err := QueryInsertReportCardAssessment(id, code, name, floatValue, subject_id, note); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstReportCardAssessment{})
			if err != nil {
				return err
			}
			if err := QueryInsertReportCardAssessment(id, code, name, floatValue, subject_id, note); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateReportCardAssessment(id string, code string, name string, value float64, subject_id string, note string) error {
	return QueryUpdateReportCardAssessment(id, code, name, value, subject_id, note)
}

func DeleteReportCardAssessment(id string) error {
	err := handlers.SPDelete("sp_mst_report_card_assessments_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashReportCardAssessments(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstReportCardAssessment, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_report_card_assessments_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstReportCardAssessment{})
	if err != nil {
		return []MstReportCardAssessment{}, 0, err
	}

	var modelResults []MstReportCardAssessment
	for _, item := range results {
		level, ok := item.(*MstReportCardAssessment)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreReportCardAssessment(id string) error {
	err := handlers.SPRestore("sp_mst_report_card_assessments_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountReportCardAssessments(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstReportCardAssessment{}, nullable)
}

/* Query */
func QueryInsertReportCardAssessment(id string, code string, name string, value float64, subject_id string, note string) error {
	query := `
		EXEC sp_mst_report_card_assessments_insert
		@id = ?,
		@code = ?,
		@name = ?,
		@value = ?,
		@subject_id = ?,
		@note = ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, code, name, value, subject_id, note)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateReportCardAssessment(id string, code string, name string, value float64, subject_id string, note string) error {
	query := `
		EXEC sp_mst_report_card_assessments_update
		@id = ?,
		@code = ?,
		@name = ?,
		@value = ?,
		@subject_id = ?,
		@note = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, code, name, value, subject_id, note)
	if err != nil {
		return err
	}

	return nil
}

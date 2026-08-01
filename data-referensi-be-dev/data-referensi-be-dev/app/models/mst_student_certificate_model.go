package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type MstStudentCertificate struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Purposes    string `json:"purposes"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstStudentCertificateDetail struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Purposes    string `json:"purposes"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type MstStudentCertificateExport struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Purposes    string `json:"purposes"`
}

type MstStudentCertificateSearch struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Purposes    string `json:"purposes"`
}

type MstStudentCertificateRelation struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Purposes    string `json:"purposes"`
}

/* Action */
func GetStudentCertificates(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstStudentCertificate, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_student_certificates_get", filter, sortBy, sortDirection, page, pageSize, &MstStudentCertificate{})
	if err != nil {
		return []MstStudentCertificate{}, 0, err
	}

	var modelResults []MstStudentCertificate
	for _, item := range results {
		level, ok := item.(*MstStudentCertificate)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchStudentCertificates(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstStudentCertificateSearch, error) {
	var results []MstStudentCertificateSearch
	err := handlers.SPGet("sp_mst_student_certificates_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstStudentCertificateSearch{}, err
	}
	return results, nil
}

func ExportStudentCertificates(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstStudentCertificateExport
	err := handlers.SPGet("sp_mst_student_certificates_get", "", "name", "asc", 1, CountStudentCertificates(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":          result.ID,
			"Name":        result.Name,
			"Description": result.Description,
			"Purposes":    result.Purposes,
		}
	}

	headers := []string{
		"ID", "Name", "Description", "Purposes",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetStudentCertificate(id string) (MstStudentCertificateDetail, error) {
	var result MstStudentCertificateDetail
	err := handlers.SPGetByID("sp_mst_student_certificates_get_by_id", id, &result)
	if err != nil {
		return MstStudentCertificateDetail{}, err
	}
	return result, nil
}

func CreateStudentCertificate(id string, name string, description string, purposes string) error {
	return QueryInsertStudentCertificate(id, name, description, purposes)
}

func ImportStudentCertificates(fileStatus string) error {
	headers := []string{
		"id", "name", "description", "purposes",
	}

	return handlers.ModelImport(fileStatus, headers, func(row map[string]string) error {
		id := row["id"]
		name := row["name"]
		description := row["description"]
		purposes := row["purposes"]

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstStudentCertificate{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateStudentCertificate(id, name, description, purposes); err != nil {
					return err
				}
			} else {
				if err := QueryInsertStudentCertificate(id, name, description, purposes); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstStudentCertificate{})
			if err != nil {
				return err
			}
			if err := QueryInsertStudentCertificate(id, name, description, purposes); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateStudentCertificate(id string, name string, description string, purposes string) error {
	return QueryUpdateStudentCertificate(id, name, description, purposes)
}

func DeleteStudentCertificate(id string) error {
	err := handlers.SPDelete("sp_mst_student_certificates_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashStudentCertificates(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstStudentCertificate, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_student_certificates_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstStudentCertificate{})
	if err != nil {
		return []MstStudentCertificate{}, 0, err
	}

	var modelResults []MstStudentCertificate
	for _, item := range results {
		level, ok := item.(*MstStudentCertificate)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreStudentCertificate(id string) error {
	err := handlers.SPRestore("sp_mst_student_certificates_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountStudentCertificates(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstStudentCertificate{}, nullable)
}

/* Query */
func QueryInsertStudentCertificate(id string, name string, description string, purposes string) error {
	query := `
		EXEC sp_mst_student_certificates_insert
		@id = ?,
		@name = ?, 
		@description = ?, 
		@purposes = ?, 
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`
	log.Print(query)

	err := handlers.SPInsert(query, id, name, description, purposes)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateStudentCertificate(id string, name string, description string, purposes string) error {
	query := `
		EXEC sp_mst_student_certificates_update
		@id = ?,
		@name = ?, 
		@description = ?, 
		@purposes = ?, 
		@updated_at = ?,
		@updated_by = ?
	`
	err := handlers.SPUpdate(query, id, name, description, purposes)
	if err != nil {
		return err
	}

	return nil
}

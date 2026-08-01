package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MstRegistrationStatus struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	IsDefault      bool   `json:"is_default"`
	AcceptedMarker bool   `json:"accepted_marker"`
	CreatedAt      int64  `json:"created_at"`
	UpdatedAt      int64  `json:"updated_at"`
}

type MstRegistrationStatusDetail struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	IsDefault      bool   `json:"is_default"`
	AcceptedMarker bool   `json:"accepted_marker"`
	CreatedAt      int64  `json:"created_at"`
	UpdatedAt      int64  `json:"updated_at"`
}

type MstRegistrationStatusExport struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	IsDefault      bool   `json:"is_default"`
	AcceptedMarker bool   `json:"accepted_marker"`
}

type MstRegistrationStatusSearch struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	IsDefault      bool   `json:"is_default"`
	AcceptedMarker bool   `json:"accepted_marker"`
}

type MstRegistrationStatusRelation struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	IsDefault      bool   `json:"is_default"`
	AcceptedMarker bool   `json:"accepted_marker"`
}

/* Action */
func GetRegistrationStatuses(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstRegistrationStatus, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_registration_statuses_get", filter, sortBy, sortDirection, page, pageSize, &MstRegistrationStatus{})
	if err != nil {
		return []MstRegistrationStatus{}, 0, err
	}

	var modelResults []MstRegistrationStatus
	for _, item := range results {
		level, ok := item.(*MstRegistrationStatus)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchRegistrationStatuses(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstRegistrationStatusSearch, error) {
	var results []MstRegistrationStatusSearch
	err := handlers.SPGet("sp_mst_registration_statuses_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstRegistrationStatusSearch{}, err
	}
	return results, nil
}

func ExportRegistrationStatuses(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstRegistrationStatusExport
	err := handlers.SPGet("sp_mst_registration_statuses_get", "", "name", "asc", 1, CountRegistrationStatuses(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":             result.ID,
			"Name":           result.Name,
			"IsDefault":      result.IsDefault,
			"AcceptedMarker": result.AcceptedMarker,
		}
	}

	headers := []string{
		"ID", "Name", "Is Default", "Accepted Marker",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetRegistrationStatus(id string) (MstRegistrationStatusDetail, error) {
	var result MstRegistrationStatusDetail
	err := handlers.SPGetByID("sp_mst_registration_statuses_get_by_id", id, &result)
	if err != nil {
		return MstRegistrationStatusDetail{}, err
	}
	return result, nil
}

func CreateRegistrationStatus(id string, name string, is_default bool, accepted_marker bool) error {
	return QueryInsertRegistrationStatus(id, name, is_default, accepted_marker)
}

func ImportRegistrationStatuses(fileStatus string) error {
	headers := []string{
		"id", "name", "is_default", "accepted_marker",
	}

	return handlers.ModelImport(fileStatus, headers, func(row map[string]string) error {
		id := row["id"]
		name := row["name"]
		is_default, err := strconv.ParseBool(row["is_default"])
		if err != nil {
			return err
		}

		accepted_marker, err := strconv.ParseBool(row["accepted_marker"])
		if err != nil {
			return err
		}

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstRegistrationStatus{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateRegistrationStatus(id, name, is_default, accepted_marker); err != nil {
					return err
				}
			} else {
				if err := QueryInsertRegistrationStatus(id, name, is_default, accepted_marker); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstRegistrationStatus{})
			if err != nil {
				return err
			}
			if err := QueryInsertRegistrationStatus(id, name, is_default, accepted_marker); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateRegistrationStatus(id string, name string, is_default bool, accepted_marker bool) error {
	return QueryUpdateRegistrationStatus(id, name, is_default, accepted_marker)
}

func DeleteRegistrationStatus(id string) error {
	err := handlers.SPDelete("sp_mst_registration_statuses_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashRegistrationStatuses(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstRegistrationStatus, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_registration_statuses_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstRegistrationStatus{})
	if err != nil {
		return []MstRegistrationStatus{}, 0, err
	}

	var modelResults []MstRegistrationStatus
	for _, item := range results {
		level, ok := item.(*MstRegistrationStatus)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreRegistrationStatus(id string) error {
	err := handlers.SPRestore("sp_mst_registration_statuses_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountRegistrationStatuses(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstRegistrationStatus{}, nullable)
}

/* Query */
func QueryInsertRegistrationStatus(id string, name string, is_default bool, accepted_marker bool) error {
	query := `
		EXEC sp_mst_registration_statuses_insert
		@id = ?,
		@name = ?,
		@is_default= ?,
		@accepted_marker= ?,
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPInsert(query, id, name, is_default, accepted_marker)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateRegistrationStatus(id string, name string, is_default bool, accepted_marker bool) error {
	query := `
		EXEC sp_mst_registration_statuses_update
		@id = ?,
		@name = ?,
		@is_default= ?,
		@accepted_marker= ?,
		@updated_at = ?,
		@updated_by = ?
	`

	err := handlers.SPUpdate(query, id, name, is_default, accepted_marker)
	if err != nil {
		return err
	}

	return nil
}

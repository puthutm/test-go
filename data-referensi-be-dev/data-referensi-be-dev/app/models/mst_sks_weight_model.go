package models

import (
	"data-referensi/handlers"
	"data-referensi/helpers"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MstSksWeight struct {
	ID        string `json:"id"`
	Amount    int    `json:"amount"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstSksWeightDetail struct {
	ID        string `json:"id"`
	Amount    int    `json:"amount"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type MstSksWeightExport struct {
	ID     string `json:"id"`
	Amount int    `json:"amount"`
}

type MstSksWeightSearch struct {
	ID     string `json:"id"`
	Amount int    `json:"amount"`
}

type MstSksWeightRelation struct {
	ID     string `json:"id"`
	Amount int    `json:"amount"`
}

/* Action */
func GetSksWeights(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstSksWeight, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_sks_weights_get", filter, sortBy, sortDirection, page, pageSize, &MstSksWeight{})
	if err != nil {
		return []MstSksWeight{}, 0, err
	}

	var modelResults []MstSksWeight
	for _, item := range results {
		level, ok := item.(*MstSksWeight)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func SearchSksWeights(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstSksWeightSearch, error) {
	var results []MstSksWeightSearch
	err := handlers.SPGet("sp_mst_sks_weights_get", filter, sortBy, sortDirection, page, pageSize, &results)
	if err != nil {
		return []MstSksWeightSearch{}, err
	}
	return results, nil
}

func ExportSksWeights(c *fiber.Ctx, fileSaveAs string) error {
	var results []MstSksWeightExport
	err := handlers.SPGet("sp_mst_sks_weights_get", "", "amount", "asc", 1, CountSksWeights(), &results)
	if err != nil {
		return fmt.Errorf("failed to get results: %v", err)
	}

	data := make([]map[string]interface{}, len(results))
	for i, result := range results {
		data[i] = map[string]interface{}{
			"ID":     result.ID,
			"Amount": result.Amount,
		}
	}

	headers := []string{
		"ID", "Amount",
	}

	if err := handlers.ModelExport(fileSaveAs, headers, data); err != nil {
		return err
	}

	return nil
}

func GetSksWeight(id string) (MstSksWeightDetail, error) {
	var result MstSksWeightDetail
	err := handlers.SPGetByID("sp_mst_sks_weights_get_by_id", id, &result)
	if err != nil {
		return MstSksWeightDetail{}, err
	}
	return result, nil
}

func CreateSksWeight(id string, amount int) error {
	return QueryInsertSksWeight(id, amount)
}

func ImportSksWeights(fileStatus string) error {
	headers := []string{
		"id", "amount",
	}

	return handlers.ModelImport(fileStatus, headers, func(row map[string]string) error {
		id := row["id"]
		amount, _ := strconv.Atoi(row["amount"])

		if id != "" {
			exist, err := helpers.CheckModelIDExist(id, &MstSksWeight{})
			if err != nil {
				return err
			}
			if exist {
				if err := QueryUpdateSksWeight(id, amount); err != nil {
					return err
				}
			} else {
				if err := QueryInsertSksWeight(id, amount); err != nil {
					return err
				}
			}
		} else {
			id, err := helpers.EnsureUUID(&MstSksWeight{})
			if err != nil {
				return err
			}
			if err := QueryInsertSksWeight(id, amount); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateSksWeight(id string, amount int) error {
	return QueryUpdateSksWeight(id, amount)
}

func DeleteSksWeight(id string) error {
	err := handlers.SPDelete("sp_mst_sks_weights_delete", id)
	if err != nil {
		return err
	}

	return nil
}

func GetTrashSksWeights(filter string, sortBy string, sortDirection string, page int, pageSize int64) ([]MstSksWeight, int64, error) {
	results, total, err := handlers.SPGetWithCount("sp_mst_sks_weights_has_deleted", filter, sortBy, sortDirection, page, pageSize, &MstSksWeight{})
	if err != nil {
		return []MstSksWeight{}, 0, err
	}

	var modelResults []MstSksWeight
	for _, item := range results {
		level, ok := item.(*MstSksWeight)
		if ok {
			modelResults = append(modelResults, *level)
		}
	}

	return modelResults, total, nil
}

func RestoreSksWeight(id string) error {
	err := handlers.SPRestore("sp_mst_sks_weights_restore", id)
	if err != nil {
		return err
	}

	return nil
}

/* Count */
func CountSksWeights(nullableDeletedAt ...bool) int64 {
	nullable := true
	if len(nullableDeletedAt) > 0 {
		nullable = nullableDeletedAt[0]
	}

	return helpers.CountModelSize(&MstSksWeight{}, nullable)
}

/* Query */
func QueryInsertSksWeight(id string, amount int) error {
	query := `
		EXEC sp_mst_sks_weights_insert
		@id = ?,
		@amount = ?, 
		@created_at = ?,
		@created_by = ?,
		@updated_at = ?,
		@updated_by = ?
	`
	log.Print(query)

	err := handlers.SPInsert(query, id, amount)
	if err != nil {
		return err
	}

	return nil
}

func QueryUpdateSksWeight(id string, amount int) error {
	query := `
		EXEC sp_mst_sks_weights_update
		@id = ?,
		@amount = ?, 
		@updated_at = ?,
		@updated_by = ?
	`
	err := handlers.SPUpdate(query, id, amount)
	if err != nil {
		return err
	}

	return nil
}

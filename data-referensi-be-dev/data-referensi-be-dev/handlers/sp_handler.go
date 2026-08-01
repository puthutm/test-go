package handlers

import (
	"data-referensi/config"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"
)

type TotalCount struct {
	Count int64 `gorm:"column:TotalCount"`
}

func tableNameFromSP(sp string) string {
	t := strings.TrimPrefix(sp, "sp_")
	suffixes := []string{
		"_get_all_trash", "_trash_get_all", "_get_all", "_get_by_id",
		"_insert", "_update", "_delete", "_restore",
	}
	for _, suf := range suffixes {
		if strings.HasSuffix(t, suf) {
			t = strings.TrimSuffix(t, suf)
			break
		}
	}
	return t
}

/* Create or Insert */
func SPInsert(query string, params ...interface{}) error {
	db := config.DB
	if db == nil {
		return fmt.Errorf("database connection is nil")
	}

	now := time.Now()
	createdAt := now.UnixMilli()
	updatedAt := now.UnixMilli()
	createdBy := GetUserId()
	updatedBy := GetUserId()

	appendParams := append(params, createdAt, createdBy, updatedAt, updatedBy)

	// If query is raw EXEC, fallback or execute
	if strings.Contains(strings.ToUpper(query), "EXEC") {
		spName := ""
		parts := strings.Fields(query)
		for i, p := range parts {
			if strings.EqualFold(p, "EXEC") && i+1 < len(parts) {
				spName = parts[i+1]
				break
			}
		}
		if spName != "" {
			table := tableNameFromSP(spName)
			log.Printf("[GORM ORM] Executing insert for table %s", table)
		}
	}

	err := db.Exec(query, appendParams...).Error
	log.Print(appendParams...)
	return err
}

func SPInsertSemesterNumber(query string, params ...interface{}) error {
	db := config.DB
	now := time.Now()
	createdAt := now.UnixMilli()
	createdBy := GetUserId()

	appendParams := append(params, createdAt, createdBy)
	return db.Exec(query, appendParams...).Error
}

/* Read */
func SPGet(sp string, filter string, sortBy string, sortDirection string, page int, pageSize int64, model interface{}) error {
	db := config.DB
	table := tableNameFromSP(sp)

	q := db.Table(table).Where("deleted_at IS NULL OR deleted_at = 0")

	if filter != "" {
		q = q.Where("name ILIKE ? OR code ILIKE ?", "%"+filter+"%", "%"+filter+"%")
	}

	if sortBy != "" {
		if sortDirection == "" {
			sortDirection = "ASC"
		}
		q = q.Order(fmt.Sprintf("%s %s", sortBy, sortDirection))
	}

	offset := (page - 1) * int(pageSize)
	if offset > 0 {
		q = q.Offset(offset)
	}
	if pageSize > 0 {
		q = q.Limit(int(pageSize))
	}

	err := q.Scan(model).Error
	if err != nil {
		// Fallback to raw query scan
		query := fmt.Sprintf("SELECT * FROM %s WHERE deleted_at IS NULL OR deleted_at = 0", table)
		return db.Raw(query).Scan(model).Error
	}
	return nil
}

func SPGetStudyProgram(sp string, filter string, sortBy string, sortDirection string, page int, pageSize int64, queryType string, model interface{}) error {
	return SPGet(sp, filter, sortBy, sortDirection, page, pageSize, model)
}

func SPGetWithCount(sp string, filter string, sortBy string, sortDirection string, page int, pageSize int64, model interface{}) ([]interface{}, int64, error) {
	db := config.DB
	table := tableNameFromSP(sp)

	q := db.Table(table).Where("deleted_at IS NULL OR deleted_at = 0")

	if filter != "" {
		q = q.Where("name ILIKE ? OR code ILIKE ?", "%"+filter+"%", "%"+filter+"%")
	}

	var totalCount int64
	_ = q.Count(&totalCount)

	if sortBy != "" {
		if sortDirection == "" {
			sortDirection = "ASC"
		}
		q = q.Order(fmt.Sprintf("%s %s", sortBy, sortDirection))
	}

	offset := (page - 1) * int(pageSize)
	if offset > 0 {
		q = q.Offset(offset)
	}
	if pageSize > 0 {
		q = q.Limit(int(pageSize))
	}

	var results []interface{}
	rows, err := q.Rows()
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		newModel := reflect.New(reflect.TypeOf(model).Elem()).Interface()
		if err := db.ScanRows(rows, newModel); err != nil {
			return nil, 0, err
		}
		results = append(results, newModel)
	}

	return results, totalCount, nil
}

func SPGetWithCountStudyProgram(sp string, filter string, sortBy string, sortDirection string, page int, pageSize int64, queryType string, model interface{}) ([]interface{}, int64, error) {
	return SPGetWithCount(sp, filter, sortBy, sortDirection, page, pageSize, model)
}

func SPGetByQuery(query string, model interface{}) error {
	db := config.DB
	return db.Raw(query).Scan(model).Error
}

func SPGetByID(sp string, id string, model interface{}) error {
	db := config.DB
	table := tableNameFromSP(sp)
	return db.Table(table).Where("id = ? AND (deleted_at IS NULL OR deleted_at = 0)", id).First(model).Error
}

/* Update */
func SPUpdate(query string, params ...interface{}) error {
	db := config.DB
	now := time.Now()
	updatedAt := now.UnixMilli()
	updatedBy := GetUserId()

	appendParams := append(params, updatedAt, updatedBy)
	return db.Exec(query, appendParams...).Error
}

func SPRestore(sp string, id string) error {
	db := config.DB
	table := tableNameFromSP(sp)
	return db.Table(table).Where("id = ?", id).Updates(map[string]interface{}{
		"deleted_at": nil,
		"deleted_by": nil,
	}).Error
}

/* Delete */
func SPDelete(sp string, id string) error {
	db := config.DB
	table := tableNameFromSP(sp)
	now := time.Now().UnixMilli()
	deletedBy := GetUserId()

	return db.Table(table).Where("id = ?", id).Updates(map[string]interface{}{
		"deleted_at": now,
		"deleted_by": deletedBy,
	}).Error
}

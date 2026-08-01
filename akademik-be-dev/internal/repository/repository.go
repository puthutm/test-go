package repository

import (
	"fmt"

	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
)

type Repository[T any] struct {
	DB *gorm.DB
}

func (r *Repository[T]) FindByGormOne(DB *gorm.DB, where string, args ...any) (T, error) {
	var model T
	err := DB.Where(where, args...).First(&model).Error
	return model, err
}

func (r *Repository[T]) FindByGormOneDeletedAt(DB *gorm.DB, deletedAt bool, where string, args ...any) (T, error) {
	var model T
	var whereD string
	if deletedAt {
		whereD = "deleted_at IS NULL OR deleted_at = 0"
	} else {
		whereD = "deleted_at > 0 "
	}
	err := DB.Where(whereD).Where(where, args...).First(&model).Error
	return model, err
}

func (r *Repository[T]) GetByGormMore(DB *gorm.DB, where string, args ...any) ([]T, error) {
	var models []T
	err := DB.Where(where, args...).Find(&models).Error
	return models, err
}

func (r *Repository[T]) CountModel(DB *gorm.DB, deletedAt bool) (count int64, err error) {
	var model T
	var where string
	if deletedAt {
		where = "deleted_at IS NULL OR deleted_at = 0"
	} else {
		where = "deleted_at > 0 "
	}
	err = DB.Model(&model).Where(where).Count(&count).Error
	return
}

func (r *Repository[T]) CountModelFilter(DB *gorm.DB, deletedAt bool, filterBy string, filter string) (count int64, err error) {
	var model T
	var teks string
	if filter != "" {
		teks = fmt.Sprintf(" AND %s LIKE '%%%s%%'", filterBy, filter)
	}
	var where string
	if deletedAt {
		where = "(deleted_at IS NULL OR deleted_at = 0)" + teks
	} else {
		where = "(deleted_at > 0)" + teks
	}
	err = DB.Model(&model).Where(where).Count(&count).Error
	return
}

func (r *Repository[T]) CountModelAll(DB *gorm.DB) (count int64, err error) {
	var model T
	err = DB.Model(&model).Count(&count).Error
	return
}

func (r *Repository[T]) CountModelWhere(DB *gorm.DB, where string, args ...any) (int64, error) {
	var model T
	var tot int64
	err := DB.Where(where, args...).Model(model).Count(&tot).Error
	if err != nil {
		return tot, err
	}
	return tot, nil
}

func (r *Repository[T]) CheckModelAll(DB *gorm.DB, where string, args ...any) (bool, error) {
	var model T
	var tot int64
	err := DB.Where(where, args...).Model(model).Count(&tot).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return tot > 0, nil
}

func (r *Repository[T]) CheckModel(DB *gorm.DB, deletedAt bool, where string, args ...any) (bool, error) {
	var model T
	var tot int64
	var whereDele string
	if deletedAt {
		whereDele = "deleted_at IS NULL OR deleted_at = 0"
	} else {
		whereDele = "deleted_at > 0"
	}
	err := DB.Model(&model).Where(whereDele).Where(where, args...).Count(&tot).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	return tot > 0, nil
}

func (r *Repository[T]) CreateModel(DB *gorm.DB, model T) error {
	err := DB.Create(&model).Error
	return err
}

func (r *Repository[T]) UpdateModel(DB *gorm.DB, model T, where string, args ...any) error {
	err := DB.Where(where, args...).Updates(&model).Error
	return err
}

func (r *Repository[T]) SaveModel(DB *gorm.DB, model *T) error {
	err := DB.Save(model).Error
	return err
}

func GetIDBySP[T any](db *gorm.DB, query string, ID string) (T, error) {
	var model T
	err := db.Raw(query, ID).Scan(&model).Error
	return model, err
}

func GetAllAndCountBySPUserId[T any](db *gorm.DB, query, userID string, pageable pageable.PageableRequest) ([]T, int64, error) {
	var model []T
	var TotalCount int64
	rows, err := db.Raw(query, userID, pageable.Search, pageable.SortBy, pageable.Sort, pageable.Page, pageable.Limit).Rows()
	if err != nil {
		return model, TotalCount, err
	}
	defer rows.Close()

	// Proses hasil pertama untuk data
	for rows.Next() {
		var data T
		if err := db.ScanRows(rows, &data); err != nil {
			return model, TotalCount, err
		}
		model = append(model, data)
	}

	// Cek jika sudah selesai dengan result set pertama
	if err := rows.Err(); err != nil {
		return model, TotalCount, err
	}

	// Pindah ke result set kedua (total count)
	if rows.NextResultSet() {
		if rows.Next() { // Pastikan ada data di result set kedua
			if err := rows.Scan(&TotalCount); err != nil {
				return model, TotalCount, err
			}
		} else {
			fmt.Println("No data found in second result set (total count). Skipping.")
		}
	} else {
		// Jika tidak ada result set kedua
		fmt.Println("No second result set found, skipping total count")
	}

	return model, TotalCount, err
}

func GetAllAndCount[T any](db *gorm.DB, query string, pageable pageable.PageableRequest) ([]T, int64, error) {
	var model []T
	var TotalCount int64
	rows, err := db.Raw(query, pageable.Search, pageable.SortBy, pageable.Sort, pageable.Page, pageable.Limit).Rows()
	if err != nil {
		return model, TotalCount, err
	}
	defer rows.Close()

	// Proses hasil pertama untuk data
	for rows.Next() {
		var data T
		if err := db.ScanRows(rows, &data); err != nil {
			return model, TotalCount, err
		}
		model = append(model, data)
	}

	// Cek jika sudah selesai dengan result set pertama
	if err := rows.Err(); err != nil {
		return model, TotalCount, err
	}

	// Pindah ke result set kedua (total count)
	if rows.NextResultSet() {
		if rows.Next() { // Pastikan ada data di result set kedua
			if err := rows.Scan(&TotalCount); err != nil {
				return model, TotalCount, err
			}
		} else {
			fmt.Println("No data found in second result set (total count). Skipping.")
		}
	} else {
		// Jika tidak ada result set kedua
		fmt.Println("No second result set found, skipping total count")
	}

	return model, TotalCount, err
}

func (r *Repository[T]) GetAllForExport(DB *gorm.DB) ([]T, error) {
	var model []T
	err := DB.Find(&model).Error
	return model, err
}

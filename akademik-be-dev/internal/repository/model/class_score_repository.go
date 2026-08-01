package repositorymodel

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	"unsia.ac.id/akademic_be/internal/model"
)

type ClassScoreRepository interface {
	GetByClassID(db *gorm.DB, pageable pageable.PageableRequestClassScore) (T []model.ClassScore, count int64, summary *model.ClassScoreSummary, err error)
	CheckSaveButton(db *gorm.DB, classID string) (bool, error)
	UpdateStatusLock(db *gorm.DB, classID string, statusLock bool, createdBy string) error
}

type classScoreRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewClassScoreRepository(
	log *logrus.Logger,
) ClassScoreRepository {
	return &classScoreRepository{
		log: log,
	}
}

func (r *classScoreRepository) GetByClassID(db *gorm.DB, pageable pageable.PageableRequestClassScore) (T []model.ClassScore, count int64, summary *model.ClassScoreSummary, err error) {
	classID := pageable.GetDefaultClassId()
	search := pageable.GetDefaultSearch()

	query := db.Table("trx_class_student_values").Where("class_id = ? AND (deleted_at IS NULL OR deleted_at = 0)", classID)
	if search != "" {
		query = query.Where("student_name ILIKE ? OR nim ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return T, count, nil, err
	}

	page := pageable.GetDefaultPage()
	limit := pageable.GetDefaultLimit()
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	if err := query.Offset(offset).Limit(limit).Scan(&T).Error; err != nil {
		return T, count, nil, err
	}

	var s model.ClassScoreSummary
	db.Table("trx_class_student_values").
		Select("COUNT(*) as total_students, SUM(CASE WHEN is_passed = true THEN 1 ELSE 0 END) as total_passed, SUM(CASE WHEN is_passed = false THEN 1 ELSE 0 END) as total_not_passed, AVG(final_score) as average_final_score").
		Where("class_id = ?", classID).
		Scan(&s)
	summary = &s

	return T, count, summary, nil
}

func (r *classScoreRepository) CheckSaveButton(db *gorm.DB, classID string) (bool, error) {
	var statusLock bool
	err := db.Table("trx_open_close_values").
		Select("status_lock").
		Where("class_id = ? AND (deleted_at IS NULL OR deleted_at = 0)", classID).
		Order("created_at DESC").
		Limit(1).
		Scan(&statusLock).Error
	if err != nil {
		return false, err
	}
	return statusLock, nil
}

func (r *classScoreRepository) UpdateStatusLock(db *gorm.DB, classID string, statusLock bool, createdBy string) error {
	record := map[string]interface{}{
		"class_id":    classID,
		"status_lock": statusLock,
		"created_at":  time.Now().UnixMilli(),
		"created_by":  createdBy,
	}
	return db.Table("trx_open_close_values").Create(record).Error
}

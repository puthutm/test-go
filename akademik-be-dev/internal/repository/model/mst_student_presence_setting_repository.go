// Package repositorymodel
package repositorymodel

import (
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	modelinternal "unsia.ac.id/akademic_be/internal/model"
)

type MstStudentPresenceSettingRepository struct {
	log *logrus.Logger
}

func NewMstStudentPresenceSettingRepository(
	log *logrus.Logger,
) *MstStudentPresenceSettingRepository {
	return &MstStudentPresenceSettingRepository{
		log: log,
	}
}

// Create

func (r *MstStudentPresenceSettingRepository) Create(
	ctx context.Context, tx *gorm.DB, studentPresenceSetttingM *modelinternal.MstStudentPresenceSetting,
) error {
	return tx.WithContext(ctx).Create(studentPresenceSetttingM).Error
}

func (r *MstStudentPresenceSettingRepository) Duplicate(
	ctx context.Context,
	tx *gorm.DB,
	param *modelinternal.MstStudentPresenceSettingDuplicateParam,
) error {
	var oldSettings []modelinternal.MstStudentPresenceSetting
	if err := tx.WithContext(ctx).Where("academic_periode_id = ? AND (deleted_at IS NULL OR deleted_at = 0)", param.AcademicPeriodeIDOld).Find(&oldSettings).Error; err != nil {
		return err
	}

	for _, setting := range oldSettings {
		newSetting := setting
		newSetting.ID = param.ID
		newSetting.AcademicPeriodeID = param.AcademicPeriodeID
		newSetting.StudyProgramID = param.StudyProgramID
		newSetting.CreatedAt = &param.CreatedAt
		newSetting.CreatedBy = &param.CreatedBy
		if err := tx.WithContext(ctx).Create(&newSetting).Error; err != nil {
			return err
		}
	}

	return nil
}

// Get

func (r *MstStudentPresenceSettingRepository) GetAllWithCount(
	ctx context.Context,
	db *gorm.DB,
	pageble pageable.PageableStudentSettingGets,
) (results []modelinternal.MstStudentPresenceSettingGetResult, count int64, err error) {
	query := db.WithContext(ctx).Table("mst_student_presence_settings").
		Where("deleted_at IS NULL OR deleted_at = 0")

	if pageble.AcademicPeriodID != nil && *pageble.AcademicPeriodID != "" {
		query = query.Where("academic_periode_id = ?", *pageble.AcademicPeriodID)
	}

	if pageble.StudyProgramID != nil && *pageble.StudyProgramID != "" {
		query = query.Where("study_program_id = ?", *pageble.StudyProgramID)
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	page := pageble.Page
	limit := pageble.Limit
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	if err := query.Offset(offset).Limit(limit).Scan(&results).Error; err != nil {
		return nil, 0, err
	}

	return results, count, nil
}

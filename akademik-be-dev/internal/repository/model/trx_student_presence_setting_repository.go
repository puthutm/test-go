// Package repositorymodel
package repositorymodel

import (
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	"unsia.ac.id/akademic_be/internal/model"
	modelinternal "unsia.ac.id/akademic_be/internal/model"
)

type TrxStudentPresenceSettingRepository struct {
	log *logrus.Logger
}

func NewTrxStudentPresenceSettingRepository(
	log *logrus.Logger,
) *TrxStudentPresenceSettingRepository {
	return &TrxStudentPresenceSettingRepository{
		log: log,
	}
}

// Create

func (r *TrxStudentPresenceSettingRepository) CreateOrUpdate(
	ctx context.Context, tx *gorm.DB, studentPresenceSettingM *modelinternal.TrxStudentPresenceSetting,
) error {
	return tx.WithContext(ctx).Table("trx_student_presence_settings").Save(studentPresenceSettingM).Error
}

func (r *TrxStudentPresenceSettingRepository) CreateOrUpdateStudentPresence(
	ctx context.Context, tx *gorm.DB, studentPresence *modelinternal.TrxStudentPresenceSaveParamBySession,
) error {
	record := map[string]interface{}{
		"id":              studentPresence.IDNew,
		"session_id":      studentPresence.SessionID,
		"student_id":      studentPresence.StudentID,
		"presence_status": studentPresence.PresenceStatus,
		"presence_type":   studentPresence.PresenceType,
		"created_at":      studentPresence.CreatedAt,
	}
	return tx.WithContext(ctx).Table("trx_student_presences").Save(record).Error
}

// Get

func (r *TrxStudentPresenceSettingRepository) GetPresenceComponentForLecturer(
	db *gorm.DB,
	param modelinternal.TrxStudentPresenceGetForLecturerParam,
	presence *modelinternal.Presence,
) error {
	return db.Table("trx_student_presence_settings").
		Where("academic_periode_id = ? AND study_program_id = ? AND subject_id = ?", param.AcademicPeriodeID, param.StudyProgramID, param.SubjectID).
		Scan(presence).Error
}

func (r *TrxStudentPresenceSettingRepository) GetPresenceComponentBySessionID(
	db *gorm.DB,
	sessionID string,
	presence *modelinternal.Presence,
) error {
	return db.Table("trx_student_presence_settings").
		Joins("JOIN mst_class_schedules ON mst_class_schedules.class_id = trx_student_presence_settings.subject_id").
		Where("mst_class_schedules.id = ?", sessionID).
		Scan(presence).Error
}

func (r *TrxStudentPresenceSettingRepository) GetSessionPresenceByClassID(
	ctx context.Context, db *gorm.DB, classID string,
) ([]model.SessionPresence, error) {
	var results []model.SessionPresence
	err := db.WithContext(ctx).Table("mst_class_schedules").
		Where("class_id = ? AND (deleted_at IS NULL OR deleted_at = 0)", classID).
		Scan(&results).Error
	return results, err
}

func (r *TrxStudentPresenceSettingRepository) GetStudentPresenceBySessionWithCount(
	ctx context.Context,
	tx *gorm.DB,
	pageble pageable.PageableStudentPresenceBySession,
) (results []model.TrxStudentPresenceBySession, count int64, err error) {
	query := tx.WithContext(ctx).Table("trx_student_presences").
		Where("session_id = ?", pageble.SessionID)

	if pageble.Status != nil && *pageble.Status != "" {
		query = query.Where("presence_status = ?", *pageble.Status)
	}

	search := pageble.GetDefaultSearch()
	if search != "" {
		query = query.Where("student_id IN (SELECT id FROM mst_student_bios WHERE name ILIKE ? OR nik ILIKE ?)", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	page := pageble.GetDefaultPage()
	limit := pageble.GetDefaultLimit()
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

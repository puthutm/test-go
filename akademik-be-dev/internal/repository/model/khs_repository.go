package repositorymodel

import (
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/model"
)

type KhsRepository interface {
	GetKHSByUserID(db *gorm.DB, ctx context.Context, userID string) ([]model.KhsSubject, []model.KhsSemester, []model.KhsNotPassed, error)
}

type khsRepository struct {
	log *logrus.Logger
}

func NewKhsRepository(log *logrus.Logger) KhsRepository {
	return &khsRepository{
		log: log,
	}
}

func (r *khsRepository) GetKHSByUserID(db *gorm.DB, ctx context.Context, userID string) ([]model.KhsSubject, []model.KhsSemester, []model.KhsNotPassed, error) {
	var subjects []model.KhsSubject
	var semesters []model.KhsSemester
	var notPassed []model.KhsNotPassed

	err := db.WithContext(ctx).Table("trx_krs_items").
		Select("trx_krs_items.*, mst_subjects.name as subject_name, mst_subjects.code as subject_code, mst_subjects.sks as subject_sks").
		Joins("JOIN mst_classes ON mst_classes.id = trx_krs_items.class_id").
		Joins("JOIN mst_subjects ON mst_subjects.id = mst_classes.subject_id").
		Joins("JOIN mst_student_bios ON mst_student_bios.id = trx_krs_items.student_id").
		Where("mst_student_bios.user_id = ? AND (trx_krs_items.deleted_at IS NULL OR trx_krs_items.deleted_at = 0)", userID).
		Scan(&subjects).Error
	if err != nil {
		return nil, nil, nil, err
	}

	err = db.WithContext(ctx).Table("trx_krs_headers").
		Select("academic_period_id, SUM(sks) as ips").
		Where("student_id IN (SELECT id FROM mst_student_bios WHERE user_id = ?) AND (deleted_at IS NULL OR deleted_at = 0)", userID).
		Group("academic_period_id").
		Scan(&semesters).Error
	if err != nil {
		return nil, nil, nil, err
	}

	err = db.WithContext(ctx).Table("trx_krs_items").
		Select("trx_krs_items.*").
		Joins("JOIN mst_student_bios ON mst_student_bios.id = trx_krs_items.student_id").
		Where("mst_student_bios.user_id = ? AND trx_krs_items.grade = 'E' AND (trx_krs_items.deleted_at IS NULL OR trx_krs_items.deleted_at = 0)", userID).
		Scan(&notPassed).Error
	if err != nil {
		return nil, nil, nil, err
	}

	return subjects, semesters, notPassed, nil
}

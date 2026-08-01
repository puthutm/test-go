package repositorymodel

import (
	"context"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	"unsia.ac.id/akademic_be/internal/model"
)

type KrsRepository interface {
	// Lecturer Students
	GetKrsLecturerStudentsWithCount(ctx context.Context, tx *gorm.DB, lecturerUserID string, pageable pageable.PageableKrsLecturerStudentsRequest) ([]model.TrxKrsLecturerStudent, int64, error)
	GetKrsLecturerStudentDetailByKrsHeaderID(ctx context.Context, tx *gorm.DB, krsHeaderID string) (model.TrxKrsLecturerStudentDetail, model.TrxKrsLecturerStudentTotalSKS, []model.TrxKrsLecturerStudentItem, error)
	UpdateKrsItemStatusByKrsItemID(ctx context.Context, tx *gorm.DB, req model.TrxKrsLecturerStudentItemUpdate) error

	// Pick Class
	GetPickClassesByUserID(db *gorm.DB, ctx context.Context, userID string, academicPeriodeID *string) ([]model.TrxKrsAcademicPeriod, []model.TrxKrsClassForPick, error)
	TakeClassByUserID(db *gorm.DB, ctx context.Context, userID string, classID string, createdAt int64, createdBy string) (model.TrxKrsTakeClassResult, error)

	// Program Head Classes
	GetKrsProgramHeadClassesWithCount(ctx context.Context, tx *gorm.DB, kaprodiUserID string, pageable pageable.PageableKrsProgramHeadClassesRequest) ([]model.TrxKrsProgramHeadClass, int64, error)

	// Saved
	GetSavedByUserID(db *gorm.DB, ctx context.Context, userID string, academicPeriodeID *string) ([]model.TrxKrsSavedItem, error)
	DeleteSavedByKrsItemID(db *gorm.DB, krsItemID string, deletedAt int64, deletedBy string) error

	// Info
	GetKrsMaxSksInfoByUserID(db *gorm.DB, ctx context.Context, userID string) (model.TrxKrsMaxSksInfo, error)
}

type krsRepository struct {
	log *logrus.Logger
}

func NewKrsRepository(log *logrus.Logger) KrsRepository {
	return &krsRepository{log: log}
}

// GetKrsLecturerStudentsWithCount - Get KRS lecturer students with pagination
func (r *krsRepository) GetKrsLecturerStudentsWithCount(
	ctx context.Context,
	tx *gorm.DB,
	lecturerUserID string,
	pageable pageable.PageableKrsLecturerStudentsRequest,
) ([]model.TrxKrsLecturerStudent, int64, error) {
	var students []model.TrxKrsLecturerStudent
	var count int64

	search := pageable.GetDefaultSearch()
	academicPeriodID := pageable.GetDefaultAcademicPeriodeId()

	query := tx.WithContext(ctx).Table("trx_krs_headers").
		Select("trx_krs_headers.id as krs_header_id, mst_student_bios.name as student_name, mst_student_bios.nik as student_nim").
		Joins("JOIN mst_student_bios ON mst_student_bios.id = trx_krs_headers.student_id").
		Where("trx_krs_headers.lecturer_user_id = ? AND (trx_krs_headers.deleted_at IS NULL OR trx_krs_headers.deleted_at = 0)", lecturerUserID)

	if academicPeriodID != nil && *academicPeriodID != "" {
		query = query.Where("trx_krs_headers.academic_period_id = ?", *academicPeriodID)
	}

	if search != "" {
		query = query.Where("mst_student_bios.name ILIKE ? OR mst_student_bios.nik ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
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

	if err := query.Offset(offset).Limit(limit).Scan(&students).Error; err != nil {
		return nil, 0, err
	}

	return students, count, nil
}

// GetKrsLecturerStudentDetailByKrsHeaderID - Get KRS student detail by KRS header ID
func (r *krsRepository) GetKrsLecturerStudentDetailByKrsHeaderID(
	ctx context.Context,
	tx *gorm.DB,
	krsHeaderID string,
) (model.TrxKrsLecturerStudentDetail, model.TrxKrsLecturerStudentTotalSKS, []model.TrxKrsLecturerStudentItem, error) {
	var studentDetail model.TrxKrsLecturerStudentDetail
	var totalSKS model.TrxKrsLecturerStudentTotalSKS
	var items []model.TrxKrsLecturerStudentItem

	err := tx.WithContext(ctx).Table("trx_krs_headers").
		Select("trx_krs_headers.student_id, mst_student_bios.nik as student_nim, mst_student_bios.name as student_name").
		Joins("JOIN mst_student_bios ON mst_student_bios.id = trx_krs_headers.student_id").
		Where("trx_krs_headers.id = ?", krsHeaderID).
		First(&studentDetail).Error
	if err != nil {
		return studentDetail, totalSKS, nil, err
	}

	tx.WithContext(ctx).Table("trx_krs_items").
		Select("COALESCE(SUM(mst_classes.sks), 0) as total_sks_taken").
		Joins("JOIN mst_classes ON mst_classes.id = trx_krs_items.class_id").
		Where("trx_krs_items.krs_header_id = ? AND (trx_krs_items.deleted_at IS NULL OR trx_krs_items.deleted_at = 0)", krsHeaderID).
		Scan(&totalSKS)

	tx.WithContext(ctx).Table("trx_krs_items").
		Select("trx_krs_items.id as krs_item_id, mst_subjects.name as subject_name_id, mst_classes.code as class_code, mst_classes.name as class_name, mst_classes.sks as total_sks").
		Joins("JOIN mst_classes ON mst_classes.id = trx_krs_items.class_id").
		Joins("JOIN mst_subjects ON mst_subjects.id = mst_classes.subject_id").
		Where("trx_krs_items.krs_header_id = ? AND (trx_krs_items.deleted_at IS NULL OR trx_krs_items.deleted_at = 0)", krsHeaderID).
		Scan(&items)

	return studentDetail, totalSKS, items, nil
}

// UpdateKrsItemStatusByKrsItemID - Update KRS item status (approved/rejected)
func (r *krsRepository) UpdateKrsItemStatusByKrsItemID(
	ctx context.Context,
	tx *gorm.DB,
	req model.TrxKrsLecturerStudentItemUpdate,
) error {
	return tx.WithContext(ctx).Table("trx_krs_items").
		Where("id = ?", req.KrsItemID).
		Updates(map[string]interface{}{
			"status":        req.ItemStatus,
			"reject_reason": req.RejectReason,
			"updated_at":    req.UpdatedAt,
			"updated_by":    req.UpdatedBy,
		}).Error
}

// GetPickClassesByUserID - Get academic periods and classes for KRS filling
func (r *krsRepository) GetPickClassesByUserID(
	db *gorm.DB, ctx context.Context, userID string, academicPeriodeID *string,
) ([]model.TrxKrsAcademicPeriod, []model.TrxKrsClassForPick, error) {
	var academicPeriods []model.TrxKrsAcademicPeriod
	var classes []model.TrxKrsClassForPick

	err := db.WithContext(ctx).Table("mst_academic_periods").
		Where("deleted_at IS NULL OR deleted_at = 0").
		Order("id DESC").
		Find(&academicPeriods).Error
	if err != nil {
		return nil, nil, err
	}

	query := db.WithContext(ctx).Table("mst_classes").
		Select("mst_classes.*, mst_subjects.name as subject_name, mst_subjects.code as subject_code, mst_subjects.sks").
		Joins("JOIN mst_subjects ON mst_subjects.id = mst_classes.subject_id").
		Where("mst_classes.deleted_at IS NULL OR mst_classes.deleted_at = 0")

	if academicPeriodeID != nil && *academicPeriodeID != "" {
		query = query.Where("mst_classes.academic_period_id = ?", *academicPeriodeID)
	}

	if err := query.Scan(&classes).Error; err != nil {
		return academicPeriods, nil, err
	}

	return academicPeriods, classes, nil
}

// TakeClassByUserID - Take a class in KRS
func (r *krsRepository) TakeClassByUserID(
	db *gorm.DB, ctx context.Context, userID string, classID string, createdAt int64, createdBy string,
) (model.TrxKrsTakeClassResult, error) {
	var studentID string
	err := db.WithContext(ctx).Table("mst_student_bios").
		Select("id").
		Where("user_id = ? AND (deleted_at IS NULL OR deleted_at = 0)", userID).
		Scan(&studentID).Error
	if err != nil || studentID == "" {
		return model.TrxKrsTakeClassResult{}, err
	}

	item := map[string]interface{}{
		"student_id": studentID,
		"class_id":   classID,
		"status":     "PENDING",
		"created_at": createdAt,
		"created_by": createdBy,
	}

	var res model.TrxKrsTakeClassResult
	if err := db.WithContext(ctx).Table("trx_krs_items").Create(item).Error; err != nil {
		return res, err
	}

	res.KrsID, _ = uuid.Parse(studentID)
	return res, nil
}

// GetKrsProgramHeadClassesWithCount - Get KRS program head classes with pagination
func (r *krsRepository) GetKrsProgramHeadClassesWithCount(
	ctx context.Context,
	tx *gorm.DB,
	kaprodiUserID string,
	pageable pageable.PageableKrsProgramHeadClassesRequest,
) ([]model.TrxKrsProgramHeadClass, int64, error) {
	var classes []model.TrxKrsProgramHeadClass
	var count int64

	search := pageable.GetDefaultSearch()

	query := tx.WithContext(ctx).Table("mst_classes").
		Select("mst_classes.*, mst_subjects.name as subject_name").
		Joins("JOIN mst_subjects ON mst_subjects.id = mst_classes.subject_id").
		Where("mst_classes.deleted_at IS NULL OR mst_classes.deleted_at = 0")

	if search != "" {
		query = query.Where("mst_classes.name ILIKE ? OR mst_subjects.name ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
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

	if err := query.Offset(offset).Limit(limit).Scan(&classes).Error; err != nil {
		return nil, 0, err
	}

	return classes, count, nil
}

// GetSavedByUserID - Get saved KRS items by user ID
func (r *krsRepository) GetSavedByUserID(
	db *gorm.DB, ctx context.Context, userID string, academicPeriodeID *string,
) ([]model.TrxKrsSavedItem, error) {
	var items []model.TrxKrsSavedItem

	query := db.WithContext(ctx).Table("trx_krs_items").
		Select("trx_krs_items.id as krs_item_id, mst_classes.code as class_code, mst_classes.name as class_name, mst_subjects.name as subject_name, mst_subjects.sks").
		Joins("JOIN mst_classes ON mst_classes.id = trx_krs_items.class_id").
		Joins("JOIN mst_subjects ON mst_subjects.id = mst_classes.subject_id").
		Joins("JOIN mst_student_bios ON mst_student_bios.id = trx_krs_items.student_id").
		Where("mst_student_bios.user_id = ? AND (trx_krs_items.deleted_at IS NULL OR trx_krs_items.deleted_at = 0)", userID)

	if academicPeriodeID != nil && *academicPeriodeID != "" {
		query = query.Where("mst_classes.academic_period_id = ?", *academicPeriodeID)
	}

	if err := query.Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

// DeleteSavedByKrsItemID - Soft delete KRS item
func (r *krsRepository) DeleteSavedByKrsItemID(
	db *gorm.DB, krsItemID string, deletedAt int64, deletedBy string,
) error {
	return db.Table("trx_krs_items").
		Where("id = ?", krsItemID).
		Updates(map[string]interface{}{
			"deleted_at": deletedAt,
			"deleted_by": deletedBy,
		}).Error
}

// GetKrsMaxSksInfoByUserID - Get max SKS limit for user
func (r *krsRepository) GetKrsMaxSksInfoByUserID(
	db *gorm.DB, ctx context.Context, userID string,
) (model.TrxKrsMaxSksInfo, error) {
	var result model.TrxKrsMaxSksInfo
	err := db.WithContext(ctx).Table("mst_sks_limits").
		Select("max_sks").
		Order("id DESC").
		Limit(1).
		Scan(&result).Error
	return result, err
}

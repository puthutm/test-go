package repositorymodel

import (
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	msg "unsia.ac.id/akademic_be/internal/config/message"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository"
	"unsia.ac.id/akademic_be/internal/repository/cached"
)

type MstClassLecturerRepository struct {
	repository.Repository[model.MstClassLecturer]
	log   *logrus.Logger
	cache cached.CacheRepository
}

func NewMstClassLecturerRepository(
	log *logrus.Logger,
	cache cached.CacheRepository,
) *MstClassLecturerRepository {
	return &MstClassLecturerRepository{
		log:   log,
		cache: cache,
	}
}

/* Create */
func (r *MstClassLecturerRepository) Create(db *gorm.DB, req *model.MstClassLecturer) error {
	err := db.Create(req).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository":  "create class lectural",
			"lectural_id": req.LecturerID,
		}).Error(msg.ErrCreate.Error())
		return err
	}
	return nil
}

/* Read */
func (r *MstClassLecturerRepository) GetByClassID(db *gorm.DB, ClassID string, data *model.MstClassLecturer) error {
	err := db.Where("class_id = ? AND (deleted_at IS NULL OR deleted_at = 0)", ClassID).First(data).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "get by class id",
			"class_id":   ClassID,
		}).Error(msg.ErrRead.Error())
		return err
	}
	return nil
}

func (r *MstClassLecturerRepository) GetClassByAcademicPeriodAndSubjectAndUserForLecturerWithCount(
	ctx context.Context,
	tx *gorm.DB, userID string,
	pageble pageable.PageableClassGetByUserAndAcademicPeriodAndSubject,
) (T []model.MstClassLecturerGetByAcademicPeriodAndSubjectAndUserForLecturer, count int64, err error) {
	query := tx.WithContext(ctx).Table("mst_class_lecturers").
		Select("mst_class_lecturers.*, mst_classes.code as class_code, mst_classes.name as class_name").
		Joins("JOIN mst_classes ON mst_classes.id = mst_class_lecturers.class_id").
		Where("mst_class_lecturers.lecturer_user_id = ? AND (mst_class_lecturers.deleted_at IS NULL OR mst_class_lecturers.deleted_at = 0)", userID)

	if pageble.AcademicPeriodID != nil && *pageble.AcademicPeriodID != "" {
		query = query.Where("mst_classes.academic_period_id = ?", *pageble.AcademicPeriodID)
	}

	if pageble.SubjectID != nil && *pageble.SubjectID != "" {
		query = query.Where("mst_classes.subject_id = ?", *pageble.SubjectID)
	}

	search := pageble.GetDefaultSearch()
	if search != "" {
		query = query.Where("mst_classes.name ILIKE ? OR mst_classes.code ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return T, count, err
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

	if err := query.Offset(offset).Limit(limit).Scan(&T).Error; err != nil {
		return T, count, err
	}

	return T, count, nil
}

func (r *MstClassLecturerRepository) GetSubjectByClassLecturerWithCount(
	ctx context.Context,
	tx *gorm.DB, userID string,
	pageble pageable.PageableRequestClass,
) (T []model.MstClassGetResultSubjectAndClassCount, count int64, err error) {
	query := tx.WithContext(ctx).Table("mst_classes").
		Select("mst_subjects.id as subject_id, mst_subjects.name as subject_name, COUNT(mst_classes.id) as class_count").
		Joins("JOIN mst_class_lecturers ON mst_class_lecturers.class_id = mst_classes.id").
		Joins("JOIN mst_subjects ON mst_subjects.id = mst_classes.subject_id").
		Where("mst_class_lecturers.lecturer_user_id = ? AND (mst_classes.deleted_at IS NULL OR mst_classes.deleted_at = 0)", userID).
		Group("mst_subjects.id, mst_subjects.name")

	if pageble.AcademicPeriodeId != "" {
		query = query.Where("mst_classes.academic_period_id = ?", pageble.AcademicPeriodeId)
	}

	search := pageble.GetDefaultSearch()
	if search != "" {
		query = query.Where("mst_subjects.name ILIKE ?", "%"+search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return T, count, err
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

	if err := query.Offset(offset).Limit(limit).Scan(&T).Error; err != nil {
		return T, count, err
	}

	return T, count, nil
}

func (r *MstClassLecturerRepository) GetByID(db *gorm.DB, ID string, data *model.MstClassLecturer) error {
	return db.Where("id = ? AND (deleted_at IS NULL OR deleted_at = 0)", ID).First(data).Error
}

/* Update */
func (r *MstClassLecturerRepository) UpdateByID(db *gorm.DB, req *model.MstClassLecturer) error {
	err := db.Model(&model.MstClassLecturer{}).
		Where("id = ?", req.ID).
		Updates(map[string]interface{}{
			"class_id":              req.ClassID,
			"lecturer_id":           req.LecturerID,
			"subtitute_lecturer_id": req.SubtituteLecturerID,
		}).Error

	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "update by class lecturer id",
			"class_id":   req.ClassID,
		}).Error(msg.ErrUpdate.Error())
	}

	return err
}

/* Delete */
func (r *MstClassLecturerRepository) DeleteByID(db *gorm.DB, ID string) error {
	return db.Where("id = ?", ID).Delete(&model.MstClassLecturer{}).Error
}

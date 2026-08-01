package repositorymodel

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	msg "unsia.ac.id/akademic_be/internal/config/message"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository"
	"unsia.ac.id/akademic_be/internal/repository/cached"
	"unsia.ac.id/akademic_be/pkg/auth"
	"unsia.ac.id/akademic_be/pkg/utils"
)

type MstSubjectRepository struct {
	log *logrus.Logger
	repository.Repository[model.MstSubject]
	cacheRepository cached.CacheRepository
}

func NewMstSubjectRepository(
	log *logrus.Logger,

	cacheRepository cached.CacheRepository,
) *MstSubjectRepository {
	return &MstSubjectRepository{
		log:             log,
		cacheRepository: cacheRepository,
	}
}

// Create
func (r *MstSubjectRepository) Create(db *gorm.DB, data *model.MstSubject) error {
	data.CreatedAt = time.Now().UnixMilli()
	err := db.Create(data).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "create subject",
			"student_id": data.CreatedBy,
		}).Error(msg.ErrCreate.Error())
		return err
	}
	return nil
}

func (r *MstSubjectRepository) CreateSupportingLecturer(db *gorm.DB, subjectID uuid.UUID, lecturerIDs []uuid.UUID) error {
	for _, lecturerID := range lecturerIDs {
		item := map[string]interface{}{
			"id":          utils.GenerateUUID(),
			"subject_id":  subjectID,
			"lecturer_id": lecturerID,
		}
		if err := db.Table("mst_subject_supporting_lecturers").Create(item).Error; err != nil {
			r.log.WithFields(logrus.Fields{
				"repository":  "create subject supporting lecturer",
				"lecturer_id": lecturerID,
			}).Error(msg.ErrCreate.Error())
			return err
		}
	}
	return nil
}

func (r *MstSubjectRepository) CreateDeveloperRPSLecuter(db *gorm.DB, subjectID uuid.UUID, lecturerIDs []uuid.UUID) error {
	for _, lecturerID := range lecturerIDs {
		item := map[string]interface{}{
			"id":          utils.GenerateUUID(),
			"subject_id":  subjectID,
			"lecturer_id": lecturerID,
		}
		if err := db.Table("mst_subject_developer_rps_lecturers").Create(item).Error; err != nil {
			r.log.WithFields(logrus.Fields{
				"repository":  "create subject developer rps lecturer",
				"lecturer_id": lecturerID,
			}).Error(msg.ErrCreate.Error())
			return err
		}
	}
	return nil
}

func (r *MstSubjectRepository) CreateSubjectCoordinatorLecuter(db *gorm.DB, subjectID uuid.UUID, lecturerIDs []uuid.UUID) error {
	for _, lecturerID := range lecturerIDs {
		item := map[string]interface{}{
			"id":          utils.GenerateUUID(),
			"subject_id":  subjectID,
			"lecturer_id": lecturerID,
		}
		if err := db.Table("mst_subject_coordinator_lecturers").Create(item).Error; err != nil {
			r.log.WithFields(logrus.Fields{
				"repository":  "create subject coordinator lecturer",
				"lecturer_id": lecturerID,
			}).Error(msg.ErrCreate.Error())
			return err
		}
	}
	return nil
}

// Update
func (r *MstSubjectRepository) UpdateByID(db *gorm.DB, data *model.MstSubject) error {
	now := time.Now().UnixMilli()
	data.UpdatedAt = &now

	err := db.Model(&model.MstSubject{}).Where("id = ?", data.ID).Updates(data).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "update subject",
			"id":         data.ID,
		}).Error(msg.ErrUpdate.Error())
		return err
	}
	return nil
}

// Delete
func (r *MstSubjectRepository) DeleteByID(db *gorm.DB, ID, deletedBy string) error {
	now := time.Now().UnixMilli()
	err := db.Model(&model.MstSubject{}).Where("id = ?", ID).Updates(map[string]interface{}{
		"deleted_at": now,
		"deleted_by": deletedBy,
	}).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "deleted subject",
			"id":         ID,
		}).Error(msg.ErrDelete.Error())
		return err
	}
	return nil
}

func (r *MstSubjectRepository) DeleteBySupportingLecturerID(db *gorm.DB, SubjectID string) error {
	return db.Table("mst_subject_supporting_lecturers").Where("subject_id = ?", SubjectID).Delete(nil).Error
}

func (r *MstSubjectRepository) DeleteByDeveloperRPSLecuterID(db *gorm.DB, SubjectID string) error {
	return db.Table("mst_subject_developer_rps_lecturers").Where("subject_id = ?", SubjectID).Delete(nil).Error
}

func (r *MstSubjectRepository) DeleteBySubjectCoordinatorLecuterID(db *gorm.DB, SubjectID string) error {
	return db.Table("mst_subject_coordinator_lecturers").Where("subject_id = ?", SubjectID).Delete(nil).Error
}

// Restore
func (r *MstSubjectRepository) RestoreByID(db *gorm.DB, ID string) error {
	return db.Model(&model.MstSubject{}).Where("id = ?", ID).Update("deleted_at", nil).Error
}

// GetByID
func (r *MstSubjectRepository) GetByID(db *gorm.DB, ID string, data *model.MstSubject) error {
	return db.Where("id = ? AND (deleted_at IS NULL OR deleted_at = 0)", ID).First(data).Error
}

func (r *MstSubjectRepository) GetSupportingLecturerBySubjectID(db *gorm.DB, SubjectID string) (supportingLecturers []model.MstSubject, err error) {
	err = db.Table("mst_subjects").
		Joins("JOIN mst_subject_supporting_lecturers ON mst_subject_supporting_lecturers.subject_id = mst_subjects.id").
		Where("mst_subject_supporting_lecturers.subject_id = ? AND (mst_subjects.deleted_at IS NULL OR mst_subjects.deleted_at = 0)", SubjectID).
		Scan(&supportingLecturers).Error
	return supportingLecturers, err
}

func (r *MstSubjectRepository) GetDeveloperRPSBySubjectID(db *gorm.DB, SubjectID string) (developerRPSs []model.MstSubject, err error) {
	err = db.Table("mst_subjects").
		Joins("JOIN mst_subject_developer_rps_lecturers ON mst_subject_developer_rps_lecturers.subject_id = mst_subjects.id").
		Where("mst_subject_developer_rps_lecturers.subject_id = ? AND (mst_subjects.deleted_at IS NULL OR mst_subjects.deleted_at = 0)", SubjectID).
		Scan(&developerRPSs).Error
	return developerRPSs, err
}

func (r *MstSubjectRepository) GetSubjectCoordinatorBySubjectID(db *gorm.DB, SubjectID string) (subjectCoordinators []model.MstSubject, err error) {
	err = db.Table("mst_subjects").
		Joins("JOIN mst_subject_coordinator_lecturers ON mst_subject_coordinator_lecturers.subject_id = mst_subjects.id").
		Where("mst_subject_coordinator_lecturers.subject_id = ? AND (mst_subjects.deleted_at IS NULL OR mst_subjects.deleted_at = 0)", SubjectID).
		Scan(&subjectCoordinators).Error
	return subjectCoordinators, err
}

// GetAllWithCount
func (r *MstSubjectRepository) GetAllWithCount(
	db *gorm.DB,
	deleted bool,
	pageble pageable.PageableRequestSubject,
) (T []model.MstSubject, count int64, err error) {
	query := db.Model(&model.MstSubject{})
	if deleted {
		query = query.Where("deleted_at IS NOT NULL AND deleted_at > 0")
	} else {
		query = query.Where("deleted_at IS NULL OR deleted_at = 0")
	}

	if pageble.Search != "" {
		query = query.Where("name_id ILIKE ? OR code ILIKE ?", "%"+pageble.Search+"%", "%"+pageble.Search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "get all subject",
		}).Error(msg.ErrMultipleRead.Error())
		return T, count, err
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

	err = query.Offset(offset).Limit(limit).Find(&T).Error
	return T, count, err
}

func (r *MstSubjectRepository) GetAllWithCountByLecuturerID(db *gorm.DB, ctx context.Context, pageable pageable.PageableRequestSubject) (T []model.MstSubject, count int64, err error) {
	user := ctx.Value("x-user-claims").(*auth.UserClaimsSpesifikRole)
	query := db.WithContext(ctx).Model(&model.MstSubject{}).
		Joins("JOIN mst_subject_supporting_lecturers ON mst_subject_supporting_lecturers.subject_id = mst_subjects.id").
		Where("mst_subject_supporting_lecturers.lecturer_id = ? AND (mst_subjects.deleted_at IS NULL OR mst_subjects.deleted_at = 0)", user.ID)

	search := pageable.GetDefaultSearch()
	if search != "" {
		query = query.Where("mst_subjects.name_id ILIKE ? OR mst_subjects.code ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return T, count, err
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

	err = query.Offset(offset).Limit(limit).Find(&T).Error
	return T, count, err
}

func (r *MstSubjectRepository) GetAllWithCountByCoordinatorLecuturerID(db *gorm.DB, ctx context.Context, pageable pageable.PageableRequestSubject) (T []model.MstSubject, count int64, err error) {
	user := ctx.Value("x-user-claims").(*auth.UserClaimsSpesifikRole)
	query := db.WithContext(ctx).Model(&model.MstSubject{}).
		Joins("JOIN mst_subject_coordinator_lecturers ON mst_subject_coordinator_lecturers.subject_id = mst_subjects.id").
		Where("mst_subject_coordinator_lecturers.lecturer_id = ? AND (mst_subjects.deleted_at IS NULL OR mst_subjects.deleted_at = 0)", user.ID)

	search := pageable.GetDefaultSearch()
	if search != "" {
		query = query.Where("mst_subjects.name_id ILIKE ? OR mst_subjects.code ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return T, count, err
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

	err = query.Offset(offset).Limit(limit).Find(&T).Error
	return T, count, err
}

func (r *MstSubjectRepository) GetByStudyProgramIDAndCurriculumYearID(
	db *gorm.DB,
	studyProgramID string,
	curriculumYearID string,
) (subjects []model.MstSubject, err error) {
	err = db.Where("study_program_id = ? AND curriculum_year_id = ? AND (deleted_at IS NULL OR deleted_at = 0)", studyProgramID, curriculumYearID).Find(&subjects).Error
	return subjects, err
}

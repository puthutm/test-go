package repositorymodel

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	msg "unsia.ac.id/akademic_be/internal/config/message"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository"
	"unsia.ac.id/akademic_be/internal/repository/cached"
	"unsia.ac.id/akademic_be/pkg/auth"
)

type MstClassRepository interface {
	Create(db *gorm.DB, ctx context.Context, req dto.MstClassRequest) error
	CreateByProgramHead(db *gorm.DB, ctx context.Context, req dto.MstClassRequestByProgramHead) error
	GetAllWithCount(db *gorm.DB, deleted bool, pageable pageable.PageableRequestClass) (T []model.MstClass, count int64, err error)
	GetAllWithCountByStudyProgramID(db *gorm.DB, pageable pageable.PageableRequestClass) (T []model.MstClass, count int64, err error)
	GetByID(db *gorm.DB, ID string, data *model.MstClass) error
	GetAllWithCountByProgramHeadID(db *gorm.DB, ctx context.Context, pageable pageable.PageableRequestClass) (T []model.MstClass, count int64, err error)
	GetByLecturerIDandActiveAcademicPeriod(db *gorm.DB, ctx context.Context) (T []model.MstClass, err error)
	UpdateByID(db *gorm.DB, ctx context.Context, req dto.MstClassRequest) error
	UpdateByIDAndProgramHead(db *gorm.DB, ctx context.Context, req dto.MstClassRequestByProgramHead) error
	RestoreByID(db *gorm.DB, ID string) error
	DeleteByID(db *gorm.DB, ctx context.Context, ID string) error
	UpdateContractByID(db *gorm.DB, req *model.MstClass) error

	CheckSaveButton(db *gorm.DB, academicPeriodID string) (bool, error)
	UpdateStatusLockedByAcademicPeriodID(db *gorm.DB, req dto.UpdateStatusLockRequest) error
}

type mstClassRepository struct {
	repository.Repository[model.MstClass]
	log             *logrus.Logger
	cacheRepository cached.CacheRepository
}

func NewMstClassRepository(
	log *logrus.Logger,
	cacheRepository cached.CacheRepository,
) MstClassRepository {
	return &mstClassRepository{
		log:             log,
		cacheRepository: cacheRepository,
	}
}

/* Create */
func (r *mstClassRepository) Create(db *gorm.DB, ctx context.Context, req dto.MstClassRequest) error {
	user := ctx.Value("x-user-claims").(*auth.UserClaimsSpesifikRole)

	item := map[string]interface{}{
		"id":                  req.ID,
		"code":                req.Code,
		"name":                req.Name,
		"academic_periode_id": req.AcademicPeriodeID,
		"subject_id":          req.SubjectID,
		"study_program_id":    req.StudyProgramID,
		"curriculum_year_id":  req.CurriculumYearID,
		"number_of_meeting":   req.NumberOfMeeting,
		"capacity":            req.Capacity,
		"created_at":          time.Now().UnixMilli(),
		"created_by":          user.ID,
	}

	return db.WithContext(ctx).Table("mst_classes").Create(item).Error
}

func (r *mstClassRepository) CreateByProgramHead(db *gorm.DB, ctx context.Context, req dto.MstClassRequestByProgramHead) error {
	user := ctx.Value("x-user-claims").(*auth.UserClaimsSpesifikRole)

	item := map[string]interface{}{
		"id":                  req.ID,
		"code":                req.Code,
		"name":                req.Name,
		"academic_periode_id": req.AcademicPeriodeID,
		"subject_id":          req.SubjectID,
		"curriculum_year_id":  req.CurriculumYearID,
		"number_of_meeting":   req.NumberOfMeeting,
		"capacity":            req.Capacity,
		"created_at":          time.Now().UnixMilli(),
		"created_by":          user.ID,
	}

	return db.WithContext(ctx).Table("mst_classes").Create(item).Error
}

/* Read */
func (r *mstClassRepository) GetAllWithCount(db *gorm.DB, deleted bool, pg pageable.PageableRequestClass) (T []model.MstClass, count int64, err error) {
	query := db.Model(&model.MstClass{})
	if deleted {
		query = query.Where("deleted_at IS NOT NULL AND deleted_at > 0")
	} else {
		query = query.Where("deleted_at IS NULL OR deleted_at = 0")
	}

	search := pg.GetDefaultSearch()
	academicPeriodID := pg.GetDefaultAcademicPeriodeId()
	studyProgramID := pg.GetDefaultStudyProgramId()

	if academicPeriodID != "" {
		query = query.Where("academic_periode_id = ?", academicPeriodID)
	}

	if studyProgramID != nil && *studyProgramID != "" {
		query = query.Where("study_program_id = ?", *studyProgramID)
	}

	if search != "" {
		query = query.Where("name ILIKE ? OR code ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return T, count, err
	}

	page := pg.GetDefaultPage()
	limit := pg.GetDefaultLimit()
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

func (r *mstClassRepository) GetAllWithCountByStudyProgramID(db *gorm.DB, pg pageable.PageableRequestClass) (T []model.MstClass, count int64, err error) {
	return r.GetAllWithCount(db, false, pg)
}

func (r *mstClassRepository) GetByID(db *gorm.DB, ID string, data *model.MstClass) error {
	return db.Where("id = ? AND (deleted_at IS NULL OR deleted_at = 0)", ID).First(data).Error
}

func (r *mstClassRepository) GetAllWithCountByProgramHeadID(db *gorm.DB, ctx context.Context, pg pageable.PageableRequestClass) (T []model.MstClass, count int64, err error) {
	user := ctx.Value("x-user-claims").(*auth.UserClaimsSpesifikRole)
	query := db.WithContext(ctx).Model(&model.MstClass{}).Where("created_by = ? AND (deleted_at IS NULL OR deleted_at = 0)", user.ID)

	search := pg.GetDefaultSearch()
	academicPeriodID := pg.GetDefaultAcademicPeriodeId()

	if academicPeriodID != "" {
		query = query.Where("academic_periode_id = ?", academicPeriodID)
	}

	if search != "" {
		query = query.Where("name ILIKE ? OR code ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return T, count, err
	}

	page := pg.GetDefaultPage()
	limit := pg.GetDefaultLimit()
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

func (r *mstClassRepository) GetByLecturerIDandActiveAcademicPeriod(db *gorm.DB, ctx context.Context) (T []model.MstClass, err error) {
	user := ctx.Value("x-user-claims").(*auth.UserClaimsSpesifikRole)
	err = db.WithContext(ctx).Table("mst_classes").
		Joins("JOIN mst_class_lecturers ON mst_class_lecturers.class_id = mst_classes.id").
		Where("mst_class_lecturers.lecturer_user_id = ? AND (mst_classes.deleted_at IS NULL OR mst_classes.deleted_at = 0)", user.ID).
		Scan(&T).Error
	return T, err
}

/* Update */
func (r *mstClassRepository) UpdateByID(db *gorm.DB, ctx context.Context, req dto.MstClassRequest) error {
	user := ctx.Value("x-user-claims").(*auth.UserClaimsSpesifikRole)

	return db.WithContext(ctx).Table("mst_classes").
		Where("id = ?", req.ID).
		Updates(map[string]interface{}{
			"code":                req.Code,
			"name":                req.Name,
			"academic_periode_id": req.AcademicPeriodeID,
			"subject_id":          req.SubjectID,
			"study_program_id":    req.StudyProgramID,
			"curriculum_year_id":  req.CurriculumYearID,
			"number_of_meeting":   req.NumberOfMeeting,
			"capacity":            req.Capacity,
			"updated_at":          time.Now().UnixMilli(),
			"updated_by":          user.ID,
		}).Error
}

func (r *mstClassRepository) UpdateByIDAndProgramHead(db *gorm.DB, ctx context.Context, req dto.MstClassRequestByProgramHead) error {
	user := ctx.Value("x-user-claims").(*auth.UserClaimsSpesifikRole)

	return db.WithContext(ctx).Table("mst_classes").
		Where("id = ?", req.ID).
		Updates(map[string]interface{}{
			"code":                req.Code,
			"name":                req.Name,
			"academic_periode_id": req.AcademicPeriodeID,
			"subject_id":          req.SubjectID,
			"curriculum_year_id":  req.CurriculumYearID,
			"number_of_meeting":   req.NumberOfMeeting,
			"capacity":            req.Capacity,
			"updated_at":          time.Now().UnixMilli(),
			"updated_by":          user.ID,
		}).Error
}

func (r *mstClassRepository) RestoreByID(db *gorm.DB, ID string) error {
	return db.Table("mst_classes").Where("id = ?", ID).Update("deleted_at", nil).Error
}

/* Delete */
func (r *mstClassRepository) DeleteByID(db *gorm.DB, ctx context.Context, ID string) error {
	user := ctx.Value("x-user-claims").(*auth.UserClaimsSpesifikRole)

	return db.WithContext(ctx).Table("mst_classes").
		Where("id = ?", ID).
		Updates(map[string]interface{}{
			"deleted_at": time.Now().UnixMilli(),
			"deleted_by": user.ID,
		}).Error
}

func (r *mstClassRepository) UpdateContractByID(db *gorm.DB, req *model.MstClass) error {
	err := db.Table("mst_classes").
		Where("id = ?", req.ID).
		Updates(map[string]interface{}{
			"contract_description": req.ContractDescription,
			"contract_file_path":   req.ContractFilePath,
		}).Error

	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "update class by id",
			"class_id":   req.ID,
		}).Error(msg.ErrUpdate.Error())
	}

	return err
}

func (r *mstClassRepository) CheckSaveButton(db *gorm.DB, academicPeriodID string) (bool, error) {
	var statusLock bool
	err := db.Table("trx_open_close_values").
		Select("status_lock").
		Where("academic_periode_id = ? AND (deleted_at IS NULL OR deleted_at = 0)", academicPeriodID).
		Order("created_at DESC").
		Limit(1).
		Scan(&statusLock).Error
	if err != nil {
		return false, err
	}
	return statusLock, nil
}

func (r *mstClassRepository) UpdateStatusLockedByAcademicPeriodID(db *gorm.DB, req dto.UpdateStatusLockRequest) error {
	item := map[string]interface{}{
		"academic_periode_id": req.AcademicPeriodID,
		"status_lock":         req.StatusLocked,
		"created_at":          req.CreatedAt,
		"created_by":          req.CreatedBy,
	}
	return db.Table("trx_open_close_values").Create(item).Error
}

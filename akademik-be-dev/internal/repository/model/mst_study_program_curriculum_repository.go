package repositorymodel

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository"
	"unsia.ac.id/akademic_be/internal/repository/cached"
	"unsia.ac.id/akademic_be/pkg/auth"
)

type MstStudyProgramCurriculumRepository struct {
	repository.Repository[model.MstStudyProgramCurriculum]
	log             *logrus.Logger
	cacheRepository cached.CacheRepository
}

func NewMstStudyProgramCurriculumRepository(
	log *logrus.Logger,
	cacheRepository cached.CacheRepository,
) *MstStudyProgramCurriculumRepository {
	return &MstStudyProgramCurriculumRepository{
		log:             log,
		cacheRepository: cacheRepository,
	}
}

/* Create */
func (r *MstStudyProgramCurriculumRepository) Create(db *gorm.DB, ctx context.Context, req dto.MstStudyProgramCurriculumRequest) error {
	user := ctx.Value("x-user-claims").(*auth.UserClaimsSpesifikRole)

	item := map[string]interface{}{
		"id":                            req.ID,
		"study_program_id":              req.StudyProgramID,
		"curriculum_year_id":            req.CurriculumYearID,
		"subject_id":                    req.SubjectID,
		"semester_number_id":            req.SemesterNumberID,
		"limit_grade_id":                req.LimitGradeID,
		"is_mandatory":                  req.IsMandatory,
		"field_study_concentration_id": req.FieldStudyConcentrationID,
		"created_at":                    time.Now().UnixMilli(),
		"created_by":                    user.ID,
	}

	return db.WithContext(ctx).Table("mst_study_program_curriculums").Create(item).Error
}

/* Read */
func (r *MstStudyProgramCurriculumRepository) GetByID(db *gorm.DB, ID string, data *model.MstStudyProgramCurriculum) error {
	return db.Where("id = ? AND (deleted_at IS NULL OR deleted_at = 0)", ID).First(data).Error
}

func (r *MstStudyProgramCurriculumRepository) GetByStudyProgramIDAndSemesterID(db *gorm.DB, deleted bool, StudyProgramID string, SemesterNumberID string, CurriculumYearID string) (T []model.MstStudyProgramCurriculum, err error) {
	query := db.Model(&model.MstStudyProgramCurriculum{}).
		Where("study_program_id = ? AND semester_number_id = ? AND curriculum_year_id = ?", StudyProgramID, SemesterNumberID, CurriculumYearID)

	if deleted {
		query = query.Where("deleted_at IS NOT NULL AND deleted_at > 0")
	} else {
		query = query.Where("deleted_at IS NULL OR deleted_at = 0")
	}

	err = query.Find(&T).Error
	return T, err
}

/* Update */
func (r *MstStudyProgramCurriculumRepository) UpdateByID(db *gorm.DB, ctx context.Context, req dto.MstStudyProgramCurriculumRequest) error {
	user := ctx.Value("x-user-claims").(*auth.UserClaimsSpesifikRole)

	return db.WithContext(ctx).Table("mst_study_program_curriculums").
		Where("id = ?", req.ID).
		Updates(map[string]interface{}{
			"study_program_id":              req.StudyProgramID,
			"curriculum_year_id":            req.CurriculumYearID,
			"subject_id":                    req.SubjectID,
			"semester_number_id":            req.SemesterNumberID,
			"limit_grade_id":                req.LimitGradeID,
			"is_mandatory":                  req.IsMandatory,
			"field_study_concentration_id": req.FieldStudyConcentrationID,
			"updated_at":                    time.Now().UnixMilli(),
			"updated_by":                    user.ID,
		}).Error
}

func (r *MstStudyProgramCurriculumRepository) UpdateBlastPackageBySemesterWithProgramStudy(
	db *gorm.DB,
	req dto.UpdatePackageMstStudyProgramCurriculumWithStudyProgramRequest,
) (err error) {
	return db.Table("mst_study_program_curriculums").
		Where("semester_number_id = ? AND curriculum_year_id = ? AND study_program_id = ?", req.SemesterNumberID, req.CurriculumYearID, req.StudyProgramID).
		Update("is_package", req.IsPackage).Error
}

func (r *MstStudyProgramCurriculumRepository) UpdateBlastPackageBySemesterWithoutProgramStudy(
	db *gorm.DB,
	req dto.UpdatePackageMstStudyProgramCurriculumWithoutStudyProgramRequest,
) (err error) {
	return db.Table("mst_study_program_curriculums").
		Where("semester_number_id = ? AND curriculum_year_id = ?", req.SemesterNumberID, req.CurriculumYearID).
		Update("is_package", req.IsPackage).Error
}

func (r *MstStudyProgramCurriculumRepository) RestoreByID(db *gorm.DB, ID string) error {
	return db.Table("mst_study_program_curriculums").Where("id = ?", ID).Update("deleted_at", nil).Error
}

/* Delete */
func (r *MstStudyProgramCurriculumRepository) DeleteByID(db *gorm.DB, ctx context.Context, ID string) error {
	user := ctx.Value("x-user-claims").(*auth.UserClaimsSpesifikRole)

	return db.WithContext(ctx).Table("mst_study_program_curriculums").
		Where("id = ?", ID).
		Updates(map[string]interface{}{
			"deleted_at": time.Now().UnixMilli(),
			"deleted_by": user.ID,
		}).Error
}

/* Read */
func (r *MstStudyProgramCurriculumRepository) GetByStudyProgramAndSemesterAndCuricullumForSubjectData(
	db *gorm.DB,
	role model.Role,
	req dto.GetStudyProgramCurriculumRequest,
) (T []model.MstStudyProgramCurriculum, err error) {
	query := db.Model(&model.MstStudyProgramCurriculum{}).
		Where("curriculum_year_id = ? AND semester_number_id = ? AND (deleted_at IS NULL OR deleted_at = 0)", req.CurriculumYearID, req.SemesterNumberID)

	if role == model.Academic && req.StudyProgramID != "" {
		query = query.Where("study_program_id = ?", req.StudyProgramID)
	}

	err = query.Find(&T).Error
	return T, err
}

package repositorymodel

import (
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository"
	"unsia.ac.id/akademic_be/internal/repository/cached"
	"unsia.ac.id/akademic_be/pkg/auth"
)

type MstStudentStudyProgramRepository struct {
	repository.Repository[model.MstStudentStudyProgram]
	log             *logrus.Logger
	cacheRepository cached.CacheRepository
}

func NewMstStudentStudyProgramRepository(log *logrus.Logger,
	cacheRepository cached.CacheRepository,
) *MstStudentStudyProgramRepository {
	return &MstStudentStudyProgramRepository{
		log:             log,
		cacheRepository: cacheRepository,
	}
}

/* Read */
func (r *MstStudentStudyProgramRepository) GetAllWithCountByProgramHeadID(db *gorm.DB, ctx context.Context, pg pageable.PageableRequestInterface) (T []model.MstStudentStudyProgram, count int64, err error) {
	user := ctx.Value("x-user-claims").(*auth.UserClaimsSpesifikRole)
	query := db.WithContext(ctx).Model(&model.MstStudentStudyProgram{}).
		Where("program_head_user_id = ? AND (deleted_at IS NULL OR deleted_at = 0)", user.ID)

	search := pg.GetDefaultSearch()
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

// Search
func (r *MstStudentStudyProgramRepository) GetAllWithCountSearchByStudyProgram(
	db *gorm.DB, pg pageable.PageableRequestByStudyProgram,
) (T []model.MstStudentStudyProgram, count int64, err error) {
	query := db.Model(&model.MstStudentStudyProgram{}).
		Where("study_program_id = ? AND (deleted_at IS NULL OR deleted_at = 0)", pg.StudyProgramID)

	search := pg.GetDefaultSearch()
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

/* Update */

/* Delete */

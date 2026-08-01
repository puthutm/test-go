package repositorymodel

import (
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository"
	"unsia.ac.id/akademic_be/internal/repository/cached"
	"unsia.ac.id/akademic_be/pkg/auth"
)

type MstStudyProgramRepository struct {
	repository.Repository[model.MstStudyProgram]
	log             *logrus.Logger
	cacheRepository cached.CacheRepository
}

func NewMstStudyProgramRepository(log *logrus.Logger,
	cacheRepository cached.CacheRepository,
) *MstStudyProgramRepository {
	return &MstStudyProgramRepository{
		log:             log,
		cacheRepository: cacheRepository,
	}
}

/* Create */
/* Read */
func (r *MstStudyProgramRepository) GetByLecturerIDandActiveAcademicPeriod(db *gorm.DB, ctx context.Context) (T []model.MstStudyProgram, err error) {
	user := ctx.Value("x-user-claims").(*auth.UserClaimsSpesifikRole)
	err = db.WithContext(ctx).Table("mst_study_programs").
		Joins("JOIN mst_class_lecturers ON mst_class_lecturers.study_program_id = mst_study_programs.id").
		Where("mst_class_lecturers.lecturer_user_id = ? AND (mst_study_programs.deleted_at IS NULL OR mst_study_programs.deleted_at = 0)", user.ID).
		Scan(&T).Error
	return T, err
}

/* Update */

/* Delete */

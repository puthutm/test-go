package repositorymodel

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository"
	"unsia.ac.id/akademic_be/internal/repository/cached"
)

type MstStudyProgramCurriculumSubjectPrerequisiteRepository struct {
	repository.Repository[model.MstStudyProgramCurriculumSubjectPrerequisite]
	log             *logrus.Logger
	cacheRepository cached.CacheRepository
}

func NewMstStudyProgramCurriculumSubjectPrerequisiteRepository(
	log *logrus.Logger,
	cacheRepository cached.CacheRepository,
) *MstStudyProgramCurriculumSubjectPrerequisiteRepository {
	return &MstStudyProgramCurriculumSubjectPrerequisiteRepository{
		log:             log,
		cacheRepository: cacheRepository,
	}
}

// Create
func (r *MstStudyProgramCurriculumSubjectPrerequisiteRepository) Create(db *gorm.DB, req model.MstStudyProgramCurriculumSubjectPrerequisite) error {
	return db.Create(&req).Error
}

func (r *MstStudyProgramCurriculumSubjectPrerequisiteRepository) DeleteByStudyProgramCurriculumID(
	db *gorm.DB,
	StudyProgramCurriculumID string,
) error {
	return db.Where("study_program_curriculum_id = ?", StudyProgramCurriculumID).Delete(&model.MstStudyProgramCurriculumSubjectPrerequisite{}).Error
}

func (r *MstStudyProgramCurriculumSubjectPrerequisiteRepository) GetByStudyProgramCurriculumID(
	db *gorm.DB,
	StudyProgramCurriculumID string,
) ([]model.MstStudyProgramCurriculumSubjectPrerequisite, error) {
	var T []model.MstStudyProgramCurriculumSubjectPrerequisite
	err := db.Where("study_program_curriculum_id = ? AND (deleted_at IS NULL OR deleted_at = 0)", StudyProgramCurriculumID).Find(&T).Error
	return T, err
}

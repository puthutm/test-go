package servicemodel

import (
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/dto/converter"
	"unsia.ac.id/akademic_be/internal/repository/cached"
	repositorymodel "unsia.ac.id/akademic_be/internal/repository/model"
	"unsia.ac.id/akademic_be/pkg/utils"
)

type MstStudyProgramService struct {
	log                       *logrus.Logger
	db                        *gorm.DB
	cache                     cached.CacheRepository
	MstStudyProgramRepository *repositorymodel.MstStudyProgramRepository
}

func NewMstStudyProgramService(
	log *logrus.Logger,
	db *gorm.DB,
	cache cached.CacheRepository,
	MstStudyProgramRepository *repositorymodel.MstStudyProgramRepository,
) *MstStudyProgramService {
	return &MstStudyProgramService{
		log:                       log,
		db:                        db,
		cache:                     cache,
		MstStudyProgramRepository: MstStudyProgramRepository,
	}
}

/* Create */

/* Read */
func (s *MstStudyProgramService) GetByLecturerIDandActiveAcademicPeriod(ctx context.Context) (res []dto.DistributionOfStudyProgramResponse, err error) {
	tx := s.db.WithContext(ctx)

	studyPrograms, err := s.MstStudyProgramRepository.GetByLecturerIDandActiveAcademicPeriod(tx, ctx)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	results := make([]dto.DistributionOfStudyProgramResponse, 0)
	for _, studyProgram := range studyPrograms {
		results = append(results, *converter.MstStudyProgramModelToDistributionOfStudyProgramResponse(studyProgram))
	}

	return results, nil
}

/* Update */

/* Delete */

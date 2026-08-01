package servicemodel

import (
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/dto/converter"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository/cached"
	repositorymodel "unsia.ac.id/akademic_be/internal/repository/model"
	"unsia.ac.id/akademic_be/pkg/utils"
)

type MstStudentEducationService struct {
	db                            *gorm.DB
	log                           *logrus.Logger
	cache                         cached.CacheRepository
	mstStudentEducationRepository *repositorymodel.MstStudentEducationRepository
}

func NewMstStudentEducationService(
	log *logrus.Logger,
	db *gorm.DB,
	cache cached.CacheRepository,
	mstStudentEducationRepository *repositorymodel.MstStudentEducationRepository,
) *MstStudentEducationService {
	return &MstStudentEducationService{
		log:                           log,
		db:                            db,
		cache:                         cache,
		mstStudentEducationRepository: mstStudentEducationRepository,
	}
}

/* Create */
/* Read */
func (s *MstStudentEducationService) GetByStudentID(ctx context.Context, StudentID string) (*dto.MstStudentEducationResponse, error) {
	tx := s.db.WithContext(ctx)

	data := new(model.MstStudentEducation)

	err := s.mstStudentEducationRepository.GetByStudentID(tx, StudentID, data)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), StudentID, utils.ErrorLocation())
		utils.PrintMsgDebuging(createMsg)
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	return converter.MstStudentEducationModelToResponse(*data), nil
}

/* Update */
func (s *MstStudentEducationService) UpdateByID(ctx context.Context, req dto.MstStudentEducationRequest) error {
	tx := s.db.WithContext(ctx)

	err := s.mstStudentEducationRepository.UpdateByID(tx, req)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		utils.PrintMsgDebuging(createMsg)
		s.log.WithError(err).Error(createMsg)

		return utils.ErrorSpToErrorFiber(err)

	}

	return nil
}

/* Delete */

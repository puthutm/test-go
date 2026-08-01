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

type MstStudentFamilyService struct {
	db                         *gorm.DB
	log                        *logrus.Logger
	cache                      cached.CacheRepository
	mstStudentFamilyRepository *repositorymodel.MstStudentFamilyRepository
}

func NewMstStudentFamilyService(
	log *logrus.Logger,
	db *gorm.DB,
	cache cached.CacheRepository,
	mstStudentFamilyRepository *repositorymodel.MstStudentFamilyRepository,
) *MstStudentFamilyService {
	return &MstStudentFamilyService{
		log:                        log,
		db:                         db,
		cache:                      cache,
		mstStudentFamilyRepository: mstStudentFamilyRepository,
	}
}

/* Create */
/* Read */
func (s *MstStudentFamilyService) GetByStudentID(ctx context.Context, StudentID string, parentType string) (*dto.MstStudentFamilyResponse, error) {
	tx := s.db.WithContext(ctx)

	data := new(model.MstStudentFamily)

	err := s.mstStudentFamilyRepository.GetByStudentID(tx, StudentID, parentType, data)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), StudentID, utils.ErrorLocation())
		utils.PrintMsgDebuging(createMsg)
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	return converter.MstStudentFamilyModelToResponse(*data), nil
}

/* Update */
func (s *MstStudentFamilyService) UpdateByStudentID(ctx context.Context, req dto.MstStudentFamilyRequest) error {
	tx := s.db.WithContext(ctx)

	err := s.mstStudentFamilyRepository.UpdateByStudentID(tx, req)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		utils.PrintMsgDebuging(createMsg)
		s.log.WithError(err).Error(createMsg)

		return utils.ErrorSpToErrorFiber(err)

	}

	return nil
}

/* Delete */

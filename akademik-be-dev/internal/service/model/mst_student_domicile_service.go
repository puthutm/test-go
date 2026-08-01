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

type MstStudentDomicileService struct {
	db                           *gorm.DB
	log                          *logrus.Logger
	cache                        cached.CacheRepository
	mstStudentDomicileRepository *repositorymodel.MstStudentDomicileRepository
}

func NewMstStudentDomicileService(
	log *logrus.Logger,
	db *gorm.DB,
	cache cached.CacheRepository,
	mstStudentDomicileRepository *repositorymodel.MstStudentDomicileRepository,
) *MstStudentDomicileService {
	return &MstStudentDomicileService{
		log:                          log,
		db:                           db,
		cache:                        cache,
		mstStudentDomicileRepository: mstStudentDomicileRepository,
	}
}

/* Create */
/* Read */
func (s *MstStudentDomicileService) GetByStudentID(ctx context.Context, StudentID string) (*dto.MstStudentDomicileResponse, error) {
	tx := s.db.WithContext(ctx)

	data := new(model.MstStudentDomicile)

	err := s.mstStudentDomicileRepository.GetByStudentID(tx, StudentID, data)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), StudentID, utils.ErrorLocation())
		utils.PrintMsgDebuging(createMsg)
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	return converter.MstStudentDomicileModelToResponse(*data), nil
}

/* Update */
func (s *MstStudentDomicileService) UpdateByStudentID(ctx context.Context, req dto.MstStudentDomicileRequest) error {
	tx := s.db.WithContext(ctx)

	err := s.mstStudentDomicileRepository.UpdateByStudentID(tx, req)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		utils.PrintMsgDebuging(createMsg)
		s.log.WithError(err).Error(createMsg)

		return utils.ErrorSpToErrorFiber(err)

	}

	return nil
}

/* Delete */

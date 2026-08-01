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

type MstBiodataOfUserService struct {
	log                        *logrus.Logger
	db                         *gorm.DB
	cache                      cached.CacheRepository
	mstBiodataOfUserRepository *repositorymodel.MstBiodataOfUserRepository
}

func NewMstBiodataOfUserService(
	log *logrus.Logger,
	db *gorm.DB,
	cache cached.CacheRepository,
	mstBiodataOfUserRepository *repositorymodel.MstBiodataOfUserRepository,
) *MstBiodataOfUserService {
	return &MstBiodataOfUserService{
		log:                        log,
		db:                         db,
		cache:                      cache,
		mstBiodataOfUserRepository: mstBiodataOfUserRepository,
	}
}

/* Create */
/* Read */
func (s *MstBiodataOfUserService) GetByUserID(ctx context.Context, UserID string) (*dto.MstBiodataOfUserResponse, error) {
	tx := s.db.WithContext(ctx)

	biodataOfUser := new(model.MstBiodataOfUser)

	err := s.mstBiodataOfUserRepository.GetByUserID(tx, UserID, biodataOfUser)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), UserID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	result := converter.MstBiodataOfUserToResponse(*biodataOfUser)

	return result, nil
}

/* Update */
/* Delete */

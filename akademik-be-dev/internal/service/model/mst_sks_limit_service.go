package servicemodel

import (
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/dto/converter"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository/cached"
	repositorymodel "unsia.ac.id/akademic_be/internal/repository/model"
	"unsia.ac.id/akademic_be/pkg/utils"
)

type MstSKSLimitService struct {
	log                   *logrus.Logger
	db                    *gorm.DB
	cache                 cached.CacheRepository
	mstSKSLimitRepository *repositorymodel.MstSKSLimitRepository
}

func NewMstSKSLimitService(
	log *logrus.Logger,
	db *gorm.DB,
	cache cached.CacheRepository,
	mstSKSLimitRepository *repositorymodel.MstSKSLimitRepository,
) *MstSKSLimitService {
	return &MstSKSLimitService{
		log:                   log,
		db:                    db,
		cache:                 cache,
		mstSKSLimitRepository: mstSKSLimitRepository,
	}
}

/* Create */
func (s *MstSKSLimitService) Create(ctx context.Context, req dto.MstSKSLimitRequest) (*dto.MstSKSLimitResponse, error) {
	tx := s.db.WithContext(ctx)

	err := s.mstSKSLimitRepository.Create(tx, ctx, req)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	sksLimit, err := s.GetByID(ctx, req.ID)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), req.ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	return sksLimit, nil
}

/* Read */
func (s *MstSKSLimitService) GetAllAndCount(ctx context.Context, pageble pageable.PageableRequestInterface) (*pageable.PageableResponse[dto.MstSKSLimitResponse], error) {
	tx := s.db.WithContext(ctx)

	sksLimits, totalData, err := s.mstSKSLimitRepository.GetAllWithCount(tx, false, pageble)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	res := make([]dto.MstSKSLimitResponse, 0, totalData)
	for _, sksLimit := range sksLimits {
		res = append(res, *converter.MstSKSLimitModelToResponse(sksLimit))
	}

	return &pageable.PageableResponse[dto.MstSKSLimitResponse]{
		Data: res,
		Metadata: pageable.Metadata{
			TotalData: totalData,
			TotalPage: utils.TotalPage(totalData, pageble.GetDefaultLimit()),
			Page:      pageble.GetDefaultPage(),
			Size:      pageble.GetDefaultLimit(),
		},
	}, nil
}
func (s *MstSKSLimitService) GetAllTrashWithCount(ctx context.Context, pageble pageable.PageableRequestInterface) (*pageable.PageableResponse[dto.MstSKSLimitResponse], error) {
	tx := s.db.WithContext(ctx)

	registrationDetails, totalData, err := s.mstSKSLimitRepository.GetAllWithCount(tx, true, pageble)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	res := make([]dto.MstSKSLimitResponse, 0, totalData)
	for _, v := range registrationDetails {
		res = append(res, *converter.MstSKSLimitModelToResponse(v))
	}

	return &pageable.PageableResponse[dto.MstSKSLimitResponse]{
		Data: res,
		Metadata: pageable.Metadata{
			TotalData: totalData,
			TotalPage: utils.TotalPage(totalData, pageble.GetDefaultLimit()),
			Page:      pageble.GetDefaultPage(),
			Size:      pageble.GetDefaultLimit(),
		},
	}, nil
}

func (s *MstSKSLimitService) GetByID(ctx context.Context, ID string) (*dto.MstSKSLimitResponse, error) {
	tx := s.db.WithContext(ctx)

	sksLimitModel := new(model.MstSKSLimit)

	err := s.mstSKSLimitRepository.GetByID(tx, ID, sksLimitModel)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	sksLimit := converter.MstSKSLimitModelToResponse(*sksLimitModel)

	return sksLimit, nil
}

/* Update */
func (s *MstSKSLimitService) UpdateByID(ctx context.Context, req dto.MstSKSLimitRequest) (*dto.MstSKSLimitResponse, error) {
	tx := s.db.WithContext(ctx)

	err := s.mstSKSLimitRepository.UpdateByID(tx, ctx, req)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), req.ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	sksLimit, err := s.GetByID(ctx, req.ID)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), req.ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	return sksLimit, nil
}

func (s *MstSKSLimitService) RestoreByID(ctx context.Context, ID string) error {
	tx := s.db.WithContext(ctx)

	err := s.mstSKSLimitRepository.RestoreByID(tx, ID)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return utils.ErrorSpToErrorFiber(err)
	}

	return nil
}

/* Delete */
func (s *MstSKSLimitService) DeleteByID(ctx context.Context, ID string) error {
	tx := s.db.WithContext(ctx)

	err := s.mstSKSLimitRepository.DeleteByID(tx, ctx, ID)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return utils.ErrorSpToErrorFiber(err)
	}

	return nil
}

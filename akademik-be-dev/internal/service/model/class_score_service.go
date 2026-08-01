package servicemodel

import (
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/dto/converter"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	repositorymodel "unsia.ac.id/akademic_be/internal/repository/model"
	"unsia.ac.id/akademic_be/pkg/utils"
)

type ClassScoreService interface {
	GetByClassID(ctx context.Context, pageable pageable.PageableRequestClassScore) (*dto.ClassScorePageableResponse, error)
	CheckSaveButton(ctx context.Context, classID string) (*dto.ClassScoreCheckSaveButtonResponse, error)
	UpdateStatusLock(ctx context.Context, request dto.UpdateStatusLockRequest) error
}

type classScoreService struct {
	log            *logrus.Logger
	db             *gorm.DB
	classScoreRepo repositorymodel.ClassScoreRepository
}

func NewClassScoreService(
	log *logrus.Logger,
	db *gorm.DB,
	classScoreRepo repositorymodel.ClassScoreRepository,
) ClassScoreService {
	return &classScoreService{
		log:            log,
		db:             db,
		classScoreRepo: classScoreRepo,
	}
}

func (s *classScoreService) GetByClassID(ctx context.Context, pageble pageable.PageableRequestClassScore) (*dto.ClassScorePageableResponse, error) {
	tx := s.db.WithContext(ctx)

	data, totalData, summary, err := s.classScoreRepo.GetByClassID(tx, pageble)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)
		return nil, utils.ErrorSpToErrorFiber(err)
	}

	res := make([]dto.ClassScoreResponse, 0, totalData)
	for _, item := range data {
		res = append(res, *converter.ClassScoreModelToResponse(item))
	}

	var summaryRes *dto.ClassScoreSummaryResponse
	if summary != nil {
		summaryRes = converter.ClassScoreSummaryModelToResponse(*summary)
	}

	return &dto.ClassScorePageableResponse{
		Data:    res,
		Summary: summaryRes,
		Metadata: dto.PageableMetadata{
			TotalData: totalData,
			TotalPage: utils.TotalPage(totalData, pageble.GetDefaultLimit()),
			Page:      pageble.GetDefaultPage(),
			Size:      pageble.GetDefaultLimit(),
		},
	}, nil
}

func (s *classScoreService) CheckSaveButton(ctx context.Context, classID string) (*dto.ClassScoreCheckSaveButtonResponse, error) {
	tx := s.db.WithContext(ctx)

	statusLock, err := s.classScoreRepo.CheckSaveButton(tx, classID)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), classID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)
		return nil, utils.ErrorSpToErrorFiber(err)
	}

	return &dto.ClassScoreCheckSaveButtonResponse{
		StatusLock: statusLock,
	}, nil
}

func (s *classScoreService) UpdateStatusLock(ctx context.Context, req dto.UpdateStatusLockRequest) error {
	tx := s.db.WithContext(ctx)

	err := s.classScoreRepo.UpdateStatusLock(tx, req.ClassID, req.StatusLocked, req.CreatedBy)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), req.ClassID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)
		return utils.ErrorSpToErrorFiber(err)
	}

	return nil
}

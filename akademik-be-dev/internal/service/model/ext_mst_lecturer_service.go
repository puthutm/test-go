package servicemodel

import (
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/dto/converter"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	"unsia.ac.id/akademic_be/internal/repository/cached"
	repositorymodel "unsia.ac.id/akademic_be/internal/repository/model"
	"unsia.ac.id/akademic_be/pkg/utils"
)

type MstLecturerService struct {
	log                   *logrus.Logger
	db                    *gorm.DB
	cache                 cached.CacheRepository
	mstLecturerRepository *repositorymodel.MstLecturerRepository
}

func NewMstLecturerService(
	log *logrus.Logger,
	db *gorm.DB,
	cache cached.CacheRepository,
	mstLecturerRepository *repositorymodel.MstLecturerRepository,
) *MstLecturerService {
	return &MstLecturerService{
		log:                   log,
		db:                    db,
		cache:                 cache,
		mstLecturerRepository: mstLecturerRepository,
	}
}

/* Create */

/* Read */
func (s *MstLecturerService) GetAllWithCountByProgramHeadID(ctx context.Context, pageble pageable.PageableRequestInterface) (*pageable.PageableResponse[dto.MstLecturerByProgramHeadResponse], error) {
	tx := s.db.WithContext(ctx)

	studentStudyPrograms, totalData, err := s.mstLecturerRepository.GetAllWithCountByProgramHeadID(tx, ctx, pageble)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	res := make([]dto.MstLecturerByProgramHeadResponse, 0, totalData)
	for _, studentStudyProgram := range studentStudyPrograms {
		res = append(res, *converter.MstLecturerModelByProgramHeadToResponse(studentStudyProgram))
	}

	return &pageable.PageableResponse[dto.MstLecturerByProgramHeadResponse]{
		Data: res,
		Metadata: pageable.Metadata{
			TotalData: totalData,
			TotalPage: utils.TotalPage(totalData, pageble.GetDefaultLimit()),
			Page:      pageble.GetDefaultPage(),
			Size:      pageble.GetDefaultLimit(),
		},
	}, nil
}

/* Update */

/* Delete */

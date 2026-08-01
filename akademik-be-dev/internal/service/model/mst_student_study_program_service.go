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

type MstStudentStudyProgramService struct {
	log                              *logrus.Logger
	db                               *gorm.DB
	cache                            cached.CacheRepository
	mstStudentStudyProgramRepository *repositorymodel.MstStudentStudyProgramRepository
}

func NewMstStudentStudyProgramService(
	log *logrus.Logger,
	db *gorm.DB,
	cache cached.CacheRepository,
	mstStudentStudyProgramRepository *repositorymodel.MstStudentStudyProgramRepository,
) *MstStudentStudyProgramService {
	return &MstStudentStudyProgramService{
		log:                              log,
		db:                               db,
		cache:                            cache,
		mstStudentStudyProgramRepository: mstStudentStudyProgramRepository,
	}
}

/* Create */

/* Read */
func (s *MstStudentStudyProgramService) GetAllWithCountByProgramHeadID(ctx context.Context, pageble pageable.PageableRequestInterface) (*pageable.PageableResponse[dto.MstStudentStudyProgramByProgramHeadResponse], error) {
	tx := s.db.WithContext(ctx)

	studentStudyPrograms, totalData, err := s.mstStudentStudyProgramRepository.GetAllWithCountByProgramHeadID(tx, ctx, pageble)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	res := make([]dto.MstStudentStudyProgramByProgramHeadResponse, 0, totalData)
	for _, studentStudyProgram := range studentStudyPrograms {
		res = append(res, *converter.MstStudentStudyProgramModelByProgramHeadToResponse(studentStudyProgram))
	}

	return &pageable.PageableResponse[dto.MstStudentStudyProgramByProgramHeadResponse]{
		Data: res,
		Metadata: pageable.Metadata{
			TotalData: totalData,
			TotalPage: utils.TotalPage(totalData, pageble.GetDefaultLimit()),
			Page:      pageble.GetDefaultPage(),
			Size:      pageble.GetDefaultLimit(),
		},
	}, nil
}

func (s *MstStudentStudyProgramService) GetAllWithCountSearchByStudyProgram(
	ctx context.Context, pageble pageable.PageableRequestByStudyProgram,
) (*pageable.PageableResponse[dto.MstStudentStudyProgramSearchResponse], error) {
	tx := s.db.WithContext(ctx)

	studentStudyPrograms, totalData, err := s.mstStudentStudyProgramRepository.GetAllWithCountSearchByStudyProgram(tx, pageble)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithFields(logrus.Fields{
			"service":  "get program study studentsearch by study program",
			"pageable": pageble,
		}).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	res := make([]dto.MstStudentStudyProgramSearchResponse, 0, len(studentStudyPrograms))
	for _, studentStudyProgram := range studentStudyPrograms {
		res = append(res, converter.MstStudentStudyProgramSearchToResponse(studentStudyProgram))
	}

	return &pageable.PageableResponse[dto.MstStudentStudyProgramSearchResponse]{
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

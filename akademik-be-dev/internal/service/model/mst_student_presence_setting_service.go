// Package servicemodel
package servicemodel

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	icemstutil "unsia.ac.id/akademic_be/pkg/icems-tools/utils"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/delivery/http/middleware"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/dto/converter"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	"unsia.ac.id/akademic_be/internal/repository/cached"
	repositorymodel "unsia.ac.id/akademic_be/internal/repository/model"
	"unsia.ac.id/akademic_be/pkg/utils"
)

type MstStudentPresenceSettingService struct {
	db                                  *gorm.DB
	log                                 *logrus.Logger
	cache                               cached.CacheRepository
	mstStudentPresenceSettingRepository *repositorymodel.MstStudentPresenceSettingRepository
}

func NewMstStudentPresenceSettingService(
	db *gorm.DB,
	log *logrus.Logger,
	cache cached.CacheRepository,
	mstStudentPresenceSettingRepository *repositorymodel.MstStudentPresenceSettingRepository,
) *MstStudentPresenceSettingService {
	return &MstStudentPresenceSettingService{
		db:                                  db,
		log:                                 log,
		cache:                               cache,
		mstStudentPresenceSettingRepository: mstStudentPresenceSettingRepository,
	}
}

func (s *MstStudentPresenceSettingService) Create(
	ctx context.Context, req dto.MstStudentPresenceSettingCreateRequest,
) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	studentPresenceSettingM := converter.ConvertMstStudentPresenceSettingRequestToMstStudentPresenceSettingModel(req, user.ID)

	err := s.mstStudentPresenceSettingRepository.Create(ctx, tx, studentPresenceSettingM)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithFields(logrus.Fields{
			"service": "create studentPresenceSetting",
			"req":     req,
		}).Error(createMsg)

		return icemstutil.ErrorSpToErrorFiber(err)
	}
	return nil
}

func (s *MstStudentPresenceSettingService) Duplicate(
	ctx context.Context, req dto.MstStudentPresenceSettingDuplicateRequest,
) error {
	tx := s.db.WithContext(ctx)

	if req.AcademicPeriodeID == req.AcademicPeriodeIDTarget {
		return fiber.NewError(fiber.StatusConflict, "Student presence already exists.")
	}

	user := middleware.GetUserClaimsCtx(ctx)

	studentPresenceSettingM := converter.ConvertMstStudentPresenceSettingDuplicateRequestToParam(req, user.ID)

	err := s.mstStudentPresenceSettingRepository.Duplicate(ctx, tx, studentPresenceSettingM)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithFields(logrus.Fields{
			"service": "duplicate studentPresenceSetting",
			"req":     req,
		}).Error(createMsg)

		return icemstutil.ErrorSpToErrorFiber(err)
	}
	return nil
}

func (s *MstStudentPresenceSettingService) GetAllWithCount(
	ctx context.Context, pageble pageable.PageableStudentSettingGets,
) (*pageable.PageableResponse[dto.MstStudentPresenceSettingGetResultResponse], error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	datas, count, err := s.mstStudentPresenceSettingRepository.GetAllWithCount(
		ctx, tx, pageble,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get all studentPresenceSetting service",
			"user-id": user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return nil, err
	}

	r := make([]dto.MstStudentPresenceSettingGetResultResponse, 0, count)

	for _, v := range datas {
		c := converter.ConvertMstStudentPresenceSettingGetResultToResponse(&v)
		r = append(r, c)
	}
	res := &pageable.PageableResponse[dto.MstStudentPresenceSettingGetResultResponse]{
		Data: r,
		Metadata: pageable.Metadata{
			TotalData: count,
			TotalPage: utils.TotalPage(count, pageble.GetDefaultLimit()),
			Page:      pageble.GetDefaultPage(),
			Size:      pageble.GetDefaultLimit(),
		},
	}
	return res, nil
}

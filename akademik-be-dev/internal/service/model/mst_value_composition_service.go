package servicemodel

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/delivery/http/middleware"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/dto/converter"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository/cached"
	repositorymodel "unsia.ac.id/akademic_be/internal/repository/model"
	"unsia.ac.id/akademic_be/pkg/utils"
)

type MstValueCompositionService struct {
	log                           *logrus.Logger
	db                            *gorm.DB
	cache                         cached.CacheRepository
	mstValueCompositionRepository repositorymodel.MstValueCompositionRepository
}

func NewMstValueCompositionService(
	log *logrus.Logger,
	db *gorm.DB,
	cache cached.CacheRepository,
	mstValueCompositionRepository repositorymodel.MstValueCompositionRepository,
) *MstValueCompositionService {
	return &MstValueCompositionService{
		log:                           log,
		db:                            db,
		cache:                         cache,
		mstValueCompositionRepository: mstValueCompositionRepository,
	}
}

// TODO: Create
func (s *MstValueCompositionService) Create(ctx context.Context, req dto.MstValueCompositionRequest) error {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)
	v, _ := utils.StringToUuid(user.ID)
	req.UserID = v

	data := new(model.MstValueComposition)
	converter.ConvertMstValueCompositionRequestToModelPointer(req, data)
	err := s.mstValueCompositionRepository.Create(tx, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "create composition value service",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

// TODO: Update
func (s *MstValueCompositionService) UpdateByID(ctx context.Context, req dto.MstValueCompositionUpdate) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)
	data := new(model.MstValueComposition)
	v, _ := utils.StringToUuid(user.ID)
	req.UserID = v

	converter.ConvertMstValueCompositionUpdateToModelPointer(req, data)
	err := s.mstValueCompositionRepository.UpdateByID(tx, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "update composition value service",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

// TODO: Delete
func (s *MstValueCompositionService) DeleteByID(ctx context.Context, ID string) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	err := s.mstValueCompositionRepository.DeleteByID(tx, ID, user.ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service":  "delete composition value service",
			"id":       ID,
			"deleteby": user.ID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

// TODO: Restore
func (s *MstValueCompositionService) RestoreByID(ctx context.Context, ID string) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	err := s.mstValueCompositionRepository.RestoreByID(tx, ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "Restore composition value service",
			"id":      ID,
			"user-id": user.ID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

// TODO: Read
func (s *MstValueCompositionService) GetByID(ctx context.Context, ID string) (res *dto.MstValueCompositionResponse, err error) {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)

	data := new(model.MstValueComposition)
	err = s.mstValueCompositionRepository.GetByID(tx, ID, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get by id composition value service",
			"id":      ID,
			"user-id": user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return
	}
	res = converter.ConvertModelToMstValueCompositionResponsePointer(data)
	return
}

func (s *MstValueCompositionService) GetAllWithCount(ctx context.Context, req pageable.PageableRequestValueComposition) (res *pageable.PageableResponse[dto.MstValueCompositionResponse], err error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	datas, count, err := s.mstValueCompositionRepository.GetAllWithCount(tx, true, req)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service":  "get all composition value service",
			"user-id":  user.ID,
			"pageable": req,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return
	}

	r := make([]dto.MstValueCompositionResponse, 0, count)

	for _, v := range datas {
		c := converter.ConvertModelToMstValueCompositionResponse(&v)
		r = append(r, c)
	}
	res = &pageable.PageableResponse[dto.MstValueCompositionResponse]{
		Data: r,
		Metadata: pageable.Metadata{
			TotalData: count,
			TotalPage: utils.TotalPage(count, req.GetDefaultLimit()),
			Page:      req.GetDefaultPage(),
			Size:      req.GetDefaultLimit(),
		},
	}
	return
}

func (s *MstValueCompositionService) GetAcademicPeriodsWithCount(ctx context.Context, req pageable.PageableRequest) (res *pageable.PageableResponse[dto.ValueCompositionGroupResponse], err error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	datas, count, err := s.mstValueCompositionRepository.GetAcademicPeriodsWithCount(tx, req)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service":  "get academic periods with count service",
			"user-id":  user.ID,
			"pageable": req,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return
	}

	r := make([]dto.ValueCompositionGroupResponse, 0, count)

	for _, v := range datas {
		c := converter.ConvertModelToValueCompositionGroupResponse(&v)
		r = append(r, c)
	}
	res = &pageable.PageableResponse[dto.ValueCompositionGroupResponse]{
		Data: r,
		Metadata: pageable.Metadata{
			TotalData: count,
			TotalPage: utils.TotalPage(count, req.GetDefaultLimit()),
			Page:      req.GetDefaultPage(),
			Size:      req.GetDefaultLimit(),
		},
	}
	return
}

func (s *MstValueCompositionService) GetAllTrashWithCount(ctx context.Context, req pageable.PageableRequestValueComposition) (res *pageable.PageableResponse[dto.MstValueCompositionResponse], err error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	datas, count, err := s.mstValueCompositionRepository.GetAllWithCount(tx, false, req)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service":  "get all trash composition value service",
			"user-id":  user.ID,
			"pageable": req,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return
	}

	r := make([]dto.MstValueCompositionResponse, 0, count)

	for _, v := range datas {
		c := converter.ConvertModelToMstValueCompositionResponse(&v)
		r = append(r, c)
	}
	res = &pageable.PageableResponse[dto.MstValueCompositionResponse]{
		Data: r,
		Metadata: pageable.Metadata{
			TotalData: count,
			TotalPage: utils.TotalPage(count, req.GetDefaultLimit()),
			Page:      req.GetDefaultPage(),
			Size:      req.GetDefaultLimit(),
		},
	}
	return
}

func (s *MstValueCompositionService) DuplicateByAcademicPeriodID(ctx context.Context, req dto.MstValueCompositionDuplicateRequest) error {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)

	err := s.mstValueCompositionRepository.DuplicateByAcademicPeriodID(tx, req, time.Now().UnixMilli(), user.ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "duplicate composition value service",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

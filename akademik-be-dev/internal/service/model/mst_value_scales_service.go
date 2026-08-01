package servicemodel

import (
	"context"

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

type MstValueScaleService struct {
	log                     *logrus.Logger
	db                      *gorm.DB
	cache                   cached.CacheRepository
	mstValueScaleRepository *repositorymodel.MstValueScaleRepository
}

func NewMstValueScaleService(
	log *logrus.Logger,
	db *gorm.DB,
	cache cached.CacheRepository,
	mstValueScaleRepository *repositorymodel.MstValueScaleRepository,
) *MstValueScaleService {
	return &MstValueScaleService{
		log:                     log,
		db:                      db,
		cache:                   cache,
		mstValueScaleRepository: mstValueScaleRepository,
	}
}

// TODO: Create
func (s *MstValueScaleService) Create(ctx context.Context, req dto.MstValueScaleRequest) error {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)
	v, _ := utils.StringToUuid(user.ID)
	req.UserID = &v

	data := new(model.MstValueScale)
	converter.ConvertMstValueScaleRequestToModelPointer(req, data)
	err := s.mstValueScaleRepository.Create(tx, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "create scale value service",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

// TODO: Update
func (s *MstValueScaleService) UpdateByID(ctx context.Context, req dto.MstValueScaleUpdate) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)
	data := new(model.MstValueScale)
	v, _ := utils.StringToUuid(user.ID)
	req.UserID = &v

	converter.ConvertMstValueScaleUpdateToModelPointer(req, data)
	err := s.mstValueScaleRepository.Update(tx, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "update scale value service",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

// TODO: Delete
func (s *MstValueScaleService) DeleteByID(ctx context.Context, ID string) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	err := s.mstValueScaleRepository.DeleteByID(tx, ID, user.ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service":  "delete scale value service",
			"id":       ID,
			"deleteby": user.ID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

// TODO: Restore
func (s *MstValueScaleService) RestoreByID(ctx context.Context, ID string) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	err := s.mstValueScaleRepository.RestoreByID(tx, ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "Restore scale value service",
			"id":      ID,
			"user-id": user.ID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

// TODO: Read
func (s *MstValueScaleService) GetByID(ctx context.Context, ID string) (res *dto.MstValueScaleResponse, err error) {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)

	data := new(model.MstValueScale)
	err = s.mstValueScaleRepository.GetByID(tx, ID, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get by id scale value service",
			"id":      ID,
			"user-id": user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return
	}
	res = converter.ConvertModelToMstValueScaleResponsePointer(data)
	return
}

func (s *MstValueScaleService) GetAllWithCount(ctx context.Context, req pageable.PageableRequestValueScale) (res *pageable.PageableResponse[dto.MstValueScaleResponse], err error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	datas, count, err := s.mstValueScaleRepository.GetAllWithCount(tx, true, req)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get all scale value service",
			"user-id": user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return
	}

	r := make([]dto.MstValueScaleResponse, 0, count)

	for _, v := range datas {
		c := converter.ConvertModelToMstValueScaleResponse(&v)
		r = append(r, c)
	}
	res = &pageable.PageableResponse[dto.MstValueScaleResponse]{
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

func (s *MstValueScaleService) GetAllTrashWithCount(ctx context.Context, req pageable.PageableRequestValueScale) (res *pageable.PageableResponse[dto.MstValueScaleResponse], err error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	datas, count, err := s.mstValueScaleRepository.GetAllWithCount(tx, false, req)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get all trash scale value service",
			"user-id": user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return
	}

	r := make([]dto.MstValueScaleResponse, 0, count)

	for _, v := range datas {
		c := converter.ConvertModelToMstValueScaleResponse(&v)
		r = append(r, c)
	}
	res = &pageable.PageableResponse[dto.MstValueScaleResponse]{
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

package servicemodel

import (
	"context"

	"github.com/google/uuid"
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

type MstClassLecturerService struct {
	log                        *logrus.Logger
	db                         *gorm.DB
	cache                      cached.CacheRepository
	mstClassLecturerRepository *repositorymodel.MstClassLecturerRepository
}

func NewMstClassLecturerService(
	log *logrus.Logger,
	db *gorm.DB,
	cache cached.CacheRepository,
	mstClassLecturerRepository *repositorymodel.MstClassLecturerRepository,
) *MstClassLecturerService {
	return &MstClassLecturerService{
		log:                        log,
		db:                         db,
		cache:                      cache,
		mstClassLecturerRepository: mstClassLecturerRepository,
	}
}

// TODO: Create
func (s *MstClassLecturerService) Create(ctx context.Context, req dto.MstClassLecturerRequest) error {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)

	data := new(model.MstClassLecturer)
	converter.ConvertMstClassLecturerRequestToModelPointer(req, data)
	err := s.mstClassLecturerRepository.Create(tx, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "create class lecturer service",
			"request": req,
			"user_id": user.ID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

// TODO: Update
func (s *MstClassLecturerService) UpdateByID(ctx context.Context, req dto.MstClassLecturerUpdate) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)
	data := new(model.MstClassLecturer)

	converter.ConvertMstClassLecturerUpdateToModelPointer(req, data)
	err := s.mstClassLecturerRepository.UpdateByID(tx, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "update class lecture service",
			"request": req,
			"user-id": user.ID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

// TODO: Delete
func (s *MstClassLecturerService) DeleteByID(ctx context.Context, ID string) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	err := s.mstClassLecturerRepository.DeleteByID(tx, ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service":  "delete class lecturer service",
			"id":       ID,
			"deleteby": user.ID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

// TODO: Restore

// TODO: Read
func (s *MstClassLecturerService) GetByID(ctx context.Context, ID string) (res *dto.MstClassLecturerResponse, err error) {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)

	data := new(model.MstClassLecturer)
	err = s.mstClassLecturerRepository.GetByID(tx, ID, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get by id class lecturer service",
			"id":      ID,
			"user-id": user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return res, err
	}
	res = converter.ConvertModelToMstClassLecturerResponsePointer(data)
	return res, err
}

func (s *MstClassLecturerService) GetByClassID(ctx context.Context, classID string) (res *dto.MstClassLecturerResponse, err error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	data := new(model.MstClassLecturer)
	err = s.mstClassLecturerRepository.GetByClassID(tx, classID, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service":  "get by class id class lecturer service",
			"class_id": classID,
			"user-id":  user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return nil, err
	}
	if data.ID == uuid.Nil {
		return nil, nil
	}
	res = converter.ConvertModelToMstClassLecturerResponsePointer(data)
	return res, nil
}

func (s *MstClassLecturerService) GetClassByAcademicPeriodAndSubjectAndUserForLecturerWithCount(
	ctx context.Context, pageble pageable.PageableClassGetByUserAndAcademicPeriodAndSubject,
) (*pageable.PageableResponse[dto.MstClassLecturerGetByAcademicPeriodAndSubjectAndUserForLecturerResponse], error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	datas, count, err := s.mstClassLecturerRepository.GetClassByAcademicPeriodAndSubjectAndUserForLecturerWithCount(
		ctx, tx, user.ID, pageble,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get class by academic and subject and user id service",
			"user-id": user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return nil, err
	}

	r := make([]dto.MstClassLecturerGetByAcademicPeriodAndSubjectAndUserForLecturerResponse, 0, count)

	for _, v := range datas {
		c := converter.ConvertMstClassLecturerGetByAcademicPeriodAndSubjectAndUserForLecturerToResponse(&v)
		r = append(r, c)
	}
	res := &pageable.PageableResponse[dto.MstClassLecturerGetByAcademicPeriodAndSubjectAndUserForLecturerResponse]{
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

func (s *MstClassLecturerService) GetSubjectByClassLecturerWithCount(
	ctx context.Context, pageble pageable.PageableRequestClass,
) (*pageable.PageableResponse[dto.MstClassGetResultSubjectAndClassCountResponse], error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	datas, count, err := s.mstClassLecturerRepository.GetSubjectByClassLecturerWithCount(
		ctx, tx, user.ID, pageble,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get class and subject by academic and user id service",
			"user-id": user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return nil, err
	}

	r := make([]dto.MstClassGetResultSubjectAndClassCountResponse, 0, count)

	for _, v := range datas {
		c := converter.ConvertMstClassGetResultSubjectAndClassCountToResponse(v)
		r = append(r, c)
	}
	res := &pageable.PageableResponse[dto.MstClassGetResultSubjectAndClassCountResponse]{
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

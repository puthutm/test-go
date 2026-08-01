package servicemodel

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/delivery/http/middleware"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/dto/converter"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	repositorymodel "unsia.ac.id/akademic_be/internal/repository/model"
	"unsia.ac.id/akademic_be/pkg/utils"
)

type KrsService interface {
	// Lecturer Students
	GetKrsLecturerStudentsWithCount(ctx context.Context, req pageable.PageableKrsLecturerStudentsRequest) (*pageable.PageableResponse[dto.TrxKrsLecturerStudentResponse], error)
	GetKrsLecturerStudentDetailByKrsHeaderID(ctx context.Context, krsHeaderID string) (*dto.TrxKrsLecturerStudentDetailResponse, error)
	UpdateKrsItemStatusByKrsItemID(ctx context.Context, req dto.TrxKrsLecturerStudentItemUpdateStatusRequest) (*dto.TrxKrsLecturerStudentItemUpdateResponse, error)

	// Pick Class
	GetPickClassesByUserID(ctx context.Context, req dto.TrxKrsPickClassGetRequest) (*dto.TrxKrsPickClassResponse, error)
	TakeClass(ctx context.Context, req dto.TrxKrsTakeClassRequest) (*dto.TrxKrsTakeClassResponse, error)

	// Program Head Classes
	GetKrsProgramHeadClassesWithCount(ctx context.Context, req pageable.PageableKrsProgramHeadClassesRequest) (*pageable.PageableResponse[dto.TrxKrsProgramHeadClassResponse], error)

	// Saved
	GetSavedByUserID(ctx context.Context, req dto.TrxKrsSavedGetRequest) ([]dto.TrxKrsSavedItemResponse, error)
	DeleteSavedByKrsItemID(ctx context.Context, krsItemID string) error

	// Info
	GetKrsMaxSksInfo(ctx context.Context) (*dto.TrxKrsMaxSksInfoResponse, error)
}

type krsService struct {
	log           *logrus.Logger
	db            *gorm.DB
	krsRepository repositorymodel.KrsRepository
}

func NewKrsService(
	log *logrus.Logger,
	db *gorm.DB,
	krsRepository repositorymodel.KrsRepository,
) KrsService {
	return &krsService{
		log:           log,
		db:            db,
		krsRepository: krsRepository,
	}
}

// GetKrsLecturerStudentsWithCount - Get KRS lecturer students with pagination
func (s *krsService) GetKrsLecturerStudentsWithCount(
	ctx context.Context,
	req pageable.PageableKrsLecturerStudentsRequest,
) (*pageable.PageableResponse[dto.TrxKrsLecturerStudentResponse], error) {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)

	students, count, err := s.krsRepository.GetKrsLecturerStudentsWithCount(
		ctx, tx, user.ID, req,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get krs lecturer students",
			"user-id": user.ID,
		}).Error(err)
		return nil, utils.ErrorSpToErrorFiber(err)
	}

	responses := make([]dto.TrxKrsLecturerStudentResponse, 0, count)
	for _, student := range students {
		responses = append(responses, converter.TrxKrsLecturerStudentModelToResponse(student))
	}

	return &pageable.PageableResponse[dto.TrxKrsLecturerStudentResponse]{
		Data: responses,
		Metadata: pageable.Metadata{
			TotalData: count,
			TotalPage: utils.TotalPage(count, req.GetDefaultLimit()),
			Page:      req.GetDefaultPage(),
			Size:      req.GetDefaultLimit(),
		},
	}, nil
}

// GetKrsLecturerStudentDetailByKrsHeaderID - Get KRS student detail by KRS header ID
func (s *krsService) GetKrsLecturerStudentDetailByKrsHeaderID(
	ctx context.Context,
	krsHeaderID string,
) (*dto.TrxKrsLecturerStudentDetailResponse, error) {
	tx := s.db.WithContext(ctx)

	studentDetail, totalSKS, items, err := s.krsRepository.GetKrsLecturerStudentDetailByKrsHeaderID(
		ctx, tx, krsHeaderID,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service":     "get krs lecturer student detail",
			"krs-header-id": krsHeaderID,
		}).Error(err)
		return nil, utils.ErrorSpToErrorFiber(err)
	}

	response := converter.TrxKrsLecturerStudentDetailModelToResponse(studentDetail, totalSKS, items)
	return &response, nil
}

// UpdateKrsItemStatusByKrsItemID - Update KRS item status
func (s *krsService) UpdateKrsItemStatusByKrsItemID(
	ctx context.Context,
	req dto.TrxKrsLecturerStudentItemUpdateStatusRequest,
) (*dto.TrxKrsLecturerStudentItemUpdateResponse, error) {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)

	// Convert request to model
	updateModel := converter.TrxKrsLecturerStudentItemUpdateRequestToModel(req, uuid.MustParse(user.ID))

	// Execute update
	err := s.krsRepository.UpdateKrsItemStatusByKrsItemID(ctx, tx, updateModel)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service":     "update krs item status",
			"krs-item-id": req.KrsItemID,
			"item-status": req.ItemStatus,
		}).Error(err)
		return nil, utils.ErrorSpToErrorFiber(err)
	}

	// Convert model to response
	response := converter.TrxKrsLecturerStudentItemUpdateModelToResponse(updateModel)
	return &response, nil
}

// GetPickClassesByUserID - Get academic periods and classes for KRS filling
func (s *krsService) GetPickClassesByUserID(
	ctx context.Context, req dto.TrxKrsPickClassGetRequest,
) (*dto.TrxKrsPickClassResponse, error) {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)

	academicPeriods, classes, err := s.krsRepository.GetPickClassesByUserID(
		tx, ctx, user.ID, req.AcademicPeriodeID,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get pick classes by user id",
			"user-id": user.ID,
		}).Error(err)
		return nil, utils.ErrorSpToErrorFiber(err)
	}

	// Convert models to DTOs
	apResponses := make([]dto.TrxKrsAcademicPeriodResponse, 0, len(academicPeriods))
	for _, ap := range academicPeriods {
		apResponses = append(apResponses, converter.TrxKrsAcademicPeriodModelToResponse(ap))
	}

	classResponses := make([]dto.TrxKrsClassForPickResponse, 0, len(classes))
	for _, c := range classes {
		classResponses = append(classResponses, converter.TrxKrsClassForPickModelToResponse(c))
	}

	return &dto.TrxKrsPickClassResponse{
		AcademicPeriods: apResponses,
		Classes:          classResponses,
	}, nil
}

// TakeClass - Mahasiswa mengambil kelas
func (s *krsService) TakeClass(
	ctx context.Context, req dto.TrxKrsTakeClassRequest,
) (*dto.TrxKrsTakeClassResponse, error) {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)

	// Get current timestamp in milliseconds
	createdAt := time.Now().UnixMilli()

	result, err := s.krsRepository.TakeClassByUserID(
		tx, ctx, user.ID, req.ClassID.String(), createdAt, user.ID,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "take class",
			"user-id": user.ID,
			"class-id": req.ClassID,
		}).Error(err)
		return nil, utils.ErrorSpToErrorFiber(err)
	}

	response := converter.TrxKrsTakeClassModelToResponse(result)
	return &response, nil
}

// GetKrsProgramHeadClassesWithCount - Get KRS program head classes with pagination
func (s *krsService) GetKrsProgramHeadClassesWithCount(
	ctx context.Context,
	req pageable.PageableKrsProgramHeadClassesRequest,
) (*pageable.PageableResponse[dto.TrxKrsProgramHeadClassResponse], error) {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)

	classes, count, err := s.krsRepository.GetKrsProgramHeadClassesWithCount(
		ctx, tx, user.ID, req,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get krs program head classes",
			"user-id": user.ID,
		}).Error(err)
		return nil, utils.ErrorSpToErrorFiber(err)
	}

	responses := make([]dto.TrxKrsProgramHeadClassResponse, 0, count)
	for _, class := range classes {
		responses = append(responses, converter.TrxKrsProgramHeadClassModelToResponse(class))
	}

	return &pageable.PageableResponse[dto.TrxKrsProgramHeadClassResponse]{
		Data: responses,
		Metadata: pageable.Metadata{
			TotalData: count,
			TotalPage: utils.TotalPage(count, req.GetDefaultLimit()),
			Page:      req.GetDefaultPage(),
			Size:      req.GetDefaultLimit(),
		},
	}, nil
}

// GetSavedByUserID - Get saved KRS items by user ID
func (s *krsService) GetSavedByUserID(
	ctx context.Context, req dto.TrxKrsSavedGetRequest,
) ([]dto.TrxKrsSavedItemResponse, error) {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)

	items, err := s.krsRepository.GetSavedByUserID(
		tx, ctx, user.ID, req.AcademicPeriodeID,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get saved krs by user id",
			"user-id": user.ID,
		}).Error(err)
		return nil, utils.ErrorSpToErrorFiber(err)
	}

	// Convert models to DTOs
	responses := make([]dto.TrxKrsSavedItemResponse, 0, len(items))
	for _, item := range items {
		responses = append(responses, converter.TrxKrsSavedItemModelToResponse(item))
	}

	return responses, nil
}

// DeleteSavedByKrsItemID - Hapus KRS item (hanya status waiting)
func (s *krsService) DeleteSavedByKrsItemID(
	ctx context.Context, krsItemID string,
) error {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)

	// Get current timestamp in milliseconds
	deletedAt := time.Now().UnixMilli()

	err := s.krsRepository.DeleteSavedByKrsItemID(
		tx, krsItemID, deletedAt, user.ID,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service":   "delete saved krs item",
			"user-id":   user.ID,
			"krs-item-id": krsItemID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}

	return nil
}

// GetKrsMaxSksInfo - Get max SKS info for current student
func (s *krsService) GetKrsMaxSksInfo(
	ctx context.Context,
) (*dto.TrxKrsMaxSksInfoResponse, error) {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)

	result, err := s.krsRepository.GetKrsMaxSksInfoByUserID(tx, ctx, user.ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get krs max sks info",
			"user-id": user.ID,
		}).Error(err)
		return nil, utils.ErrorSpToErrorFiber(err)
	}

	response := converter.TrxKrsMaxSksInfoModelToResponse(result)
	return &response, nil
}
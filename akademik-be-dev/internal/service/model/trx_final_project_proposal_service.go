package servicemodel

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/delivery/http/middleware"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/dto/converter"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository/cached"
	repositorymodel "unsia.ac.id/akademic_be/internal/repository/model"
	"unsia.ac.id/akademic_be/internal/service"
	"unsia.ac.id/akademic_be/pkg/utils"
)

type TrxFinalProjectProposalService struct {
	log                               *logrus.Logger
	db                                *gorm.DB
	cache                             cached.CacheRepository
	trxFinalProjectProposalRepository *repositorymodel.TrxFinalProjectProposalRepository
	mstStudentBioRepository           *repositorymodel.MstStudentBioRepository
	storageService                    *service.StorageService
}

func NewTrxFinalProjectProposalService(
	log *logrus.Logger,
	db *gorm.DB,
	cache cached.CacheRepository,
	trxFinalProjectProposalRepository *repositorymodel.TrxFinalProjectProposalRepository,
	mstStudentBioRepository *repositorymodel.MstStudentBioRepository,
	storageService *service.StorageService,
) *TrxFinalProjectProposalService {
	return &TrxFinalProjectProposalService{
		log:                               log,
		db:                                db,
		cache:                             cache,
		trxFinalProjectProposalRepository: trxFinalProjectProposalRepository,
		mstStudentBioRepository:           mstStudentBioRepository,
		storageService:                    storageService,
	}
}

/* Create */
func (s *TrxFinalProjectProposalService) Create(ctx context.Context, req dto.TrxFinalProjectProposalRequest) (*dto.TrxFinalProjectProposalResponseForAdmin, error) {
	tx := s.db.WithContext(ctx)

	now := time.Now().UnixMilli()
	user := middleware.GetUserClaimsCtx(ctx)
	data := new(model.TrxFinalProjectProposal)
	data.ID = req.ID
	data.TitleID = req.TitleID
	data.TitleEN = req.TitleEN
	data.Topic = req.Topic
	data.Abstract = req.Abstract
	data.Date = &now

	validExtension := []string{".pdf"}
	if req.File != nil {
		filePath, err := s.storageService.UploadFileV3(
			ctx, req.File, true,
			"student/final-project/proposal", "",
			validExtension, 3,
		)
		if err != nil {
			s.log.WithFields(logrus.Fields{
				"service": "update final project proposal for UploadFileV3 failed",
				"request": req,
				"user-id": user.ID,
			}).Error(err)
			return nil, utils.ErrorSpToErrorFiber(err)
		}
		if filePath == "" {
			s.log.WithFields(logrus.Fields{
				"service": "update final project proposal for UploadFileV3 failed",
				"request": req,
				"user-id": user.ID,
			}).Error(err)
			return nil, fiber.NewError(fiber.StatusInternalServerError, "failed upload file")
		}
		data.FilePath = filePath
	}

	err := s.trxFinalProjectProposalRepository.Create(tx, ctx, data)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	finalProjectProposal, err := s.GetByID(ctx, req.ID.String())
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), req.ID.String(), utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	return finalProjectProposal, nil
}

func (s *TrxFinalProjectProposalService) AsignAcademicSupervisor(ctx context.Context, req dto.TrxFinalProjectProposalAssignAcademicSupervisorRequest) error {
	tx := s.db.WithContext(ctx)

	err := s.trxFinalProjectProposalRepository.AsignAcademicSupervisor(tx, ctx, req)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

/* Read */
func (s *TrxFinalProjectProposalService) GetAllForStudent(ctx context.Context) (T []*dto.TrxFinalProjectProposalByUserIDResponse, err error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	finalProjectProposals, err := s.trxFinalProjectProposalRepository.GetAllByStudent(tx, ctx)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get all final project proposal service",
			"user-id": user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return nil, err
	}

	results := make([]*dto.TrxFinalProjectProposalByUserIDResponse, 0)

	for _, finalProjectProposal := range finalProjectProposals {
		results = append(results, converter.TrxFinalProjectProposalModelByUserIDToResponse(finalProjectProposal))
	}

	return results, nil
}

func (s *TrxFinalProjectProposalService) GetByID(ctx context.Context, ID string) (*dto.TrxFinalProjectProposalResponseForAdmin, error) {
	tx := s.db.WithContext(ctx)

	trxFinalProjectProposal := new(model.TrxFinalProjectProposal)

	err := s.trxFinalProjectProposalRepository.GetByID(tx, ID, trxFinalProjectProposal)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	trxFinalProjectProposalMentorLecturers, err := s.trxFinalProjectProposalRepository.GetMentorLecturerByID(
		tx, ID,
	)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	result := converter.TrxFinalProjectProposalModelPointerToResponseWithMentorLecturerResponse(trxFinalProjectProposal, trxFinalProjectProposalMentorLecturers)

	return result, nil
}

func (s *TrxFinalProjectProposalService) GetByIDForStudent(ctx context.Context, ID string) (*dto.TrxFinalProjectProposalResponse, error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	trxFinalProjectProposal := new(model.TrxFinalProjectProposal)

	err := s.trxFinalProjectProposalRepository.GetByID(tx, ID, trxFinalProjectProposal)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	studentModel := new(model.MstStudentBio)
	err = s.mstStudentBioRepository.CheckByUserID(tx, user.ID, studentModel)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service":     "get student by user id for get id final project proposal service",
			"user-id":     user.ID,
			"id-proposal": ID,
		}).Error(err)
	}

	if studentModel.ID != trxFinalProjectProposal.StudentID {
		return nil, fiber.NewError(fiber.StatusNotFound, "Specified ID was not found.")
	}

	result := converter.TrxFinalProjectProposalModelPointerToResponse(trxFinalProjectProposal)

	return result, nil
}

func (s *TrxFinalProjectProposalService) GetByIDGroupByStudent(ctx context.Context, ID string) (T []*dto.TrxFinalProjectProposalByUserIDResponse, err error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	finalProjectProposals, err := s.trxFinalProjectProposalRepository.GetByIDGroupByStudent(tx, ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get all final project proposal service",
			"user-id": user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return nil, err
	}

	results := make([]*dto.TrxFinalProjectProposalByUserIDResponse, 0)

	for _, finalProjectProposal := range finalProjectProposals {
		results = append(results, converter.TrxFinalProjectProposalModelByUserIDToResponse(finalProjectProposal))
	}

	return results, nil
}

func (s *TrxFinalProjectProposalService) GetProposalStudentByUser(ctx context.Context, ID string) (T []*dto.TrxFinalProjectProposalByUserIDResponse, err error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	finalProjectProposals, err := s.trxFinalProjectProposalRepository.GetProposalStudentByUser(tx, ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get all final project proposal service",
			"user-id": user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return nil, err
	}

	results := make([]*dto.TrxFinalProjectProposalByUserIDResponse, 0)

	for _, finalProjectProposal := range finalProjectProposals {
		results = append(results, converter.TrxFinalProjectProposalModelByUserIDToResponse(finalProjectProposal))
	}

	return results, nil
}

func (s *TrxFinalProjectProposalService) GetAllWithCountByProgramHeadID(
	ctx context.Context, pageble pageable.PageableRequestFinalProjectProposal,
) (*pageable.PageableResponse[dto.TrxFinalProjectProposalsResponseForAdmin], error) {
	tx := s.db.WithContext(ctx)

	Classes, totalData, err := s.trxFinalProjectProposalRepository.GetAllWithCountProgramHeadIDORLecturer(tx, ctx, pageble)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	res := make([]dto.TrxFinalProjectProposalsResponseForAdmin, 0, totalData)
	for _, class := range Classes {
		res = append(res, *converter.TrxFinalProjectProposalModelProgramHeadAllToResponse(class))
	}

	return &pageable.PageableResponse[dto.TrxFinalProjectProposalsResponseForAdmin]{
		Data: res,
		Metadata: pageable.Metadata{
			TotalData: totalData,
			TotalPage: utils.TotalPage(totalData, pageble.GetDefaultLimit()),
			Page:      pageble.GetDefaultPage(),
			Size:      pageble.GetDefaultLimit(),
		},
	}, nil
}

func (s *TrxFinalProjectProposalService) GetByStudentIDandStudyProgramID(ctx context.Context, StudentID string, StudyProgramID string) (*dto.TrxFinalProjectProposalByStudenIDandStudyProgramIDResponse, error) {
	tx := s.db.WithContext(ctx)

	trxFinalProjectProposal := new(model.TrxFinalProjectProposal)

	err := s.trxFinalProjectProposalRepository.GetByStudentIDandStudyProgramID(tx, StudentID, StudyProgramID, trxFinalProjectProposal)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), StudentID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	result := converter.TrxFinalProjectProposalModelByStudenIDandStudyProgramIDToResponse(*trxFinalProjectProposal)

	return result, nil
}

/* Update */
func (s *TrxFinalProjectProposalService) UpdateByID(ctx context.Context, req dto.TrxFinalProjectProposalUpdateStatusRequest) error {
	tx := s.db.WithContext(ctx)

	err := s.trxFinalProjectProposalRepository.UpdateStatusByID(tx, ctx, req)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), req.ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return utils.ErrorSpToErrorFiber(err)
	}

	return nil
}

/* Delete */

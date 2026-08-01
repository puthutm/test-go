package servicemodel

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	modelsdm "gitlab.unsia.ac.id/icems/icems-tools/gateway/model/sdm"
	restapisdm "gitlab.unsia.ac.id/icems/icems-tools/gateway/rest-api/sdm"
	utilicems "gitlab.unsia.ac.id/icems/icems-tools/utils"
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

type MstClassService interface {
	Create(ctx context.Context, req dto.MstClassRequest) (*dto.MstClassResponse, error)
	CreateByProgramHead(ctx context.Context, req dto.MstClassRequestByProgramHead) (*dto.MstClassResponse, error)
	GetAllWithCount(ctx context.Context, pageble pageable.PageableRequestClass) (*pageable.PageableResponse[dto.MstClassResponse], error)
	GetAllWithCountByStudyProgramID(ctx context.Context, pageble pageable.PageableRequestClass) (*pageable.PageableResponse[dto.MstClassResponse], error)
	GetAllTrashWithCount(ctx context.Context, pageble pageable.PageableRequestClass) (*pageable.PageableResponse[dto.MstClassResponse], error)
	GetByID(ctx context.Context, ID string) (*dto.MstClassResponse, error)
	GetAllWithCountByProgramHeadID(ctx context.Context, pageble pageable.PageableRequestClass) (*pageable.PageableResponse[dto.MstClassResponse], error)
	GetByLecturerIDandActiveAcademicPeriod(ctx context.Context) (res []dto.MstClassStudentDistributionResponse, err error)
	UpdateByID(ctx context.Context, req dto.MstClassRequest) (*dto.MstClassResponse, error)
	UpdateByIDAndProgramHead(ctx context.Context, req dto.MstClassRequestByProgramHead) (*dto.MstClassResponse, error)
	RestoreByID(ctx context.Context, ID string) error
	DeleteByID(ctx context.Context, ID string) error
	DeleteByIDAndProgramHead(ctx context.Context, ID string) error
	UpdateContractByID(ctx context.Context, req dto.MstClassContractRequest) error
	UpdateContractByIDAndProgramHead(ctx context.Context, req dto.MstClassContractRequest) error

	CheckSaveButton(ctx context.Context, academicPeriodID string) (*dto.MstClassCheckSaveButtonResponse, error)
	UpdateStatusLockedByAcademicPeriodID(ctx context.Context, req dto.UpdateStatusLockRequest) error
}

type mstClassService struct {
	log                  *logrus.Logger
	db                   *gorm.DB
	cache                cached.CacheRepository
	mstClassRepository   repositorymodel.MstClassRepository
	mstSubjectRepository *repositorymodel.MstSubjectRepository
	storageService       *service.StorageService

	// external gatewat or service
	generalInformationRest *restapisdm.GeneralInformationRest
}

func NewMstClassService(
	log *logrus.Logger,
	db *gorm.DB,
	cache cached.CacheRepository,
	mstClassRepository repositorymodel.MstClassRepository,
	mstSubjectRepository *repositorymodel.MstSubjectRepository,
	storageService *service.StorageService,
	generalInformationRest *restapisdm.GeneralInformationRest,
) MstClassService {
	return &mstClassService{
		log:                    log,
		db:                     db,
		cache:                  cache,
		mstClassRepository:     mstClassRepository,
		mstSubjectRepository:   mstSubjectRepository,
		storageService:         storageService,
		generalInformationRest: generalInformationRest,
	}
}

/* Create */
func (s *mstClassService) Create(ctx context.Context, req dto.MstClassRequest) (*dto.MstClassResponse, error) {
	tx := s.db.WithContext(ctx)

	data := new(model.MstSubject)
	err := s.mstSubjectRepository.GetByID(tx, req.SubjectID, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"function":   "CreateClass",
			"service":    "get by id subject in insert class",
			"subject id": req.SubjectID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return nil, err
	}

	req.StudyProgramID = data.StudyProgramID
	req.CurriculumYearID = data.CurriculumYearID.String()

	err = s.mstClassRepository.Create(tx, ctx, req)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	Class, err := s.GetByID(ctx, req.ID)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), req.ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	return Class, nil
}

func (s *mstClassService) CreateByProgramHead(ctx context.Context, req dto.MstClassRequestByProgramHead) (*dto.MstClassResponse, error) {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)

	data := new(model.MstSubject)
	err := s.mstSubjectRepository.GetByID(tx, req.SubjectID, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"function":   "CreateClass by program head",
			"service":    "get by id subject in insert class",
			"subject id": req.SubjectID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return nil, err
	}

	generalInformation, err := s.generalInformationRest.GetDataWithParamsOrQuery(nil, nil, "/"+user.ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class service createByProgramHead",
			"user-id": user.ID,
		}).Error(err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "get data sdm error "+err.Error())
	}

	if generalInformation.Error {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class service createByProgramHead",
			"user-id": user.ID,
		}).Error(generalInformation.Message)
		return nil, fiber.NewError(generalInformation.Status, generalInformation.Message)
	}

	if generalInformation.Data == nil {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class service createByProgramHead",
			"user-id": user.ID,
		}).Error(err)
		return nil, fiber.NewError(fiber.StatusNotFound, "generalInformation not found wit user -> "+user.ID)
	}

	studyProgramID := generalInformation.Data.StudyProgramID
	if studyProgramID == nil {
		s.log.WithFields(logrus.Fields{
			"service":        "get general information from sdm in class service createByProgramHead",
			"user-id":        user.ID,
			"studyProgramID": "studyProgramID not found",
		}).Info(err)
		return nil, fiber.NewError(fiber.StatusNotFound, "studyProgramID not found wit user -> "+user.ID)
	}

	if generalInformation.Data.StudyProgramID.String() != data.StudyProgramID {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class service createByProgramHead",
			"user-id": user.ID,
		}).Info("create class by program head not access")
		return nil, fiber.NewError(fiber.StatusForbidden, "ProgramHead not access data")
	}

	req.CurriculumYearID = data.CurriculumYearID.String()

	err = s.mstClassRepository.CreateByProgramHead(tx, ctx, req)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	Class, err := s.GetByID(ctx, req.ID)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), req.ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	return Class, nil
}

/* Read */
func (s *mstClassService) GetAllWithCount(ctx context.Context, pageble pageable.PageableRequestClass) (*pageable.PageableResponse[dto.MstClassResponse], error) {
	tx := s.db.WithContext(ctx)

	Classes, totalData, err := s.mstClassRepository.GetAllWithCount(tx, false, pageble)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	res := make([]dto.MstClassResponse, 0, totalData)
	for _, class := range Classes {
		res = append(res, *converter.MstClassModelToResponse(class))
	}

	return &pageable.PageableResponse[dto.MstClassResponse]{
		Data: res,
		Metadata: pageable.Metadata{
			TotalData: totalData,
			TotalPage: utils.TotalPage(totalData, pageble.GetDefaultLimit()),
			Page:      pageble.GetDefaultPage(),
			Size:      pageble.GetDefaultLimit(),
		},
	}, nil
}

func (s *mstClassService) GetAllWithCountByStudyProgramID(ctx context.Context, pageble pageable.PageableRequestClass) (*pageable.PageableResponse[dto.MstClassResponse], error) {
	tx := s.db.WithContext(ctx)

	Classes, totalData, err := s.mstClassRepository.GetAllWithCountByStudyProgramID(tx, pageble)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	res := make([]dto.MstClassResponse, 0, totalData)
	for _, class := range Classes {
		res = append(res, *converter.MstClassModelToResponse(class))
	}

	return &pageable.PageableResponse[dto.MstClassResponse]{
		Data: res,
		Metadata: pageable.Metadata{
			TotalData: totalData,
			TotalPage: utils.TotalPage(totalData, pageble.GetDefaultLimit()),
			Page:      pageble.GetDefaultPage(),
			Size:      pageble.GetDefaultLimit(),
		},
	}, nil
}

func (s *mstClassService) GetAllTrashWithCount(ctx context.Context, pageble pageable.PageableRequestClass) (*pageable.PageableResponse[dto.MstClassResponse], error) {
	tx := s.db.WithContext(ctx)

	classes, totalData, err := s.mstClassRepository.GetAllWithCount(tx, true, pageble)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	res := make([]dto.MstClassResponse, 0, totalData)
	for _, v := range classes {
		res = append(res, *converter.MstClassModelToResponse(v))
	}

	return &pageable.PageableResponse[dto.MstClassResponse]{
		Data: res,
		Metadata: pageable.Metadata{
			TotalData: totalData,
			TotalPage: utils.TotalPage(totalData, pageble.GetDefaultLimit()),
			Page:      pageble.GetDefaultPage(),
			Size:      pageble.GetDefaultLimit(),
		},
	}, nil
}

func (s *mstClassService) GetByID(ctx context.Context, ID string) (*dto.MstClassResponse, error) {
	tx := s.db.WithContext(ctx)

	classModel := new(model.MstClass)

	err := s.mstClassRepository.GetByID(tx, ID, classModel)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	classes := converter.MstClassModelToResponse(*classModel)

	return classes, nil
}

func (s *mstClassService) GetAllWithCountByProgramHeadID(ctx context.Context, pageble pageable.PageableRequestClass) (*pageable.PageableResponse[dto.MstClassResponse], error) {
	tx := s.db.WithContext(ctx)

	Classes, totalData, err := s.mstClassRepository.GetAllWithCountByProgramHeadID(tx, ctx, pageble)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	res := make([]dto.MstClassResponse, 0, totalData)
	for _, class := range Classes {
		res = append(res, *converter.MstClassModelToResponse(class))
	}

	return &pageable.PageableResponse[dto.MstClassResponse]{
		Data: res,
		Metadata: pageable.Metadata{
			TotalData: totalData,
			TotalPage: utils.TotalPage(totalData, pageble.GetDefaultLimit()),
			Page:      pageble.GetDefaultPage(),
			Size:      pageble.GetDefaultLimit(),
		},
	}, nil
}

func (s *mstClassService) GetByLecturerIDandActiveAcademicPeriod(ctx context.Context) (res []dto.MstClassStudentDistributionResponse, err error) {
	tx := s.db.WithContext(ctx)

	classes, err := s.mstClassRepository.GetByLecturerIDandActiveAcademicPeriod(tx, ctx)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	results := make([]dto.MstClassStudentDistributionResponse, 0)
	for _, class := range classes {
		results = append(results, *converter.MstClassStudentDistributionToResponse(class))
	}

	return results, nil
}

/* Update */
func (s *mstClassService) UpdateByID(ctx context.Context, req dto.MstClassRequest) (*dto.MstClassResponse, error) {
	tx := s.db.WithContext(ctx)

	data := new(model.MstSubject)
	err := s.mstSubjectRepository.GetByID(tx, req.SubjectID, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"function":   "UpdateClass",
			"service":    "get by id subject in Update class",
			"subject id": req.SubjectID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return nil, err
	}

	req.StudyProgramID = data.StudyProgramID
	req.CurriculumYearID = data.CurriculumYearID.String()

	err = s.mstClassRepository.UpdateByID(tx, ctx, req)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), req.ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	classes, err := s.GetByID(ctx, req.ID)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), req.ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	return classes, nil
}

func (s *mstClassService) UpdateByIDAndProgramHead(ctx context.Context, req dto.MstClassRequestByProgramHead) (*dto.MstClassResponse, error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	data := new(model.MstSubject)
	err := s.mstSubjectRepository.GetByID(tx, req.SubjectID, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"function":   "UpdateClass by program head",
			"service":    "get by id subject in Update class",
			"subject id": req.SubjectID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return nil, err
	}

	generalInformation, err := s.generalInformationRest.GetDataWithParamsOrQuery(nil, nil, "/"+user.ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class service updateByProgramHead",
			"user-id": user.ID,
		}).Error(err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "get data sdm error "+err.Error())
	}

	if generalInformation.Error {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class service updateByProgramHead",
			"user-id": user.ID,
		}).Error(generalInformation.Error)
		return nil, fiber.NewError(generalInformation.Status, generalInformation.Message)
	}

	if generalInformation.Data == nil {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class service updateByProgramHead",
			"user-id": user.ID,
		}).Error(err)
		return nil, fiber.NewError(fiber.StatusNotFound, "generalInformation not found wit user -> "+user.ID)
	}

	studyProgramID := generalInformation.Data.StudyProgramID
	if studyProgramID == nil {
		s.log.WithFields(logrus.Fields{
			"service":        "get general information from sdm in class service updateByProgramHead",
			"user-id":        user.ID,
			"studyProgramID": "studyProgramID not found",
		}).Info(err)
		return nil, fiber.NewError(fiber.StatusNotFound, "studyProgramID not found wit user -> "+user.ID)
	}

	if generalInformation.Data.StudyProgramID.String() != data.StudyProgramID {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class service createByProgramHead",
			"user-id": user.ID,
		}).Info("update class by program head not access")
		return nil, fiber.NewError(fiber.StatusForbidden, "ProgramHead not access data")
	}

	req.CurriculumYearID = data.CurriculumYearID.String()

	err = s.mstClassRepository.UpdateByIDAndProgramHead(tx, ctx, req)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), req.ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	classes, err := s.GetByID(ctx, req.ID)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), req.ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	return classes, nil
}

func (s *mstClassService) RestoreByID(ctx context.Context, ID string) error {
	tx := s.db.WithContext(ctx)

	err := s.mstClassRepository.RestoreByID(tx, ID)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return utils.ErrorSpToErrorFiber(err)
	}

	return nil
}

/* Delete */
func (s *mstClassService) DeleteByID(ctx context.Context, ID string) error {
	tx := s.db.WithContext(ctx)

	err := s.mstClassRepository.DeleteByID(tx, ctx, ID)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return utils.ErrorSpToErrorFiber(err)
	}

	return nil
}

func (s *mstClassService) DeleteByIDAndProgramHead(ctx context.Context, ID string) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	data := new(model.MstClass)
	err := s.mstClassRepository.GetByID(tx, ID, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"function":   "UpdateClass by program head",
			"service":    "get by id class in delete class",
			"subject id": ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return err
	}

	generalInformation, err := s.generalInformationRest.GetDataWithParamsOrQuery(nil, nil, "/"+user.ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class service deleteByProgramHead",
			"user-id": user.ID,
		}).Error(err)
		return fiber.NewError(fiber.StatusInternalServerError, "get data sdm error "+err.Error())
	}

	if generalInformation.Error {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class service deleteByProgramHead",
			"user-id": user.ID,
		}).Error(generalInformation.Message)
		return fiber.NewError(generalInformation.Status, generalInformation.Message)
	}

	if generalInformation.Data == nil {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class service deleteByProgramHead",
			"user-id": user.ID,
		}).Error(err)
		return fiber.NewError(fiber.StatusNotFound, "generalInformation in sdm not found with user -> "+user.ID)
	}

	studyProgramID := generalInformation.Data.StudyProgramID
	if studyProgramID == nil {
		s.log.WithFields(logrus.Fields{
			"service":        "get general information from sdm in class service deleteByProgramHead",
			"user-id":        user.ID,
			"studyProgramID": "studyProgramID not found",
		}).Info(err)
		return fiber.NewError(fiber.StatusNotFound, "studyProgramID in sdm not found wit user -> "+user.ID)
	}

	if generalInformation.Data.StudyProgramID.String() != data.StudyProgramID && data.StudyProgramID != "0" {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class service deleteByProgramHead",
			"user-id": user.ID,
		}).Info("update class by program head not access")
		return fiber.NewError(fiber.StatusForbidden, "ProgramHead not access data")
	}

	err = s.mstClassRepository.DeleteByID(tx, ctx, ID)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return utils.ErrorSpToErrorFiber(err)
	}

	return nil
}

func (s *mstClassService) UpdateContractByID(ctx context.Context, req dto.MstClassContractRequest) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)
	data := new(model.MstClass)

	data.ID = req.ID
	data.ContractDescription = req.ContractDescription

	validExtension := []string{".jpg", ".pdf", ".png"}

	contractFilePath, err := s.storageService.UploadFileV3(
		ctx, req.ContractFile, true,
		"akademic/classes", "",
		validExtension, 3,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "update class contract service for UploadFileV3 failed",
			"request": req,
			"user-id": user.ID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	if contractFilePath == "" {
		s.log.WithFields(logrus.Fields{
			"service": "update class contract service for UploadFileV3 failed",
			"request": req,
			"user-id": user.ID,
		}).Error(err)
		return fiber.NewError(fiber.StatusInternalServerError, "failed upload file")
	}

	data.ContractFilePath = &contractFilePath

	err = s.mstClassRepository.UpdateContractByID(tx, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "update class contract service",
			"request": req,
			"user-id": user.ID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

func (s *mstClassService) UpdateContractByIDAndProgramHead(ctx context.Context, req dto.MstClassContractRequest) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	dataCheck := new(model.MstClass)
	err := s.mstClassRepository.GetByID(tx, req.ID.String(), dataCheck)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"function":   "UpdateClass by program head",
			"service":    "get by id class in update contract class",
			"subject id": req.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return err
	}

	generalInformation, err := s.generalInformationRest.GetDataWithParamsOrQuery(nil, nil, "/"+user.ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class service updateContractByProgramHead",
			"user-id": user.ID,
		}).Error(err)
		return fiber.NewError(fiber.StatusInternalServerError, "get data sdm error "+err.Error())
	}

	if generalInformation.Error {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class service updateContractByProgramHead",
			"user-id": user.ID,
		}).Error(generalInformation.Message)
		return fiber.NewError(generalInformation.Status, generalInformation.Message)
	}

	if generalInformation.Data == nil {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class service updateContractByProgramHead",
			"user-id": user.ID,
		}).Error(err)
		return fiber.NewError(fiber.StatusNotFound, "generalInformation in sdm not found with user -> "+user.ID)
	}

	studyProgramID := generalInformation.Data.StudyProgramID
	if studyProgramID == nil {
		s.log.WithFields(logrus.Fields{
			"service":        "get general information from sdm in class service updateContractByProgramHead",
			"user-id":        user.ID,
			"studyProgramID": "studyProgramID not found",
		}).Info(err)
		return fiber.NewError(fiber.StatusNotFound, "studyProgramID in sdm not found wit user -> "+user.ID)
	}

	if generalInformation.Data.StudyProgramID.String() != dataCheck.StudyProgramID && dataCheck.StudyProgramID != "0" {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class service updateContractByProgramHead",
			"user-id": user.ID,
		}).Info("update classctract by program head not access")
		return fiber.NewError(fiber.StatusForbidden, "user not access data")
	}

	data := new(model.MstClass)

	data.ID = req.ID
	data.ContractDescription = req.ContractDescription

	validExtension := []string{".jpg", ".pdf", ".png"}

	contractFilePath, err := s.storageService.UploadFileV3(
		ctx, req.ContractFile, true,
		"akademic/classes", "",
		validExtension, 3,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "update class contract service for UploadFileV3 failed",
			"request": req,
			"user-id": user.ID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	if contractFilePath == "" {
		s.log.WithFields(logrus.Fields{
			"service": "update class contract service for UploadFileV3 failed",
			"request": req,
			"user-id": user.ID,
		}).Error(err)
		return fiber.NewError(fiber.StatusInternalServerError, "failed upload file")
	}

	data.ContractFilePath = &contractFilePath

	err = s.mstClassRepository.UpdateContractByID(tx, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "update class contract service",
			"request": req,
			"user-id": user.ID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

// helper for this service
func (s *mstClassService) checkStudyProgramUseProgramHead_GeneralInformation(
	userID string, classData *model.MstClass, funcAction string,
) (*modelsdm.GeneralInformationResponse, error) {
	generalInformation, err := s.generalInformationRest.GetDataWithParamsOrQuery(
		nil, nil, "/"+userID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class service " + funcAction + " by programhead",
			"user-id": userID,
		}).Error(err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "get data sdm error "+err.Error())
	}

	err = utilicems.CheckRestError(
		s.log, generalInformation, userID,
		funcAction+" class by programhead", "get general information from sdm in class service "+funcAction+" by program head",
	)
	if err != nil {
		return nil, err
	}

	studyProgramID := generalInformation.Data.StudyProgramID
	if studyProgramID == nil {
		s.log.WithFields(logrus.Fields{
			"service":        "get general information from sdm in class service " + funcAction + "ByProgramHead",
			"user-id":        userID,
			"studyProgramID": "studyProgramID not found",
		}).Info(err)
		return nil, fiber.NewError(fiber.StatusNotFound, "studyProgramID in sdm not found wit user -> "+userID)
	}

	if classData.StudyProgramID != studyProgramID.String() && classData.StudyProgramID != "0" {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class service " + funcAction + "ByProgramHead",
			"user-id": userID,
		}).Info(funcAction + " class by program head not access")
		return nil, fiber.NewError(fiber.StatusForbidden, "ProgramHead not access data")
	}

	return generalInformation.Data, nil
}

func (s *mstClassService) CheckSaveButton(ctx context.Context, academicPeriodID string) (*dto.MstClassCheckSaveButtonResponse, error) {
	tx := s.db.WithContext(ctx)

	statusLock, err := s.mstClassRepository.CheckSaveButton(tx, academicPeriodID)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), academicPeriodID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)
		return nil, utils.ErrorSpToErrorFiber(err)
	}

	return &dto.MstClassCheckSaveButtonResponse{
		StatusLock: statusLock,
	}, nil
}

func (s *mstClassService) UpdateStatusLockedByAcademicPeriodID(ctx context.Context, req dto.UpdateStatusLockRequest) error {
	tx := s.db.WithContext(ctx)

	req.CreatedAt = time.Now().UnixMilli()

	if err := s.mstClassRepository.UpdateStatusLockedByAcademicPeriodID(tx, req); err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), req.ClassID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)
		return utils.ErrorSpToErrorFiber(err)
	}

	return nil
}

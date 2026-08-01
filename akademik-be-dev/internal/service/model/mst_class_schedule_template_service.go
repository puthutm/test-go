package servicemodel

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	msg "unsia.ac.id/akademic_be/pkg/icems-tools/dto/message"
	modelsdm "unsia.ac.id/akademic_be/pkg/icems-tools/gateway/model/sdm"
	restapisdm "unsia.ac.id/akademic_be/pkg/icems-tools/gateway/rest-api/sdm"
	utilicems "unsia.ac.id/akademic_be/pkg/icems-tools/utils"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/delivery/http/middleware"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/dto/converter"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository/cached"
	repositorymodel "unsia.ac.id/akademic_be/internal/repository/model"
	"unsia.ac.id/akademic_be/pkg/utils"
)

type MstClassScheduleTemplateService struct {
	log                                *logrus.Logger
	db                                 *gorm.DB
	cache                              cached.CacheRepository
	mstClassScheduleTemplateRepository *repositorymodel.MstClassScheduleTemplateRepository
	mstClassRepository                 repositorymodel.MstClassRepository
	mstClassScheduleService            *MstClassScheduleService

	// external
	generalInformationGatewayRest *restapisdm.GeneralInformationRest
}

func NewMstClassScheduleTemplateService(
	log *logrus.Logger,
	db *gorm.DB,
	cache cached.CacheRepository,
	mstClassScheduleTemplateRepository *repositorymodel.MstClassScheduleTemplateRepository,
	mstClassRepository repositorymodel.MstClassRepository,

	mstClassScheduleService *MstClassScheduleService,

	// external
	generalInformationGatewayRest *restapisdm.GeneralInformationRest,
) *MstClassScheduleTemplateService {
	return &MstClassScheduleTemplateService{
		log:                                log,
		db:                                 db,
		cache:                              cache,
		mstClassScheduleTemplateRepository: mstClassScheduleTemplateRepository,
		mstClassRepository:                 mstClassRepository,
		mstClassScheduleService:            mstClassScheduleService,

		// external
		generalInformationGatewayRest: generalInformationGatewayRest,
	}
}

// TODO: Create
func (s *MstClassScheduleTemplateService) Create(ctx context.Context, req dto.MstClassScheduleTemplateCreateRequest) error {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)

	userID, err := utils.StringToUuid(user.ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "create class schdule template service",
			"request": req,
			"user_id": user.ID,
		}).Error(err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	check, err := s.mstClassScheduleTemplateRepository.CheckModelAll(tx, "class_id = ?", req.ClassID.String())
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "create class schdule template service",
			"request": req,
			"user_id": user.ID,
		}).Error(err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	if check {
		return fiber.NewError(fiber.StatusConflict, "data already exist")
	}

	data := converter.ConvertMstClassScheduleTemplateCreateRequestToModel(req)
	data.CreatedBy = &userID
	err = s.mstClassScheduleTemplateRepository.Create(tx, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "create class schdule template service",
			"request": req,
			"user_id": user.ID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

func (s *MstClassScheduleTemplateService) CreateByProgramHead(ctx context.Context, req dto.MstClassScheduleTemplateCreateRequest) error {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)

	userID, err := utils.StringToUuid(user.ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "create class schdule template service",
			"request": req,
			"user_id": user.ID,
		}).Error(err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	check, err := s.mstClassScheduleTemplateRepository.CheckModelAll(tx, "class_id = ?", req.ClassID.String())
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "create class schdule template service",
			"request": req,
			"user_id": user.ID,
		}).Error(err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	if check {
		return fiber.NewError(fiber.StatusConflict, "data already exist")
	}

	classData := new(model.MstClass)
	err = s.mstClassRepository.GetByID(tx, req.ClassID.String(), classData)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get class  by id in create by programhead class schedule template service",
			"user-id": user.ID,
		}).Error(err)
		return utilicems.ErrorSpToErrorFiber(err)
	}

	generalInformation, err := s.generalInformationGatewayRest.GetDataWithParamsOrQuery(
		nil, nil, "/"+user.ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class schedule template service create by programhead",
			"user-id": user.ID,
		}).Error(err)
		return fiber.NewError(fiber.StatusInternalServerError, "get data sdm error "+err.Error())
	}

	err = utilicems.CheckRestError(
		s.log, generalInformation, user.ID,
		"create class schedule template by programhead", "get general information from sdm in class schedule template service create by program head",
	)
	if err != nil {
		return err
	}

	studyProgramID := generalInformation.Data.StudyProgramID
	if studyProgramID == nil {
		s.log.WithFields(logrus.Fields{
			"service":        "get general information from sdm in class service createByProgramHead",
			"user-id":        user.ID,
			"studyProgramID": "studyProgramID not found",
		}).Info(err)
		return fiber.NewError(fiber.StatusNotFound, "studyProgramID in sdm not found wit user -> "+user.ID)
	}

	if classData.StudyProgramID != studyProgramID.String() && classData.StudyProgramID != "0" {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class schedule template service createaByProgramHead",
			"user-id": user.ID,
		}).Info("create class schedule templatet by program head not access")
		return fiber.NewError(fiber.StatusForbidden, "ProgramHead not access data")
	}

	data := converter.ConvertMstClassScheduleTemplateCreateRequestToModel(req)
	data.CreatedBy = &userID
	err = s.mstClassScheduleTemplateRepository.Create(tx, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "create class schdule template service",
			"request": req,
			"user_id": user.ID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

// TODO: Update
func (s *MstClassScheduleTemplateService) UpdateByID(ctx context.Context, req dto.MstClassScheduleTemplateUpdateRequest) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	userID, err := utils.StringToUuid(user.ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "update class schdule template service",
			"request": req,
			"user_id": user.ID,
		}).Error(err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	data := converter.ConvertMstClassScheduleTemplateUpdateRequestToModel(req)
	data.UpdatedBy = &userID
	err = s.mstClassScheduleTemplateRepository.UpdateByID(tx, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "update class schedule template service",
			"request": req,
			"user-id": user.ID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

func (s *MstClassScheduleTemplateService) UpdateByIDAndProgramHead(ctx context.Context, req dto.MstClassScheduleTemplateUpdateRequest) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	userID, err := utils.StringToUuid(user.ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "update class schdule template service",
			"request": req,
			"user_id": user.ID,
		}).Error(err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	classData := new(model.MstClass)
	err = s.mstClassRepository.GetByID(tx, req.ClassID.String(), classData)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get class  by id in update by programhead class schedule template service",
			"user-id": user.ID,
		}).Error(err)
		return utilicems.ErrorSpToErrorFiber(err)
	}

	generalInformation, err := s.generalInformationGatewayRest.GetDataWithParamsOrQuery(
		nil, nil, "/"+user.ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class schedule template service update by programhead",
			"user-id": user.ID,
		}).Error(err)
		return fiber.NewError(fiber.StatusInternalServerError, "get data sdm error "+err.Error())
	}

	err = utilicems.CheckRestError(
		s.log, generalInformation, user.ID,
		"update class schedule template by programhead", "get general information from sdm in class schedule template service update by program head",
	)
	if err != nil {
		return err
	}

	studyProgramID := generalInformation.Data.StudyProgramID
	if studyProgramID == nil {
		s.log.WithFields(logrus.Fields{
			"service":        "get general information from sdm in class service updateByProgramHead",
			"user-id":        user.ID,
			"studyProgramID": "studyProgramID not found",
		}).Info(err)
		return fiber.NewError(fiber.StatusNotFound, "studyProgramID in sdm not found wit user -> "+user.ID)
	}

	if classData.StudyProgramID != studyProgramID.String() && classData.StudyProgramID != "0" {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class schedule template service updateaByProgramHead",
			"user-id": user.ID,
		}).Info("update class schedule templatet by program head not access")
		return fiber.NewError(fiber.StatusForbidden, "ProgramHead not access data")
	}

	data := converter.ConvertMstClassScheduleTemplateUpdateRequestToModel(req)
	data.UpdatedBy = &userID
	err = s.mstClassScheduleTemplateRepository.UpdateByID(tx, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "update class schedule template service",
			"request": req,
			"user-id": user.ID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

// TODO: Delete
func (s *MstClassScheduleTemplateService) DeleteByID(ctx context.Context, ID string) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	err := s.mstClassScheduleService.DeleteByScheduleTemplate(ctx, ID)
	if err != nil {
		return err
	}

	err = s.mstClassScheduleTemplateRepository.DeleteByID(tx, ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service":  "delete class schedule template service",
			"id":       ID,
			"deleteby": user.ID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

func (s *MstClassScheduleTemplateService) DeleteByIDAndProgramHead(ctx context.Context, ID string) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	classScheduleTemplateData := new(model.MstClassScheduleTemplate)
	err := s.mstClassScheduleTemplateRepository.GetByID(tx, ID, classScheduleTemplateData)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get class schedule template  by id in delete by programhead class schedule template service",
			"user-id": user.ID,
		}).Error(err)
		return utilicems.ErrorSpToErrorFiber(err)
	}

	classData := new(model.MstClass)
	err = s.mstClassRepository.GetByID(tx, classScheduleTemplateData.ClassID.String(), classData)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get class  by id in delete by programhead class schedule template service",
			"user-id": user.ID,
		}).Error(err)
		return utilicems.ErrorSpToErrorFiber(err)
	}

	generalInformation, err := s.generalInformationGatewayRest.GetDataWithParamsOrQuery(
		nil, nil, "/"+user.ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class schedule template service delete by programhead",
			"user-id": user.ID,
		}).Error(err)
		return fiber.NewError(fiber.StatusInternalServerError, "get data sdm error "+err.Error())
	}

	err = utilicems.CheckRestError(
		s.log, generalInformation, user.ID,
		"delete class schedule template by programhead", "get general information from sdm in class schedule template service delete by program head",
	)
	if err != nil {
		return err
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

	if classData.StudyProgramID != studyProgramID.String() && classData.StudyProgramID != "0" {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class schedule template service deleteaByProgramHead",
			"user-id": user.ID,
		}).Info("delete class schedule templatet by program head not access")
		return fiber.NewError(fiber.StatusForbidden, "ProgramHead not access data")
	}

	err = s.mstClassScheduleService.DeleteByScheduleTemplate(ctx, ID)
	if err != nil {
		return err
	}

	err = s.mstClassScheduleTemplateRepository.DeleteByID(tx, ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service":  "delete class schedule template service",
			"id":       ID,
			"deleteby": user.ID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

// TODO: Restore

// TODO: Read
func (s *MstClassScheduleTemplateService) GetByID(ctx context.Context, ID string) (res *dto.MstClassScheduleTemplateResponse, err error) {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)

	data := new(model.MstClassScheduleTemplate)
	err = s.mstClassScheduleTemplateRepository.GetByID(tx, ID, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get by id class schedule template service",
			"id":      ID,
			"user-id": user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return
	}
	res = converter.ConvertModelToMstClassSchedulePointerTemplateResponsePointer(data)
	return
}

func (s *MstClassScheduleTemplateService) GetByIDAndProgramHead(ctx context.Context, ID string) (res *dto.MstClassScheduleTemplateResponse, err error) {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)

	data := new(model.MstClassScheduleTemplate)
	err = s.mstClassScheduleTemplateRepository.GetByID(tx, ID, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get by id class schedule template service",
			"id":      ID,
			"user-id": user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return
	}

	classData := new(model.MstClass)
	err = s.mstClassRepository.GetByID(tx, data.ClassID.String(), classData)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get class  by id in getid by programhead class schedule template service",
			"user-id": user.ID,
		}).Error(err)
		err = utilicems.ErrorSpToErrorFiber(err)
		return
	}

	generalInformation, err := s.generalInformationGatewayRest.GetDataWithParamsOrQuery(
		nil, nil, "/"+user.ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class schedule template service getid by programhead",
			"user-id": user.ID,
		}).Error(err)
		err = fiber.NewError(fiber.StatusInternalServerError, "get data sdm error "+err.Error())
		return
	}

	err = utilicems.CheckRestError(
		s.log, generalInformation, user.ID,
		"getid class schedule template by programhead", "get general information from sdm in class schedule template service getid by program head",
	)
	if err != nil {
		return
	}

	studyProgramID := generalInformation.Data.StudyProgramID
	if studyProgramID == nil {
		s.log.WithFields(logrus.Fields{
			"service":        "get general information from sdm in class service updateByProgramHead",
			"user-id":        user.ID,
			"studyProgramID": "studyProgramID not found",
		}).Info(err)
		err = fiber.NewError(fiber.StatusNotFound, "studyProgramID in sdm not found wit user -> "+user.ID)
		return
	}

	if classData.StudyProgramID != studyProgramID.String() && classData.StudyProgramID != "0" {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class schedule template service updateaByProgramHead",
			"user-id": user.ID,
		}).Info("getid class schedule templatet by program head not access")
		err = fiber.NewError(fiber.StatusForbidden, "ProgramHead not access data")
		return
	}

	res = converter.ConvertModelToMstClassSchedulePointerTemplateResponsePointer(data)
	return
}

func (s *MstClassScheduleTemplateService) GetByClassID(ctx context.Context, classID string) (*dto.MstClassScheduleTemplateResponse, error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	scheduleTemplateData := new(model.MstClassScheduleTemplate)
	err := s.mstClassScheduleTemplateRepository.GetByClassID(tx, classID, scheduleTemplateData)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service":  "get by class id class schedule template service",
			"class_id": classID,
			"user-id":  user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return nil, err
	}

	if scheduleTemplateData.ID == uuid.Nil {
		return nil, fiber.NewError(fiber.StatusNotFound, msg.ErrNotFound.Error())
	}

	res := converter.ConvertModelToMstClassSchedulePointerTemplateResponsePointer(scheduleTemplateData)
	return res, nil
}

func (s *MstClassScheduleTemplateService) GetByClassIDAndProgramHead(ctx context.Context, classID string) (*dto.MstClassScheduleTemplateResponse, error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	classData := new(model.MstClass)
	err := s.mstClassRepository.GetByID(tx, classID, classData)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get class  by id in readbyclass by programhead class schedule template service",
			"user-id": user.ID,
		}).Error(err)
		return nil, utilicems.ErrorSpToErrorFiber(err)
	}

	generalInformation, err := s.generalInformationGatewayRest.GetDataWithParamsOrQuery(
		nil, nil, "/"+user.ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class schedule template service readbyclass by programhead",
			"user-id": user.ID,
		}).Error(err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "get data sdm error "+err.Error())
	}

	err = utilicems.CheckRestError(
		s.log, generalInformation, user.ID,
		"readbyclass class schedule template by programhead", "get general information from sdm in class schedule template service readbyclass by program head",
	)
	if err != nil {
		return nil, err
	}

	studyProgramID := generalInformation.Data.StudyProgramID
	if studyProgramID == nil {
		s.log.WithFields(logrus.Fields{
			"service":        "get general information from sdm in class service readbyclassByProgramHead",
			"user-id":        user.ID,
			"studyProgramID": "studyProgramID not found",
		}).Info(err)
		return nil, fiber.NewError(fiber.StatusNotFound, "studyProgramID in sdm not found wit user -> "+user.ID)
	}

	if classData.StudyProgramID != studyProgramID.String() && classData.StudyProgramID != "0" {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class schedule template service readbyclassaByProgramHead",
			"user-id": user.ID,
		}).Info("readbyclass class schedule templatet by program head not access")
		return nil, fiber.NewError(fiber.StatusForbidden, "ProgramHead not access data")
	}

	scheduleTemplateData := new(model.MstClassScheduleTemplate)
	err = s.mstClassScheduleTemplateRepository.GetByClassID(tx, classID, scheduleTemplateData)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service":  "get by class id class schedule template service",
			"class_id": classID,
			"user-id":  user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return nil, err
	}

	if scheduleTemplateData.ID == uuid.Nil {
		return nil, fiber.NewError(fiber.StatusNotFound, msg.ErrNotFound.Error())
	}
	res := converter.ConvertModelToMstClassSchedulePointerTemplateResponsePointer(scheduleTemplateData)
	return res, nil
}

// helper for this service
func (s *MstClassScheduleTemplateService) checkStudyProgramUseProgramHead_InClassAndGeneralInformation(
	tx *gorm.DB, userID, classID, funcAction string,
) (*model.MstClass, *modelsdm.GeneralInformationResponse, error) {
	classData := new(model.MstClass)
	err := s.mstClassRepository.GetByID(tx, classID, classData)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get class by id in " + funcAction + "by programhead class schedule template service",
			"user-id": userID,
		}).Error(err)
		return nil, nil, utilicems.ErrorSpToErrorFiber(err)
	}
	generalInformation, err := s.generalInformationGatewayRest.GetDataWithParamsOrQuery(
		nil, nil, "/"+userID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class schedule templatet service " + funcAction + " by programhead",
			"user-id": userID,
		}).Error(err)
		return nil, nil, fiber.NewError(fiber.StatusInternalServerError, "get data sdm error "+err.Error())
	}

	err = utilicems.CheckRestError(
		s.log, generalInformation, userID,
		funcAction+" class schedule templatet by programhead", "get general information from sdm in class schedule templatet service "+funcAction+" by program head",
	)
	if err != nil {
		return nil, nil, err
	}

	studyProgramID := generalInformation.Data.StudyProgramID
	if studyProgramID == nil {
		s.log.WithFields(logrus.Fields{
			"service":        "get general information from sdm in class service " + funcAction + "ByProgramHead",
			"user-id":        userID,
			"studyProgramID": "studyProgramID not found",
		}).Info(err)
		return nil, nil, fiber.NewError(fiber.StatusNotFound, "studyProgramID in sdm not found wit user -> "+userID)
	}

	if classData.StudyProgramID != studyProgramID.String() && classData.StudyProgramID != "0" {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class schedule templatet service " + funcAction + "ByProgramHead",
			"user-id": userID,
		}).Info(funcAction + " class schedule templatet by program head not access")
		return nil, nil, fiber.NewError(fiber.StatusForbidden, "ProgramHead not access data")
	}

	return classData, generalInformation.Data, nil
}

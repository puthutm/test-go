package servicemodel

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	modelsdm "unsia.ac.id/akademic_be/pkg/icems-tools/gateway/model/sdm"
	restapisdm "unsia.ac.id/akademic_be/pkg/icems-tools/gateway/rest-api/sdm"
	utilicems "unsia.ac.id/akademic_be/pkg/icems-tools/utils"
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

type MstClassParticipantService struct {
	log                           *logrus.Logger
	db                            *gorm.DB
	cache                         cached.CacheRepository
	mstClassParticipantRepository *repositorymodel.MstClassParticipantRepository
	mstStudentBioRepository       *repositorymodel.MstStudentBioRepository
	mstClassRepository            repositorymodel.MstClassRepository

	// external
	generalInformationGatewayRest *restapisdm.GeneralInformationRest
}

func NewMstClassParticipantService(
	log *logrus.Logger,
	db *gorm.DB,
	cache cached.CacheRepository,
	mstClassParticipantRepository *repositorymodel.MstClassParticipantRepository,
	mstStudentBioRepository *repositorymodel.MstStudentBioRepository,
	mstClassRepository repositorymodel.MstClassRepository,

	// external
	generalInformationGatewayRest *restapisdm.GeneralInformationRest,
) *MstClassParticipantService {
	return &MstClassParticipantService{
		log:                           log,
		db:                            db,
		cache:                         cache,
		mstClassParticipantRepository: mstClassParticipantRepository,
		mstStudentBioRepository:       mstStudentBioRepository,
		mstClassRepository:            mstClassRepository,

		// external
		generalInformationGatewayRest: generalInformationGatewayRest,
	}
}

// TODO: Create
func (s *MstClassParticipantService) Create(ctx context.Context, req dto.MstClassParticipantRequest) error {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)
	v, _ := utils.StringToUuid(user.ID)

	data := new(model.MstClassParticipant)
	converter.ConvertMstClassParticipantRequestToModelPointer(req, data)
	err := s.mstClassParticipantRepository.Create(tx, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "create class participant service",
			"request": req,
			"user_id": v,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

func (s *MstClassParticipantService) CreateByProgramHead(ctx context.Context, req dto.MstClassParticipantRequest) error {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)
	v, _ := utils.StringToUuid(user.ID)

	// generalInformation, err := s.generalInformationGatewayRest.GetDataWithParamsOrQuery(
	// 	nil, nil, "/"+user.ID)
	// if err != nil {
	// 	s.log.WithFields(logrus.Fields{
	// 		"service": "get general information from sdm in class participant service create by programhead",
	// 		"user-id": user.ID,
	// 	}).Error(err)
	// 	return fiber.NewError(fiber.StatusInternalServerError, "get data sdm error "+err.Error())
	// }
	//
	// err = utilicems.CheckRestError(
	// 	s.log, generalInformation, user.ID,
	// 	"create class participant by programhead", "get general information from sdm in class participant service create by program head",
	// )
	// if err != nil {
	// 	return err
	// }
	//
	// studyProgramID := generalInformation.Data.StudyProgramID
	// if studyProgramID == nil {
	// 	s.log.WithFields(logrus.Fields{
	// 		"service":        "get general information from sdm in class service createByProgramHead",
	// 		"user-id":        user.ID,
	// 		"studyProgramID": "studyProgramID not found",
	// 	}).Info(err)
	// 	return fiber.NewError(fiber.StatusNotFound, "studyProgramID in sdm not found wit user -> "+user.ID)
	// }
	//
	// classData := new(model.MstClass)
	// err = s.mstClassRepository.GetByID(tx, req.ClassID.String(), classData)
	// if err != nil {
	// 	s.log.WithFields(logrus.Fields{
	// 		"service": "get class  by id in create by programhead class participan service",
	// 		"user-id": user.ID,
	// 	}).Error(err)
	// 	return utilicems.ErrorSpToErrorFiber(err)
	// }
	//
	// if classData.StudyProgramID.String() != studyProgramID.String() {
	// 	s.log.WithFields(logrus.Fields{
	// 		"service": "get general information from sdm in class participant service createaByProgramHead",
	// 		"user-id": user.ID,
	// 	}).Info("create class participant by program head not access")
	// 	return fiber.NewError(fiber.StatusForbidden, "ProgramHead not access data")
	// }

	_, _, err := s.checkStudyProgramUseProgramHead_InClassAndGeneralInformation(
		tx, user.ID, req.ClassID.String(), "create",
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "create by programhead class participan service",
			"user-id": user.ID,
		}).Error(err)
		// error use fiber.NewError
		return err
	}

	data := new(model.MstClassParticipant)
	converter.ConvertMstClassParticipantRequestToModelPointer(req, data)
	err = s.mstClassParticipantRepository.Create(tx, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "create class participant service",
			"request": req,
			"user_id": v,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

// TODO: Update

// TODO: Delete
func (s *MstClassParticipantService) DeleteByID(ctx context.Context, ID string) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	err := s.mstClassParticipantRepository.DeleteByID(tx, ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service":  "delete class participant service",
			"id":       ID,
			"deleteby": user.ID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

func (s *MstClassParticipantService) DeleteByIDAndProgramHead(ctx context.Context, ID string) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	classParticipantData := new(model.MstClassParticipant)
	err := s.mstClassParticipantRepository.GetByID(tx, ID, classParticipantData)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get class participant by id in delete by programhead class participan service",
			"user-id": user.ID,
		}).Error(err)
		return utilicems.ErrorSpToErrorFiber(err)
	}

	// generalInformation, err := s.generalInformationGatewayRest.GetDataWithParamsOrQuery(
	// 	nil, nil, "/"+user.ID)
	// if err != nil {
	// 	s.log.WithFields(logrus.Fields{
	// 		"service": "get general information from sdm in class lecutere service delete by programhead and id",
	// 		"user-id": user.ID,
	// 	}).Error(err)
	// 	return fiber.NewError(fiber.StatusInternalServerError, "get data sdm error "+err.Error())
	// }
	//
	// err = utilicems.CheckRestError(
	// 	s.log, generalInformation, user.ID,
	// 	"create class participant by programhead", "get general information from sdm in class lecturer service delete by program head",
	// )
	// if err != nil {
	// 	return err
	// }
	//
	// studyProgramID := generalInformation.Data.StudyProgramID
	// if studyProgramID == nil {
	// 	s.log.WithFields(logrus.Fields{
	// 		"service":        "get general information from sdm in class service deleteByProgramHead",
	// 		"user-id":        user.ID,
	// 		"studyProgramID": "studyProgramID not found",
	// 	}).Info(err)
	// 	return fiber.NewError(fiber.StatusNotFound, "studyProgramID in sdm not found wit user -> "+user.ID)
	// }
	//
	// classData := new(model.MstClass)
	// err = s.mstClassRepository.GetByID(tx, classParticipantData.ClassID.String(), classData)
	// if err != nil {
	// 	s.log.WithFields(logrus.Fields{
	// 		"service": "get class  by id in delete by programhead class participan service",
	// 		"user-id": user.ID,
	// 	}).Error(err)
	// 	return utilicems.ErrorSpToErrorFiber(err)
	// }
	//
	// if classData.StudyProgramID.String() != studyProgramID.String() {
	// 	s.log.WithFields(logrus.Fields{
	// 		"service": "get general information from sdm in class participant service deleteByProgramHead",
	// 		"user-id": user.ID,
	// 	}).Info("delete class participan by program head not access")
	// 	return fiber.NewError(fiber.StatusForbidden, "ProgramHead not access data")
	// }

	_, _, err = s.checkStudyProgramUseProgramHead_InClassAndGeneralInformation(
		tx, user.ID, classParticipantData.ClassID.String(), "delete",
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "delete by programhead class participan service",
			"user-id": user.ID,
		}).Error(err)
		// error use fiber.NewError
		return err
	}

	err = s.mstClassParticipantRepository.DeleteByID(tx, ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service":  "delete class participant service",
			"id":       ID,
			"deleteby": user.ID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

// TODO: Restore

// TODO: Read
func (s *MstClassParticipantService) GetByID(ctx context.Context, ID string) (res *dto.MstClassParticipantResponse, err error) {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)

	data := new(model.MstClassParticipant)
	err = s.mstClassParticipantRepository.GetByID(tx, ID, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get by id class participant service",
			"id":      ID,
			"user-id": user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return
	}
	res = converter.ConvertModelToMstClassParticipantResponsePointer(data)
	return
}

func (s *MstClassParticipantService) GetAllWithCount(ctx context.Context, req pageable.PageableRequestClassParticipant) (res *pageable.PageableResponse[dto.MstClassParticipantResponse], err error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	datas, count, err := s.mstClassParticipantRepository.GetAllWithCount(
		tx, true, req,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get all class participant service",
			"user-id": user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return
	}

	r := make([]dto.MstClassParticipantResponse, 0, count)

	for _, v := range datas {
		c := converter.ConvertModelToMstClassParticipantResponse(v)
		r = append(r, c)
	}
	res = &pageable.PageableResponse[dto.MstClassParticipantResponse]{
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

func (s *MstClassParticipantService) GetAllWithCountByClassIDForLecturer(ctx context.Context, req pageable.PageableRequestClassParticipant) (res *pageable.PageableResponse[dto.MstClassParticipantResponse], err error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	datas, count, err := s.mstClassParticipantRepository.GetAllWithCountByClassIDForLecturer(
		tx, true, req,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get all class participant service",
			"user-id": user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return
	}

	r := make([]dto.MstClassParticipantResponse, 0, count)

	for _, v := range datas {
		c := converter.ConvertModelToMstClassParticipantResponse(v)
		r = append(r, c)
	}
	res = &pageable.PageableResponse[dto.MstClassParticipantResponse]{
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

// helper for this service
func (s *MstClassParticipantService) checkStudyProgramUseProgramHead_InClassAndGeneralInformation(
	tx *gorm.DB, userID, classID, funcAction string,
) (*model.MstClass, *modelsdm.GeneralInformationResponse, error) {
	generalInformation, err := s.generalInformationGatewayRest.GetDataWithParamsOrQuery(
		nil, nil, "/"+userID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class participant service " + funcAction + " by programhead",
			"user-id": userID,
		}).Error(err)
		return nil, nil, fiber.NewError(fiber.StatusInternalServerError, "get data sdm error "+err.Error())
	}

	err = utilicems.CheckRestError(
		s.log, generalInformation, userID,
		funcAction+" class participant by programhead", "get general information from sdm in class participant service "+funcAction+" by program head",
	)
	if err != nil {
		return nil, nil, err
	}

	studyProgramID := generalInformation.Data.StudyProgramID
	if studyProgramID == nil {
		s.log.WithFields(logrus.Fields{
			"service":        "get general information from sdm in class participant service " + funcAction + "ByProgramHead",
			"user-id":        userID,
			"studyProgramID": "studyProgramID not found",
		}).Info(err)
		return nil, nil, fiber.NewError(fiber.StatusNotFound, "studyProgramID in sdm not found wit user -> "+userID)
	}

	classData := new(model.MstClass)
	err = s.mstClassRepository.GetByID(tx, classID, classData)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get class  by id in " + funcAction + "by programhead class participan service",
			"user-id": userID,
		}).Error(err)
		return nil, nil, utilicems.ErrorSpToErrorFiber(err)
	}

	if classData.StudyProgramID != studyProgramID.String() && classData.StudyProgramID != "0" {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class participant service " + funcAction + "ByProgramHead",
			"user-id": userID,
		}).Info(funcAction + " class participant by program head not access")
		return nil, nil, fiber.NewError(fiber.StatusForbidden, "ProgramHead not access data")
	}

	return classData, generalInformation.Data, nil
}

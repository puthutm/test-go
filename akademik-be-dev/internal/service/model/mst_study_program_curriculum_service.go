package servicemodel

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	modelsdm "gitlab.unsia.ac.id/icems/icems-tools/gateway/model/sdm"
	restapisdm "gitlab.unsia.ac.id/icems/icems-tools/gateway/rest-api/sdm"
	utilicems "gitlab.unsia.ac.id/icems/icems-tools/utils"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/delivery/http/middleware"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/dto/converter"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository/cached"
	repositorymodel "unsia.ac.id/akademic_be/internal/repository/model"
	"unsia.ac.id/akademic_be/pkg/utils"
)

type MstStudyProgramCurriculumService struct {
	log                                                    *logrus.Logger
	db                                                     *gorm.DB
	cache                                                  cached.CacheRepository
	mstStudyProgramCurriculumRepository                    *repositorymodel.MstStudyProgramCurriculumRepository
	mstClassRepository                                     repositorymodel.MstClassRepository
	mstSubjectRepository                                   *repositorymodel.MstSubjectRepository
	mstStudyProgramCurriculumSubjectPrerequisiteRepository *repositorymodel.MstStudyProgramCurriculumSubjectPrerequisiteRepository

	// external
	generalInformationGatewayRest *restapisdm.GeneralInformationRest
}

func NewMstStudyProgramCurriculumService(
	log *logrus.Logger,
	db *gorm.DB,
	cache cached.CacheRepository,
	mstStudyProgramCurriculumRepository *repositorymodel.MstStudyProgramCurriculumRepository,
	mstClassRepository repositorymodel.MstClassRepository,
	mstSubjectRepository *repositorymodel.MstSubjectRepository,
	mstStudyProgramCurriculumSubjectPrerequisiteRepository *repositorymodel.MstStudyProgramCurriculumSubjectPrerequisiteRepository,

	// external
	generalInformationGatewayRest *restapisdm.GeneralInformationRest,
) *MstStudyProgramCurriculumService {
	return &MstStudyProgramCurriculumService{
		log:                                 log,
		db:                                  db,
		cache:                               cache,
		mstStudyProgramCurriculumRepository: mstStudyProgramCurriculumRepository,
		mstClassRepository:                  mstClassRepository,
		mstSubjectRepository:                mstSubjectRepository,
		mstStudyProgramCurriculumSubjectPrerequisiteRepository: mstStudyProgramCurriculumSubjectPrerequisiteRepository,

		// external
		generalInformationGatewayRest: generalInformationGatewayRest,
	}
}

/* Create */

func (s *MstStudyProgramCurriculumService) Create(ctx context.Context, req dto.MstStudyProgramCurriculumRequest) (*dto.MstStudyProgramCurriculumResponseDetail, error) {
	tx := s.db.WithContext(ctx)

	err := s.mstStudyProgramCurriculumRepository.Create(tx, ctx, req)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	var reqSubjectPrerequisite model.MstStudyProgramCurriculumSubjectPrerequisite

	IDUUID, _ := utilicems.StringToUuid(req.ID)
	reqSubjectPrerequisite.StudyProgramCurriculumID = IDUUID
	for _, v := range req.SubjectPrerequisite {

		reqSubjectPrerequisite.ID = utilicems.GenerateUUID()
		reqSubjectPrerequisite.StudyProgramCurriculumSubjectID = v

		err := s.mstStudyProgramCurriculumSubjectPrerequisiteRepository.Create(tx, reqSubjectPrerequisite)
		if err != nil {
			createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
			s.log.WithError(err).Error(createMsg)

			return nil, utils.ErrorSpToErrorFiber(err)
		}
	}

	studyProgramCurriculum, err := s.GetByID(ctx, req.ID)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), req.ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	return studyProgramCurriculum, nil
}

func (s *MstStudyProgramCurriculumService) CreateByProgramHead(ctx context.Context, req dto.MstStudyProgramCurriculumRequest) (*dto.MstStudyProgramCurriculumResponseDetail, error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	subjectData, _, err := s.checkStudyProgramUseProgramHead_GeneralInformationAndSubject(
		tx, user.ID, req.SubjectID, "create",
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "create by programhead study program curiculum service",
			"user-id": user.ID,
		}).Error(err)
		// error use fiber.NewError
		return nil, err
	}

	req.StudyProgramID = &subjectData.StudyProgramID

	return s.Create(ctx, req)
}

/* Read */
func (s *MstStudyProgramCurriculumService) GetByID(ctx context.Context, ID string) (*dto.MstStudyProgramCurriculumResponseDetail, error) {
	tx := s.db.WithContext(ctx)

	classModel := new(model.MstStudyProgramCurriculum)

	err := s.mstStudyProgramCurriculumRepository.GetByID(tx, ID, classModel)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	studyProgramCurriculum := converter.MstStudyProgramCurriculumModelToResponseDetail(*classModel)
	subjectPrerequisiteData, err := s.mstStudyProgramCurriculumSubjectPrerequisiteRepository.GetByStudyProgramCurriculumID(tx, classModel.ID.String())
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

	}
	subjectPrerequisiteDto := make([]dto.MstStudyProgramCurriculumSubjectPrerequisiteResponse, 0, len(subjectPrerequisiteData))
	for _, v := range subjectPrerequisiteData {
		subjectPrerequisiteDto = append(subjectPrerequisiteDto, converter.MstStudyProgramCurriculumSubjectPrerequisiteModelToResponse(v))
	}

	studyProgramCurriculum.SubjectPrerequisite = subjectPrerequisiteDto

	return studyProgramCurriculum, nil
}

func (s *MstStudyProgramCurriculumService) GetByIDAndProgramHead(ctx context.Context, ID string) (*dto.MstStudyProgramCurriculumResponseDetail, error) {
	tx := s.db.WithContext(ctx)

	classModel := new(model.MstStudyProgramCurriculum)
	err := s.mstStudyProgramCurriculumRepository.GetByID(tx, ID, classModel)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	user := middleware.GetUserClaimsCtx(ctx)

	generalInformationGatewayRest, err := s.checkStudyProgramUseProgramHead_GeneralInformation(
		user.ID, "read",
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "create by programhead study program curiculum service",
			"user-id": user.ID,
		}).Error(err)
		// error use fiber.NewError
		return nil, err
	}

	if classModel.StudyProgramID.String() != generalInformationGatewayRest.StudyProgramID.String() {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in studyProgramCurriculum service read ByProgramHead",
			"user-id": user.ID,
		}).Info("read class by program head not access")
		return nil, fiber.NewError(fiber.StatusForbidden, "ProgramHead not access data")
	}

	studyProgramCurriculum := converter.MstStudyProgramCurriculumModelToResponseDetail(*classModel)
	subjectPrerequisiteData, err := s.mstStudyProgramCurriculumSubjectPrerequisiteRepository.GetByStudyProgramCurriculumID(tx, classModel.ID.String())
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

	}
	subjectPrerequisiteDto := make([]dto.MstStudyProgramCurriculumSubjectPrerequisiteResponse, 0, len(subjectPrerequisiteData))
	for _, v := range subjectPrerequisiteData {
		subjectPrerequisiteDto = append(subjectPrerequisiteDto, converter.MstStudyProgramCurriculumSubjectPrerequisiteModelToResponse(v))
	}

	studyProgramCurriculum.SubjectPrerequisite = subjectPrerequisiteDto

	return studyProgramCurriculum, nil
}

func (s *MstStudyProgramCurriculumService) GetByStudyProgramIDAndSemesterID(
	ctx context.Context, deleted bool,
	StudyProgramID, SemesterNumberID, CurriculumYearID string,
) (*dto.MstStudyProgramCurriculumWithTotalResponse, error) {
	tx := s.db.WithContext(ctx)

	studyProgramCurriculums, err := s.mstStudyProgramCurriculumRepository.GetByStudyProgramIDAndSemesterID(tx, deleted, StudyProgramID, SemesterNumberID, CurriculumYearID)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	res := new(dto.MstStudyProgramCurriculumWithTotalResponse)

	results := make([]dto.MstStudyProgramCurriculumResponse, 0, len(studyProgramCurriculums))
	for _, class := range studyProgramCurriculums {
		subjectPrerequisiteData, err := s.mstStudyProgramCurriculumSubjectPrerequisiteRepository.GetByStudyProgramCurriculumID(tx, class.ID.String())
		if err != nil {
			createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
			s.log.WithError(err).Error(createMsg)

		}
		subjectPrerequisiteDto := make([]dto.MstStudyProgramCurriculumSubjectPrerequisiteResponse, 0, len(subjectPrerequisiteData))
		for _, v := range subjectPrerequisiteData {
			subjectPrerequisiteDto = append(subjectPrerequisiteDto, converter.MstStudyProgramCurriculumSubjectPrerequisiteModelToResponse(v))
		}

		studyProgramCurriculumDto := *converter.MstStudyProgramCurriculumModelToResponse(class)
		studyProgramCurriculumDto.SubjectPrerequisite = subjectPrerequisiteDto

		results = append(results, studyProgramCurriculumDto)

		if class.SubjectTotalSKS == nil {
			continue
		}
		res.Total.SKS += *class.SubjectTotalSKS
		if class.SubjectTotalSKS == nil {
			continue
		}
		if *class.IsMandatory {
			res.Total.Mandatory++
		} else {
			res.Total.NoMandatory++
		}
	}

	res.MstStudyProgramCurriculumsResponse = results
	return res, nil
}

func (s *MstStudyProgramCurriculumService) GetByStudyProgramIDAndSemesterIDAndProgramHead(
	ctx context.Context,
	deleted bool, SemesterNumberID string, CurriculumYearID string,
) (*dto.MstStudyProgramCurriculumWithTotalResponse, error) {
	user := middleware.GetUserClaimsCtx(ctx)

	generalInformationGatewayRest, err := s.checkStudyProgramUseProgramHead_GeneralInformation(
		user.ID, "create",
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "read by programhead study program curiculum service",
			"user-id": user.ID,
		}).Error(err)
		// error use fiber.NewError
		return nil, err
	}

	// studyProgramCurriculums, err := s.mstStudyProgramCurriculumRepository.GetByStudyProgramIDAndSemesterID(tx, deleted, generalInformationGatewayRest.StudyProgramID.String(), SemesterNumberID, CurriculumYearID)
	// if err != nil {
	// 	createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
	// 	s.log.WithError(err).Error(createMsg)
	//
	// 	return nil, utils.ErrorSpToErrorFiber(err)
	// }
	//
	// results := make([]dto.MstStudyProgramCurriculumResponse, 0)
	// for _, class := range studyProgramCurriculums {
	// 	results = append(results, *converter.MstStudyProgramCurriculumModelToResponse(class))
	// }
	//
	// return results, nil
	return s.GetByStudyProgramIDAndSemesterID(ctx, deleted, generalInformationGatewayRest.StudyProgramID.String(), SemesterNumberID, CurriculumYearID)
}

/* Update */
func (s *MstStudyProgramCurriculumService) UpdateByID(ctx context.Context, req dto.MstStudyProgramCurriculumRequest) (*dto.MstStudyProgramCurriculumResponseDetail, error) {
	tx := s.db.WithContext(ctx)

	err := s.mstStudyProgramCurriculumRepository.UpdateByID(tx, ctx, req)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), req.ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	err = s.mstStudyProgramCurriculumSubjectPrerequisiteRepository.DeleteByStudyProgramCurriculumID(tx, req.ID)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), req.ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}
	var reqSubjectPrerequisite model.MstStudyProgramCurriculumSubjectPrerequisite

	IDUUID, _ := utilicems.StringToUuid(req.ID)
	reqSubjectPrerequisite.StudyProgramCurriculumID = IDUUID
	for _, v := range req.SubjectPrerequisite {

		reqSubjectPrerequisite.ID = utilicems.GenerateUUID()
		reqSubjectPrerequisite.StudyProgramCurriculumSubjectID = v

		err := s.mstStudyProgramCurriculumSubjectPrerequisiteRepository.Create(tx, reqSubjectPrerequisite)
		if err != nil {
			createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
			s.log.WithError(err).Error(createMsg)

			return nil, utils.ErrorSpToErrorFiber(err)
		}
	}

	studyProgramCurriculum, err := s.GetByID(ctx, req.ID)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), req.ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	return studyProgramCurriculum, nil
}

func (s *MstStudyProgramCurriculumService) UpdateByIDAndProgramHead(ctx context.Context, req dto.MstStudyProgramCurriculumRequest) (*dto.MstStudyProgramCurriculumResponseDetail, error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	subjectData, _, err := s.checkStudyProgramUseProgramHead_GeneralInformationAndSubject(
		tx, user.ID, req.SubjectID, "update",
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "update by programhead study program curiculum service",
			"user-id": user.ID,
		}).Error(err)
		// error use fiber.NewError
		return nil, err
	}

	req.StudyProgramID = &subjectData.StudyProgramID

	// err = s.mstStudyProgramCurriculumRepository.UpdateByID(tx, ctx, req)
	// if err != nil {
	// 	createMsg := utils.CreateMsgDebuging(err.Error(), req.ID, utils.ErrorLocation())
	// 	s.log.WithError(err).Error(createMsg)
	//
	// 	return nil, utils.ErrorSpToErrorFiber(err)
	// }
	//
	// studyProgramCurriculum, err := s.GetByID(ctx, req.ID)
	// if err != nil {
	// 	createMsg := utils.CreateMsgDebuging(err.Error(), req.ID, utils.ErrorLocation())
	// 	s.log.WithError(err).Error(createMsg)
	//
	// 	return nil, utils.ErrorSpToErrorFiber(err)
	// }

	return s.UpdateByID(ctx, req)
}

func (s *MstStudyProgramCurriculumService) UpdateBlastPackageBySemesterWithProgramStudy(
	ctx context.Context,
	req dto.UpdatePackageMstStudyProgramCurriculumWithStudyProgramRequest,
) error {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)

	err := s.mstStudyProgramCurriculumRepository.UpdateBlastPackageBySemesterWithProgramStudy(tx, req)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithFields(logrus.Fields{
			"service": "update package study program curiculum service with programStudy",
			"user-id": user.ID,
		}).Error(createMsg)
		// error use fiber.NewError
		return utils.ErrorSpToErrorFiber(err)
	}

	return nil
}

func (s *MstStudyProgramCurriculumService) UpdateBlastPackageBySemesterWithoutProgramStudy(
	ctx context.Context,
	req dto.UpdatePackageMstStudyProgramCurriculumWithoutStudyProgramRequest,
) error {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)

	req.UserID = user.ID

	err := s.mstStudyProgramCurriculumRepository.UpdateBlastPackageBySemesterWithoutProgramStudy(tx, req)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithFields(logrus.Fields{
			"service": "update package study program curiculum service without programStudy",
			"user-id": user.ID,
		}).Error(createMsg)
		// error use fiber.NewError
		return utils.ErrorSpToErrorFiber(err)
	}

	return nil
}

func (s *MstStudyProgramCurriculumService) RestoreByID(ctx context.Context, ID string) error {
	tx := s.db.WithContext(ctx)

	err := s.mstStudyProgramCurriculumRepository.RestoreByID(tx, ID)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return utils.ErrorSpToErrorFiber(err)
	}

	return nil
}

func (s *MstStudyProgramCurriculumService) RestoreByIDAndProgramHead(ctx context.Context, ID string) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	_, err := s.checkStudyProgramUseProgramHead_GeneralInformation(
		user.ID, "restore",
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "restore by programhead study program curiculum service",
			"user-id": user.ID,
		}).Error(err)
		// error use fiber.NewError
		return err
	}

	err = s.mstStudyProgramCurriculumRepository.RestoreByID(tx, ID)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return utils.ErrorSpToErrorFiber(err)
	}

	return nil
}

/* Delete */
func (s *MstStudyProgramCurriculumService) DeleteByID(ctx context.Context, ID string) error {
	tx := s.db.WithContext(ctx)

	err := s.mstStudyProgramCurriculumRepository.DeleteByID(tx, ctx, ID)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return utils.ErrorSpToErrorFiber(err)
	}

	err = s.mstStudyProgramCurriculumSubjectPrerequisiteRepository.DeleteByStudyProgramCurriculumID(tx, ID)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return utils.ErrorSpToErrorFiber(err)
	}

	return nil
}

func (s *MstStudyProgramCurriculumService) DeleteByIDAndProgramHead(ctx context.Context, ID string) error {
	user := middleware.GetUserClaimsCtx(ctx)

	_, err := s.checkStudyProgramUseProgramHead_GeneralInformation(
		user.ID, "delete",
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "delete by programhead study program curiculum service",
			"user-id": user.ID,
		}).Error(err)
		// error use fiber.NewError
		return err
	}

	return s.DeleteByID(ctx, ID)
}

func (s *MstStudyProgramCurriculumService) getByStudyProgramAndSemesterAndCuricullumForSubjectData(
	ctx context.Context,
	req dto.GetStudyProgramCurriculumRequest,
	role model.Role,
) (res []dto.MstStudyProgramCurriculumResponse, err error) {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)
	req.UserID = user.ID

	studyProgramCurriculums, err := s.mstStudyProgramCurriculumRepository.GetByStudyProgramAndSemesterAndCuricullumForSubjectData(
		tx, role, req)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)
		s.log.WithFields(logrus.Fields{
			"repsoitory from":      "studyProgramCurriculums -> GetByStudyProgramAndSemesterAndCuricullumForSubjectData",
			"service":              "GetByStudyProgramAndSemesterAndCuricullumForSubjectData",
			"error-message-custom": "sp bisa di cek",
		}).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	results := make([]dto.MstStudyProgramCurriculumResponse, 0, len(studyProgramCurriculums))
	for _, class := range studyProgramCurriculums {
		studyProgramCurriculumDto := *converter.MstStudyProgramCurriculumModelToResponse(class)
		results = append(results, studyProgramCurriculumDto)
	}

	return results, nil
}

func (s *MstStudyProgramCurriculumService) GetByStudyProgramAndSemesterAndCuricullumForSubjectDataAcademic(
	ctx context.Context,
	req dto.GetStudyProgramCurriculumRequest,
) (res []dto.MstStudyProgramCurriculumResponse, err error) {
	// logic for academic
	return s.getByStudyProgramAndSemesterAndCuricullumForSubjectData(ctx, req, model.Academic)
}

func (s *MstStudyProgramCurriculumService) GetByStudyProgramAndSemesterAndCuricullumForSubjectDataProgramHead(
	ctx context.Context,
	req dto.GetStudyProgramCurriculumRequest,
) (res []dto.MstStudyProgramCurriculumResponse, err error) {
	// logic for programHead
	return s.getByStudyProgramAndSemesterAndCuricullumForSubjectData(ctx, req, model.ProgramHead)
}

// helper for this service
func (s *MstStudyProgramCurriculumService) checkStudyProgramUseProgramHead_GeneralInformationAndSubject(
	tx *gorm.DB, userID, subjectID, funcAction string,
) (*model.MstSubject, *modelsdm.GeneralInformationResponse, error) {
	generalInformation, err := s.generalInformationGatewayRest.GetDataWithParamsOrQuery(
		nil, nil, "/"+userID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in studyProgramCurriculum service " + funcAction + " by programhead",
			"user-id": userID,
		}).Error(err)
		return nil, nil, fiber.NewError(fiber.StatusInternalServerError, "get data sdm error "+err.Error())
	}

	err = utilicems.CheckRestError(
		s.log, generalInformation, userID,
		funcAction+" studyProgramCurriculum  by programhead", "get general information from sdm in studyProgramCurriculum service "+funcAction+" by program head",
	)
	if err != nil {
		return nil, nil, err
	}

	studyProgramID := generalInformation.Data.StudyProgramID
	if studyProgramID == nil {
		s.log.WithFields(logrus.Fields{
			"service":        "get general information from sdm in studyProgramCurriculum service" + funcAction + "ByProgramHead",
			"user-id":        userID,
			"studyProgramID": "studyProgramID not found",
		}).Info(err)
		return nil, nil, fiber.NewError(fiber.StatusNotFound, "studyProgramID in sdm not found wit user -> "+userID)
	}

	subjectData := new(model.MstSubject)
	err = s.mstSubjectRepository.GetByID(tx, subjectID, subjectData)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), userID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, nil, utils.ErrorSpToErrorFiber(err)
	}

	if subjectData.StudyProgramID != studyProgramID.String() {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in studyProgramCurriculum service " + funcAction + "ByProgramHead",
			"user-id": userID,
		}).Info(funcAction + "studyProgramCurriculum by program head not access")
		return nil, nil, fiber.NewError(fiber.StatusForbidden, "ProgramHead not access data")
	}

	return subjectData, generalInformation.Data, nil
}

// helper for this service
func (s *MstStudyProgramCurriculumService) checkStudyProgramUseProgramHead_GeneralInformation(
	userID string, funcAction string,
) (*modelsdm.GeneralInformationResponse, error) {
	generalInformation, err := s.generalInformationGatewayRest.GetDataWithParamsOrQuery(
		nil, nil, "/"+userID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in studyProgramCurriculum service " + funcAction + " by programhead",
			"user-id": userID,
		}).Error(err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "get data sdm error "+err.Error())
	}

	err = utilicems.CheckRestError(
		s.log, generalInformation, userID,
		funcAction+" class by programhead", "get general information from sdm in studyProgramCurriculum service "+funcAction+" by program head",
	)
	if err != nil {
		return nil, err
	}

	studyProgramID := generalInformation.Data.StudyProgramID
	if studyProgramID == nil {
		s.log.WithFields(logrus.Fields{
			"service":        "get general information from sdm in studyProgramCurriculum sservice " + funcAction + "ByProgramHead",
			"user-id":        userID,
			"studyProgramID": "studyProgramID not found",
		}).Info(err)
		return nil, fiber.NewError(fiber.StatusNotFound, "studyProgramID in sdm not found wit user -> "+userID)
	}

	return generalInformation.Data, nil
}

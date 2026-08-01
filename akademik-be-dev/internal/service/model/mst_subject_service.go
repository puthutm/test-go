package servicemodel

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	restapisdm "unsia.ac.id/akademic_be/pkg/icems-tools/gateway/rest-api/sdm"
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

type MstSubjectService struct {
	log                    *logrus.Logger
	db                     *gorm.DB
	cache                  cached.CacheRepository
	mstSubjectRepository   *repositorymodel.MstSubjectRepository
	generalInformationRest *restapisdm.GeneralInformationRest
}

func NewMstSubjectService(
	log *logrus.Logger,
	db *gorm.DB,
	cache cached.CacheRepository,
	mstSubjectRepository *repositorymodel.MstSubjectRepository,
	generalInformationRest *restapisdm.GeneralInformationRest,
) *MstSubjectService {
	return &MstSubjectService{
		log:                    log,
		db:                     db,
		cache:                  cache,
		mstSubjectRepository:   mstSubjectRepository,
		generalInformationRest: generalInformationRest,
	}
}

// TODO: Create
func (s *MstSubjectService) Create(ctx context.Context, req dto.MstSubjectRequest) error {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)
	v, _ := utils.StringToUuid(user.ID)
	req.UserID = v

	data := new(model.MstSubject)
	converter.ConvertMstSubjectRequestToModelPointer(req, data)
	data.TotalSKS = req.SimulationSKS + req.FaceToFaceSKS + req.PracticumSKS + req.FieldPracticeSKS
	err := s.mstSubjectRepository.Create(tx, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "create subject service",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}

	SubjectID := data.ID

	/* Supporting Lecturer*/
	err = s.mstSubjectRepository.CreateSupportingLecturer(tx, SubjectID, req.SupportingLecturerID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "create subject service",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}

	/* Developer RPS Lecuter */
	err = s.mstSubjectRepository.CreateDeveloperRPSLecuter(tx, SubjectID, req.DeveloperRPSLecuterID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "create subject service",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}

	/* SubjectCoordinatorLecuterID */
	err = s.mstSubjectRepository.CreateSubjectCoordinatorLecuter(tx, SubjectID, req.SubjectCoordinatorLecuterID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "create subject service",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}

	return nil
}

// TODO: Update
func (s *MstSubjectService) UpdateByID(ctx context.Context, req dto.MstSubjectUpdate) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)
	data := new(model.MstSubject)
	v, _ := utils.StringToUuid(user.ID)
	req.UserID = v

	converter.ConvertMstSubjectUpdateToModelPointer(req, data)
	data.TotalSKS = req.SimulationSKS + req.FaceToFaceSKS + req.PracticumSKS + req.FieldPracticeSKS
	err := s.mstSubjectRepository.UpdateByID(tx, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "update subject service",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}

	SubjectID := data.ID

	/* Supporting Lecturer*/
	err = s.mstSubjectRepository.DeleteBySupportingLecturerID(tx, SubjectID.String())
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "create subject service",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}

	err = s.mstSubjectRepository.CreateSupportingLecturer(tx, SubjectID, req.SupportingLecturerID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "create subject service",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}

	/* Developer RPS Lecuter */
	err = s.mstSubjectRepository.DeleteByDeveloperRPSLecuterID(tx, SubjectID.String())
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "create subject service",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}

	err = s.mstSubjectRepository.CreateDeveloperRPSLecuter(tx, SubjectID, req.DeveloperRPSLecuterID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "create subject service",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}

	/* SubjectCoordinatorLecuterID */
	err = s.mstSubjectRepository.DeleteBySubjectCoordinatorLecuterID(tx, SubjectID.String())
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "create subject service",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}

	err = s.mstSubjectRepository.CreateSubjectCoordinatorLecuter(tx, SubjectID, req.SubjectCoordinatorLecuterID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "create subject service",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}

	return nil
}

// TODO: Delete
func (s *MstSubjectService) DeleteByID(ctx context.Context, ID string) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	err := s.mstSubjectRepository.DeleteByID(tx, ID, user.ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service":  "delete subject service",
			"id":       ID,
			"deleteby": user.ID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

// TODO: Restore
func (s *MstSubjectService) RestoreByID(ctx context.Context, ID string) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	err := s.mstSubjectRepository.RestoreByID(tx, ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "Restore subject service",
			"id":      ID,
			"user-id": user.ID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

// TODO: Read
func (s *MstSubjectService) GetByID(ctx context.Context, ID string) (res *dto.MstSubjectResponse, err error) {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)

	data := new(model.MstSubject)
	err = s.mstSubjectRepository.GetByID(tx, ID, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get by id subject service",
			"id":      ID,
			"user-id": user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return
	}
	res = converter.ConvertModelToMstSubjectResponsePointer(data)
	return
}

func (s *MstSubjectService) GetSupportingLecturerBySubjectID(ctx context.Context, ID string) (res []model.MstSubject, err error) {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)

	supportingLecturers, err := s.mstSubjectRepository.GetSupportingLecturerBySubjectID(tx, ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get by id subject service",
			"id":      ID,
			"user-id": user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return
	}
	return supportingLecturers, nil
}

func (s *MstSubjectService) GetDeveloperRPSBySubjectID(ctx context.Context, ID string) (res []model.MstSubject, err error) {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)

	devel, err := s.mstSubjectRepository.GetDeveloperRPSBySubjectID(tx, ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get by id subject service",
			"id":      ID,
			"user-id": user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return
	}
	return devel, nil
}

func (s *MstSubjectService) GetSubjectCoordinatorBySubjectID(ctx context.Context, ID string) (res []model.MstSubject, err error) {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)

	devel, err := s.mstSubjectRepository.GetSubjectCoordinatorBySubjectID(tx, ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get by id subject service",
			"id":      ID,
			"user-id": user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return
	}
	return devel, nil
}

func (s *MstSubjectService) GetAllWithCount(ctx context.Context, req pageable.PageableRequestSubject) (res *pageable.PageableResponse[dto.MstSubjectResponse], err error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	datas, count, err := s.mstSubjectRepository.GetAllWithCount(tx, true, req)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get all subject service",
			"user-id": user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return
	}

	r := make([]dto.MstSubjectResponse, 0, count)

	for _, v := range datas {
		c := converter.ConvertModelToMstSubjectResponse(&v)
		r = append(r, c)
	}
	res = &pageable.PageableResponse[dto.MstSubjectResponse]{
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

func (s *MstSubjectService) GetAllTrashWithCount(ctx context.Context, req pageable.PageableRequestSubject) (res *pageable.PageableResponse[dto.MstSubjectResponse], err error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	datas, count, err := s.mstSubjectRepository.GetAllWithCount(tx, false, req)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get all trash subject service",
			"user-id": user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return
	}

	r := make([]dto.MstSubjectResponse, 0, count)

	for _, v := range datas {
		c := converter.ConvertModelToMstSubjectResponse(&v)
		r = append(r, c)
	}
	res = &pageable.PageableResponse[dto.MstSubjectResponse]{
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

func (s *MstSubjectService) GetAllWithCountByLecuturerID(ctx context.Context, req pageable.PageableRequestSubject) (res *pageable.PageableResponse[dto.MstSubjectResponse], err error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	datas, count, err := s.mstSubjectRepository.GetAllWithCountByLecuturerID(tx, ctx, req)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get all trash subject service",
			"user-id": user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return
	}

	r := make([]dto.MstSubjectResponse, 0, count)

	for _, v := range datas {
		c := converter.ConvertModelToMstSubjectResponse(&v)
		r = append(r, c)
	}
	res = &pageable.PageableResponse[dto.MstSubjectResponse]{
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

func (s *MstSubjectService) GetAllWithCountByCoordinatorLecuturerID(ctx context.Context, req pageable.PageableRequestSubject) (res *pageable.PageableResponse[dto.MstSubjectResponse], err error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	datas, count, err := s.mstSubjectRepository.GetAllWithCountByCoordinatorLecuturerID(tx, ctx, req)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get subject service coordinatorlecturer",
			"user-id": user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return
	}

	r := make([]dto.MstSubjectResponse, 0, count)

	for _, v := range datas {
		c := converter.ConvertModelToMstSubjectResponse(&v)
		r = append(r, c)
	}
	res = &pageable.PageableResponse[dto.MstSubjectResponse]{
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

func (s *MstSubjectService) GetByStudyProgramAndCurriculumYear(
	ctx context.Context,
	studyProgramID, curriculumYearID string,
) ([]dto.MstSubjectResponseForSearch, error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	datas, err := s.mstSubjectRepository.GetByStudyProgramIDAndCurriculumYearID(
		tx,
		studyProgramID, curriculumYearID,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get subject service by studyProgramID and curriculumYearID",
			"user-id": user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return nil, err
	}

	r := make([]dto.MstSubjectResponseForSearch, 0, len(datas))

	for _, v := range datas {
		c := converter.ConvertModelToMstSubjectResponseForSearch(v)
		r = append(r, c)
	}
	return r, err
}

// get subject by kaprodi an curriculumYearID
func (s *MstSubjectService) GetByProgramHeadAndCurriculumYear(
	ctx context.Context,
	curriculumYearID string,
) ([]dto.MstSubjectResponseForSearch, error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	generalInformation, err := s.generalInformationRest.GetDataWithParamsOrQuery(nil, nil, "/"+user.ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in subject service GetByProgramHeadAndCurriculumYear",
			"user-id": user.ID,
		}).Error(err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "get data sdm error "+err.Error())
	}

	if generalInformation.Error {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in subject service GetByProgramHeadAndCurriculumYear",
			"user-id": user.ID,
		}).Error(err)
		return nil, fiber.NewError(generalInformation.Status, generalInformation.Message)
	}

	if generalInformation.Data == nil {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in subject service GetByProgramHeadAndCurriculumYear",
			"user-id": user.ID,
		}).Error(err)
		return nil, fiber.NewError(fiber.StatusNotFound, "generalInformation not found wit user -> "+user.ID)
	}

	studyProgramID := generalInformation.Data.StudyProgramID

	if studyProgramID == nil {
		s.log.WithFields(logrus.Fields{
			"service":        "get general information from sdm in subject service GetByProgramHeadAndCurriculumYear",
			"user-id":        user.ID,
			"studyProgramID": "studyProgramID not found",
		}).Info(err)
		return nil, nil
	}

	datas, err := s.mstSubjectRepository.GetByStudyProgramIDAndCurriculumYearID(
		tx,
		studyProgramID.String(), curriculumYearID,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get subject service by kaprodi and curriculum year",
			"user-id": user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return nil, err
	}

	r := make([]dto.MstSubjectResponseForSearch, 0, len(datas))

	for _, v := range datas {
		c := converter.ConvertModelToMstSubjectResponseForSearch(v)
		r = append(r, c)
	}
	if len(r) == 0 {
		return nil, nil
	}
	return r, err
}

package servicemodel

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/teambition/rrule-go"
	modelsdm "gitlab.unsia.ac.id/icems/icems-tools/gateway/model/sdm"
	restapidatareferensi "gitlab.unsia.ac.id/icems/icems-tools/gateway/rest-api/datareferensi"
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

type MstClassScheduleService struct {
	log                                *logrus.Logger
	db                                 *gorm.DB
	cache                              cached.CacheRepository
	mstClassScheduleRepository         *repositorymodel.MstClassScheduleRepository
	storageService                     *service.StorageService
	mstClassRepository                 repositorymodel.MstClassRepository
	mstClassScheduleTemplateRepository *repositorymodel.MstClassScheduleTemplateRepository

	// external
	generalInformationGatewayRest *restapisdm.GeneralInformationRest
	academicPeriodgatewayRest     *restapidatareferensi.AcademicPeriodDetailWithSessionRest
}

func NewMstClassScheduleService(
	log *logrus.Logger,
	db *gorm.DB,
	cache cached.CacheRepository,
	mstClassScheduleRepository *repositorymodel.MstClassScheduleRepository,
	storageService *service.StorageService,
	mstClassRepository repositorymodel.MstClassRepository,

	mstClassScheduleTemplateRepository *repositorymodel.MstClassScheduleTemplateRepository,

	// external
	generalInformationGatewayRest *restapisdm.GeneralInformationRest,
	academicPeriodgatewayRest *restapidatareferensi.AcademicPeriodDetailWithSessionRest,
) *MstClassScheduleService {
	return &MstClassScheduleService{
		log:                                log,
		db:                                 db,
		cache:                              cache,
		mstClassScheduleRepository:         mstClassScheduleRepository,
		storageService:                     storageService,
		mstClassRepository:                 mstClassRepository,
		mstClassScheduleTemplateRepository: mstClassScheduleTemplateRepository,

		// external
		generalInformationGatewayRest: generalInformationGatewayRest,
		academicPeriodgatewayRest:     academicPeriodgatewayRest,
	}
}

// TODO: Create
func (s *MstClassScheduleService) Create(ctx context.Context, req dto.MstClassScheduleRequest) error {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)

	data := new(model.MstClassSchedule)
	converter.ConvertMstClassScheduleRequestToModelPointer(req, data)

	err := s.mstClassScheduleRepository.Create(tx, data)
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

func (s *MstClassScheduleService) CreateByProgramHead(ctx context.Context, req dto.MstClassScheduleRequest) error {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)
	_, _, err := s.checkStudyProgramUseProgramHead_GeneralInformation(
		tx, user.ID, req.ClassID.String(), "create",
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "create by programhead class schedule service",
			"user-id": user.ID,
		}).Error(err)
		// error use fiber.NewError
		return err
	}
	data := new(model.MstClassSchedule)
	converter.ConvertMstClassScheduleRequestToModelPointer(req, data)

	err = s.mstClassScheduleRepository.Create(tx, data)
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
func (s *MstClassScheduleService) UpdateByDayTime(ctx context.Context, req dto.MstClassScheduleUpdateRequest) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	data := converter.ConvertMstClassScheduleUpdateRequestToCommand(req)
	err := s.mstClassScheduleRepository.UpdateByDayTime(tx, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "update class class schedule service",
			"request": req,
			"user-id": user.ID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

func (s *MstClassScheduleService) UpdateByIDForLecturer(ctx context.Context, req dto.MstClassScheduleUpdateForLecturerRequest) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)
	data := new(model.MstClassSchedule)

	data.ID = req.ID
	data.MaterialPlan = req.MaterialPlan
	data.MaterialRealization = req.MaterialRealization

	validExtension := []string{".jpg", ".pdf"}
	if req.MaterialAttachmentFile != nil {
		materialAttachmentFilePath, err := s.storageService.UploadFileV3(
			ctx, req.MaterialAttachmentFile, true,
			"lecturer/academic/class-schedule/material-attachment-file", "",
			validExtension, 3,
		)
		if err != nil {
			s.log.WithFields(logrus.Fields{
				"service": "update class schedule service for UploadFileV3 failed",
				"request": req,
				"user-id": user.ID,
			}).Error(err)
			return utils.ErrorSpToErrorFiber(err)
		}
		if materialAttachmentFilePath == "" {
			s.log.WithFields(logrus.Fields{
				"service": "update class schedule service for UploadFileV3 failed",
				"request": req,
				"user-id": user.ID,
			}).Error(err)
			return fiber.NewError(fiber.StatusInternalServerError, "failed upload file")
		}
		data.MaterialAttachmentFilePath = &materialAttachmentFilePath
	}

	if req.AttendanceDocumentFile != nil {
		attendanceDocumentFilePath, err := s.storageService.UploadFileV3(
			ctx, req.AttendanceDocumentFile, true,
			"lecturer/academic/class-schedule/attendance-document-file", "",
			validExtension, 3,
		)
		if err != nil {
			s.log.WithFields(logrus.Fields{
				"service": "update class schedule service for UploadFileV3 failed",
				"request": req,
				"user-id": user.ID,
			}).Error(err)
			return utils.ErrorSpToErrorFiber(err)
		}
		if attendanceDocumentFilePath == "" {
			s.log.WithFields(logrus.Fields{
				"service": "update class schedule service for UploadFileV3 failed",
				"request": req,
				"user-id": user.ID,
			}).Error(err)
			return fiber.NewError(fiber.StatusInternalServerError, "failed upload file")
		}
		data.AttendanceDocumentFilePath = &attendanceDocumentFilePath
	}

	if req.JournalDocumentFile != nil {
		journalDocumentFilePath, err := s.storageService.UploadFileV3(
			ctx, req.JournalDocumentFile, true,
			"lecturer/academic/class-schedule/attendance-document-file", "",
			validExtension, 3,
		)
		if err != nil {
			s.log.WithFields(logrus.Fields{
				"service": "update class schedule service for UploadFileV3 failed",
				"request": req,
				"user-id": user.ID,
			}).Error(err)
			return utils.ErrorSpToErrorFiber(err)
		}
		if journalDocumentFilePath == "" {
			s.log.WithFields(logrus.Fields{
				"service": "update class schedule service for UploadFileV3 failed",
				"request": req,
				"user-id": user.ID,
			}).Error(err)
			return fiber.NewError(fiber.StatusInternalServerError, "failed upload file")
		}
		data.JournalDocumentFilePath = &journalDocumentFilePath
	}

	err := s.mstClassScheduleRepository.UpdateByIDForLecturer(tx, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "update class schedule service",
			"request": req,
			"user-id": user.ID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

// TODO: Delete
func (s *MstClassScheduleService) DeleteByID(ctx context.Context, ID string) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	err := s.mstClassScheduleRepository.DeleteByID(tx, ID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service":  "delete class schedule service",
			"id":       ID,
			"deleteby": user.ID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

func (s *MstClassScheduleService) DeleteDayTime(
	ctx context.Context,
	req dto.MstClassScheduleGetByRequest,
) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	data := converter.ConvertMstClassScheduleGetByRequestToCommand(req)

	err := s.mstClassScheduleRepository.DeleteByDayTime(tx, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service":  "delete class schedule service",
			"deleteby": user.ID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

func (s *MstClassScheduleService) DeleteByScheduleTemplate(
	ctx context.Context,
	classScheduletemplateID string,
) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	err := s.mstClassScheduleRepository.DeleteByScheduleTemplate(tx, classScheduletemplateID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service":  "delete class schedule service",
			"deleteby": user.ID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

func (s *MstClassScheduleService) GetAllWithCountByLecturerID(
	ctx context.Context, pageble pageable.PageableRequestClassScheduleLecturer,
) (*pageable.PageableResponse[dto.MstClassResponseForSchedule], error) {
	tx := s.db.WithContext(ctx)

	Classes, totalData, err := s.mstClassScheduleRepository.GetAllWithCountByLecturerID(tx, ctx, pageble)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	res := make([]dto.MstClassResponseForSchedule, 0, totalData)
	for _, class := range Classes {
		res = append(res, converter.MstClassModelToResponseForSchedule(class))
	}

	return &pageable.PageableResponse[dto.MstClassResponseForSchedule]{
		Data: res,
		Metadata: pageable.Metadata{
			TotalData: totalData,
			TotalPage: utils.TotalPage(totalData, pageble.GetDefaultLimit()),
			Page:      pageble.GetDefaultPage(),
			Size:      pageble.GetDefaultLimit(),
		},
	}, nil
}

// TODO: Restore

// TODO: Read
func (s *MstClassScheduleService) GetByClassID(ctx context.Context, classID string) ([]dto.MstClassScheduleResponse, error) {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)

	datas, err := s.mstClassScheduleRepository.GetByClassID(tx, classID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service":  "get by class_id class schedule service",
			"class_id": classID,
			"user-id":  user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return nil, err
	}
	res := make([]dto.MstClassScheduleResponse, 0, len(datas))
	for _, v := range datas {
		res = append(res, converter.ConvertModelNoPointerToMstClassScheduleResponse(v))
	}
	return res, err
}

func (s *MstClassScheduleService) GetByLecturerIDandActiveAcademicPeriod(ctx context.Context) (res []dto.MstClassScheduleAcademicSystemDistributionResponse, err error) {
	tx := s.db.WithContext(ctx)

	classes, err := s.mstClassScheduleRepository.GetByLecturerIDandActiveAcademicPeriod(tx, ctx)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	results := make([]dto.MstClassScheduleAcademicSystemDistributionResponse, 0)
	for _, class := range classes {
		results = append(results, *converter.MstClassScheduleAcademicSystemDistributionToResponse(class))
	}

	return results, nil
}

func (s *MstClassScheduleService) GetByID(ctx context.Context, ID string) (*dto.MstClassScheduleResponse, error) {
	tx := s.db.WithContext(ctx)

	classSchedule := new(model.MstClassSchedule)

	err := s.mstClassScheduleRepository.GetByID(tx, ID, classSchedule)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	result := converter.ConvertModelToMstClassScheduleResponsePointer(classSchedule)

	return result, nil
}

func (s *MstClassScheduleService) GetByIDForPresence(
	ctx context.Context, ID string,
) (*dto.MstClassScheduleForClassSessionPresenceResponse, error) {
	tx := s.db.WithContext(ctx)

	classSchedule := new(model.MstClassScheduleForClassSessionPresence)

	err := s.mstClassScheduleRepository.GetByIDForPresence(tx, ID, classSchedule)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	result := converter.ConvertModelToClassSessionPresenceResponse(classSchedule)

	return &result, nil
}

func (s *MstClassScheduleService) GetByDayTime(
	ctx context.Context, req dto.MstClassScheduleGetByRequest,
) (res *dto.MstClassScheduleResponse, err error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	reqCommand := converter.ConvertMstClassScheduleGetByRequestToCommand(req)

	data := new(model.MstClassSchedule)
	err = s.mstClassScheduleRepository.GetByDayTime(tx, reqCommand, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get by class id class schedule service",
			"request": reqCommand,
			"user-id": user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return res, err
	}
	res = converter.ConvertModelToMstClassScheduleResponsePointer(data)
	return res, err
}

func (s *MstClassScheduleService) GetByClassAsDate(
	ctx context.Context,
	req pageable.PageableRequestClassParticipant,
) (res *pageable.PageableResponse[dto.MstClassScheduleResponse], err error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	datas, count, err := s.mstClassScheduleRepository.GetByClassAsDate(
		tx, true, req,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get by class as date in class schedule service",
			"user-id": user.ID,
			"req":     req,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return res, err
	}

	r := make([]dto.MstClassScheduleResponse, 0, count)

	for _, v := range datas {
		c := converter.ConvertModelNoPointerToMstClassScheduleResponse(v)
		r = append(r, c)
	}
	res = &pageable.PageableResponse[dto.MstClassScheduleResponse]{
		Data: r,
		Metadata: pageable.Metadata{
			TotalData: count,
			TotalPage: utils.TotalPage(count, req.GetDefaultLimit()),
			Page:      req.GetDefaultPage(),
			Size:      req.GetDefaultLimit(),
		},
	}
	return res, err
}

func (s *MstClassScheduleService) GenerateByAcademic(
	ctx context.Context,
	req dto.GenerateScheduleRequest,
) error {
	return s.generate(ctx, req)
}

func (s *MstClassScheduleService) GenerateByProgramHead(
	ctx context.Context,
	req dto.GenerateScheduleRequest,
) error {
	user := middleware.GetUserClaimsCtx(ctx)
	tx := s.db.WithContext(ctx)

	_, _, err := s.checkStudyProgramUseProgramHead_GeneralInformation(tx, user.ID, req.ClassID.String(), "Generate")
	if err != nil {
		return err
	}
	return s.generate(ctx, req)
}

// Generate
func (s *MstClassScheduleService) generate(
	ctx context.Context,
	req dto.GenerateScheduleRequest,
) error {
	tx := s.db.WithContext(ctx)
	user := middleware.GetUserClaimsCtx(ctx)

	// get schedule template
	classScheduleTemplateData := new(model.MstClassScheduleTemplate)
	err := s.mstClassScheduleTemplateRepository.GetByClassID(tx, req.ClassID.String(), classScheduleTemplateData)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get class template by id in class schedule service",
			"user-id": user.ID,
			"req":     req,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return err
	}

	classData := new(model.MstClass)
	err = s.mstClassRepository.GetByID(tx, req.ClassID.String(), classData)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get class by id in class schedule service",
			"user-id": user.ID,
			"req":     req,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return err
	}

	// check academic period data
	academicPeriodData, err := s.academicPeriodgatewayRest.GetDataWithParamsValue(
		nil, nil,
		classData.AcademicPeriodeID,
	)
	if err != nil {
		return err
	}
	err = utilicems.CheckRestError(
		s.log, academicPeriodData, user.ID,
		"generate class schedule", fmt.Sprintf("get academicPeriodData %v from datareferensi in class schedule service", classData.AcademicPeriodeID),
	)
	if err != nil {
		return err
	}

	if academicPeriodData.Data.StartDateOfCollege == "" {
		s.log.WithFields(logrus.Fields{
			"service":            "academic not set startDate in class schedule service",
			"user-id":            user.ID,
			"academicPeriodData": academicPeriodData,
		}).Error(err)
		return fiber.NewError(fiber.StatusFailedDependency, "academicPerio not set StartDateOfCollege")
	}

	tot, err := s.mstClassScheduleRepository.CountModelWhere(tx, "schedule_template_id = ?", classScheduleTemplateData.ID.String())
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "count class schedule by template id  in class schedule service",
			"user-id": user.ID,
			"req":     req,
		}).Error(err)
		return err
	}
	if tot >= int64(academicPeriodData.Data.UasSession) {
		err = s.mstClassScheduleRepository.DeleteByScheduleTemplate(tx, classScheduleTemplateData.ID.String())
		if err != nil {
			s.log.WithFields(logrus.Fields{
				"service": "delete by template id  in class schedule service",
				"user-id": user.ID,
				"req":     req,
			}).Error(err)
			return err
		}
	}

	startDate := academicPeriodData.Data.StartDateOfCollege
	datePeriode, err := time.Parse(time.RFC3339, startDate)
	if err != nil {
		return err
	}
	countSession := academicPeriodData.Data.NumberOfLectureMeeting

	countSessionInt := utilicems.StringToInt(countSession)

	// generate date
	set := rrule.Set{}
	// generate real session
	r, _ := rrule.NewRRule(rrule.ROption{
		Freq: rrule.WEEKLY,
		// Byweekday: []rrule.Weekday{wd},
		Count:   countSessionInt,
		Dtstart: datePeriode,
	})
	set.RRule(r)
	sessionReal := set.All()

	// generate fake for lecturer
	dayMap := map[string]rrule.Weekday{
		"monday":    rrule.MO,
		"tuesday":   rrule.TU,
		"wednesday": rrule.WE,
		"thursday":  rrule.TH,
		"friday":    rrule.FR,
		"saturday":  rrule.SA,
		"sunday":    rrule.SU,
	}
	dayName := strings.ToLower(classScheduleTemplateData.DayName)
	wd, ok := dayMap[dayName]
	if !ok {
		logrus.Warnf("Invalid day name: %s", classScheduleTemplateData.DayName)
		// wd = nil, atau kasih fallback
		wd = rrule.MO
	}
	r2, _ := rrule.NewRRule(rrule.ROption{
		Freq:      rrule.WEEKLY,
		Byweekday: []rrule.Weekday{wd},
		Count:     countSessionInt,
		Dtstart:   datePeriode,
	})
	set.RRule(r2)

	// var generateSchedule func(
	// 	se int,
	// 	error chan error,
	// )
	//
	// insert class schedule
	for i, v := range set.All() {
		var dateSessionReal time.Time
		se := i + 1

		if int64(se) < tot {
			continue
		}

		req := new(model.MstClassSchedule)
		req.ID = utilicems.GenerateUUID()
		req.ClassID = classData.ID
		req.ScheduleTemplateID = classScheduleTemplateData.ID
		req.Session = se
		req.DayName = classScheduleTemplateData.DayName
		req.StartTime = classScheduleTemplateData.StartTime
		req.EndTime = classScheduleTemplateData.EndTime
		req.Date = v
		req.TypeOfMeeting = classScheduleTemplateData.TypeOfMeeting

		if se == academicPeriodData.Data.UtsSession {
			dateSessionReal = sessionReal[i]
			req.Date = dateSessionReal
			req.DayName = utilicems.DayEN[int(dateSessionReal.Weekday())]
			req.IsUTS = true
		}
		if se == academicPeriodData.Data.UasSession {
			dateSessionReal = sessionReal[i]
			req.Date = dateSessionReal
			req.DayName = utilicems.DayEN[int(dateSessionReal.Weekday())]
			req.IsUAS = true
		}
		err := s.mstClassScheduleRepository.Create(tx, req)
		if err != nil {
			s.log.WithFields(logrus.Fields{
				"service":                   "create failed class schedule service",
				"user-id":                   user.ID,
				"classScheduleTemplateData": classScheduleTemplateData,
			}).Error(err)
			_ = s.mstClassScheduleRepository.DeleteByScheduleTemplate(tx, classScheduleTemplateData.ID.String())
			return fiber.NewError(fiber.StatusInternalServerError, "classSchedule failed generate")
		}
	}

	return nil
}

// helper function in this service
func (s *MstClassScheduleService) checkStudyProgramUseProgramHead_GeneralInformation(
	tx *gorm.DB, userID string, classID string, funcAction string,
) (*model.MstClass, *modelsdm.GeneralInformationResponse, error) {
	generalInformation, err := s.generalInformationGatewayRest.GetDataWithParamsOrQuery(
		nil, nil, "/"+userID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class schedule service " + funcAction + " by programhead",
			"user-id": userID,
		}).Error(err)
		return nil, nil, fiber.NewError(fiber.StatusInternalServerError, "get data sdm error "+err.Error())
	}

	err = utilicems.CheckRestError(
		s.log, generalInformation, userID,
		funcAction+" class schedule by programhead", "get general information from sdm in class schedule service "+funcAction+" by program head",
	)
	if err != nil {
		return nil, nil, err
	}

	studyProgramID := generalInformation.Data.StudyProgramID
	if studyProgramID == nil {
		s.log.WithFields(logrus.Fields{
			"service":        "get general information from sdm in class schedule service " + funcAction + "ByProgramHead",
			"user-id":        userID,
			"studyProgramID": "studyProgramID not found",
		}).Info(err)
		return nil, nil, fiber.NewError(fiber.StatusNotFound, "studyProgramID in sdm not found wit user -> "+userID)
	}

	classData := new(model.MstClass)
	err = s.mstClassRepository.GetByID(tx, classID, classData)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get class by id in " + funcAction + "by programhead class schedulue service",
			"user-id": userID,
		}).Error(err)
		return nil, nil, utilicems.ErrorSpToErrorFiber(err)
	}

	if classData.StudyProgramID != studyProgramID.String() && classData.StudyProgramID != "0" {
		s.log.WithFields(logrus.Fields{
			"service": "get general information from sdm in class schedule service " + funcAction + "ByProgramHead",
			"user-id": userID,
		}).Info(funcAction + " class schedule by program head not access")
		return nil, nil, fiber.NewError(fiber.StatusForbidden, "ProgramHead not access data")
	}

	return classData, generalInformation.Data, nil
}

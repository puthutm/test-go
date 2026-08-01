// Package servicemodel
package servicemodel

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	icemstutil "gitlab.unsia.ac.id/icems/icems-tools/utils"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/delivery/http/middleware"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/dto/converter"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository/cached"
	repositorymodel "unsia.ac.id/akademic_be/internal/repository/model"
	"unsia.ac.id/akademic_be/pkg/auth"
	"unsia.ac.id/akademic_be/pkg/utils"
)

type TrxStudentPresenceSettingService struct {
	db                                  *gorm.DB
	log                                 *logrus.Logger
	cache                               cached.CacheRepository
	trxStudentPresenceSettingRepository *repositorymodel.TrxStudentPresenceSettingRepository
}

func NewTrxStudentPresenceSettingService(
	db *gorm.DB,
	log *logrus.Logger,
	cache cached.CacheRepository,
	trxStudentPresenceSettingRepository *repositorymodel.TrxStudentPresenceSettingRepository,
) *TrxStudentPresenceSettingService {
	return &TrxStudentPresenceSettingService{
		db:                                  db,
		log:                                 log,
		cache:                               cache,
		trxStudentPresenceSettingRepository: trxStudentPresenceSettingRepository,
	}
}

func (s *TrxStudentPresenceSettingService) CreateOrUpdate(
	ctx context.Context, req dto.TrxStudentPresenceSettingCreateOrUpdateRequest,
) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	// model
	now := time.Now().UnixMilli()
	userID, _ := icemstutil.StringToUuid(user.ID)
	studentPresenceSettingM := &model.TrxStudentPresenceSetting{
		ID:                icemstutil.GenerateUUID(),
		AcademicPeriodeID: req.AcademicPeriodeID,
		StudyProgramID:    req.StudyProgramID,
		SubjectID:         req.SubjectID,

		Presence: model.Presence{
			UseOpenSession:        req.UseOpenSession,
			OpenSessionPercentage: req.OpenSessionPercentage,

			UseDocumentMaterial:        req.UseDocumentMaterial,
			DocumentMaterialPercentage: req.DocumentMaterialPercentage,

			UseQuiz:        req.UseQuiz,
			QuizPercentage: req.QuizPercentage,

			UseTask:        req.UseTask,
			TaskPercentage: req.TaskPercentage,

			UseVideo:        req.UseVideo,
			VideoPercentage: req.VideoPercentage,

			UseUTS:        req.UseUTS,
			UTSPercentage: req.UTSPercentage,

			UseUAS:        req.UseUAS,
			UASPercentage: req.UASPercentage,

			UseComment:        req.UseComment,
			CommentPercentage: req.CommentPercentage,
		},
		CreatedAt: &now,
		CreatedBy: &userID,
	}

	err := s.trxStudentPresenceSettingRepository.CreateOrUpdate(ctx, tx, studentPresenceSettingM)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithFields(logrus.Fields{
			"service": "create studentPresenceSetting",
			"req":     req,
		}).Error(createMsg)

		return icemstutil.ErrorSpToErrorFiber(err)
	}
	return nil
}

func (s *TrxStudentPresenceSettingService) createOrUpdateStudentPresenceRoot(
	ctx context.Context,
	tx *gorm.DB, user *auth.UserClaimsSpesifikRole,
	req dto.TrxStudentPresenceSaveParamBySessionRequest,
) error {
	// model
	now := time.Now().UnixMilli()
	studentPresence := &model.TrxStudentPresenceSaveParamBySession{
		IDNew:          icemstutil.GenerateUUID(),
		SessionID:      req.SessionID,
		StudentID:      req.StudentID,
		PresenceStatus: req.PresenceStatus,
		PresenceType:   req.PresenceType,
		CreatedAt:      now,
	}

	err := s.trxStudentPresenceSettingRepository.CreateOrUpdateStudentPresence(ctx, tx, studentPresence)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithFields(logrus.Fields{
			"service": "save studentPresence",
			"req":     req,
			"user":    user,
		}).Error(createMsg)

		return icemstutil.ErrorSpToErrorFiber(err)
	}
	return nil
}

func (s *TrxStudentPresenceSettingService) CreateOrUpdateStudentPresence(
	ctx context.Context, req dto.TrxStudentPresenceSaveParamBySessionRequest,
) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	// model
	return s.createOrUpdateStudentPresenceRoot(
		ctx, tx, user, req,
	)
}

func (s *TrxStudentPresenceSettingService) CreateOrUpdateStudentPresenceSlice(
	ctx context.Context, req dto.TrxStudentPresenceSliceSaveParamBySessionRequest,
) error {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	// model
	for _, v := range req.PresenceParams {
		re := dto.TrxStudentPresenceSaveParamBySessionRequest{
			SessionID:     req.SessionID,
			StudentID:     req.StudentID,
			PresenceParam: v,
		}
		if err := s.createOrUpdateStudentPresenceRoot(ctx, tx, user, re); err != nil {
			return err
		}
	}
	return nil
}

func (s *TrxStudentPresenceSettingService) GetPresenceComponentForLecturer(
	ctx context.Context, req dto.TrxStudentPresenceGetForLecturerRequest,
) (*dto.PresenceRes, error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)
	param := model.TrxStudentPresenceGetForLecturerParam{
		UserID:            user.ID,
		AcademicPeriodeID: req.AcademicPeriodeID,
		StudyProgramID:    req.StudyProgramID,
		SubjectID:         req.SubjectID,
	}

	var presence model.Presence

	err := s.trxStudentPresenceSettingRepository.GetPresenceComponentForLecturer(tx, param, &presence)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithFields(logrus.Fields{
			"service": "get studentPresenceSetting for lecturer",
			"req":     req,
		}).Error(createMsg)

		return nil, icemstutil.ErrorSpToErrorFiber(err)
	}

	result := converter.ConvertPresenceToResponse(presence)

	return result, nil
}

func (s *TrxStudentPresenceSettingService) GetPresenceComponentBySession(
	ctx context.Context, sessionID string,
) (*dto.PresenceRes, error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	var presence model.Presence

	err := s.trxStudentPresenceSettingRepository.GetPresenceComponentBySessionID(tx, sessionID, &presence)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithFields(logrus.Fields{
			"service":    "get studentPresenceSetting for lecturer",
			"session_id": sessionID,
			"user":       user,
		}).Error(createMsg)

		return nil, icemstutil.ErrorSpToErrorFiber(err)
	}

	result := converter.ConvertPresenceToResponse(presence)

	return result, nil
}

func (s *TrxStudentPresenceSettingService) GetSessionPresenceByClassID(
	ctx context.Context, classID string,
) ([]dto.SessionPresenceResponse, error) {
	tx := s.db.WithContext(ctx)

	sessions, err := s.trxStudentPresenceSettingRepository.GetSessionPresenceByClassID(ctx, tx, classID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "GetSessionPresenceByClassID",
			"classID": classID,
		}).Error(err.Error())
		return nil, icemstutil.ErrorSpToErrorFiber(err)
	}

	var resp []dto.SessionPresenceResponse
	for _, s := range sessions {
		resp = append(resp, dto.SessionPresenceResponse{
			SessionID:          s.SessionID,
			Session:            s.Session,
			SessionDate:        s.SessionDate,
			PresencePercentage: s.PresencePercentage,
		})
	}

	return resp, nil
}

func (s *TrxStudentPresenceSettingService) GetStudentPresenceBySessionWithCount(
	ctx context.Context, pageble pageable.PageableStudentPresenceBySession,
) (*pageable.PageableResponse[dto.TrxStudentPresenceBySessionResponse], error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	datas, count, err := s.trxStudentPresenceSettingRepository.GetStudentPresenceBySessionWithCount(
		ctx, tx, pageble,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "GetStudentPresenceBySessionWithCount",
			"user-id": user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return nil, err
	}

	r := make([]dto.TrxStudentPresenceBySessionResponse, 0, count)

	for _, v := range datas {
		c := converter.ConvertTrxStudentPresenceBySessionToResponse(v)
		r = append(r, c)
	}
	res := &pageable.PageableResponse[dto.TrxStudentPresenceBySessionResponse]{
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

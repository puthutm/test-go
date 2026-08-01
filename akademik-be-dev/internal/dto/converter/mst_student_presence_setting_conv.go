// Package converter
package converter

import (
	"time"

	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/pkg/utils"
)

func ConvertMstStudentPresenceSettingRequestToMstStudentPresenceSettingModel(
	req dto.MstStudentPresenceSettingCreateRequest, userID string,
) *model.MstStudentPresenceSetting {
	now := time.Now().UnixMilli()

	userIDR, err := utils.StringToUuid(userID)
	if err != nil {
		panic(err)
	}

	var (
		uasPercentage int
		utsPercentage int
	)

	if req.UseUAS {
		uasPercentage = 100
	}
	if req.UseUTS {
		utsPercentage = 100
	}

	return &model.MstStudentPresenceSetting{
		ID:                    utils.GenerateUUID(),
		AcademicPeriodeID:     req.AcademicPeriodeID,
		StudyProgramID:        req.StudyProgramID,
		UseOpenSession:        req.UseOpenSession,
		OpenSessionPercentage: req.OpenSessionPercentage,

		UseDocumentMaterial:     req.UseDocumentMaterial,
		DocumentMaterialPercent: req.DocumentMaterialPercent,

		UseQuiz:        req.UseQuiz,
		QuizPercentage: req.QuizPercentage,

		UseTask:        req.UseTask,
		TaskPercentage: req.TaskPercentage,

		UseVideo:        req.UseVideo,
		VideoPercentage: req.VideoPercentage,

		UseUTS:        req.UseUTS,
		UTSPercentage: utsPercentage,

		UseUAS:        req.UseUAS,
		UASPercentage: uasPercentage,

		UseComment:        req.UseComment,
		CommentPercentage: req.CommentPercentage,

		CreatedAt: &now,
		CreatedBy: &userIDR,
	}
}

func ConvertMstStudentPresenceSettingDuplicateRequestToParam(
	req dto.MstStudentPresenceSettingDuplicateRequest, userID string,
) *model.MstStudentPresenceSettingDuplicateParam {
	now := time.Now().UnixMilli()

	userIDR, err := utils.StringToUuid(userID)
	if err != nil {
		panic(err) // Disarankan handle dengan lebih baik di production
	}

	return &model.MstStudentPresenceSettingDuplicateParam{
		ID:                   utils.GenerateUUID(),
		AcademicPeriodeID:    req.AcademicPeriodeID,
		StudyProgramID:       req.StudyProgramID,
		AcademicPeriodeIDOld: req.AcademicPeriodeIDTarget,
		CreatedAt:            now,
		CreatedBy:            userIDR,
	}
}

func ConvertMstStudentPresenceSettingGetResultToResponse(
	data *model.MstStudentPresenceSettingGetResult,
) dto.MstStudentPresenceSettingGetResultResponse {
	return dto.MstStudentPresenceSettingGetResultResponse{
		MstStudentPresenceSettingGetResponse: dto.MstStudentPresenceSettingGetResponse{
			ID:                data.ID,
			AcademicPeriodeID: data.AcademicPeriodeID,
			StudyProgramID:    data.StudyProgramID,
		},
		AcademicPeriodeName: data.AcademicPeriodeName,
		StudyProgramName:    data.StudyProgramName,
	}
}

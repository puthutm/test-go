package converter

import (
	"time"

	"github.com/google/uuid"
	utilicems "unsia.ac.id/akademic_be/pkg/icems-tools/utils"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/pkg/utils"
)

func ConvertMstClassScheduleTemplateCreateRequestToModel(request dto.MstClassScheduleTemplateCreateRequest) *model.MstClassScheduleTemplate {
	return &model.MstClassScheduleTemplate{
		ID:            utils.GenerateUUID(),
		ClassID:       request.ClassID,
		DayName:       request.DayName,
		StartTime:     request.StartTime,
		EndTime:       request.EndTime,
		TypeOfMeeting: request.TypeOfMeeting,
		CreatedAt:     time.Now().UnixMilli(),
	}
}

func ConvertMstClassScheduleTemplateUpdateRequestToModel(request dto.MstClassScheduleTemplateUpdateRequest) *model.MstClassScheduleTemplate {
	nowT := time.Now().UnixMilli()
	return &model.MstClassScheduleTemplate{
		ID:            request.ID,
		ClassID:       request.ClassID,
		DayName:       request.DayName,
		StartTime:     request.StartTime,
		EndTime:       request.EndTime,
		TypeOfMeeting: request.TypeOfMeeting,
		UpdatedAt:     &nowT,
	}
}

func ConvertModelToMstClassSchedulePointerTemplateResponsePointer(model *model.MstClassScheduleTemplate) *dto.MstClassScheduleTemplateResponse {
	if model == nil {
		return nil
	}
	startTimeNew := utilicems.StringDateTimeToStringTime(model.StartTime)
	endTimeNew := utilicems.StringDateTimeToStringTime(model.EndTime)

	return &dto.MstClassScheduleTemplateResponse{
		ID:            model.ID,
		ClassID:       model.ClassID,
		DayName:       model.DayName,
		StartTime:     startTimeNew,
		EndTime:       endTimeNew,
		TypeOfMeeting: model.TypeOfMeeting,
		CreatedAt:     model.CreatedAt,
		CreatedBy:     model.CreatedBy,
		UpdatedAt:     model.UpdatedAt,
		UpdatedBy:     model.UpdatedBy,
	}
}

func ConvertModelToMstClassSchedulePointerTemplateResponse(model *model.MstClassScheduleTemplate) dto.MstClassScheduleTemplateResponse {
	if model == nil {
		return dto.MstClassScheduleTemplateResponse{}
	}
	return *ConvertModelToMstClassSchedulePointerTemplateResponsePointer(model)
}

func ConvertModelToMstClassScheduleTemplateResponse(model model.MstClassScheduleTemplate) dto.MstClassScheduleTemplateResponse {
	if model.ID == uuid.Nil {
		return dto.MstClassScheduleTemplateResponse{}
	}
	return *ConvertModelToMstClassSchedulePointerTemplateResponsePointer(&model)
}

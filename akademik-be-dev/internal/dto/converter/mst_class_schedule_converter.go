package converter

import (
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	utilicems "unsia.ac.id/akademic_be/pkg/icems-tools/utils"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/service/command"
	"unsia.ac.id/akademic_be/pkg/utils"
)

// ConvertMstClassScheduleRequestToModelPointer converts request DTO to model pointer
func ConvertMstClassScheduleRequestToModelPointer(
	request dto.MstClassScheduleRequest, model *model.MstClassSchedule,
) error {
	session, err := strconv.Atoi(request.Session)
	if err != nil {
		return fmt.Errorf("invalid session format: %w", err)
	}

	status, err := strconv.Atoi(request.Status)
	if err != nil {
		return fmt.Errorf("invalid status format: %w", err)
	}

	date, err := time.Parse("2006-01-02", request.Date)
	if err != nil {
		return fmt.Errorf("invalid date format (expected YYYY-MM-DD): %w", err)
	}

	model.ID = utils.GenerateUUID()
	model.ClassID = request.ClassID
	model.Session = session
	model.DayName = request.DayName
	model.Date = date
	model.StartTime = request.StartTime
	model.EndTime = request.EndTime
	model.TypeOfMeeting = request.TypeOfMeeting
	model.Status = status

	return nil
}

// ConvertModelToMstClassScheduleResponsePointer converts model to response DTO pointer
func ConvertModelToMstClassScheduleResponsePointer(model *model.MstClassSchedule) *dto.MstClassScheduleResponse {
	var ID *uuid.UUID
	if model.ID != uuid.Nil {
		ID = &model.ID
	}
	var ClassID *uuid.UUID
	if model.ClassID != uuid.Nil {
		ClassID = &model.ClassID
	}
	startTimeNew := utilicems.StringDateTimeToStringTime(model.StartTime)
	endTimeNew := utilicems.StringDateTimeToStringTime(model.EndTime)

	return &dto.MstClassScheduleResponse{
		ID:                         ID,
		ClassID:                    ClassID,
		ScheduleTemplateID:         model.ScheduleTemplateID,
		Session:                    model.Session,
		DayName:                    model.DayName,
		Date:                       model.Date,
		StartTime:                  startTimeNew,
		EndTime:                    endTimeNew,
		TypeOfMeeting:              model.TypeOfMeeting,
		Status:                     model.Status,
		MaterialAttachmentFilePath: model.MaterialAttachmentFilePath,
		AttendanceDocumentFilePath: model.AttendanceDocumentFilePath,
		JournalDocumentFilePath:    model.JournalDocumentFilePath,
		MaterialPlan:               model.MaterialPlan,
		MaterialRealization:        model.MaterialRealization,
		IsUTS:                      model.IsUTS,
		IsUAS:                      model.IsUAS,
	}
}

// ConvertModelToMstClassScheduleResponse converts model to response DTO
func ConvertModelToMstClassScheduleResponse(model *model.MstClassSchedule) dto.MstClassScheduleResponse {
	return *ConvertModelToMstClassScheduleResponsePointer(model)
}

func ConvertModelNoPointerToMstClassScheduleResponse(model model.MstClassSchedule) dto.MstClassScheduleResponse {
	return *ConvertModelToMstClassScheduleResponsePointer(&model)
}

func ConvertModelToScheduleResponseForListByClass(model model.MstClassSchedule) dto.MstClassScheduleResponseForListByClass {
	return dto.MstClassScheduleResponseForListByClass{
		ID:                 model.ID,
		ScheduleTemplateID: model.ScheduleTemplateID,
		Session:            model.Session,
		DayName:            model.DayName,
		Date:               model.Date,
		StartTime:          model.StartTime,
		EndTime:            model.EndTime,
		TypeOfMeeting:      model.TypeOfMeeting,
		Status:             model.Status,
		IsUTS:              model.IsUTS,
		IsUAS:              model.IsUAS,
	}
}

func ConvertModelToClassSessionPresenceResponse(model *model.MstClassScheduleForClassSessionPresence) dto.MstClassScheduleForClassSessionPresenceResponse {
	return dto.MstClassScheduleForClassSessionPresenceResponse{
		SessionID:                 model.SessionID,
		SessionName:               model.SessionName,
		SessionDate:               model.SessionDate,
		ClassName:                 model.ClassName,
		ClassCode:                 model.ClassCode,
		ScheduleTemplateDayName:   model.ScheduleTemplateDayName,
		ScheduleTemplateStartTime: model.ScheduleTemplateStartTime,
		ScheduleTemplateEndTime:   model.ScheduleTemplateEndTime,
		AcademicPeriodName:        model.AcademicPeriodName,
		TotalParticipant:          model.TotalParticipant,
		StudentPresenceCount:      model.StudentPresenceCount,
		StudentAbsentCount:        model.StudentAbsentCount,
	}
}

// Command
// ConvertMstClassScheduleUpdateRequestToCommand converts update request DTO to command
func ConvertMstClassScheduleUpdateRequestToCommand(
	request dto.MstClassScheduleUpdateRequest,
) command.MstClassScheduleUpdateCommand {
	return command.MstClassScheduleUpdateCommand{
		ID:            request.ID,
		ClassID:       request.ClassID,
		DayNameOld:    request.DayNameOld,
		StartTimeOld:  request.StartTimeOld,
		EndTimeOld:    request.EndTimeOld,
		DayName:       request.DayName,
		Date:          request.Date,
		StartTime:     request.StartTime,
		EndTime:       request.EndTime,
		TypeOfMeeting: request.TypeOfMeeting,
	}
}

// ConvertMstClassScheduleGetByRequestToCommand converts get-by request DTO to command
func ConvertMstClassScheduleGetByRequestToCommand(
	request dto.MstClassScheduleGetByRequest,
) command.MstClassScheduleGetByCommand {
	return command.MstClassScheduleGetByCommand{
		ID:        request.ID,
		ClassID:   request.ClassID,
		DayName:   request.DayName,
		StartTime: request.StartTime,
		EndTime:   request.EndTime,
	}
}

func ConvertMstClassScheduleUpdateRequestToCommandPointer(request dto.MstClassScheduleUpdateRequest, cmd *command.MstClassScheduleUpdateCommand) {
	cmd.ID = request.ID
	cmd.ClassID = request.ClassID
	cmd.DayNameOld = request.DayNameOld
	cmd.StartTimeOld = request.StartTimeOld
	cmd.EndTimeOld = request.EndTimeOld
	cmd.DayName = request.DayName
	cmd.Date = request.Date
	cmd.StartTime = request.StartTime
	cmd.EndTime = request.EndTime
	cmd.TypeOfMeeting = request.TypeOfMeeting
}

func ConvertMstClassScheduleGetByRequestToCommandPointer(request dto.MstClassScheduleGetByRequest, cmd *command.MstClassScheduleGetByCommand) {
	cmd.ID = request.ID
	cmd.ClassID = request.ClassID
	cmd.DayName = request.DayName
	cmd.StartTime = request.StartTime
	cmd.EndTime = request.EndTime
}

func MstClassScheduleAcademicSystemDistributionToResponse(model model.MstClassSchedule) *dto.MstClassScheduleAcademicSystemDistributionResponse {
	return &dto.MstClassScheduleAcademicSystemDistributionResponse{
		TypeOfMeeting: model.TypeOfMeeting,
	}
}

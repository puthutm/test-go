package converter

import (
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
)

func MstClassScheduleTaskToResponse(model model.MstClassScheduleTask) *dto.MstClassScheduleTaskResponse {
	return &dto.MstClassScheduleTaskResponse{
		ID:                             model.ID,
		ScheduleID:                     model.ScheduleID,
		Title:                          model.Title,
		Description:                    model.Description,
		IsGradeable:                    model.IsGradeable,
		IsUseDeadline:                  model.IsUseDeadline,
		DeadlineOfAssignmentSubmission: model.DeadlineOfAssignmentSubmission,
		IsSharing:                      model.IsSharing,
		SharingDate:                    model.SharingDate,
		TimeToOpen:                     model.TimeToOpen,
		Retake:                         model.Retake,
		Views:                          model.Views,
		CreatedAt:                      model.CreatedAt,
		UpdatedAt:                      model.UpdatedAt,
		DeletedAt:                      model.DeletedAt,
		SessionSchedule:                model.SessionSchedule,
		TotalCollect:                   model.TotalCollect,
	}
}

package converter

import (
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/pkg/utils"
)

func ConvertMstClassParticipantRequestToModelPointer(request dto.MstClassParticipantRequest, model *model.MstClassParticipant) {
	model.ID = utils.GenerateUUID()
	model.ClassID = request.ClassID
	model.StudentID = request.StudentID
}

func ConvertModelToMstClassParticipantResponsePointer(model *model.MstClassParticipant) *dto.MstClassParticipantResponse {
	return &dto.MstClassParticipantResponse{
		ID:          model.ID,
		ClassID:     model.ClassID,
		StudentID:   model.StudentID,
		StudentNIM:  model.StudentNIM,
		StudentName: model.StudentName,
	}
}

func ConvertModelPointerToMstClassParticipantResponse(model *model.MstClassParticipant) dto.MstClassParticipantResponse {
	return dto.MstClassParticipantResponse{
		ID:          model.ID,
		ClassID:     model.ClassID,
		StudentID:   model.StudentID,
		StudentNIM:  model.StudentNIM,
		StudentName: model.StudentName,
	}
}

func ConvertModelToMstClassParticipantResponse(model model.MstClassParticipant) dto.MstClassParticipantResponse {
	return dto.MstClassParticipantResponse{
		ID:               model.ID,
		ClassID:          model.ClassID,
		StudentID:        model.StudentID,
		StudentNIM:       model.StudentNIM,
		StudentName:      model.StudentName,
		StudyProgramName: model.StudyProgramName,
		YearOfEntry:      model.YearOfEntry,
	}
}

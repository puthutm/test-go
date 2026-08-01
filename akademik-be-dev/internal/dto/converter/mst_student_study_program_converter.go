package converter

import (
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
)

func MstStudentStudyProgramModelToResponse(model model.MstStudentStudyProgram) *dto.MstStudentStudyProgramResponse {
	return &dto.MstStudentStudyProgramResponse{
		ID:                      model.ID,
		StudentID:               model.StudentID,
		EnrollmentBatchDetailID: model.EnrollmentBatchDetailID,
		NIM:                     model.NIM,
		StudyProgramID:          model.StudyProgramID,
		Status:                  model.Status,
		YearOfEntry:             model.YearOfEntry,
		CreatedAt:               model.CreatedAt,
		UpdatedAt:               model.UpdatedAt,
		DeletedAt:               model.DeletedAt,
	}
}

func MstStudentStudyProgramModelByProgramHeadToResponse(model model.MstStudentStudyProgram) *dto.MstStudentStudyProgramByProgramHeadResponse {
	return &dto.MstStudentStudyProgramByProgramHeadResponse{
		StudentID:        model.StudentID,
		StudentNIM:       model.StudentNIM,
		StudentName:      model.StudentName,
		StudyProgramName: model.StudyProgramName,
		StudentStatus:    model.StudentStatus,
	}
}

func MstStudentStudyProgramSearchToResponse(model model.MstStudentStudyProgram) dto.MstStudentStudyProgramSearchResponse {
	return dto.MstStudentStudyProgramSearchResponse{
		StudentID:   model.StudentID,
		StudentNIM:  model.StudentNIM,
		StudentName: model.StudentName,
	}
}

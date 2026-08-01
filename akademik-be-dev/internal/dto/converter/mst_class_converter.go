package converter

import (
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
)

func MstClassModelToResponse(model model.MstClass) *dto.MstClassResponse {
	return &dto.MstClassResponse{
		ID:                  model.ID,
		Code:                model.Code,
		Name:                model.Name,
		AcademicPeriodeID:   model.AcademicPeriodeID,
		SubjectID:           model.SubjectID,
		StudyProgramID:      model.StudyProgramID,
		CurriculumYearID:    model.CurriculumYearID,
		Capacity:            model.Capacity,
		NumberOfMeeting:     model.NumberOfMeeting,
		ContractDescription: model.ContractDescription,
		ContractFilePath:    model.ContractFilePath,
		CreatedAt:           model.CreatedAt,
		CreatedBy:           model.CreatedBy,
		UpdatedAt:           model.UpdatedAt,
		UpdatedBy:           model.UpdatedBy,
		DeletedAt:           model.DeletedAt,
		TotalParticipant:    model.TotalParticipant,
		SubjectNameID:       model.SubjectNameID,
		SubjectNameEN:       model.SubjectNameEN,
		LecturerctName:      model.LecturerName,

		StudyProgramName: model.StudyProgramName,

		LecturerSystem:          model.LecturerSystem,
		SubjectTotalSks:         model.SubjectTotalSks,
		CurriculumYearName:      model.CurriculumYearName,
		StartDateOfCollege:      model.StartDateOfCollege,
		EndDateOfCollege:        model.EndDateOfCollege,
		AcademicPeriodeFullname: model.AcademicPeriodeFullname,
	}
}

func MstClassStudentDistributionToResponse(model model.MstClass) *dto.MstClassStudentDistributionResponse {
	return &dto.MstClassStudentDistributionResponse{
		ClassID:   model.ID,
		ClassCode: model.Code,
		ClassName: model.Name,
	}
}

func MstClassModelToResponseForSchedule(model model.MstClass) dto.MstClassResponseForSchedule {
	return dto.MstClassResponseForSchedule{
		ID:               model.ID,
		Code:             model.Code,
		Name:             model.Name,
		Capacity:         model.Capacity,
		TotalParticipant: model.TotalParticipant,
		SubjectNameID:    model.SubjectNameID,
		SubjectNameEN:    model.SubjectNameEN,
		LecturerctName:   model.LecturerName,

		StudyProgramName:   model.StudyProgramName,
		CurriculumYearName: model.CurriculumYearName,
		DayName:            model.DayName,
		StartTime:          model.StartTime,
		EndTime:            model.EndTime,
	}
}

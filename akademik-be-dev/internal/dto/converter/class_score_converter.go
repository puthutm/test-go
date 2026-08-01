package converter

import (
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
)

func ClassScoreModelToResponse(data model.ClassScore) *dto.ClassScoreResponse {
	return &dto.ClassScoreResponse{
		StudentID:        data.StudentID,
		Nim:              data.Nim,
		StudentName:      data.StudentName,
		PresenceScore:    data.PresenceScore,
		TaskScore:        data.TaskScore,
		UtsScore:         data.UtsScore,
		UasScore:         data.UasScore,
		FinalScore:       data.FinalScore,
		QualityValue:     data.QualityValue,
		GradeID:          data.GradeID,
		GradeCode:        data.GradeCode,
		GradeName:        data.GradeName,
		GradeDescription: data.GradeDescription,
		IsPassed:         data.IsPassed,
		PassNote:         data.PassNote,
		LimitGradeID:     data.LimitGradeID,
		LimitGradeCode:   data.LimitGradeCode,
		LimitGradeLower:  data.LimitGradeLower,
	}
}

func ClassScoreSummaryModelToResponse(data model.ClassScoreSummary) *dto.ClassScoreSummaryResponse {
	return &dto.ClassScoreSummaryResponse{
		TotalStudents:     data.TotalStudents,
		TotalPassed:       data.TotalPassed,
		TotalNotPassed:    data.TotalNotPassed,
		AverageFinalScore: data.AverageFinalScore,
		LimitGradeCode:    data.LimitGradeCode,
		LimitGradeLower:   data.LimitGradeLower,
	}
}

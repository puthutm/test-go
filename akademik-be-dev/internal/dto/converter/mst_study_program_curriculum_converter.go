package converter

import (
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
)

func MstStudyProgramCurriculumModelToResponse(model model.MstStudyProgramCurriculum) *dto.MstStudyProgramCurriculumResponse {
	return &dto.MstStudyProgramCurriculumResponse{
		ID:                        model.ID,
		SubjectCode:               model.SubjectCode,
		CurriculumYearID:          model.CurriculumYearID,
		SubjectID:                 model.SubjectID,
		SemesterNumberID:          model.SemesterNumberID,
		LimitGradeID:              model.LimitGradeID,
		StudyProgramID:            model.StudyProgramID,
		IsMandatory:               model.IsMandatory,
		IsPackage:                 model.IsPackage,
		FieldStudyConcentrationID: model.FieldStudyConcentrationID,
		// CreatedAt:        model.CreatedAt,
		// CreatedBy:        model.CreatedBy,
		// UpdatedAt:        model.UpdatedAt,
		// UpdatedBy:        model.UpdatedBy,
		// DeletedAt:        model.DeletedAt,
		// DeletedBy:        model.DeletedBy,

		SubjectNameID:               model.SubjectNameID,
		SubjectNameEN:               model.SubjectNameEN,
		SubjectTotalSKS:             model.SubjectTotalSKS,
		GradeName:                   model.GradeName,
		GradeCode:                   model.GradeCode,
		FieldStudyConcentrationName: model.FieldStudyConcentrationName,
		FieldStudyConcentrationCode: model.FieldStudyConcentrationCode,
	}
}

func MstStudyProgramCurriculumModelToResponseDetail(model model.MstStudyProgramCurriculum) *dto.MstStudyProgramCurriculumResponseDetail {
	return &dto.MstStudyProgramCurriculumResponseDetail{
		ID:                        model.ID,
		SubjectCode:               model.SubjectCode,
		CurriculumYearID:          model.CurriculumYearID,
		SubjectID:                 model.SubjectID,
		SemesterNumberID:          model.SemesterNumberID,
		LimitGradeID:              model.LimitGradeID,
		IsPackage:                 model.IsPackage,
		StudyProgramID:            model.StudyProgramID,
		IsMandatory:               model.IsMandatory,
		CreatedAt:                 model.CreatedAt,
		CreatedBy:                 model.CreatedBy,
		UpdatedAt:                 model.UpdatedAt,
		UpdatedBy:                 model.UpdatedBy,
		DeletedAt:                 model.DeletedAt,
		DeletedBy:                 model.DeletedBy,
		FieldStudyConcentrationID: model.FieldStudyConcentrationID,

		SubjectNameID:               model.SubjectNameID,
		SubjectNameEN:               model.SubjectNameEN,
		SubjectTotalSKS:             model.SubjectTotalSKS,
		GradeName:                   model.GradeName,
		GradeCode:                   model.GradeCode,
		FieldStudyConcentrationName: model.FieldStudyConcentrationName,
		FieldStudyConcentrationCode: model.FieldStudyConcentrationCode,
	}
}

package converter

import (
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/pkg/utils"
)

func ConvertMstSubjectRequestToModelPointer(request dto.MstSubjectRequest, model *model.MstSubject) {
	model.ID = utils.GenerateUUID()
	model.CurriculumYearID = request.CurriculumYearID
	model.StudyProgramID = request.StudyProgramID
	model.CourseTypeID = request.CourseTypeID
	model.CourseGroupID = request.CourseGroupID
	model.Code = request.Code
	model.NameID = request.NameID
	model.NameEN = request.NameEN
	model.FaceToFaceSKS = request.FaceToFaceSKS
	model.PracticumSKS = request.PracticumSKS
	model.FieldPracticeSKS = request.FieldPracticeSKS
	model.SimulationSKS = request.SimulationSKS
	model.TotalSKS = request.TotalSKS
	model.FieldOfStudiesID = request.FieldOfStudiesID
	model.IsMKU = request.IsMKU
	model.IsSAP = request.IsSAP
	model.IsSilabus = request.IsSilabus
	model.IsTeachingMaterial = request.IsTeachingMaterial
	model.IsDiktat = request.IsDiktat
	model.CreatedBy = &request.UserID
}

func ConvertMstSubjectUpdateToModelPointer(request dto.MstSubjectUpdate, model *model.MstSubject) {
	model.ID = request.ID
	model.CurriculumYearID = request.CurriculumYearID
	model.StudyProgramID = request.StudyProgramID
	model.CourseTypeID = request.CourseTypeID
	model.CourseGroupID = request.CourseGroupID
	model.Code = request.Code
	model.NameID = request.NameID
	model.NameEN = request.NameEN
	model.FaceToFaceSKS = request.FaceToFaceSKS
	model.PracticumSKS = request.PracticumSKS
	model.FieldPracticeSKS = request.FieldPracticeSKS
	model.SimulationSKS = request.SimulationSKS
	model.TotalSKS = request.TotalSKS
	model.FieldOfStudiesID = request.FieldOfStudiesID
	model.IsMKU = request.IsMKU
	model.IsSAP = request.IsSAP
	model.IsSilabus = request.IsSilabus
	model.IsTeachingMaterial = request.IsTeachingMaterial
	model.IsDiktat = request.IsDiktat
	model.UpdatedBy = &request.UserID
}

func ConvertModelToMstSubjectResponse(model *model.MstSubject) dto.MstSubjectResponse {
	return dto.MstSubjectResponse{
		ID:                 model.ID,
		CurriculumYearID:   model.CurriculumYearID,
		StudyProgramID:     model.StudyProgramID,
		Code:               model.Code,
		NameID:             model.NameID,
		NameEN:             model.NameEN,
		CourseTypeID:       model.CourseTypeID,
		CourseGroupID:      model.CourseGroupID,
		FaceToFaceSKS:      model.FaceToFaceSKS,
		PracticumSKS:       model.PracticumSKS,
		FieldPracticeSKS:   model.FieldPracticeSKS,
		SimulationSKS:      model.SimulationSKS,
		TotalSKS:           model.TotalSKS,
		FieldOfStudiesID:   model.FieldOfStudiesID,
		IsMKU:              model.IsMKU,
		IsSAP:              model.IsSAP,
		IsSilabus:          model.IsSilabus,
		IsTeachingMaterial: model.IsTeachingMaterial,
		IsDiktat:           model.IsDiktat,
		CreatedAt:          model.CreatedAt,
		UpdatedAt:          model.UpdatedAt,
		DeletedAt:          model.DeletedAt,
		StudyProgramName:   model.StudyProgramName,
		CurriculumYearName: model.CurriculumYearName,
		CourseTypeName:     model.CourseTypeName,
		CourseGroupName:    model.CourseGroupName,
		FieldStudyName:     model.FieldStudyName,
	}
}

func ConvertModelToMstSubjectResponsePointer(model *model.MstSubject) *dto.MstSubjectResponse {
	return &dto.MstSubjectResponse{
		ID:                 model.ID,
		CurriculumYearID:   model.CurriculumYearID,
		StudyProgramID:     model.StudyProgramID,
		Code:               model.Code,
		NameID:             model.NameID,
		NameEN:             model.NameEN,
		CourseTypeID:       model.CourseTypeID,
		CourseGroupID:      model.CourseGroupID,
		FaceToFaceSKS:      model.FaceToFaceSKS,
		PracticumSKS:       model.PracticumSKS,
		FieldPracticeSKS:   model.FieldPracticeSKS,
		SimulationSKS:      model.SimulationSKS,
		TotalSKS:           model.TotalSKS,
		FieldOfStudiesID:   model.FieldOfStudiesID,
		IsMKU:              model.IsMKU,
		IsSAP:              model.IsSAP,
		IsSilabus:          model.IsSilabus,
		IsTeachingMaterial: model.IsTeachingMaterial,
		IsDiktat:           model.IsDiktat,
		CreatedAt:          model.CreatedAt,
		UpdatedAt:          model.UpdatedAt,
		DeletedAt:          model.DeletedAt,
		StudyProgramName:   model.StudyProgramName,
		CurriculumYearName: model.CurriculumYearName,
		CourseTypeName:     model.CourseTypeName,
		CourseGroupName:    model.CourseGroupName,
		FieldStudyName:     model.FieldStudyName,
	}
}

// for search
func ConvertModelToMstSubjectResponseForSearch(model model.MstSubject) dto.MstSubjectResponseForSearch {
	return dto.MstSubjectResponseForSearch{
		ID:     model.ID,
		Code:   model.Code,
		NameID: model.NameID,
		NameEN: model.NameEN,
	}
}

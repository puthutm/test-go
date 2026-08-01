package converter

import (
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/pkg/utils"
)

func ConvertMstValueScaleRequestToModelPointer(request dto.MstValueScaleRequest, model *model.MstValueScale) {
	model.ID = utils.GenerateUUID()
	model.StudyProgramID = request.StudyProgramID
	model.GradeID = request.GradeID
	model.WeightValue = request.WeightValue
	model.LowerValue = request.LowerValue
	model.UpperValue = request.UpperValue
	model.Description = request.Description
	model.CreatedBy = request.UserID
}

func ConvertMstValueScaleUpdateToModelPointer(request dto.MstValueScaleUpdate, model *model.MstValueScale) {
	model.ID = request.ID
	model.StudyProgramID = request.StudyProgramID
	model.GradeID = request.GradeID
	model.WeightValue = request.WeightValue
	model.LowerValue = request.LowerValue
	model.UpperValue = request.UpperValue
	model.Description = request.Description
	model.UpdatedBy = request.UserID
}

func ConvertModelToMstValueScaleResponsePointer(model *model.MstValueScale) *dto.MstValueScaleResponse {
	return &dto.MstValueScaleResponse{
		ID:               model.ID,
		StudyProgramID:   model.StudyProgramID,
		GradeID:          model.GradeID,
		WeightValue:      model.WeightValue,
		LowerValue:       model.LowerValue,
		UpperValue:       model.UpperValue,
		Description:      model.Description,
		CreatedAt:        model.CreatedAt,
		CreatedBy:        model.CreatedBy,
		UpdatedAt:        model.UpdatedAt,
		UpdatedBy:        model.UpdatedBy,
		DeletedAt:        model.DeletedAt,
		DeletedBy:        model.DeletedBy,
		StudyProgramName: model.StudyProgramName,
		GradeName:        model.GradeName,
	}
}

func ConvertModelToMstValueScaleResponse(model *model.MstValueScale) dto.MstValueScaleResponse {
	return dto.MstValueScaleResponse{
		ID:               model.ID,
		StudyProgramID:   model.StudyProgramID,
		GradeID:          model.GradeID,
		WeightValue:      model.WeightValue,
		LowerValue:       model.LowerValue,
		UpperValue:       model.UpperValue,
		Description:      model.Description,
		CreatedAt:        model.CreatedAt,
		CreatedBy:        model.CreatedBy,
		UpdatedAt:        model.UpdatedAt,
		UpdatedBy:        model.UpdatedBy,
		DeletedAt:        model.DeletedAt,
		DeletedBy:        model.DeletedBy,
		StudyProgramName: model.StudyProgramName,
		GradeName:        model.GradeName,
	}
}

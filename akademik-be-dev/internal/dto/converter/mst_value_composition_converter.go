package converter

import (
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/pkg/utils"
)

func ConvertMstValueCompositionRequestToModelPointer(request dto.MstValueCompositionRequest, model *model.MstValueComposition) {
	model.ID = utils.GenerateUUID()
	model.ValueElementID = request.ValueElementID
	model.AcademicPeriodeID = request.AcademicPeriodeID
	model.Percentage = request.Percentage
	model.IsPassingRequirement = *request.IsPassingRequirement
	model.CreatedBy = &request.UserID
}

func ConvertMstValueCompositionUpdateToModelPointer(request dto.MstValueCompositionUpdate, model *model.MstValueComposition) {
	model.ID = request.ID
	model.ValueElementID = request.ValueElementID
	model.AcademicPeriodeID = request.AcademicPeriodeID
	model.Percentage = request.Percentage
	model.IsPassingRequirement = *request.IsPassingRequirement
	model.UpdatedBy = &request.UserID
}

func ConvertModelToMstValueCompositionResponse(model *model.MstValueComposition) dto.MstValueCompositionResponse {
	return dto.MstValueCompositionResponse{
		ID:                   model.ID,
		ValueElementID:       model.ValueElementID,
		AcademicPeriodeID:    model.AcademicPeriodeID,
		Percentage:           model.Percentage,
		IsPassingRequirement: model.IsPassingRequirement,
		CreatedAt:            model.CreatedAt,
		CreatedBy:            model.CreatedBy,
		UpdatedAt:            model.UpdatedAt,
		UpdatedBy:            model.UpdatedBy,
		DeletedAt:            model.DeletedAt,
		DeletedBy:            model.DeletedBy,
		ValueElementName:     model.ValueElementName,
		AcademicPeriodeName:  model.AcademicPeriodeName,
	}
}

func ConvertModelToMstValueCompositionResponsePointer(model *model.MstValueComposition) *dto.MstValueCompositionResponse {
	return &dto.MstValueCompositionResponse{
		ID:                   model.ID,
		ValueElementID:       model.ValueElementID,
		AcademicPeriodeID:    model.AcademicPeriodeID,
		Percentage:           model.Percentage,
		IsPassingRequirement: model.IsPassingRequirement,
		CreatedAt:            model.CreatedAt,
		CreatedBy:            model.CreatedBy,
		UpdatedAt:            model.UpdatedAt,
		UpdatedBy:            model.UpdatedBy,
		DeletedAt:            model.DeletedAt,
		DeletedBy:            model.DeletedBy,
		ValueElementName:     model.ValueElementName,
		AcademicPeriodeName:  model.AcademicPeriodeName,
	}
}

func ConvertModelToValueCompositionGroupResponse(model *model.ValueCompositionGroup) dto.ValueCompositionGroupResponse {
	return dto.ValueCompositionGroupResponse{
		AcademicPeriodeID:   model.AcademicPeriodeID,
		AcademicPeriodeName: model.AcademicPeriodeName,
		TotalComponents:     model.TotalComponents,
		TotalPercentage:     model.TotalPercentage,
		IsValid:             model.IsValid == 1,
	}
}

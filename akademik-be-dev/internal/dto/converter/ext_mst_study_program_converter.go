package converter

import (
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
)

func MstStudyProgramModelToResponse(model model.MstStudyProgram) *dto.MstStudyProgramResponse {
	return &dto.MstStudyProgramResponse{
		ID:        model.ID,
		Code:      model.Code,
		Name:      model.Name,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		DeletedAt: model.DeletedAt,
	}
}
func MstStudyProgramModelToDistributionOfStudyProgramResponse(model model.MstStudyProgram) *dto.DistributionOfStudyProgramResponse {
	return &dto.DistributionOfStudyProgramResponse{
		ID:   model.ID,
		Code: model.Code,
		Name: model.Name,
	}
}

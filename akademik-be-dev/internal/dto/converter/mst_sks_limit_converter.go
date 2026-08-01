package converter

import (
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
)

func MstSKSLimitModelToResponse(model model.MstSKSLimit) *dto.MstSKSLimitResponse {
	return &dto.MstSKSLimitResponse{
		ID:        model.ID,
		IPSMin:    model.IPSMin,
		IPSMax:    model.IPSMax,
		SKSLimit:  model.SKSLimit,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		DeletedAt: model.DeletedAt,
	}
}

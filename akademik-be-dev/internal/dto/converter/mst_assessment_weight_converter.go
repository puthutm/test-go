package converter

import (
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
)

func MstAssessmentWeightModelToResponse(m *model.MstAssessmentWeight) *dto.MstAssessmentWeightResponse {
	return &dto.MstAssessmentWeightResponse{
		ID:                         m.ID,
		AttitudeBehaviorPercentage: m.AttitudeBehaviorPercentage,
		TaskPercentage:             m.TaskPercentage,
		UTSPercentage:              m.UTSPercentage,
		UASPercentage:              m.UASPercentage,
	}
}

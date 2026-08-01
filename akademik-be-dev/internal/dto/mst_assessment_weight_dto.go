package dto

import "github.com/google/uuid"

/* Request */
type MstAssessmentWeightRequest struct {
	ID                         uuid.UUID `json:"-"`
	AttitudeBehaviorPercentage float64   `json:"attitude_behavior_percentage" validate:"required,numeric"`
	TaskPercentage             float64   `json:"task_percentage" validate:"required,numeric"`
	UTSPercentage              float64   `json:"uts_percentage" validate:"required,numeric"`
	UASPercentage              float64   `json:"uas_percentage" validate:"required,numeric"`
}

/* Response */
type MstAssessmentWeightResponse struct {
	ID                         uuid.UUID `json:"id"`
	AttitudeBehaviorPercentage float64   `json:"attitude_behavior_percentage"`
	TaskPercentage             float64   `json:"task_percentage"`
	UTSPercentage              float64   `json:"uts_percentage"`
	UASPercentage              float64   `json:"uas_percentage"`
}

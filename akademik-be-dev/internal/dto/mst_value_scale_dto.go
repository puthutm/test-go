package dto

import "github.com/google/uuid"

type MstValueScaleRequest struct {
	StudyProgramID uuid.UUID  `json:"study_program_id" validate:"required,uuid"`
	GradeID        uuid.UUID  `json:"grade_id" validate:"required,uuid"`
	WeightValue    float64    `json:"weight_value" validate:"required,gt=0,lt=100"`
	LowerValue     float64    `json:"lower_value" validate:"required"`
	UpperValue     float64    `json:"upper_value" validate:"required"`
	Description    *string    `json:"description" validate:"required,stringMax=1000"`
	UserID         *uuid.UUID `json:"-"`
}

type MstValueScaleUpdate struct {
	ID uuid.UUID `json:"-" validate:"omitempty"`
	MstValueScaleRequest
}

type MstValueScaleResponse struct {
	ID             uuid.UUID  `json:"id"`
	StudyProgramID uuid.UUID  `json:"study_program_id"`
	GradeID        uuid.UUID  `json:"grade_id"`
	WeightValue    float64    `json:"weight_value"`
	LowerValue     float64    `json:"lower_value"`
	UpperValue     float64    `json:"upper_value"`
	Description    *string    `json:"description"`
	CreatedAt      int64      `json:"created_at"`
	CreatedBy      *uuid.UUID `json:"created_by"`
	UpdatedAt      *int64     `json:"updated_at"`
	UpdatedBy      *uuid.UUID `json:"updated_by"`
	DeletedAt      *int64     `json:"deleted_at"`
	DeletedBy      *uuid.UUID `json:"deleted_by"`

	// Additional fields for joined data
	StudyProgramName string `json:"study_program_name"`
	GradeName        string `json:"grade_name"`
}

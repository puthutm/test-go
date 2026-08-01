package dto

import "github.com/google/uuid"

type MstValueCompositionRequest struct {
	ValueElementID    uuid.UUID `json:"value_element_id" validate:"required,uuid"`
	AcademicPeriodeID uuid.UUID `json:"academic_periode_id" validate:"required,uuid"`

	Percentage           float64   `json:"percentage" validate:"required,gt=0,lt=100"`
	IsPassingRequirement *bool     `json:"is_passing_requirement" validate:"required"`
	UserID               uuid.UUID `json:"-"`
}

type MstValueCompositionUpdate struct {
	ID                uuid.UUID `json:"-" validate:"required,uuid"`
	ValueElementID    uuid.UUID `json:"value_element_id" validate:"required,uuid"`
	AcademicPeriodeID uuid.UUID `json:"academic_periode_id" validate:"required,uuid"`

	Percentage           float64   `json:"percentage" validate:"required,gt=0,lt=100"`
	IsPassingRequirement *bool     `json:"is_passing_requirement" validate:"required"`
	UserID               uuid.UUID `json:"-"`
}

type MstValueCompositionResponse struct {
	ID                   uuid.UUID  `json:"id"`
	ValueElementID       uuid.UUID  `json:"value_element_id"`
	AcademicPeriodeID    uuid.UUID  `json:"academic_periode_id"`
	Percentage           float64    `json:"percentage"`
	IsPassingRequirement bool       `json:"is_passing_requirement"`
	CreatedAt            int64      `json:"created_at"`
	CreatedBy            *uuid.UUID `json:"created_by"`
	UpdatedAt            *int64     `json:"updated_at"`
	UpdatedBy            *uuid.UUID `json:"updated_by"`
	DeletedAt            *int64     `json:"deleted_at"`
	DeletedBy            *uuid.UUID `json:"deleted_by"`
	ValueElementName     string     `json:"value_element_name"`
	AcademicPeriodeName  string     `json:"academic_periode_name"`
}

type ValueCompositionGroupResponse struct {
	AcademicPeriodeID   string  `json:"academic_periode_id"`
	AcademicPeriodeName string  `json:"academic_periode_name"`
	TotalComponents     int     `json:"total_components"`
	TotalPercentage     float64 `json:"total_percentage"`
	IsValid             bool    `json:"is_valid"`
}

type MstValueCompositionDuplicateRequest struct {
	AcademicPeriodIDSource uuid.UUID `json:"academic_period_id_source" validate:"required,uuid"`
	AcademicPeriodIDTarget uuid.UUID `json:"academic_period_id_target" validate:"required,uuid"`
	IsOverwrite            *bool     `json:"is_overwrite" validate:"required"`
}

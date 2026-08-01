package dto

import (
	"time"

	"github.com/google/uuid"
)

/* Request */
type MstStudentFamilyRequest struct {
	ID               string  `json:"-"`
	StudentID        string  `json:"-"`
	Name             string  `json:"name" validate:"required,stringMax=255"`
	NIK              string  `json:"nik" validate:"required,stringMax=16"`
	EducationLevelID *string `json:"education_level_id" validate:"omitempty,uuid"`
	Type             string  `json:"type" validate:"required,parent"`
	Phone            *string `json:"phone" validate:"omitempty,valid_number_phone"`
	Phone2           *string `json:"phone2" validate:"omitempty,valid_number_phone"`
	Email            *string `json:"email" validate:"omitempty,email"`
	Kinship          *string `json:"kinship" validate:"omitempty"`
	StatusKinship    *string `json:"status_kinship" validate:"omitempty"`
	LifeStatus       *string `json:"life_status" validate:"omitempty"`
	Address          *string `json:"address" validate:"omitempty"`
	BirthPlaceID     *string `json:"birth_place_id" validate:"omitempty,uuid"`
	BirthDate        *string `json:"birth_date" validate:"omitempty,datetime=2006-01-02"`
	JobID            *string `json:"job_id" validate:"omitempty,uuid"`
	Income           *string `json:"income" validate:"omitempty,numeric"`
}

/* Response */
type MstStudentFamilyResponse struct {
	ID                   uuid.UUID  `json:"id"`
	StudentID            uuid.UUID  `json:"student_id"`
	Name                 string     `json:"name"`
	NIK                  string     `json:"nik"`
	EducationalLevelID   *uuid.UUID `json:"educational_level_id"`
	EducationalLevelName *string    `json:"educational_level_name"`
	Type                 *string    `json:"type"`
	Phone                *string    `json:"phone"`
	Phone2               *string    `json:"phone2"`
	Email                *string    `json:"email"`
	Kinship              *string    `json:"kinship"`
	StatusKinship        *string    `json:"status_kinship"`
	LifeStatus           *string    `json:"life_status"`
	Address              *string    `json:"address"`
	BirthPlaceID         *uuid.UUID `json:"birth_place_id"`
	BirthPlaceName       *string    `json:"birth_place_name"`
	BirthDate            *time.Time `json:"birth_date"`
	JobID                *uuid.UUID `json:"job_id"`
	JobName              *string    `json:"job_name"`
	Income               *float64   `json:"income"`
	CreatedAt            int64      `json:"created_at"`
	UpdatedAt            *int64     `json:"updated_at"`
}

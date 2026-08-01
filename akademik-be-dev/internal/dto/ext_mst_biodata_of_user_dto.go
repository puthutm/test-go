package dto

import (
	"time"

	"github.com/google/uuid"
)

/* Request */
/* Response */
type MstBiodataOfUserResponse struct {
	ID               uuid.UUID  `json:"id"`
	NipNim           string     `json:"nip_nim"`
	FrontTitle       *string    `json:"front_title"`
	BackTitle        *string    `json:"back_title"`
	BirthPlace       *string    `json:"birth_place"`
	BirthDate        *time.Time `json:"birth_date"`
	Gender           string     `json:"gender"`
	MarriageStatusID uuid.UUID  `json:"marriage_status_id"`
	ReligionID       uuid.UUID  `json:"religion_id"`
	EthnicID         uuid.UUID  `json:"ethnic_id"`
	Weight           float64    `json:"weight"`
	Height           float64    `json:"height"`
	BloodTypeID      *uuid.UUID `json:"blood_type_id"`
	CountryID        *uuid.UUID `json:"country_id"`
	Status           *string    `json:"status"`
	UserID           uuid.UUID  `json:"user_id"`
	CreatedAt        *int64     `json:"created_at"`
	UpdatedAt        *int64     `json:"updated_at"`
	DeletedAt        *int64     `json:"deleted_at"`
}

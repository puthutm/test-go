package dto

import "github.com/google/uuid"

/* Request */
/* Response */
type MstStudyProgramResponse struct {
	ID        uuid.UUID `json:"id"`
	Code      string    `json:"code"`
	Name      *string   `json:"name"`
	CreatedAt int64     `json:"created_at"`
	UpdatedAt *int64    `json:"updated_at"`
	DeletedAt *int64    `json:"deleted_at"`
}

type DistributionOfStudyProgramResponse struct {
	ID   uuid.UUID `json:"id"`
	Code string    `json:"code"`
	Name *string   `json:"name"`
}

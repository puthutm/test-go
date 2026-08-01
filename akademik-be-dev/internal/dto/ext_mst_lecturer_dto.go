package dto

import "github.com/google/uuid"

/* Request */
/* Response */
type MstLecturerByProgramHeadResponse struct {
	ID               uuid.UUID `json:"id"`
	NIP              *string   `json:"nip"`
	Name             *string   `json:"name"`
	NIDN             *string   `json:"nidn"`
	Gender           *string   `json:"gender"`
	Phone            *string   `json:"phone"`
	Email            *string   `json:"email"`
	Status           *string   `json:"status"`
	StudyProgramName *string   `json:"study_program_name"`
}

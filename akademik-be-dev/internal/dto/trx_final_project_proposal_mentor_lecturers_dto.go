package dto

import "github.com/google/uuid"

type TrxFinalProjectProposalMentorLecturerRequest struct {
	FinalProjectProposalID uuid.UUID  `json:"final_project_proposal_id" validate:"required"`
	LecturerID             uuid.UUID  `json:"lecturer_id" validate:"required"`
	AssignDate             int64      `json:"assign_date" validate:"required"`
	AssignBy               *uuid.UUID `json:"assign_by,omitempty"`
}

type TrxFinalProjectProposalMentorLecturerResponse struct {
	ID                     uuid.UUID  `json:"id"`
	FinalProjectProposalID uuid.UUID  `json:"final_project_proposal_id"`
	LecturerID             uuid.UUID  `json:"lecturer_id"`
	AssignDate             int64      `json:"assign_date"`
	AssignBy               *uuid.UUID `json:"assign_by,omitempty"`

	LecturerName string `json:"lecturer_name,omitempty"`
}

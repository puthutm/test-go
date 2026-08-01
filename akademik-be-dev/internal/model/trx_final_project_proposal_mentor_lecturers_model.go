package model

import (
	"github.com/google/uuid"
)

type TrxFinalProjectProposalMentorLecturer struct {
	ID                     uuid.UUID  `gorm:"type:char(36);primaryKey"`
	FinalProjectProposalID uuid.UUID  `gorm:"type:char(36);not null"`
	LecturerID             uuid.UUID  `gorm:"type:char(36);not null"`
	AssignDate             int64      `gorm:"not null" ` // Unix timestamp
	AssignBy               *uuid.UUID `gorm:"type:char(36)" `

	LecturerName string `gorm:"column:lecturer_name"`
}

func (TrxFinalProjectProposalMentorLecturer) TableName() string {
	return "trx_final_project_proposal_mentor_lecturers"
}

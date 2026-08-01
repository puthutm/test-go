package dto

import "github.com/google/uuid"

/* Request */
type MstClassParticipantRequest struct {
	ID        uuid.UUID `json:"-"`
	ClassID   uuid.UUID `json:"-"`
	StudentID uuid.UUID `json:"student_id" validate:"required,uuid"`
}

/* Response */
type MstClassParticipantResponse struct {
	ID        uuid.UUID `json:"id" gorm:"column:id"`
	ClassID   uuid.UUID `json:"class_id" gorm:"column:class_id"`
	StudentID uuid.UUID `json:"student_id" gorm:"column:student_id"`

	// relation
	StudentNIM       string `json:"student_nim"`
	StudentName      string `json:"student_name"`
	StudyProgramName string `json:"study_program_name"`
	YearOfEntry      string `json:"year_of_entry"`
}

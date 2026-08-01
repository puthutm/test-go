package model

import (
	"github.com/google/uuid"
)

type MstClassParticipant struct {
	ID        uuid.UUID `gorm:"type:char(36);column:id;primaryKey"`
	ClassID   uuid.UUID `gorm:"type:char(36);column:class_id"`
	StudentID uuid.UUID `gorm:"type:char(36);column:student_id"`

	// Additional Fields
	StudentNIM         string `gorm:"column:student_nim"`
	StudentName        string `gorm:"column:student_name"`
	StudentProgramName string `gorm:"column:student_program_name"`
	StudyProgramName   string `gorm:"column:study_program_name"`
	YearOfEntry        string `gorm:"column:year_of_entry"`
}

func (MstClassParticipant) TableName() string {
	return "mst_class_participants"
}

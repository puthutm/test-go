package model

import "github.com/google/uuid"

type MstStudyProgram struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	Code      string    `gorm:"type:varchar(5)"`
	Name      *string   `gorm:"type:nvarchar(255);not null"`
	CreatedAt int64     `gorm:"type:bigint"`
	UpdatedAt *int64    `gorm:"type:bigint"`
	DeletedAt *int64    `gorm:"type:bigint"`

	StudentNIM       *string `gorm:"column:student_nim"`
	StudentName      *string `gorm:"column:student_name"`
	StudyProgramName *string `gorm:"column:study_program_name"`
	StudentStatus    *string `gorm:"column:student_status"`
}

package model

import (
	"github.com/google/uuid"
)

type MstStudentStudyProgram struct {
	ID                      uuid.UUID `gorm:"column:id;type:char(36);primaryKey"`
	StudentID               uuid.UUID `gorm:"column:student_id;type:char(36);not null"`
	EnrollmentBatchDetailID uuid.UUID `gorm:"column:enrollment_batch_detail_id;type:char(36);not null"`
	NIM                     *string   `gorm:"column:nim;type:varchar(50)"`
	StudyProgramID          uuid.UUID `gorm:"column:study_program_id;type:char(36);not null"`
	Status                  string    `gorm:"column:status;type:nvarchar(225);not null"`
	YearOfEntry             *string   `gorm:"column:year_of_entry;type:varchar(4)"`
	CreatedAt               int64     `gorm:"column:created_at;type:bigint"`
	UpdatedAt               *int64    `gorm:"column:updated_at;type:bigint"`
	DeletedAt               *int64    `gorm:"column:deleted_at;type:bigint"`

	// Additional Fields
	StudentNIM       *string `gorm:"column:student_nim"`
	StudentName      *string `gorm:"column:student_name"`
	StudyProgramName *string `gorm:"column:study_program_name"`
	StudentStatus    *string `gorm:"column:student_status"`
}

func (MstStudentStudyProgram) TableName() string {
	return "mst_student_study_programs"
}

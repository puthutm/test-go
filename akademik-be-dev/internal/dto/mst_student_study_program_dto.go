package dto

import "github.com/google/uuid"

/* Request */
type MstStudentStudyProgramRequest struct {
	ID                      string  `json:"-"`
	StudentID               string  `json:"student_id" validate:"required,uuid"`
	EnrollmentBatchDetailID string  `json:"enrollment_batch_detail_id" validate:"required,uuid"`
	NIM                     *string `json:"nim" validate:"omitempty,stringMax=50"`
	StudyProgramID          string  `json:"study_program_id" validate:"required,uuid"`
	Status                  string  `json:"status" validate:"required,stringMax=225"`
	YearOfEntry             *string `json:"year_of_entry" validate:"omitempty,numeric,stringMax=4"`
}

/* Response */
type MstStudentStudyProgramResponse struct {
	ID                      uuid.UUID `json:"id" gorm:"column:id"`
	StudentID               uuid.UUID `json:"student_id" gorm:"column:student_id;not null"`
	EnrollmentBatchDetailID uuid.UUID `json:"enrollment_batch_detail_id" gorm:"column:enrollment_batch_detail_id;not null"`
	NIM                     *string   `json:"nim" gorm:"column:nim"`
	StudyProgramID          uuid.UUID `json:"study_program_id" gorm:"column:study_program_id;not null"`
	Status                  string    `json:"status" gorm:"column:status;not null"`
	YearOfEntry             *string   `json:"year_of_entry" gorm:"column:year_of_entry"`
	CreatedAt               int64     `json:"created_at" gorm:"column:created_at"`
	UpdatedAt               *int64    `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt               *int64    `json:"deleted_at" gorm:"column:deleted_at"`
}

type MstStudentStudyProgramByProgramHeadResponse struct {
	StudentID        uuid.UUID `json:"student_id" gorm:"column:student_id;not null"`
	StudentNIM       *string   `json:"student_nim" gorm:"column:student_nim"`
	StudentName      *string   `json:"student_name" gorm:"column:student_name"`
	StudyProgramName *string   `json:"study_program_name" gorm:"column:study_program_name"`
	StudentStatus    *string   `json:"student_status" gorm:"column:student_status"`
}

type MstStudentStudyProgramSearchResponse struct {
	StudentID   uuid.UUID `json:"student_id"`
	StudentNIM  *string   `json:"student_nim"`
	StudentName *string   `json:"student_name"`
}

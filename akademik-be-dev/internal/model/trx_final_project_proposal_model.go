package model

import "github.com/google/uuid"

type TrxFinalProjectProposal struct {
	ID                     uuid.UUID  `gorm:"type:char(36);column:id;not null"`
	StudentID              uuid.UUID  `gorm:"type:char(36);column:student_id;not null"`
	TitleID                string     `gorm:"type:nvarchar(255);column:title_id;not null"`
	TitleEN                string     `gorm:"type:nvarchar(255);column:title_en;not null"`
	Topic                  string     `gorm:"type:nvarchar(255);column:topic;not null"`
	AcademicPeriodeID      uuid.UUID  `gorm:"type:char(36);column:academic_periode_id;not null"`
	StudyProgramID         uuid.UUID  `gorm:"type:char(36);column:study_program_id;not null"`
	Abstract               string     `gorm:"type:nvarchar(max);column:abstract;not null"`
	FilePath               string     `gorm:"type:varchar(max);column:file_path;not null"`
	Status                 int        `gorm:"type:int;column:status;default:0;not null"`
	Date                   *int64     `gorm:"type:bigint;column:date"`
	ConfirmationStatusDate *int64     `gorm:"type:bigint;column:confirmation_status_date"`
	ConfirmationBy         *uuid.UUID `gorm:"type:char(36);column:confirmation_by"`
	Feedback               *string    `gorm:"type:nvarchar(max);column:feedback"`
	CreatedAt              *int64     `gorm:"type:bigint;column:created_at"`
	UpdatedAt              *int64     `gorm:"type:bigint;column:updated_at"`
	DeletedAt              *int64     `gorm:"type:bigint;column:deleted_at"`

	StudentName          *string `gorm:"column:student_name"`
	StudentNIM           *string `gorm:"column:student_nim"`
	StudyProgramName     *string `gorm:"column:study_program_name"`
	AcademicPeriodeName  *string `gorm:"column:academic_periode_name"`
	RegistrationPathName *string `gorm:"column:registration_path_name"`
	LecturerName         *string `gorm:"column:lecturer_name"`
}

func (TrxFinalProjectProposal) TableName() string {
	return "trx_final_project_proposals"
}

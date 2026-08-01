package model

import (
	"github.com/google/uuid"
)

type MstClass struct {
	ID                  uuid.UUID  `gorm:"type:char(36);column:id;primaryKey"`
	Code                string     `gorm:"type:varchar(50);column:code"`
	Name                string     `gorm:"type:nvarchar(225);column:name"`
	AcademicPeriodeID   uuid.UUID  `gorm:"type:char(36);column:academic_periode_id"`
	SubjectID           uuid.UUID  `gorm:"type:char(36);column:subject_id"`
	StudyProgramID      string     `gorm:"column:study_program_id"`
	CurriculumYearID    uuid.UUID  `gorm:"type:char(36);column:curriculum_year_id"`
	Capacity            int        `gorm:"type:int;column:capacity"`
	NumberOfMeeting     int        `gorm:"type:int;column:number_of_meeting"`
	ContractDescription *string    `gorm:"type:nvarchar(max);column:contract_description"`
	ContractFilePath    *string    `gorm:"type:nvarchar(max);column:contract_file_path"`
	CreatedAt           int64      `gorm:"type:bigint;column:created_at"`
	CreatedBy           *uuid.UUID `gorm:"type:char(36);column:created_by"`
	UpdatedAt           *int64     `gorm:"type:bigint;column:updated_at"`
	UpdatedBy           *uuid.UUID `gorm:"type:char(36);column:updated_by"`
	DeletedAt           *int64     `gorm:"type:bigint;column:deleted_at"`
	DeletedBy           *uuid.UUID `gorm:"type:char(36);column:deleted_by"`

	// Additional Fields
	TotalParticipant int    `gorm:"column:total_participants"`
	SubjectNameID    string `gorm:"column:subject_name_id"`
	SubjectNameEN    string `gorm:"column:subject_name_en"`
	LecturerName     string `gorm:"column:lecturer_name"`
	StudyProgramName string `gorm:"column:study_program_name"`
	DayName          string `gorm:"column:day_name"`
	StartTime        string `gorm:"column:start_time"`
	EndTime          string `gorm:"column:end_time"`

	// relataion detail id
	LecturerSystem          string `gorm:"column:lecturer_system"`
	SubjectTotalSks         int    `gorm:"column:subject_total_sks"`
	CurriculumYearName      string `gorm:"column:curriculum_year_name"`
	StartDateOfCollege      string `gorm:"column:start_date_of_college"`
	EndDateOfCollege        string `gorm:"column:end_date_of_college"`
	AcademicPeriodeFullname string `gorm:"column:academic_periode_fullname"`
}

func (MstClass) TableName() string {
	return "mst_classes"
}

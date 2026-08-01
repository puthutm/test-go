package dto

import (
	"mime/multipart"

	"github.com/google/uuid"
)

/* Request */
type MstClassRequest struct {
	ID                string `json:"-"`
	Code              string `json:"code" validate:"required,stringMax=50"`
	Name              string `json:"name" validate:"required,stringMax=225"`
	AcademicPeriodeID string `json:"academic_periode_id" validate:"required,uuid"`
	SubjectID         string `json:"subject_id" validate:"required,uuid"`
	StudyProgramID    string `json:"-" `
	CurriculumYearID  string `json:"-"`
	Capacity          int    `json:"capacity" validate:"required,gt=0,lte=50"`
	NumberOfMeeting   int    `json:"number_of_meeting" validate:"required,gt=0,lt=20"`
}

type MstClassRequestByProgramHead struct {
	ID                string `json:"-"`
	Code              string `json:"code" validate:"required,stringMax=50"`
	Name              string `json:"name" validate:"required,stringMax=225"`
	AcademicPeriodeID string `json:"academic_periode_id" validate:"required,uuid"`
	SubjectID         string `json:"subject_id" validate:"required,uuid"`
	CurriculumYearID  string `json:"-"`
	Capacity          int    `json:"capacity" validate:"required,gt=0,lt=50"`
	NumberOfMeeting   int    `json:"number_of_meeting" validate:"required,gt=0,lt=20"`
}

type MstClassContractRequest struct {
	ID                  uuid.UUID `form:"-"`
	ContractDescription *string   `form:"contract_description"`
	ContractFile        *multipart.FileHeader
}

/* Response */
type MstClassResponse struct {
	ID                  uuid.UUID  `json:"id" gorm:"column:id"`
	Code                string     `json:"code" gorm:"column:code"`
	Name                string     `json:"name" gorm:"column:name"`
	AcademicPeriodeID   uuid.UUID  `json:"academic_periode_id" gorm:"column:academic_periode_id"`
	SubjectID           uuid.UUID  `json:"subject_id" gorm:"column:subject_id"`
	StudyProgramID      string     `json:"study_program_id" gorm:"column:study_program_id"`
	CurriculumYearID    uuid.UUID  `json:"curriculum_year_id" gorm:"column:curriculum_year_id"`
	Capacity            int        `json:"capacity" gorm:"column:capacity"`
	NumberOfMeeting     int        `json:"number_of_meeting"`
	ContractDescription *string    `json:"contract_description" gorm:"column:contract_description"`
	ContractFilePath    *string    `json:"contract_file_path" gorm:"column:contract_file_path"`
	CreatedAt           int64      `json:"created_at" gorm:"column:created_at"`
	CreatedBy           *uuid.UUID `json:"created_by" gorm:"column:created_by"`
	UpdatedAt           *int64     `json:"updated_at" gorm:"column:updated_at"`
	UpdatedBy           *uuid.UUID `json:"updated_by" gorm:"column:updated_by"`
	DeletedAt           *int64     `json:"deleted_at" gorm:"column:deleted_at"`
	DeletedBy           *uuid.UUID `json:"deleted_by" gorm:"column:deleted_by"`

	TotalParticipant int    `json:"total_participant" gorm:"column:total_participants"`
	SubjectNameID    string `json:"subject_name_id" gorm:"column:subject_name_id"`
	SubjectNameEN    string `json:"subject_name_en" gorm:"column:subject_name_en"`
	LecturerctName   string `json:"lecturer_name" gorm:"column:lecturer_name"`
	StudyProgramName string `json:"study_program_name"`

	// relataion detail id
	LecturerSystem          string `json:"lecturer_system"`
	SubjectTotalSks         int    `json:"subject_total_sks"`
	CurriculumYearName      string `json:"curriculum_year_name"`
	StartDateOfCollege      string `json:"start_date_of_college"`
	EndDateOfCollege        string `json:"end_date_of_college"`
	AcademicPeriodeFullname string `json:"academic_periode_fullname"`
}
type MstClassStudentDistributionResponse struct {
	ClassID   uuid.UUID `json:"id"`
	ClassCode string    `json:"code"`
	ClassName string    `json:"name"`
}

type MstClassResponseForSchedule struct {
	ID                 uuid.UUID `json:"id" gorm:"column:id"`
	Code               string    `json:"code" gorm:"column:code"`
	Name               string    `json:"name" gorm:"column:name"`
	Capacity           int       `json:"capacity" gorm:"column:capacity"`
	TotalParticipant   int       `json:"total_participant" gorm:"column:total_participants"`
	SubjectNameID      string    `json:"subject_name_id" gorm:"column:subject_name_id"`
	SubjectNameEN      string    `json:"subject_name_en" gorm:"column:subject_name_en"`
	LecturerctName     string    `json:"lecturer_name" gorm:"column:lecturer_name"`
	StudyProgramName   string    `json:"study_program_name"`
	CurriculumYearName string    `json:"curriculum_year_name"`
	DayName            string    `json:"day_name"`
	StartTime          string    `json:"start_time"`
	EndTime            string    `json:"end_time"`
}

type MstClassCheckSaveButtonResponse struct {
	StatusLock bool `json:"status_lock"`
}

type MstClassUpdateStatusLockedRequest struct {
	AcademicPeriodeID string `json:"-"`
	ClassID           string `json:"-"`
	StatusLocked      bool   `json:"status_locked"`
	CreatedBy         string `json:"-"`
	CreatedAt         int64  `json:"-"`
}

package dto

import (
	"mime/multipart"

	"github.com/google/uuid"
)

/* Request */
type TrxFinalProjectProposalRequest struct {
	ID       uuid.UUID             `form:"-"`
	TitleID  string                `form:"title_id" validate:"required"`
	TitleEN  string                `form:"title_en" validate:"required"`
	Topic    string                `form:"topic" validate:"required"`
	Abstract string                `form:"abstract" validate:"required"`
	File     *multipart.FileHeader `form:"file"`
}

type TrxFinalProjectProposalUpdateStatusRequest struct {
	ID       string `json:"-"`
	Status   string `json:"status" validate:"required,oneof=1 2 3"`
	Feedback string `json:"feedback" validate:"required"`
}
type TrxFinalProjectProposalAssignAcademicSupervisorRequest struct {
	ID                     string `json:"-"`
	FinalProjectProposalID string `json:"-"`
	LecturerID             string `json:"lecturer_id" validate:"required,uuid"`
}

/* Response */
type TrxFinalProjectProposalResponse struct {
	ID                     uuid.UUID  `json:"id"`
	StudentID              uuid.UUID  `json:"student_id"`
	TitleID                string     `json:"title_id"`
	TitleEN                string     `json:"title_en"`
	Topic                  string     `json:"topic"`
	AcademicPeriodeID      uuid.UUID  `json:"academic_periode_id"`
	StudyProgramID         uuid.UUID  `json:"study_program_id"`
	Abstract               string     `json:"abstract"`
	FilePath               string     `json:"file_path"`
	Status                 int        `json:"status"`
	Date                   *int64     `json:"date"`
	ConfirmationStatusDate *int64     `json:"confirmation_status_date"`
	ConfirmationBy         *uuid.UUID `json:"confirmation_by"`
	Feedback               *string    `json:"feedback"`
}

type TrxFinalProjectProposalResponseForAdmin struct {
	TrxFinalProjectProposalResponse
	MentorLecturers []TrxFinalProjectProposalMentorLecturerResponse `json:"mentor_lecturers"`
}

type TrxFinalProjectProposalByUserIDResponse struct {
	ID      uuid.UUID `json:"id"`
	TitleID string    `json:"title_id"`
	TitleEN string    `json:"title_en"`
	Status  int       `json:"status"`
	Date    *int64    `json:"date"`
}

type TrxFinalProjectProposalByStudenIDandStudyProgramIDResponse struct {
	StudentName          *string `json:"student_name"`
	StudentNIM           *string `json:"student_nim"`
	StudyProgramName     *string `json:"study_program_name"`
	AcademicPeriodeName  *string `json:"academic_periode_name"`
	RegistrationPathName *string `json:"registration_path_name"`
}

type TrxFinalProjectProposalsResponseForAdmin struct {
	TrxFinalProjectProposalByUserIDResponse
	Topic            string  `json:"topic"`
	StudentName      *string `json:"student_name"`
	StudyProgramName *string `json:"study_program_name"`
	LecturerName     *string `json:"lecturer_name"`
}

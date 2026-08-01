package dto

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

/* Request */
type MstClassScheduleRequest struct {
	ID            uuid.UUID `json:"-"`
	ClassID       uuid.UUID `json:"-"`
	Session       string    `json:"session" validate:"required,numeric"`
	DayName       string    `json:"day_name" validate:"required,dayNameEN"`
	Date          string    `json:"date" validate:"required,datetime=2006-01-02"`
	StartTime     string    `json:"start_time" validate:"required,datetime=15:04:05"`
	EndTime       string    `json:"end_time" validate:"required,datetime=15:04:05"`
	TypeOfMeeting string    `json:"type_of_meeting" validate:"required"`
	Status        string    `json:"status" validate:"required,numeric"`
}

type GenerateScheduleRequest struct {
	ClassID uuid.UUID
}

// Update
type MstClassScheduleUpdateRequest struct {
	ID            uuid.UUID `json:"-"`
	ClassID       uuid.UUID `json:"-"`
	DayNameOld    string    `json:"day_name_old" validate:"required,dayNameEN"`
	StartTimeOld  string    `json:"start_time_old" validate:"required,datetime=15:04:05"`
	EndTimeOld    string    `json:"end_time_old" validate:"required,datetime=15:04:05"`
	DayName       string    `json:"day_name" validate:"required,dayNameEN"`
	Date          string    `json:"date" validate:"required,datetime=2006-01-02"`
	StartTime     string    `json:"start_time" validate:"required,datetime=15:04:05"`
	EndTime       string    `json:"end_time" validate:"required,datetime=15:04:05"`
	TypeOfMeeting string    `json:"type_of_meeting" validate:"required"`
}

type MstClassScheduleGetByRequest struct {
	ID        uuid.UUID `json:"-"`
	ClassID   uuid.UUID `json:"-"`
	DayName   string    `query:"day_name" validate:"required,dayNameEN"`
	Date      string    `query:"date" validate:"required,datetime=2006-01-02"`
	StartTime string    `query:"start_time" validate:"required,datetime=15:04:05"`
	EndTime   string    `query:"end_time" validate:"required,datetime=15:04:05"`
}

type MstClassScheduleUpdateForLecturerRequest struct {
	ID                     uuid.UUID             `form:"-"`
	MaterialAttachmentFile *multipart.FileHeader `form:"material_attachment_file"`
	AttendanceDocumentFile *multipart.FileHeader `form:"attendance_document_file"`
	JournalDocumentFile    *multipart.FileHeader `form:"journal_document_file"`
	MaterialPlan           *string               `form:"material_plan"`
	MaterialRealization    *string               `form:"material_realization"`
}

/* Response */
type MstClassScheduleResponse struct {
	ID                         *uuid.UUID `json:"id,omitempty"`
	ClassID                    *uuid.UUID `json:"class_id,omitempty"`
	ScheduleTemplateID         uuid.UUID  `json:"schedule_template_id"`
	Session                    int        `json:"session,omitempty" `
	DayName                    string     `json:"day_name"`
	Date                       time.Time  `json:"date" `
	StartTime                  string     `json:"start_time"`
	EndTime                    string     `json:"end_time" `
	TypeOfMeeting              string     `json:"type_of_meeting" `
	Status                     int        `json:"status,omitempty" `
	MaterialAttachmentFilePath *string    `json:"material_attachment_file_path" `
	AttendanceDocumentFilePath *string    `json:"attendance_document_file_path" `
	JournalDocumentFilePath    *string    `json:"journal_document_file_path" `
	MaterialPlan               *string    `json:"material_plan" `
	MaterialRealization        *string    `json:"material_realization"`
	IsUTS                      bool       `json:"is_uts"`
	IsUAS                      bool       `json:"is_uas"`
}

type MstClassScheduleResponseForListByClass struct {
	ID                 uuid.UUID `json:"id"`
	ScheduleTemplateID uuid.UUID `json:"schedule_template_id"`
	Session            int       `json:"session,omitempty" `
	DayName            string    `json:"day_name"`
	Date               time.Time `json:"date" `
	StartTime          string    `json:"start_time"`
	EndTime            string    `json:"end_time" `
	TypeOfMeeting      string    `json:"type_of_meeting" `
	Status             int       `json:"status,omitempty" `
	IsUTS              bool      `json:"is_uts"`
	IsUAS              bool      `json:"is_uas"`
}

type MstClassScheduleAcademicSystemDistributionResponse struct {
	TypeOfMeeting string `json:"type_of_meeting" `
}

// for spesisik

type MstClassScheduleForClassSessionPresenceResponse struct {
	SessionID                 uuid.UUID `json:"session_id"`
	SessionName               string    `json:"session_name"`
	SessionDate               time.Time `json:"session_date"`
	ClassName                 string    `json:"class_name"`
	ClassCode                 string    `json:"class_code"`
	ScheduleTemplateDayName   string    `json:"class_schedule_template_day_name"`
	ScheduleTemplateStartTime string    `json:"class_schedule_template_start_time"`
	ScheduleTemplateEndTime   string    `json:"class_schedule_template_end_time"`
	AcademicPeriodName        string    `json:"academic_periode_name"`
	TotalParticipant          int       `json:"total_participant"`
	StudentPresenceCount      int       `json:"student_presence_count"`
	StudentAbsentCount        int       `json:"student_absent_count"`
}

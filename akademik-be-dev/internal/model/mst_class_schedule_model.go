package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MstClassSchedule struct {
	ID                         uuid.UUID `gorm:"type:char(36);column:id;primaryKey"`
	ClassID                    uuid.UUID `gorm:"type:char(36);column:class_id"`
	ScheduleTemplateID         uuid.UUID `gorm:"type:char(36);column:schedule_template_id"`
	Session                    int       `gorm:"type:int;column:session"`
	DayName                    string    `gorm:"type:nvarchar(100);column:day_name"`
	Date                       time.Time `gorm:"type:date;column:date"`
	StartTime                  string    `gorm:"type:time;column:start_time"`
	EndTime                    string    `gorm:"type:time;column:end_time"`
	TypeOfMeeting              string    `gorm:"type:nvarchar(max);column:type_of_meeting"`
	Status                     int       `gorm:"type:int;column:status"`
	IsUTS                      bool      `gorm:"column:is_uts"`
	IsUAS                      bool      `gorm:"column:is_uas"`
	MaterialAttachmentFilePath *string   `gorm:"type:nvarchar(max);column:material_attachment_file_path"`
	AttendanceDocumentFilePath *string   `gorm:"type:nvarchar(max);column:attendance_document_file_path"`
	JournalDocumentFilePath    *string   `gorm:"type:nvarchar(max);column:journal_document_file_path"`
	MaterialPlan               *string        `gorm:"type:nvarchar(max);column:material_plan"`
	MaterialRealization        *string        `gorm:"type:nvarchar(max);column:material_realization"`
	DeletedAt                  gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (MstClassSchedule) TableName() string {
	return "mst_class_schedules"
}

// for spesifik

type MstClassScheduleForClassSessionPresence struct {
	SessionID                 uuid.UUID `gorm:"column:session_id"`
	SessionName               string    `gorm:"column:session_name"`
	SessionDate               time.Time `gorm:"column:session_date"`
	ClassName                 string    `gorm:"column:class_name"`
	ClassCode                 string    `gorm:"column:class_code"`
	ScheduleTemplateDayName   string    `gorm:"column:class_schedule_template_day_name"`
	ScheduleTemplateStartTime string    `gorm:"column:class_schedule_template_start_time"`
	ScheduleTemplateEndTime   string    `gorm:"column:class_schedule_template_end_time"`
	AcademicPeriodName        string    `gorm:"column:academic_periode_name"`
	TotalParticipant          int       `gorm:"column:total_participant"`
	StudentPresenceCount      int       `gorm:"column:student_presence_count"`
	StudentAbsentCount        int       `gorm:"column:student_absent_count"`
}

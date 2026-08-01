// Package model
package model

import (
	"time"

	"github.com/google/uuid"
)

// Presence
// UseOpenSession        bool `gorm:"type:bit;column:use_open_session;not null"`
// OpenSessionPercentage int  `gorm:"type:int(10);column:open_session_percentage;not null"`
//
// UseDocumentMaterial     bool `gorm:"type:bit;column:use_document_material;not null"`
// DocumentMaterialPercent int  `gorm:"type:int(10);column:document_material_percentage;not null"`
//
// UseQuiz        bool `gorm:"type:bit;column:use_quiz;not null"`
// QuizPercentage int  `gorm:"type:int(10);column:quiz_percentage;not null"`
//
// UseTask        bool `gorm:"type:bit;column:use_task;not null"`
// TaskPercentage int  `gorm:"type:int(10);column:task_percentage;not null"`
//
// UseVideo        bool `gorm:"type:bit;column:use_video;not null"`
// VideoPercentage int  `gorm:"type:int(10);column:video_percentage;not null"`
//
// UseUTS        bool `gorm:"type:bit;column:use_uts;not null"`
// UTSPercentage int  `gorm:"type:int(10);column:uts_percentage;not null"`
//
// UseUAS        bool `gorm:"type:bit;column:use_uas;not null"`
// UASPercentage int  `gorm:"type:int(10);column:uas_percentage;not null"`
//
// UseComment        bool `gorm:"type:bit;column:use_comment;not null"`
// CommentPercentage int  `gorm:"type:int(10);column:comment_percentage;not null"`

type Presence struct {
	UseOpenSession        bool `gorm:"column:use_open_session;not null"`
	OpenSessionPercentage int  `gorm:"column:open_session_percentage;not null"`

	UseDocumentMaterial        bool `gorm:"column:use_document_material;not null"`
	DocumentMaterialPercentage int  `gorm:"column:document_material_percentage;not null"`

	UseQuiz        bool `gorm:"column:use_quiz;not null"`
	QuizPercentage int  `gorm:"column:quiz_percentage;not null"`

	UseTask        bool `gorm:"column:use_task;not null"`
	TaskPercentage int  `gorm:"column:task_percentage;not null"`

	UseVideo        bool `gorm:"column:use_video;not null"`
	VideoPercentage int  `gorm:"column:video_percentage;not null"`

	UseUTS        bool `gorm:"column:use_uts;not null"`
	UTSPercentage int  `gorm:"column:uts_percentage;not null"`

	UseUAS        bool `gorm:"column:use_uas;not null"`
	UASPercentage int  `gorm:"column:uas_percentage;not null"`

	UseComment        bool `gorm:"column:use_comment;not null"`
	CommentPercentage int  `gorm:"column:comment_percentage;not null"`
}

type TrxStudentPresence struct {
	ID              uuid.UUID `gorm:"type:char(36);column:id;primaryKey"`
	StudentID       uuid.UUID `gorm:"column:student_id"`
	ClassScheduleID uuid.UUID `gorm:"column:class_schedule_id"`

	Presence

	CreatedAt int64 `gorm:"type:bigint(19);column:created_at"`
}

func (TrxStudentPresence) TableName() string {
	return "trx_student_presences"
}

// for spesifik

type TrxStudentPresenceGetForLecturerParam struct {
	UserID            string
	AcademicPeriodeID string
	StudyProgramID    string
	SubjectID         string
}

type TrxStudentPresenceSaveParamBySession struct {
	IDNew          uuid.UUID
	SessionID      uuid.UUID
	StudentID      uuid.UUID
	PresenceStatus bool
	PresenceType   string
	CreatedAt      int64
}

// ---- Get

type SessionPresence struct {
	SessionID          uuid.UUID `gorm:"column:session_id"`
	Session            int       `gorm:"column:session"`
	SessionDate        time.Time `gorm:"column:session_date"`
	PresencePercentage float64   `gorm:"column:presence_percentage"`
}

type TrxStudentPresenceBySession struct {
	StudentID               string  `gorm:"column:student_id"`
	StudentName             string  `gorm:"column:student_name"`
	StudentNIM              string  `gorm:"column:student_nim"`
	OpenSessionPercentage   float64 `gorm:"column:open_session_percentage"`
	DocumentMaterialPercent float64 `gorm:"column:document_material_percentage"`
	QuizPercentage          float64 `gorm:"column:quiz_percentage"`
	TaskPercentage          float64 `gorm:"column:task_percentage"`
	VideoPercentage         float64 `gorm:"column:video_percentage"`
	UTSPercentage           float64 `gorm:"column:uts_percentage"`
	UASPercentage           float64 `gorm:"column:uas_percentage"`
	CommentPercentage       float64 `gorm:"column:comment_percentage"`
	TotalPercentage         float64 `gorm:"column:total_percentage"`
	PresenceFlag            int     `gorm:"column:presence_flag"` // 0 = tidak hadir, 1 = hadir
}

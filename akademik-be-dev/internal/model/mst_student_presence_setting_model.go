// Package model
package model

import (
	"github.com/google/uuid"
)

type MstStudentPresenceSetting struct {
	ID                uuid.UUID `gorm:"type:char(36);column:id;primaryKey"`
	AcademicPeriodeID uuid.UUID `gorm:"type:char(36);column:academic_periode_id;not null"`
	StudyProgramID    string    `gorm:"column:study_program_id;not null"`

	UseOpenSession        bool `gorm:"type:bit;column:use_open_session;not null"`
	OpenSessionPercentage int  `gorm:"type:int(10);column:open_session_percentage;not null"`

	UseDocumentMaterial     bool `gorm:"type:bit;column:use_document_material;not null"`
	DocumentMaterialPercent int  `gorm:"type:int(10);column:document_material_percentage;not null"`

	UseQuiz        bool `gorm:"type:bit;column:use_quiz;not null"`
	QuizPercentage int  `gorm:"type:int(10);column:quiz_percentage;not null"`

	UseTask        bool `gorm:"type:bit;column:use_task;not null"`
	TaskPercentage int  `gorm:"type:int(10);column:task_percentage;not null"`

	UseVideo        bool `gorm:"type:bit;column:use_video;not null"`
	VideoPercentage int  `gorm:"type:int(10);column:video_percentage;not null"`

	UseUTS        bool `gorm:"type:bit;column:use_uts;not null"`
	UTSPercentage int  `gorm:"type:int(10);column:uts_percentage;not null"`

	UseUAS        bool `gorm:"type:bit;column:use_uas;not null"`
	UASPercentage int  `gorm:"type:int(10);column:uas_percentage;not null"`

	UseComment        bool `gorm:"type:bit;column:use_comment;not null"`
	CommentPercentage int  `gorm:"type:int(10);column:comment_percentage;not null"`

	CreatedAt *int64     `gorm:"type:bigint(19);column:created_at"`
	CreatedBy *uuid.UUID `gorm:"type:char(36);column:created_by"`
	UpdatedAt *int64     `gorm:"type:bigint(19);column:updated_at"`
	UpdatedBy *uuid.UUID `gorm:"type:char(36);column:updated_by"`
	DeletedAt *int64     `gorm:"type:bigint(19);column:deleted_at"`
	DeletedBy *uuid.UUID `gorm:"type:char(36);column:deleted_by"`
}

func (MstStudentPresenceSetting) TableName() string {
	return "mst_student_presence_settings"
}

// for request spesifik

type MstStudentPresenceSettingDuplicateParam struct {
	ID                   uuid.UUID `gorm:"type:char(36);column:id"`
	AcademicPeriodeID    uuid.UUID `gorm:"type:char(36);column:academic_periode_id"`
	StudyProgramID       string    `gorm:"column:study_program_id"`
	AcademicPeriodeIDOld uuid.UUID `gorm:"type:char(36)"`
	CreatedAt            int64     `gorm:"type:bigint(19);column:created_at"`
	CreatedBy            uuid.UUID `gorm:"type:char(36);column:created_by"`
}

type MstStudentPresenceSettingGetResult struct {
	MstStudentPresenceSetting
	AcademicPeriodeName string `gorm:"column:academic_periode_name"`
	StudyProgramName    string `gorm:"column:study_program_name"`
}

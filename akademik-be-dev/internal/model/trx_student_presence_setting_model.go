// Package model
package model

import (
	"github.com/google/uuid"
)

type TrxStudentPresenceSetting struct {
	ID                uuid.UUID `gorm:"type:char(36);column:id;primaryKey"`
	AcademicPeriodeID uuid.UUID `gorm:"type:char(36);column:academic_periode_id;not null"`
	StudyProgramID    uuid.UUID `gorm:"type:char(36);column:study_program_id;not null"`
	LecturerID        uuid.UUID `gorm:"column:lecturer_id"`
	SubjectID         uuid.UUID `gorm:"column:subject_id"`

	Presence

	CreatedAt *int64     `gorm:"type:bigint(19);column:created_at"`
	CreatedBy *uuid.UUID `gorm:"type:char(36);column:created_by"`
	UpdatedAt *int64     `gorm:"type:bigint(19);column:updated_at"`
	UpdatedBy *uuid.UUID `gorm:"type:char(36);column:updated_by"`
	DeletedAt *int64     `gorm:"type:bigint(19);column:deleted_at"`
	DeletedBy *uuid.UUID `gorm:"type:char(36);column:deleted_by"`
}

func (TrxStudentPresenceSetting) TableName() string {
	return "trx_student_presence_settings"
}

// for request spesifik

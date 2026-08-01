package model

import (
	"github.com/google/uuid"
)

type MstClassLecturer struct {
	ID                  uuid.UUID  `gorm:"type:char(36);column:id;primaryKey"`
	ClassID             uuid.UUID  `gorm:"type:char(36);column:class_id"`
	LecturerID          uuid.UUID  `gorm:"type:char(36);column:lecturer_id"`
	SubtituteLecturerID *uuid.UUID `gorm:"type:nvarchar(50);column:subtitute_lecturer_id"`

	// Additional Fields
	LecturerName          string `gorm:"column:lecturer_name"`
	SubtituteLecturerName string `gorm:"column:subtitute_lecturer_name"`
}

func (MstClassLecturer) TableName() string {
	return "mst_class_lecturers"
}

// for spesifik req

type MstClassLecturerGetByAcademicPeriodAndSubjectAndUserForLecturer struct {
	ClassID   uuid.UUID `gorm:"type:char(36);column:class_id"`
	ClassName string    `gorm:"column:class_name"`
	ClassCode string    `gorm:"column:class_code"`

	StudyProgramID string `gorm:"column:study_program_id"`
}

type MstClassGetResultSubjectAndClassCount struct {
	AcademicPeriodeID   uuid.UUID `gorm:"column:academic_periode_id"`
	AcademicPeriodeName string    `gorm:"column:academic_periode_name"`
	SubjectID           uuid.UUID `gorm:"column:subject_id"`
	SubjectNameID       string    `gorm:"column:subject_name_id"`
	SubjectNameEN       string    `gorm:"column:subject_name_en"`
	StudyProgramID      uuid.UUID `gorm:"column:study_program_id"`
	LecturerID          uuid.UUID `gorm:"column:lecturer_id"`
	ClassCount          int       `gorm:"column:class_count"`
}

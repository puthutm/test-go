// Package dto
package dto

import "github.com/google/uuid"

/* Request */

type MstClassLecturerRequest struct {
	ID                  uuid.UUID  `json:"-"`
	ClassID             uuid.UUID  `json:"-" `
	LecturerID          uuid.UUID  `json:"lecturer_id"  validate:"required,uuid"`
	SubtituteLecturerID *uuid.UUID `json:"subtitute_lecturer_id"  validate:"omitempty,uuid"`
}
type MstClassLecturerUpdate struct {
	ID                  uuid.UUID  `json:"-"`
	ClassID             uuid.UUID  `json:"-" `
	LecturerID          uuid.UUID  `json:"lecturer_id"  validate:"required,uuid"`
	SubtituteLecturerID *uuid.UUID `json:"subtitute_lecturer_id"  validate:"omitempty,uuid"`
}

/* Response */

type MstClassLecturerResponse struct {
	ID                  uuid.UUID  `json:"id"`
	ClassID             uuid.UUID  `json:"class_id"`
	LecturerID          uuid.UUID  `json:"lecturer_id"`
	SubtituteLecturerID *uuid.UUID `json:"subtitute_lecturer_id"`

	// realtion
	LecturerName          string `gorm:"column:lecturer_name"`
	SubtituteLecturerName string `gorm:"column:subtitute_lecturer_name"`
}

type MstClassLecturerGetByAcademicPeriodAndSubjectAndUserForLecturerResponse struct {
	ClassID   uuid.UUID `json:"class_id"`
	ClassName string    `json:"class_name"`
	ClassCode string    `json:"class_code"`

	StudyProgramID string `json:"study_program_id"`
}

type MstClassGetResultSubjectAndClassCountResponse struct {
	AcademicPeriodeID   uuid.UUID `json:"academic_periode_id"`
	AcademicPeriodeName string    `json:"academic_periode_name"`
	SubjectID           uuid.UUID `json:"subject_id"`
	SubjectNameID       string    `json:"subject_name_id"`
	SubjectNameEN       string    `json:"subject_name_en"`
	StudyProgramID      uuid.UUID `json:"study_program_id"`
	LecturerID          uuid.UUID `json:"lecturer_id"`
	ClassCount          int       `json:"class_count"`
}

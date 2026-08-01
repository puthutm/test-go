// Package dto
package dto

import "github.com/google/uuid"

// type MstStudentPresenceSettingCreateRequest struct {
// 	AcademicPeriodeID uuid.UUID `json:"academic_periode_id" validate:"required"`
// 	StudyProgramID    uuid.UUID `json:"study_program_id" validate:"required"`
//
// 	UseOpenSession        bool `json:"use_open_session" validate:"bool"`
// 	OpenSessionPercentage int  `json:"open_session_percentage" validate:"required"`
//
// 	UseDocumentMaterial     bool `json:"use_document_material" validate:"bool"`
// 	DocumentMaterialPercent int  `json:"document_material_percentage" validate:"required"`
//
// 	UseQuiz        bool `json:"use_quiz" validate:"bool"`
// 	QuizPercentage int  `json:"quiz_percentage" validate:"required"`
//
// 	UseTask        bool `json:"use_task" validate:"bool"`
// 	TaskPercentage int  `json:"task_percentage" validate:"required"`
//
// 	UseVideo        bool `json:"use_video" validate:"bool"`
// 	VideoPercentage int  `json:"video_percentage" validate:"required"`
//
// 	UseUTS        bool `json:"use_uts" validate:"bool"`
// 	UTSPercentage int  `json:"uts_percentage" validate:"required"`
//
// 	UseUAS        bool `json:"use_uas" validate:"bool"`
// 	UASPercentage int  `json:"uas_percentage" validate:"required"`
//
// 	UseComment        bool `json:"use_comment" validate:"bool"`
// 	CommentPercentage int  `json:"comment_percentage" validate:"required"`
// }

type MstStudentPresenceSettingCreateRequest struct {
	AcademicPeriodeID uuid.UUID `json:"academic_periode_id" validate:"required"`
	StudyProgramID    string    `json:"study_program_id" validate:"required"`

	UseOpenSession        bool `json:"use_open_session" validate:"bool"`
	OpenSessionPercentage int  `json:"open_session_percentage" validate:"gte=0,lte=100"`

	UseDocumentMaterial     bool `json:"use_document_material" validate:"bool"`
	DocumentMaterialPercent int  `json:"document_material_percentage" validate:"gte=0,lte=100"`

	UseQuiz        bool `json:"use_quiz" validate:"bool"`
	QuizPercentage int  `json:"quiz_percentage" validate:"gte=0,lte=100"`

	UseTask        bool `json:"use_task" validate:"bool"`
	TaskPercentage int  `json:"task_percentage" validate:"gte=0,lte=100"`

	UseVideo        bool `json:"use_video" validate:"bool"`
	VideoPercentage int  `json:"video_percentage" validate:"gte=0,lte=100"`

	UseUTS bool `json:"use_uts" validate:"bool"`

	UseUAS bool `json:"use_uas" validate:"bool"`

	UseComment        bool `json:"use_comment" validate:"bool"`
	CommentPercentage int  `json:"comment_percentage" validate:"gte=0,lte=100"`
}

type MstStudentPresenceSettingDuplicateRequest struct {
	AcademicPeriodeID       uuid.UUID `json:"academic_periode_id" validate:"required"`
	StudyProgramID          string    `json:"study_program_id" validate:"required"`
	AcademicPeriodeIDTarget uuid.UUID `json:"academic_periode_id_target" validate:"required"`
}

type MstStudentPresenceSettingGetResponse struct {
	ID                uuid.UUID `json:"id"`
	AcademicPeriodeID uuid.UUID `json:"academic_periode_id"`
	StudyProgramID    string    `json:"study_program_id"`
}

type MstStudentPresenceSettingGetResultResponse struct {
	MstStudentPresenceSettingGetResponse
	AcademicPeriodeName string `json:"academic_periode_name"`
	StudyProgramName    string `json:"study_program_name"`
}

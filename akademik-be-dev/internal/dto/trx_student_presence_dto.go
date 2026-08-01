// Package dto
package dto

import (
	"time"

	"github.com/google/uuid"
)

type TrxStudentPresenceSettingCreateOrUpdateRequest struct {
	AcademicPeriodeID uuid.UUID `json:"academic_periode_id" validate:"required,uuid_valid_not_nil"`
	StudyProgramID    uuid.UUID `json:"study_program_id" validate:"required,uuid_valid_not_nil"`
	// LecturerID        uuid.UUID `json:"lecturer_id" validate:"required"`
	SubjectID uuid.UUID `json:"subject_id" validate:"required"`

	PresenceRes
}

type TrxStudentPresenceGetForLecturerRequest struct {
	AcademicPeriodeID string `json:"academic_periode_id" validate:"required,uuid_valid_not_nil"`
	StudyProgramID    string `json:"study_program_id" query:"study_program_id" validate:"required,uuid_valid_not_nil"`
	SubjectID         string `json:"subject_id" validate:"required,uuid_valid_not_nil"`
}

type PresenceParam struct {
	PresenceStatus bool   `json:"presence_status"`
	PresenceType   string `json:"presence_type"`
}
type TrxStudentPresenceSaveParamBySessionRequest struct {
	SessionID uuid.UUID `json:"session_id" validate:"required,uuid_valid_not_nil"`
	StudentID uuid.UUID `json:"student_id" validate:"required,uuid_valid_not_nil"`
	PresenceParam
}

type TrxStudentPresenceSliceSaveParamBySessionRequest struct {
	SessionID      uuid.UUID       `json:"session_id" validate:"required,uuid_valid_not_nil"`
	StudentID      uuid.UUID       `json:"student_id" validate:"required,uuid_valid_not_nil"`
	PresenceParams []PresenceParam `json:"presence_params"`
}

//---- res

type PresenceRes struct {
	UseOpenSession        bool `json:"use_open_session"`
	OpenSessionPercentage int  `json:"open_session_percentage"`

	UseDocumentMaterial        bool `json:"use_document_material"`
	DocumentMaterialPercentage int  `json:"document_material_percentage"`

	UseQuiz        bool `json:"use_quiz"`
	QuizPercentage int  `json:"quiz_percentage"`

	UseTask        bool `json:"use_task"`
	TaskPercentage int  `json:"task_percentage"`

	UseVideo        bool `json:"use_video"`
	VideoPercentage int  `json:"video_percentage"`

	UseUTS        bool `json:"use_uts"`
	UTSPercentage int  `json:"uts_percentage"`

	UseUAS        bool `json:"use_uas"`
	UASPercentage int  `json:"uas_percentage"`

	UseComment        bool `json:"use_comment"`
	CommentPercentage int  `json:"comment_percentage"`
}

type SessionPresenceResponse struct {
	SessionID          uuid.UUID `json:"session_id"`
	Session            int       `json:"session"`
	SessionDate        time.Time `json:"session_date"`
	PresencePercentage float64   `json:"presence_percentage"`
}

type TrxStudentPresenceBySessionResponse struct {
	StudentID               string  `json:"student_id"`
	StudentName             string  `json:"student_name"`
	StudentNIM              string  `json:"student_nim"`
	OpenSessionPercentage   float64 `json:"open_session_percentage"`
	DocumentMaterialPercent float64 `json:"document_material_percentage"`
	QuizPercentage          float64 `json:"quiz_percentage"`
	TaskPercentage          float64 `json:"task_percentage"`
	VideoPercentage         float64 `json:"video_percentage"`
	UTSPercentage           float64 `json:"uts_percentage"`
	UASPercentage           float64 `json:"uas_percentage"`
	CommentPercentage       float64 `json:"comment_percentage"`
	TotalPercentage         float64 `json:"total_percentage"`
	PresenceFlag            int     `json:"presence_flag"` // 0 = tidak hadir, 1 = hadir
}

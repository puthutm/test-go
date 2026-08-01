package dto

import (
	"github.com/google/uuid"
)

/* Request */
type MstClassScheduleTaskRequest struct {
	ID                             string    `json:"-"`
	ScheduleID                     uuid.UUID `json:"schedule_id" validate:"required,uuid"`
	Title                          string    `json:"title" validate:"required"`
	Description                    string    `json:"description" validate:"required"`
	IsGradeable                    string    `json:"is_gradeable" validate:"required,trueFalse"`
	IsUseDeadline                  string    `json:"is_use_deadline" validate:"required,trueFalse"`
	DeadlineOfAssignmentSubmission string    `json:"deadline_of_assignment_submission" validate:"required,numeric"`
	TimeToOpen                     string    `json:"time_to_open" validate:"required,numeric"`
	Retake                         string    `json:"retake" validate:"required,numeric"`
}

/* Response */
type MstClassScheduleTaskResponse struct {
	ID                             uuid.UUID `json:"id" gorm:"column:id"`
	ScheduleID                     uuid.UUID `json:"schedule_id" gorm:"column:schedule_id"`
	Title                          string    `json:"title" gorm:"column:title"`
	Description                    string    `json:"description" gorm:"column:description"`
	IsGradeable                    bool      `json:"is_gradeable" gorm:"column:is_gradeable"`
	IsUseDeadline                  bool      `json:"is_use_deadline" gorm:"column:is_use_deadline"`
	DeadlineOfAssignmentSubmission int64     `json:"deadline_of_assignment_submission" gorm:"column:deadline_of_assignment_submission"`
	IsSharing                      bool      `json:"is_sharing" gorm:"column:is_sharing"`
	SharingDate                    *int64    `json:"sharing_date" gorm:"column:sharing_date"`
	TimeToOpen                     *int      `json:"time_to_open" gorm:"column:time_to_open"`
	Retake                         *int      `json:"retake" gorm:"column:retake"`
	Views                          *int      `json:"views" gorm:"column:views"`

	// StartOfAssignmentSubmission int64     `json:"start_of_assignment_submission" gorm:"column:start_of_assignment_submission"`
	// EndOfAssignmentSubmission   int64     `json:"end_of_assignment_submission" gorm:"column:end_of_assignment_submission"`
	CreatedAt       int64  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt       *int64 `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt       *int64 `json:"deleted_at" gorm:"column:deleted_at"`
	SessionSchedule *int64 `json:"session_schedule" gorm:"column:session_schedule"`
	TotalCollect    *int64 `json:"total_collect" gorm:"column:total_collect"`
}

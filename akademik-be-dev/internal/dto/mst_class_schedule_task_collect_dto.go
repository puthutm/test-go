package dto

import (
	"github.com/google/uuid"
)

/* Request */
type MstClassScheduleTaskCollectRequest struct {
	ID             string `json:"-"`
	ScheduleTaskID string `json:"schedule_task_id" validate:"required,uuid"`
	StudentID      string `json:"student_id" validate:"required,uuid"`
}

/* Response */
type MstClassScheduleTaskCollectResponse struct {
	ID             uuid.UUID `json:"id" gorm:"column:id"`
	ScheduleTaskID uuid.UUID `json:"schedule_task_id" gorm:"column:schedule_task_id"`
	StudentID      uuid.UUID `json:"student_id" gorm:"column:student_id"`
	CreatedAt      int64     `json:"created_at" gorm:"column:created_at"`
}

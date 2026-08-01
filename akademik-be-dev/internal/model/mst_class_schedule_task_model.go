package model

import (
	"github.com/google/uuid"
)

type MstClassScheduleTask struct {
	ID                             uuid.UUID  `gorm:"type:char(36);column:id;primaryKey"`
	ScheduleID                     uuid.UUID  `gorm:"type:char(36);column:schedule_id"`
	Title                          string     `gorm:"type:nvarchar(max);column:title"`
	Description                    string     `gorm:"type:nvarchar(max);column:description"`
	IsGradeable                    bool       `gorm:"type:bit;column:is_gradeable"`
	IsUseDeadline                  bool       `gorm:"type:bit;column:is_use_deadline"`
	DeadlineOfAssignmentSubmission int64      `gorm:"type:bigint;column:deadline_of_assignment_submission"`
	IsSharing                      bool       `gorm:"type:bit;column:is_sharing"`
	SharingDate                    *int64     `gorm:"type:bigint;column:sharing_date"`
	TimeToOpen                     *int       `gorm:"type:int;column:time_to_open"`
	Retake                         *int       `gorm:"type:int;column:retake"`
	Views                          *int       `gorm:"type:int;column:views"`
	CreatedAt                      int64      `gorm:"type:bigint;column:created_at"`
	CreatedBy                      *uuid.UUID `gorm:"type:char(36);column:created_by"`
	UpdatedAt                      *int64     `gorm:"type:bigint;column:updated_at"`
	UpdatedBy                      *uuid.UUID `gorm:"type:char(36);column:updated_by"`
	DeletedAt                      *int64     `gorm:"type:bigint;column:deleted_at"`
	DeletedBy                      *uuid.UUID `gorm:"type:char(36);column:deleted_by"`

	// Additional Fields
	SessionSchedule *int64 `gorm:"column:session_schedule"`
	TotalCollect    *int64 `gorm:"column:total_collect"`
}

func (MstClassScheduleTask) TableName() string {
	return "mst_class_schedule_tasks"
}

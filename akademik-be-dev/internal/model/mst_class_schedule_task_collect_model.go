package model

import (
	"github.com/google/uuid"
)

type MstClassScheduleTaskCollect struct {
	ID             uuid.UUID `gorm:"type:char(36);column:id;primaryKey"`
	ScheduleTaskID uuid.UUID `gorm:"type:char(36);column:schedule_task_id"`
	StudentID      uuid.UUID `gorm:"type:char(36);column:student_id"`
	CreatedAt      int64     `gorm:"type:bigint;column:created_at"`
}

func (MstClassScheduleTaskCollect) TableName() string {
	return "mst_class_schedule_task_collects"
}

package model

import (
	"github.com/google/uuid"
)

type MstClassScheduleTemplate struct {
	ID            uuid.UUID `gorm:"type:char(36);primaryKey"`
	ClassID       uuid.UUID `gorm:"type:char(36);not null"`
	DayName       string    `gorm:"type:nvarchar(100);not null"`
	StartTime     string    `gorm:"type:time;not null"`
	EndTime       string    `gorm:"type:time;not null"`
	TypeOfMeeting string    `gorm:"type:nvarchar(max);not null"`

	CreatedAt int64      `gorm:"column:created_at"`
	CreatedBy *uuid.UUID `gorm:"type:char(36)"`
	UpdatedAt *int64     `gorm:"column:updated_at"`
	UpdatedBy *uuid.UUID `gorm:"type:char(36)"`
}

func (MstClassScheduleTemplate) TableName() string {
	return "mst_class_schedule_templates"
}

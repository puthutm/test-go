package model

import (
	"github.com/google/uuid"
)

type MstClassScheduleTeachingMaterial struct {
	ID          uuid.UUID  `gorm:"type:char(36);column:id;primaryKey"`
	ScheduleID  string     `gorm:"type:char(36);column:schedule_id;not null"`
	Title       string     `gorm:"type:nvarchar(225);column:title;not null"`
	Description *string    `gorm:"type:nvarchar(max);column:description"`
	IsSharing   bool       `gorm:"type:bit;column:is_sharing;not null"`
	SharingDate string     `gorm:"type:date;column:sharing_date;not null"`
	Views       int        `gorm:"type:int;column:views"`
	CreatedAt   *int64     `gorm:"type:bigint;column:created_at"`
	CreatedBy   *uuid.UUID `gorm:"type:char(36);column:created_by"`
	UpdatedAt   *int64     `gorm:"type:bigint;column:updated_at"`
	UpdatedBy   *uuid.UUID `gorm:"type:char(36);column:updated_by"`
	DeletedAt   *int64     `gorm:"type:bigint;column:deleted_at"`
	DeletedBy   *uuid.UUID `gorm:"type:char(36);column:deleted_by"`
}

func (MstClassScheduleTeachingMaterial) TableName() string {
	return "mst_class_schedule_teaching_materials"
}

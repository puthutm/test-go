package model

import (
	"github.com/google/uuid"
)

type MstSKSLimit struct {
	ID        uuid.UUID  `gorm:"type:char(36);column:id;primaryKey"`
	IPSMin    float64    `gorm:"column:ips_min;type:decimal(3,2)"`
	IPSMax    float64    `gorm:"column:ips_max;type:decimal(3,2)"`
	SKSLimit  int        `gorm:"column:sks_limit;type:int"`
	CreatedAt int64      `gorm:"type:bigint;column:created_at"`
	CreatedBy *uuid.UUID `gorm:"type:char(36);column:created_by"`
	UpdatedAt *int64     `gorm:"type:bigint;column:updated_at"`
	UpdatedBy *uuid.UUID `gorm:"type:char(36);column:updated_by"`
	DeletedAt *int64     `gorm:"type:bigint;column:deleted_at"`
	DeletedBy *uuid.UUID `gorm:"type:char(36);column:deleted_by"`
}

func (MstSKSLimit) TableName() string {
	return "mst_sks_limits"
}

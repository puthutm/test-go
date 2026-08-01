package model

import (
	"github.com/google/uuid"
)

type MstValueComposition struct {
	ID                uuid.UUID `gorm:"type:char(36);column:id;primaryKey"`
	ValueElementID    uuid.UUID `gorm:"type:char(36);column:value_element_id"`
	AcademicPeriodeID uuid.UUID `gorm:"type:char(36);column:academic_periode_id"`

	Percentage           float64    `gorm:"type:decimal(3,2);column:percentage"`
	IsPassingRequirement bool       `gorm:"type:bit;column:is_passing_requirement"`
	CreatedAt            int64      `gorm:"type:bigint;column:created_at"`
	CreatedBy            *uuid.UUID `gorm:"type:char(36);column:created_by"`
	UpdatedAt            *int64     `gorm:"type:bigint;column:updated_at"`
	UpdatedBy            *uuid.UUID `gorm:"type:char(36);column:updated_by"`
	DeletedAt            *int64     `gorm:"type:bigint;column:deleted_at"`
	DeletedBy            *uuid.UUID `gorm:"type:char(36);column:deleted_by"`

	// Additional Fields
	ValueElementName    string `gorm:"column:value_element_name"`
	AcademicPeriodeName string `gorm:"column:academic_periode_name"`
}

func (MstValueComposition) TableName() string {
	return "mst_value_compositions"
}

type ValueCompositionGroup struct {
	RowNo               int64   `gorm:"column:row_no"`
	AcademicPeriodeID   string  `gorm:"column:academic_periode_id"`
	AcademicPeriodeName string  `gorm:"column:academic_periode_name"`
	TotalComponents     int     `gorm:"column:total_components"`
	TotalPercentage     float64 `gorm:"column:total_percentage"`
	IsValid             int     `gorm:"column:is_valid"`
}

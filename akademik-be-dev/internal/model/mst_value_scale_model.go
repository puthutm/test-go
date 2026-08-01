package model

import (
	"github.com/google/uuid"
)

type MstValueScale struct {
	ID             uuid.UUID  `gorm:"type:char(36);column:id;primaryKey"`
	StudyProgramID uuid.UUID  `gorm:"type:char(36);column:study_program_id"`
	GradeID        uuid.UUID  `gorm:"type:char(36);column:grade_id"`
	WeightValue    float64    `gorm:"type:decimal(3,2);column:weight_value"`
	LowerValue     float64    `gorm:"type:decimal(3,2);column:lower_value"`
	UpperValue     float64    `gorm:"type:decimal(3,2);column:upper_value"`
	Description    *string    `gorm:"type:nvarchar(max);column:description"`
	CreatedAt      int64      `gorm:"type:bigint;column:created_at"`
	CreatedBy      *uuid.UUID `gorm:"type:char(36);column:created_by"`
	UpdatedAt      *int64     `gorm:"type:bigint;column:updated_at"`
	UpdatedBy      *uuid.UUID `gorm:"type:char(36);column:updated_by"`
	DeletedAt      *int64     `gorm:"type:bigint;column:deleted_at"`
	DeletedBy      *uuid.UUID `gorm:"type:char(36);column:deleted_by"`

	// Additional Fields
	StudyProgramName string `gorm:"column:study_program_name"`
	GradeName        string `gorm:"column:grade_name"`
}

func (MstValueScale) TableName() string {
	return "mst_value_scales"
}

package model

import "github.com/google/uuid"

type MstAssessmentWeight struct {
	ID                         uuid.UUID `gorm:"type:char(36);column:id;primaryKey"`
	AttitudeBehaviorPercentage float64   `gorm:"column:attitude_behavior_percentage;type:decimal(5,2)"`
	TaskPercentage             float64   `gorm:"column:task_percentage;type:decimal(5,2)"`
	UTSPercentage              float64   `gorm:"column:uts_percentage;type:decimal(5,2)"`
	UASPercentage              float64   `gorm:"column:uas_percentage;type:decimal(5,2)"`
}

func (MstAssessmentWeight) TableName() string {
	return "mst_assessment_weights"
}

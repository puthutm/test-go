// Package model
package model

import "github.com/google/uuid"

type TrxStudentPresenceSettingDetail struct {
	StudentPresenceSettingID uuid.UUID `gorm:"column:student_presence_setting_id"`
	ClassID                  uuid.UUID `gorm:"column:class_id"`
}

func (TrxStudentPresenceSettingDetail) TableName() string {
	return "trx_student_presence_setting_details"
}

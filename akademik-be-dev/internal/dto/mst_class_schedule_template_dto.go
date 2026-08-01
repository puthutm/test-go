package dto

import "github.com/google/uuid"

type MstClassScheduleTemplateCreateRequest struct {
	ClassID       uuid.UUID `json:"-"`
	DayName       string    `json:"day_name" validate:"required,dayNameEN"`
	StartTime     string    `json:"start_time" validate:"required,datetime=15:04:05"`
	EndTime       string    `json:"end_time" validate:"required,datetime=15:04:05"`
	TypeOfMeeting string    `json:"type_of_meeting" validate:"required"`
}

type MstClassScheduleTemplateUpdateRequest struct {
	ID            uuid.UUID `json:"-"`
	ClassID       uuid.UUID `json:"-"`
	DayName       string    `json:"day_name" validate:"required,dayNameEN"`
	StartTime     string    `json:"start_time" validate:"required,datetime=15:04:05"`
	EndTime       string    `json:"end_time" validate:"required,datetime=15:04:05"`
	TypeOfMeeting string    `json:"type_of_meeting" validate:"required"`
}

type MstClassScheduleTemplateResponse struct {
	ID            uuid.UUID `json:"id"`
	ClassID       uuid.UUID `json:"class_id"`
	DayName       string    `json:"day_name"`
	StartTime     string    `json:"start_time"`
	EndTime       string    `json:"end_time"`
	TypeOfMeeting string    `json:"type_of_meeting"`

	CreatedAt int64      `json:"created_at"`
	CreatedBy *uuid.UUID `json:"created_by"`
	UpdatedAt *int64     `json:"updated_at"`
	UpdatedBy *uuid.UUID `json:"updated_by"`
}

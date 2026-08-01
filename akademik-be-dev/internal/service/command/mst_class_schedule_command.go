package command

import "github.com/google/uuid"

type MstClassScheduleUpdateCommand struct {
	ID            uuid.UUID
	ClassID       uuid.UUID
	DayNameOld    string
	StartTimeOld  string
	EndTimeOld    string
	DayName       string
	Date          string
	StartTime     string
	EndTime       string
	TypeOfMeeting string
}

type MstClassScheduleGetByCommand struct {
	ID        uuid.UUID
	ClassID   uuid.UUID
	DayName   string
	StartTime string
	EndTime   string
}

package models

import (
	"time"
)

type DayOfWeek int

type WorkingHours struct {
	ID        int       `json:"id" validate:"required"`
	Day       DayOfWeek `json:"day" validate:"required,min=0,max=6"`
	OpenTime  time.Time `json:"open_time" validate:"required"`
	CloseTime time.Time `json:"close_time" validate:"required,gtfield=Open"`
}

func NewWorkingHours(
	day DayOfWeek, open, close time.Time) *WorkingHours {

	return &WorkingHours{
		Day:       day,
		OpenTime:  open,
		CloseTime: close,
	}
}

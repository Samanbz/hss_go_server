package base

import "time"

type DayOfWeek uint

const (
	Monday DayOfWeek = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

type WorkingHours struct {
	Day   DayOfWeek  `json:"day" validate:"required,min=0,max=6"`
	Open  *time.Time `json:"open" validate:"required"`
	Close *time.Time `json:"close" validate:"required,gtfield=Open"`
}

func NewWorkingHourss(
	mondayOpen, mondayClose,
	tuesdayOpen, tuesdayClose,
	wednesdayOpen, wednesdayClose,
	thursdayOpen, thursdayClose,
	fridayOpen, fridayClose,
	saturdayOpen, saturdayClose,
	sundayOpen, sundayClose *time.Time) *[7]*WorkingHours {

	var workingHours [7]*WorkingHours
	workingHours[0] = &WorkingHours{Monday, mondayOpen, mondayClose}
	workingHours[1] = &WorkingHours{Tuesday, tuesdayOpen, tuesdayClose}
	workingHours[2] = &WorkingHours{Wednesday, wednesdayOpen, wednesdayClose}
	workingHours[3] = &WorkingHours{Thursday, thursdayOpen, thursdayClose}
	workingHours[4] = &WorkingHours{Friday, fridayOpen, fridayClose}
	workingHours[5] = &WorkingHours{Saturday, saturdayOpen, saturdayClose}
	workingHours[6] = &WorkingHours{Sunday, sundayOpen, sundayClose}

	return &workingHours
}

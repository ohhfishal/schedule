package recurrence

import (
	"fmt"
	"time"
)

type Frequency string
type WeekDay string

const (
	SUNDAY    = "SU"
	MONDAY    = "MO"
	TUESDAY   = "TU"
	WEDNESDAY = "WE"
	THURSDAY  = "TH"
	FRIDAY    = "FR"
	SATURDAY  = "SA"
)

const (
	MINUTELY = "MINUTELY"
	HOURLY   = "HOURLY"
	DAILY    = "DAILY"
	WEEKLY   = "WEEKLY"
	MONTHLY  = "MONTHLY"
	YEARLY   = "YEARLY"
)

var weekDays = map[WeekDay]string{
	SUNDAY:    "Sunday",
	MONDAY:    "Monday",
	TUESDAY:   "Tuesday",
	WEDNESDAY: "Wednesday",
	THURSDAY:  "Thursday",
	FRIDAY:    "Friday",
	SATURDAY:  "Saturday",
}

var frequencies = map[Frequency]time.Duration{
	MINUTELY: time.Minute,
	HOURLY:   time.Hour,
	DAILY:    DAY,
	WEEKLY:   WEEK,
	MONTHLY:  DAY * 28, // TODO: Months are weird...
	YEARLY:   DAY * 365,
}

func (f Frequency) Valid() error {
	_, ok := frequencies[f]
	if !ok {
		return fmt.Errorf(`invalid frequency: %s`, f)
	}
	return nil
}

func (day WeekDay) Valid() error {
	_, ok := weekDays[day]
	if !ok {
		return fmt.Errorf(`invalid week day: %s`, day)
	}
	return nil
}

package recurrence

import (
	"fmt"
	"slices"
	"time"
)

type ByFilter uint8
type WeekDay string
type Frequency string

//go:generate stringer -type=ByFilter

const (
	BYSETPOS ByFilter = 1 << iota
	BYWEEKNO
	BYYEARDAY
	BYMONTHDAY
	BYHOUR
	BYMONTH
	BYDAY
	BYMINUTE
)

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

var filters []ByFilter = []ByFilter{
	BYSETPOS,
	BYWEEKNO,
	BYYEARDAY,
	BYMONTHDAY,
	BYHOUR,
	BYMONTH,
	BYDAY,
	BYMINUTE,
}

var weekDays = map[WeekDay]time.Weekday{
	SUNDAY:    time.Sunday,
	MONDAY:    time.Monday,
	TUESDAY:   time.Tuesday,
	WEDNESDAY: time.Wednesday,
	THURSDAY:  time.Thursday,
	FRIDAY:    time.Friday,
	SATURDAY:  time.Saturday,
}

var frequencies = map[Frequency]time.Duration{
	MINUTELY: time.Minute,
	HOURLY:   time.Hour,
	DAILY:    DAY,
	WEEKLY:   WEEK,
	MONTHLY:  DAY * 28, // TODO: Months are weird...
	YEARLY:   DAY * 365,
}

func (filter ByFilter) Valid() error {
	ok := slices.Contains(filters, filter)
	if !ok {
		return fmt.Errorf(`invalid filter: %s`, filter)
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

func (f Frequency) Valid() error {
	_, ok := frequencies[f]
	if !ok {
		return fmt.Errorf(`invalid frequency: %s`, f)
	}
	return nil
}

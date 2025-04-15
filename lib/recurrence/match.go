package recurrence

import (
	"errors"
	"fmt"
	"slices"
	"time"
)

type Match func(time.Time) error

// One for each
// [X] BYHOUR
// [X] BYMINUTE // TODO Test
// [ ] BYMONTH // Look up table to time.Month
// [ ] BYMONTHDAY
// [ ] BYDAY
// [ ] BYYEARDAY
// --- Harder ones
// [ ] BYSETPOS ByFilter = 1 << iota
// [ ] BYWEEKNO

func NewByHour(initial []int) (Match, error) {
	if len(initial) == 0 {
		return nil, errors.New(`must include at least one hour to match`)
	}
	for _, hour := range initial {
		if hour < 0 || hour >= 24 {
			return nil, fmt.Errorf(`invalid hour: %d`, hour)
		}
	}
	hours := slices.Clone(initial)
	return Match(func(date time.Time) error {
		hour := date.Hour()
		if !slices.Contains(hours, hour) {
			return fmt.Errorf(`BYHOUR: hour %d not included in %v`, hour, hours)
		}
		return nil
	}), nil
}

func NewByMinute(initial []int) (Match, error) {
	if len(initial) == 0 {
		return nil, errors.New(`must include at least one hour to match`)
	}
	for _, minute := range initial {
		if minute < 0 || minute >= 60 {
			return nil, fmt.Errorf(`invalid minute: %d`, minute)
		}
	}
	minutes := slices.Clone(initial)
	return Match(func(date time.Time) error {
		minute := date.Minute()
		if !slices.Contains(minutes, minute) {
			return fmt.Errorf(`BYMINUTE: minute %d not included in %v`, minute, minutes)
		}
		return nil
	}), nil
}

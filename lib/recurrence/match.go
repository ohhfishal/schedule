package recurrence

import (
	"errors"
	"fmt"
	"math"
	"slices"
	"time"
)

type Match func(time.Time) error

// One for each
// [X] BYHOUR
// [X] BYMINUTE
// [X] BYMONTH
// [X] BYMONTHDAY
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

func NewByMonth(initial []int) (Match, error) {
	if len(initial) == 0 {
		return nil, errors.New(`must include at least one hour to match`)
	}

	months := []time.Month{}
	for _, month := range initial {
		if month < 1 || month > 12 {
			return nil, fmt.Errorf(`invalid month: %d`, month)
		}
		months = append(months, (time.Month)(month))
	}

	return Match(func(date time.Time) error {
		month := date.Month()
		if !slices.Contains(months, month) {
			return fmt.Errorf(`BYMONTH: month %d not included in %v`, month, months)
		}
		return nil
	}), nil
}

func NewByMonthDay(initial []int) (Match, error) {
	if len(initial) == 0 {
		return nil, errors.New(`must include at least one hour to match`)
	}

	days := []int{}
	for _, day := range initial {
		if day == 0 || day > 31 || day < -31 {
			return nil, fmt.Errorf(`invalid month: %d`, day)
		}
		days = append(days, day)
	}

	return Match(func(date time.Time) error {
		day := date.Day()
		daysInMonth := daysIn(date)

		helper := func(match int) bool {
			return (day == match) || (match < 0 && (daysInMonth)+match+1 == day) ||
				(match > daysInMonth && day == daysInMonth) ||
				(int(math.Abs(float64(match))) >= daysInMonth && day == 1)
		}

		if !slices.ContainsFunc(days, helper) {
			return fmt.Errorf(`BYMONTHDAY: day %d not included in %v`, day, days)
		}
		return nil
	}), nil
}

func NewByDay(initial []ByDay) (Match, error) {
	if len(initial) == 0 {
		return nil, errors.New(`must include at least one day to match`)
	}

	entries := []ByDay{}
	for _, entry := range initial {
		// TODO: Validate input
		entries = append(entries, entry)
	}

	return Match(func(date time.Time) error {
		day := date.Weekday()
		for _, entry := range entries {
			goDay, ok := weekDays[entry.Day]
			if !ok {
				return fmt.Errorf(`invalid day: %v`, entry.Day)
			}
			if day == goDay {
				return nil
			}
			// TODO: Handle cases formatted as: -1MU
		}
		// TODO: Better format entries?
		return fmt.Errorf(`BYDAY: day %d not included in %v`, day, entries)
	}), nil
}

// From: https://cs.opensource.google/go/go/+/refs/tags/go1.24.2:src/time/time.go;l=1679-1689
func isLeap(year int) bool {
	// year%4 == 0 && (year%100 != 0 || year%400 == 0)
	// Bottom 2 bits must be clear.
	// For multiples of 25, bottom 4 bits must be clear.
	// Thanks to Cassio Neri for this trick.
	mask := 0xf
	if year%25 != 0 {
		mask = 3
	}
	return year&mask == 0
}

// From:https://cs.opensource.google/go/go/+/refs/tags/go1.24.2:src/time/time.go;l=1285-1298s://cs.opensource.google/go/go/+/refs/tags/go1.24.2:src/time/time.go;l=1285-1298
func daysIn(t time.Time) int {
	month := t.Month()
	if month == time.February {
		if isLeap(t.Year()) {
			return 29
		}
		return 28
	}
	// With the special case of February eliminated, the pattern is
	//	31 30 31 30 31 30 31 31 30 31 30 31
	// Adding m&1 produces the basic alternation;
	// adding (m>>3)&1 inverts the alternation starting in August.
	return 30 + int((month+month>>3)&1)
}

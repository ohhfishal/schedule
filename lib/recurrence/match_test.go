package recurrence_test

import (
	"testing"
	"time"

	assert "github.com/alecthomas/assert/v2"
	"github.com/ohhfishal/schedule/lib/recurrence"
)

var NewYears = ThisYear()
var MidnightToday = Midnight()
var now = time.Now()
var LeapYear = 2024

// {RRule: "RRULE:FREQ=DAILY;BYHOUR=9,10,11,12,13,14,15,16;BYMINUTE=0,20,40"},

func TestByMonth(t *testing.T) {
	tests := []struct {
		Name       string
		Months     []int
		BuildFails bool
		Matches    []time.Time
		NoMatches  []time.Time
	}{
		{
			Name:       `empty input`,
			Months:     []int{},
			BuildFails: true,
		},
		{
			Name:       `invalid input fails build`,
			Months:     []int{-1},
			BuildFails: true,
		},
		{
			Name:   `valid examples`,
			Months: []int{1, 3, 5, 7, 9, 11},
			Matches: []time.Time{
				time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location()),
				time.Date(now.Year(), 1, 2, 0, 0, 0, 0, now.Location()),
				time.Date(now.Year(), 3, 14, 1, 5, 9, 2, now.Location()),
			},
			NoMatches: []time.Time{
				time.Date(now.Year(), 2, 1, 0, 0, 0, 0, now.Location()),
				time.Date(now.Year(), 6, 15, 0, 0, 0, 0, now.Location()),
			},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			matcher, err := recurrence.NewByMonth(test.Months)
			if test.BuildFails {
				assert.Error(t, err, `expected build to fail`)
				return
			}
			assert.NoError(t, err, `expected build to succeed`)

			for _, match := range test.Matches {
				err := matcher(match)
				assert.NoError(t, err, `expected %v to match`, match)
			}

			for _, noMatch := range test.NoMatches {
				err := matcher(noMatch)
				assert.Error(t, err, `expected %v to not match`, noMatch)
			}
		})
	}
}

func TestByMonthDay(t *testing.T) {
	tests := []struct {
		Name       string
		Days       []int
		BuildFails bool
		Matches    []time.Time
		NoMatches  []time.Time
	}{
		{
			Name:       `empty input`,
			Days:       []int{},
			BuildFails: true,
		},
		{
			Name:       `invalid input fails build`,
			Days:       []int{-32},
			BuildFails: true,
		},
		{
			Name: `valid examples`,
			Days: []int{1, 3, 5, 7, 9, 11, 13},
			Matches: []time.Time{
				time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, now.Location()),
				time.Date(now.Year(), time.March, 3, 0, 0, 0, 0, now.Location()),
				time.Date(now.Year(), time.April, 5, 0, 0, 0, 0, now.Location()),
				time.Date(now.Year(), time.May, 7, 0, 0, 0, 0, now.Location()),
				time.Date(now.Year(), time.December, 9, 0, 0, 0, 0, now.Location()),
			},
			NoMatches: []time.Time{
				time.Date(now.Year(), time.January, 2, 0, 0, 0, 0, now.Location()),
				time.Date(now.Year(), time.March, 4, 0, 0, 0, 0, now.Location()),
				time.Date(now.Year(), time.April, 6, 0, 0, 0, 0, now.Location()),
				time.Date(now.Year(), time.May, 8, 0, 0, 0, 0, now.Location()),
				time.Date(now.Year(), time.December, 10, 0, 0, 0, 0, now.Location()),
			},
		},
		{
			Name: `clamping positive days`,
			Days: []int{30},
			Matches: []time.Time{
				// Not a leap year
				time.Date(2025, time.February, 28, 0, 0, 0, 0, now.Location()),
				time.Date(LeapYear, time.February, 29, 0, 0, 0, 0, now.Location()),
			},
			NoMatches: []time.Time{
				time.Date(now.Year(), time.January, 31, 0, 0, 0, 0, now.Location()),
			},
		},
		{
			Name: `clamping negative days`,
			Days: []int{-30},
			Matches: []time.Time{
				time.Date(2025, time.February, 1, 0, 0, 0, 0, now.Location()),
				time.Date(LeapYear, time.February, 1, 0, 0, 0, 0, now.Location()),
				time.Date(LeapYear, time.April, 1, 0, 0, 0, 0, now.Location()),
				time.Date(now.Year(), time.January, 2, 0, 0, 0, 0, now.Location()),
			},
		},
		{
			Name: `clamping negative days/2`,
			Days: []int{-6},
			Matches: []time.Time{
				time.Date(2025, time.February, 23, 0, 0, 0, 0, now.Location()),
				time.Date(LeapYear, time.February, 24, 0, 0, 0, 0, now.Location()),
			},
		},
		{
			Name: `negative days`,
			Days: []int{-1},
			Matches: []time.Time{
				// Not a leap year
				time.Date(2025, time.February, 28, 0, 0, 0, 0, now.Location()),
				// Leap year
				time.Date(LeapYear, time.February, 29, 0, 0, 0, 0, now.Location()),
				time.Date(now.Year(), time.January, 31, 0, 0, 0, 0, now.Location()),
			},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			matcher, err := recurrence.NewByMonthDay(test.Days)
			if test.BuildFails {
				assert.Error(t, err, `expected build to fail`)
				return
			}
			assert.NoError(t, err, `expected build to succeed`)

			for _, match := range test.Matches {
				err := matcher(match)
				assert.NoError(t, err, `expected %v to match`, match)
			}

			for _, noMatch := range test.NoMatches {
				err := matcher(noMatch)
				assert.Error(t, err, `expected %v to not match`, noMatch)
			}
		})
	}
}

func TestByDay(t *testing.T) {
	tests := []struct {
		Name       string
		Days       []recurrence.ByDay
		BuildFails bool
		Matches    []time.Time
		NoMatches  []time.Time
	}{
		{
			Name:       `empty input`,
			BuildFails: true,
		},
		{
			Name: `valid usecase (no ints)`,
			Days: []recurrence.ByDay{
				{Day: recurrence.SUNDAY},
				{Day: recurrence.TUESDAY},
				{Day: recurrence.WEDNESDAY},
			},
			Matches: []time.Time{
				time.Date(2025, time.January, 5, 0, 0, 0, 0, now.Location()),
				time.Date(2025, time.January, 7, 0, 0, 0, 0, now.Location()),
				time.Date(2025, time.January, 8, 0, 0, 0, 0, now.Location()),
			},
			NoMatches: []time.Time{
				time.Date(2025, time.January, 3, 0, 0, 0, 0, now.Location()),
				time.Date(2025, time.January, 4, 0, 0, 0, 0, now.Location()),
				time.Date(2025, time.January, 9, 0, 0, 0, 0, now.Location()),
			},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			matcher, err := recurrence.NewByDay(test.Days)
			if test.BuildFails {
				assert.Error(t, err, `expected build to fail`)
				return
			}
			assert.NoError(t, err, `expected build to succeed`)

			for _, match := range test.Matches {
				err := matcher(match)
				assert.NoError(t, err, `expected %v to match`, match)
			}

			for _, noMatch := range test.NoMatches {
				err := matcher(noMatch)
				assert.Error(t, err, `expected %v to not match`, noMatch)
			}
		})
	}
}

func TestByYearDay(t *testing.T) {
	tests := []struct {
		Name       string
		Days       []int
		BuildFails bool
		Matches    []time.Time
		NoMatches  []time.Time
	}{
		{
			Name:       `empty input`,
			Days:       []int{},
			BuildFails: true,
		},
		{
			Name:       `invalid input fails build`,
			Days:       []int{-1},
			BuildFails: true,
		},
		{
			Name: `valid examples`,
			Days: []int{1, 31, 62, 365},
			Matches: []time.Time{
				time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, now.Location()),
				time.Date(now.Year(), time.January, 31, 0, 0, 0, 0, now.Location()),
				time.Date(now.Year(), time.January, 62, 0, 0, 0, 0, now.Location()),
				time.Date(now.Year(), time.January, 365, 0, 0, 0, 0, now.Location()),
			},
			NoMatches: []time.Time{
				time.Date(now.Year(), time.January, 2, 0, 0, 0, 0, now.Location()),
				time.Date(now.Year(), time.March, 4, 0, 0, 0, 0, now.Location()),
				time.Date(now.Year(), time.April, 6, 0, 0, 0, 0, now.Location()),
				time.Date(now.Year(), time.May, 8, 0, 0, 0, 0, now.Location()),
				time.Date(now.Year(), time.December, 10, 0, 0, 0, 0, now.Location()),
			},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			matcher, err := recurrence.NewByYearDay(test.Days)
			if test.BuildFails {
				assert.Error(t, err, `expected build to fail`)
				return
			}
			assert.NoError(t, err, `expected build to succeed`)

			for _, match := range test.Matches {
				err := matcher(match)
				assert.NoError(t, err, `expected %v to match`, match)
			}

			for _, noMatch := range test.NoMatches {
				err := matcher(noMatch)
				assert.Error(t, err, `expected %v to not match`, noMatch)
			}
		})
	}
}

func TestByHour(t *testing.T) {
	tests := []struct {
		Name       string
		Hours      []int
		BuildFails bool
		Matches    []time.Time
		NoMatches  []time.Time
	}{
		{
			Name:       `empty input`,
			Hours:      []int{},
			BuildFails: true,
		},
		{
			Name:       `invalid input fails build`,
			Hours:      []int{-1},
			BuildFails: true,
		},
		{
			Name:  `valid input build`,
			Hours: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23},
			Matches: []time.Time{
				MidnightToday.Add(time.Hour * 0),
				MidnightToday.Add(time.Hour * 1),
				MidnightToday.Add(time.Hour * 2),
				MidnightToday.Add(time.Hour * 3),
				MidnightToday.Add(time.Hour * 4),
				MidnightToday.Add(time.Hour * 5),
				MidnightToday.Add(time.Hour * 6),
				MidnightToday.Add(time.Hour * 7),
				MidnightToday.Add(time.Hour * 8),
				MidnightToday.Add(time.Hour * 9),
				MidnightToday.Add(time.Hour * 10),
				MidnightToday.Add(time.Hour * 11),
				MidnightToday.Add(time.Hour * 12),
				MidnightToday.Add(time.Hour * 13),
				MidnightToday.Add(time.Hour * 14),
				MidnightToday.Add(time.Hour * 15),
				MidnightToday.Add(time.Hour * 16),
				MidnightToday.Add(time.Hour * 17),
				MidnightToday.Add(time.Hour * 18),
				MidnightToday.Add(time.Hour * 19),
				MidnightToday.Add(time.Hour * 20),
				MidnightToday.Add(time.Hour * 21),
				MidnightToday.Add(time.Hour * 22),
				MidnightToday.Add(time.Hour * 23),
			},
		},
		{
			Name:  `user example`,
			Hours: []int{0, 2, 4, 6},
			Matches: []time.Time{
				MidnightToday.Add(time.Hour * 0),
				MidnightToday.Add(time.Hour*2 + time.Minute*30),
				MidnightToday.Add(time.Hour*4 + time.Minute*30 + time.Second*45),
				MidnightToday.Add(time.Hour * 6),
			},
			NoMatches: []time.Time{
				MidnightToday.Add(time.Hour * 1),
				MidnightToday.Add(time.Hour*3 + time.Minute*30),
				MidnightToday.Add(time.Hour*5 + time.Minute*30 + time.Second*45),
				MidnightToday.Add(time.Hour * 5),
				MidnightToday.Add(time.Hour * 7),
			},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			matcher, err := recurrence.NewByHour(test.Hours)
			if test.BuildFails {
				assert.Error(t, err, `expected build to fail`)
				return
			}
			assert.NoError(t, err, `expected build to succeed`)

			for _, match := range test.Matches {
				err := matcher(match)
				assert.NoError(t, err, `expected %v to match`, match)
			}

			for _, noMatch := range test.NoMatches {
				err := matcher(noMatch)
				assert.Error(t, err, `expected %v to not match`, noMatch)
			}
		})
	}
}

func TestByMinute(t *testing.T) {
	tests := []struct {
		Name       string
		Minutes    []int
		BuildFails bool
		Matches    []time.Time
		NoMatches  []time.Time
	}{
		{
			Name:       `empty input`,
			Minutes:    []int{},
			BuildFails: true,
		},
		{
			Name:       `invalid input fails build`,
			Minutes:    []int{-1},
			BuildFails: true,
		},
		{
			Name:    `valid input build`,
			Minutes: []int{0, 5, 10},
			Matches: []time.Time{
				MidnightToday.Add(time.Minute * 0),
				MidnightToday.Add(time.Minute * 5),
				MidnightToday.Add(time.Hour + time.Minute*10),
			},
			NoMatches: []time.Time{
				MidnightToday.Add(time.Hour*1 + time.Minute*3),
				MidnightToday.Add(time.Hour*3 + time.Minute*30),
				MidnightToday.Add(time.Hour*5 + time.Minute*30 + time.Second*45),
			},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			matcher, err := recurrence.NewByMinute(test.Minutes)
			if test.BuildFails {
				assert.Error(t, err, `expected build to fail`)
				return
			}
			assert.NoError(t, err, `expected build to succeed`)

			for _, match := range test.Matches {
				err := matcher(match)
				assert.NoError(t, err, `expected %v to match`, match)
			}

			for _, noMatch := range test.NoMatches {
				err := matcher(noMatch)
				assert.Error(t, err, `expected %v to not match`, noMatch)
			}
		})
	}
}

func ThisYear() time.Time {
	now := time.Now()
	return time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
}

func Midnight() time.Time {
	now := time.Now()
	return time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0, 0, 0, 0,
		now.Location(),
	)
}

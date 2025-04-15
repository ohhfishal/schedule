package recurrence_test

import (
	"testing"
	"time"

	assert "github.com/alecthomas/assert/v2"
	"github.com/ohhfishal/schedule/lib/recurrence"
)

func TestByHour(t *testing.T) {
	tests := []struct {
		Name       string
		Hours      []int
		BuildFails bool
		Matches    []time.Time
	}{
		{
			Name:       `empty input`,
			Hours:      []int{},
			BuildFails: true,
		},
		{
			Name:  `valid input build`,
			Hours: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
		},
		// {RRule: "RRULE:FREQ=DAILY;BYHOUR=9,10,11,12,13,14,15,16;BYMINUTE=0,20,40"},
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
		})
	}
}

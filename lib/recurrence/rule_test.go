package recurrence_test

import (
	"fmt"
	"iter"
	"testing"
	"time"

	assert "github.com/alecthomas/assert/v2"
	"github.com/ohhfishal/schedule/lib/recurrence"
)

var testTime = time.Date(2006, time.January, 2, 0, 0, 0, 0, now.Location())

func TestFull(t *testing.T) {
	var tests = []struct {
		RRule      string
		Times      []time.Time
		Terminates bool
	}{
		{RRule: "RRULE:FREQ=DAILY", Times: []time.Time{
			time.Date(2006, time.January, 2, 0, 0, 0, 0, now.Location()),
			time.Date(2006, time.January, 3, 0, 0, 0, 0, now.Location()),
			time.Date(2006, time.January, 4, 0, 0, 0, 0, now.Location()),
			time.Date(2006, time.January, 5, 0, 0, 0, 0, now.Location()),
			time.Date(2006, time.January, 6, 0, 0, 0, 0, now.Location()),
		}},
		{RRule: "RRULE:FREQ=DAILY;INTERVAL=2", Times: []time.Time{
			time.Date(2006, time.January, 2, 0, 0, 0, 0, now.Location()),
			time.Date(2006, time.January, 4, 0, 0, 0, 0, now.Location()),
			time.Date(2006, time.January, 6, 0, 0, 0, 0, now.Location()),
			time.Date(2006, time.January, 8, 0, 0, 0, 0, now.Location()),
			time.Date(2006, time.January, 10, 0, 0, 0, 0, now.Location()),
		}},
		{RRule: "RRULE:FREQ=DAILY;COUNT=10", Times: []time.Time{
			time.Date(2006, time.January, 2, 0, 0, 0, 0, now.Location()),
			time.Date(2006, time.January, 3, 0, 0, 0, 0, now.Location()),
			time.Date(2006, time.January, 4, 0, 0, 0, 0, now.Location()),
			time.Date(2006, time.January, 5, 0, 0, 0, 0, now.Location()),
			time.Date(2006, time.January, 6, 0, 0, 0, 0, now.Location()),
			time.Date(2006, time.January, 7, 0, 0, 0, 0, now.Location()),
			time.Date(2006, time.January, 8, 0, 0, 0, 0, now.Location()),
			time.Date(2006, time.January, 9, 0, 0, 0, 0, now.Location()),
			time.Date(2006, time.January, 10, 0, 0, 0, 0, now.Location()),
			time.Date(2006, time.January, 11, 0, 0, 0, 0, now.Location()),
		}, Terminates: true},
		{RRule: "RRULE:FREQ=DAILY;UNTIL=19971224T000000Z", Times: []time.Time{}, Terminates: true},
		{RRule: "RRULE:FREQ=DAILY;COUNT=3;BYDAY=MO,WE,TH", Times: []time.Time{
			time.Date(2006, time.January, 2, 0, 0, 0, 0, now.Location()),
			time.Date(2006, time.January, 4, 0, 0, 0, 0, now.Location()),
			time.Date(2006, time.January, 5, 0, 0, 0, 0, now.Location()),
		}, Terminates: true},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d/%s", i, test.RRule), func(t *testing.T) {
			rule, err := recurrence.ParseRRule(test.RRule)
			assert.NoError(t, err, "Failed to parse RRULE")

			iterator, err := rule.Iter(testTime)
			assert.NoError(t, err, "Could not create iterator")
			index := 0
			next, stop := iter.Pull[time.Time](iterator)
			defer stop()
			for i, expected := range test.Times {
				result, _ := next()
				t.Logf(`result: %v`, result)
				assert.Equal(t, result, expected, fmt.Sprintf("Yielded incorrect time (%d)", i))
				index++
			}

			_, ok := next()
			if test.Terminates {
				assert.Equal(t, false, ok)
				assert.Equal(t, len(test.Times), index, "Differing number of elements yielded")
			} else {
				assert.Equal(t, true, ok)
			}
		})
	}
}

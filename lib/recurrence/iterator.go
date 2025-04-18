package recurrence

import (
	"errors"
	"time"
)

const _MAX_ITERATIONS = 100

type ruleIter struct {
	Rule  Rule
	Start time.Time

	cursor time.Time
	count  int
}

func (iter *ruleIter) Next() (time.Time, error) {
	// First iteration case
	if iter.cursor.IsZero() && iter.count == 0 {
		iter.cursor = iter.Start
	}

	for range _MAX_ITERATIONS {
		// Base cases
		if iter.Rule.Count != NONE && iter.count >= iter.Rule.Count {
			return time.Time{}, errors.New(`done: reached count`)
		} else if !iter.Rule.Until.IsZero() && iter.cursor.After(iter.Rule.Until) {
			return time.Time{}, errors.New(`done: after until`)
		}
		cur := iter.cursor

		// TODO: This does not work for months since they vary based on the month
		frequency := frequencies[iter.Rule.Frequency]
		delta := frequency
		if iter.Rule.Interval != NONE {
			delta = delta * time.Duration(iter.Rule.Interval)
		}
		iter.cursor = iter.cursor.Add(delta)
		if err := contains(iter.Rule.By, cur); err == nil {
			iter.count++
			return cur, nil
		}
	}
	return time.Time{}, errors.New(`max iterations reached`)
}

func contains(matchers []Match, date time.Time) error {
	for _, matcher := range matchers {
		if err := matcher(date); err != nil {
			return err
		}
	}
	// TODO: Apply all the ByDay...
	// See: https://icalendar.org/iCalendar-RFC-5545/3-3-10-recurrence-rule.html
	// NOTES:
	// If multiple BYxxx rule parts are specified, then after evaluating the specified FREQ and INTERVAL rule parts, the BYxxx rule parts are applied to the current set of evaluated occurrences in the following order: BYMONTH, BYWEEKNO, BYYEARDAY, BYMONTHDAY, BYDAY, BYHOUR, BYMINUTE, BYSECOND and BYSETPOS; then COUNT and UNTIL are evaluated.
	return nil
}

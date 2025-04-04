package recurrence

import (
	"errors"
	"fmt"
	"time"

	"iter"
)

const NONE = -1
const DAY = time.Hour * 24
const WEEK = DAY * 7

type Match interface {
	Match(Rule, time.Time)
}

type Rule struct {
	Count     int       // Default NONE
	Frequency Frequency // Required
	Until     time.Time // Default zero
	Interval  int       // Default NONE
	WeekStart WeekDay   // Default Monday
	By        []Match   // Default empty
}

func DefaultRule() Rule {
	return Rule{
		Interval:  NONE,
		Count:     NONE,
		WeekStart: SUNDAY,
	}
}

func (r Rule) Valid() error {
	if r.Count < NONE {
		return fmt.Errorf(`count: %d invalid. Must be >= -1`, r.Interval)
	}
	if err := r.Frequency.Valid(); err != nil {
		return err
	}
	if r.Interval == 0 || r.Interval < NONE {
		return fmt.Errorf(`interval: %d invalid. Must be -1 or > 0`, r.Interval)
	}
	if err := r.WeekStart.Valid(); err != nil {
		return err
	}
	// TODO: Dependeds on how Match is implemented. If stays an interface it's already valid.
	return nil
}

func (r Rule) Terminates() bool {
	// TODO: Implement
	return true
}

func (r Rule) Iter(start time.Time) (iter.Seq[time.Time], error) {
	if err := r.Valid(); err != nil {
		return nil, fmt.Errorf(`invalid state: %w`, err)
	}
	// See: https://icalendar.org/iCalendar-RFC-5545/3-3-10-recurrence-rule.html
	// NOTES:
	// If multiple BYxxx rule parts are specified, then after evaluating the specified FREQ and INTERVAL rule parts, the BYxxx rule parts are applied to the current set of evaluated occurrences in the following order: BYMONTH, BYWEEKNO, BYYEARDAY, BYMONTHDAY, BYDAY, BYHOUR, BYMINUTE, BYSECOND and BYSETPOS; then COUNT and UNTIL are evaluated.
	curTime := start

	// TODO: This does not work for months since they vary based on the month
	frequency := frequencies[r.Frequency]
	delta := frequency
	if r.Interval != NONE {
		delta = delta * time.Duration(r.Interval)
	}

	count := 1
	return iter.Seq[time.Time](func(yield func(time.Time) bool) {
		for {
			if !r.Until.IsZero() && curTime.After(r.Until) {
				break
			}
			// TODO: Further test this seems wrong
			// TODO: Move this check to the end of the loop? IE Don't return start?
			//       This way does address the case where start > r.Until or count = 0
			if !yield(curTime) || (r.Count != NONE && count >= r.Count) {
				return
			}
			curTime = curTime.Add(delta)
			count++
			// TODO: Implement BYXXXX
		}

	}), nil
}

func (r *Rule) All(start time.Time) ([]time.Time, error) {
	if !r.Terminates() {
		return nil, errors.New(`infinite matches exist`)
	}
	iter, err := r.Iter(start)
	if err != nil {
		return nil, err
	}

	matches := []time.Time{}
	for t := range iter {
		matches = append(matches, t)
	}
	return matches, nil
}

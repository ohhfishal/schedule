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

type Rule struct {
	Count     int          // Default NONE
	Frequency Frequency    // Required
	Until     time.Time    // Default zero
	Interval  int          // Default NONE
	WeekStart time.Weekday // Default Monday
	By        []Match      // Default empty
}

func DefaultRule() Rule {
	return Rule{
		Interval:  NONE,
		Count:     NONE,
		WeekStart: time.Sunday,
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
	if r.WeekStart < time.Sunday || r.WeekStart > time.Saturday {
		return fmt.Errorf(`week start: %d invalid`, r.WeekStart)
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
	iterator := ruleIter{
		Rule:  r,
		Start: start,
	}
	return iter.Seq[time.Time](func(yield func(time.Time) bool) {
		// TODO: This assumes it is the right one...
		for {
			next, err := iterator.Next()
			if err != nil {
				return
			}

			if !yield(next) {
				return
			}
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

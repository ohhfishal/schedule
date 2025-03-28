package recurrence

import (
	"errors"
	"time"

	"iter"
)

type Frequency string
type Day string
type Match interface {
	Match(Rule, time.Time)
}

func DefaultRule() Rule { return Rule{} }

type Rule struct {
	Count     int
	Frequency Frequency
	Until     time.Time
	Interval  int
	WeekStart Day // (Enum w/default Monday)
	By        []Match
}

func (r Rule) Valid() error {
	// TODO: Implement
	return nil
}

func (r Rule) Terminates() bool {
	// TODO: Implement
	return true
}

func (r *Rule) Iter(start time.Time) iter.Seq[time.Time] {
	// TODO: Implement
	return nil
}

func (r *Rule) All(start time.Time) ([]time.Time, error) {
	if !r.Terminates() {
		return nil, errors.New(`infinite matches exist`)
	}
	matches := []time.Time{}
	iter := r.Iter(start)
	for t := range iter {
		matches = append(matches, t)
	}
	return matches, nil
}

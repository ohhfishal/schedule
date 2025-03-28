package recurrence

import (
	"time"

	"iter"
)

type Frequency string
type Day string
type Match string

type Rule struct {
	Frequency Frequency
	Count     int
	Until     time.Time
	Interval  int
	WKST      Day // (Enum w/default Monday)
	By        []Match
}

func (r *Rule) All(start time.Time) iter.Seq[Rule] {
	// TODO: Implement
	return nil
}

func (r Rule) RRule() string {
	// TODO: Implement
	return "IMPLEMENT"

}

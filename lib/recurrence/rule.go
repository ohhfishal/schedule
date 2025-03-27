package recurrence

import (
	"time"

	"iter"
)

type Frequency string

type Rule struct {
	Frequency Frequency
	Count     int
}

func (r *Rule) All(start time.Time) iter.Seq[Rule] {
	// TODO: Implement
	return nil
}

func (r Rule) RRule() string {
	// TODO: Implement
	return "IMPLEMENT"

}

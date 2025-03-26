package recurrence

import (
	"errors"
	"time"

	"iter"
)

type Rule struct {
}

func (r *Rule) All(start time.Time) iter.Seq[Rule] {
	// TODO: Implement
	return nil
}

func ParseRRule(input string) (Rule, error) {
	// TODO: Implement
	/*
	 rrule      = "RRULE" rrulparam ":" recur CRLF
	 rrulparam  = *(";" other-param)
	*/
	/* NOTE:
	- recur appears to be all the rruleparams
	- CLRF means the CR/end of line
	*/
	rule := Rule{}
	return rule, errors.New("Not Implemented")
}

func (r Rule) RRule() string {
	// TODO: Implement
	return "RRULE()"

}

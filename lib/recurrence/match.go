package recurrence

import "errors"

// One for each
// BYSETPOS ByFilter = 1 << iota
// BYWEEKNO
// BYYEARDAY
// BYMONTHDAY
// BYMONTH
// BYDAY
// BYMINUTE

// BYHOUR
// {RRule: "RRULE:FREQ=DAILY;BYHOUR=9,10,11,12,13,14,15,16;BYMINUTE=0,20,40"},
func NewByHour(hours []int) (Match, error) {
	// TODO: Validate all the hours
	// TODO: Use time.Time.Hour() to get hour
	return nil, errors.New(`not implemented`)
}

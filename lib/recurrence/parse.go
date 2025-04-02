package recurrence

import (
	"errors"
	"fmt"
	"sync"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var ErrBadValue = errors.New(`invalid value`)

var (
	ruleLexerInit sync.Once
	ruleLexer     lexer.Definition
	ruleLexerErr  error
)

func ParseRRule(input string) (*Rule, error) {
	lexer, err := getRuleLexer()
	if err != nil {
		return nil, fmt.Errorf("lexer: %w", err)
	}
	ruleParser, err := participle.Build[ruleGrammar](
		participle.Lexer(lexer),
	)
	if err != nil {
		return nil, fmt.Errorf("grammar: %w", err)
	}
	parsed, err := ruleParser.ParseString("", input)
	if err != nil {
		return nil, fmt.Errorf("format: %w", err)
	}
	return parsed.Rule()
}

type ruleGrammar struct {
	Parameters []parameter `parser:"'RRULE' ':' @@ (';' @@ )*"`
}

type parameter struct {
	Count      *int       `parser:"'COUNT' '=' @Int"`
	Freq       *Frequency `parser:"| ('FREQ' '=' @Frequency)"`
	Until      *string    `parser:"| ('UNTIL' '=' @Time)"`
	Interval   *int       `parser:"| ('INTERVAL' '=' @Int)"`
	WeekStart  *Day       `parser:"| ('WKST' '=' @Day)"`
	BySetPos   *int       `parser:"| ('BYSETPOS' '=' @Int)"`
	ByYearDay  *int       `parser:"| ('BYYEARDAY' '=' @Int (',' @Int)*)"`
	ByMonthDay *[]int     `parser:"| ('BYMONTHDAY' '=' @Int (',' @Int)*)"`
	ByWeekNo   *int       `parser:"| ('BYWEEKNO' '=' @Int)"` // TODO: Make a list?
	ByHour     *[]int     `parser:"| ('BYHOUR' '=' @Int (',' @Int)*)"`
	ByMonth    *[]int     `parser:"| ('BYMONTH' '=' @Int (',' @Int)*)"`
	ByDay      *ByDay     `parser:"| ('BYDAY' '=' @@ (',' @@)*)"`
	ByMinute   *[]int     `parser:"| ('BYMINUTE' '=' @Int (',' @Int)*)"`
}

func (rg ruleGrammar) Rule() (*Rule, error) {
	rule := DefaultRule()
	for _, p := range rg.Parameters {
		if err := p.Apply(&rule); err != nil {
			return nil, err
		}
	}
	if err := rule.Valid(); err != nil {
		return nil, fmt.Errorf(`%w state: %w`, ErrBadValue, err)
	}
	return &rule, nil
}

func (p parameter) Apply(rule *Rule) error {
	switch {
	case p.Count != nil:
		rule.Count = *p.Count
	case p.Freq != nil:
		rule.Frequency = *p.Freq
	case p.Interval != nil:
		rule.Interval = *p.Interval
	case p.WeekStart != nil:
		rule.WeekStart = *p.WeekStart
	case p.Until != nil:
		// TODO: Covert to time.Time
		// return errors.New("not implemented")
		fallthrough
	case p.BySetPos != nil:
		fallthrough
	case p.ByWeekNo != nil:
		fallthrough
	case p.ByYearDay != nil:
		fallthrough
	case p.ByMonthDay != nil:
		fallthrough
	case p.ByHour != nil:
		fallthrough
	case p.ByMonth != nil:
		fallthrough
	case p.ByDay != nil:
		fallthrough
	case p.ByMinute != nil:
		// return errors.New("not implemented")
		return nil
	default:
		// NOTE: This should never happen
		return errors.New(`no parameter set`)
	}
	return nil
}

type ByDay struct {
	Offset *int `parser:"@Int?"`
	Day    Day  `parser:"@Day"`
}

var _TIME_REGEX = `\d{8}T\d{6}Z`

func getRuleLexer() (lexer.Definition, error) {
	ruleLexerInit.Do(func() {
		ruleLexer, ruleLexerErr = lexer.NewSimple([]lexer.SimpleRule{
			{Name: "RRULE", Pattern: `RRULE`},
			{Name: "FREQ", Pattern: `FREQ`},
			{Name: "COUNT", Pattern: `COUNT`},
			{Name: "UNTIL", Pattern: `UNTIL`},
			{Name: "INTERVAL", Pattern: `INTERVAL`},
			{Name: "WKST", Pattern: `WKST`},
			{Name: "BYYEARDAY", Pattern: `BYYEARDAY`},
			{Name: "BYWEEKNO", Pattern: `BYWEEKNO`},
			{Name: "BYSETPOS", Pattern: `BYSETPOS`},
			{Name: "BYMONTHDAY", Pattern: `BYMONTHDAY`},
			{Name: "BYHOUR", Pattern: `BYHOUR`},
			{Name: "BYDAY", Pattern: `BYDAY`},
			{Name: "BYMONTH", Pattern: `BYMONTH`},
			{Name: "BYMINUTE", Pattern: `BYMINUTE`},
			{Name: ":", Pattern: `:`},
			{Name: ",", Pattern: `,`},
			{Name: ";", Pattern: `;`},
			{Name: "=", Pattern: `=`},
			{Name: "Time", Pattern: _TIME_REGEX},
			{Name: "Frequency", Pattern: `MINUTELY|HOURLY|DAILY|WEEKLY|MONTHLY|YEARLY`},
			{Name: "Day", Pattern: `SU|MO|TU|WE|TH|FR|SA`},
			{Name: "Int", Pattern: `-?\d+`},
		})
	})
	return ruleLexer, ruleLexerErr
}

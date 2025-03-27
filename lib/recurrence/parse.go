package recurrence

import (
	"fmt"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var _TIME_REGEX = `\d{8}T\d{6}Z`

var (
	ruleLexer = lexer.MustSimple([]lexer.SimpleRule{
		{Name: "RRULE", Pattern: `RRULE`},
		{Name: ":", Pattern: `:`},
		{Name: ";", Pattern: `;`},
		{Name: "Keyword", Pattern: `(?i)FREQ|COUNT|UNTIL|INTERVAL|BYMONTH`},
		{Name: "=", Pattern: `=`},
		{Name: "Time", Pattern: _TIME_REGEX},
		{Name: "Frequency", Pattern: `(?i)MINUTELY|HOURLY|DAILY|WEEKLY|MONTHLY|YEARLY`},
		{Name: "Int", Pattern: `\d+`},
	})
)

type ruleGrammar struct {
	Parameters []parameter `parser:"'RRULE' ':' @@ (';' @@ )*"`
}

func (rg ruleGrammar) Rule() (*Rule, error) {
	// TODO: Implement
	return nil, nil
}

type parameter struct {
	Name  string `parser:"@Keyword '='"`
	Value value  `parser:"@@"`
}

type value struct {
	// TODO: See if I can have a function to convert Time to time.Time
	Time      *string    `parser:"@Time"`
	Int       *int       `parser:"| @Int"`
	Frequency *Frequency `parser:"| @Frequency"`
}

func ParseRRule(input string) (*Rule, error) {
	/* Summarized from https://icalendar.org/iCalendar-RFC-5545/3-8-5-3-recurrence-rule.html
	rrule      = "RRULE" rrulparam ":" recur CRLF
	rrulparam  = *(";" other-param)
	*/
	ruleParser, err := participle.Build[ruleGrammar](
		participle.Lexer(ruleLexer),
		// Do we need to look ahead? Maybe for determining type of value?
		participle.UseLookahead(participle.MaxLookahead),
	)
	if err != nil {
		return nil, fmt.Errorf("bad grammar: %w", err)
	}

	parsed, err := ruleParser.ParseString("", input)
	if err != nil {
		return nil, fmt.Errorf("bad format: %w", err)
	}
	return parsed.Rule()
}

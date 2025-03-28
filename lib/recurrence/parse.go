package recurrence

import (
	"fmt"
	"sync"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var (
	ruleLexerInit sync.Once
	ruleLexer     lexer.Definition
	ruleLexerErr  error
)

func ParseRRule(input string) (*Rule, error) {
	lexer, err := getRuleLexer()
	if err != nil {
		return nil, fmt.Errorf("bad lexer: %w", err)
	}
	ruleParser, err := participle.Build[ruleGrammar](
		participle.Lexer(lexer),
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

func (rg ruleGrammar) Rule() (*Rule, error) {
	// TODO: Implement
	return nil, nil
}

type ruleGrammar struct {
	Parameters []parameter `parser:"'RRULE' ':' @@ (';' @@ )*"`
}

type parameter struct {
	Name  string  `parser:"(@Keyword | @Match) '='"`
	Value []value `parser:"@@ (',' @@)*"`
}

type value struct {
	// TODO: See if I can have a function to convert Time to time.Time
	Int       *int       `parser:"@Int?"`
	Day       *string    `parser:"@Day?"`
	Time      *string    `parser:"(@Time"`
	Frequency *Frequency `parser:"| @Frequency)?"`
}

var _TIME_REGEX = `\d{8}T\d{6}Z`

func getRuleLexer() (lexer.Definition, error) {
	ruleLexerInit.Do(func() {
		ruleLexer, ruleLexerErr = lexer.NewSimple([]lexer.SimpleRule{
			{Name: "RRULE", Pattern: `RRULE`},
			{Name: ":", Pattern: `:`},
			{Name: ",", Pattern: `,`},
			{Name: ";", Pattern: `;`},
			{Name: "Keyword", Pattern: `FREQ|COUNT|UNTIL|INTERVAL|WKST`},
			{Name: "Match", Pattern: `BYYEARDAY|BYWEEKNO|BYSETPOS|BYMONTHDAY|BYMINUTE|BYHOUR|BYDAY|BYMONTH`},
			{Name: "=", Pattern: `=`},
			{Name: "Time", Pattern: _TIME_REGEX},
			{Name: "Frequency", Pattern: `MINUTELY|HOURLY|DAILY|WEEKLY|MONTHLY|YEARLY`},
			{Name: "Day", Pattern: `SU|MO|TU|WE|TH|FR|SA`},
			{Name: "Int", Pattern: `-?\d+`},
		})
	})
	return ruleLexer, ruleLexerErr
}

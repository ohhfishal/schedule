package recurrence

import (
	"fmt"
	"time"

	"iter"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var (
	ruleLexer = lexer.MustSimple([]lexer.SimpleRule{
		// {Name: "Frequency", Pattern: `(?i)MINUTELY|HOURLY|DAILY|WEEKLY|MONTHLY|YEARLY`},
		{Name: "EOL", Pattern: `[\n\r]+`},
		{Name: "Int", Pattern: `\d+`},
		{Name: "Number", Pattern: `[-+]?(\d*\.)?\d+`},
		{Name: "Whitespace", Pattern: `[ \t]+`},
	})
)

type Frequency string

type Rule struct {
	Frequency Frequency `parser:"'RRULE'"`
	// Frequency Frequency `parser:"'RRULE:FREQ='@Frequency"`
	// Count int `parser:"';COUNT=' @Int"`
}

func (r *Rule) All(start time.Time) iter.Seq[Rule] {
	// TODO: Implement
	return nil
}

func ParseRRule(input string) (*Rule, error) {
	// TODO: Implement
	/*
	 rrule      = "RRULE" rrulparam ":" recur CRLF
	 rrulparam  = *(";" other-param)
	*/
	/* NOTE:
	- recur appears to be all the rruleparams
	- CLRF means the CR/end of line
	*/
	ruleParser, err := participle.Build[Rule](
		participle.Lexer(ruleLexer),
		participle.Elide("EOL", "Whitespace"),
		participle.UseLookahead(participle.MaxLookahead), // TODO: Do we need this? (Maybe with KEY=)
	)
	if err != nil {
		return nil, fmt.Errorf("bad grammar: %w", err)
	}
	return ruleParser.ParseString("", input)
}

func (r Rule) RRule() string {
	// TODO: Implement
	return "IMPLEMENT"

}

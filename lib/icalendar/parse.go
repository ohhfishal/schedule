package icalendar

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var ErrBadValue = errors.New(`invalid value`)

var _TIME_REGEX = `\d{8}T\d{6}Z`
var TIME_FORMAT = `20060102T150405Z`

var (
	calendarLexerInit sync.Once
	calendarLexer     lexer.Definition
	calendarLexerErr  error
)

func Parse(input string) (*Calendar, error) {
	lexer, err := getCalendarLexer()
	if err != nil {
		return nil, fmt.Errorf("lexer: %w", err)
	}
	parser, err := participle.Build[calendarGrammer](
		participle.Lexer(lexer),
	)
	if err != nil {
		return nil, fmt.Errorf("grammar: %w", err)
	}
	parsed, err := parser.ParseString("", input)
	if err != nil {
		return nil, fmt.Errorf("format: %w", err)
	}
	return parsed.Calendar()
}

type calendarGrammer struct {
	Properties []Property `parser:"@@*"`
}

type Property struct {
	Name       Name        `parser:"@@"`
	Parameters []Parameter `parser:"(';' @@)*"`
	Value      Value       `parser:"':' @@ 'EOL'"`
}

type Name struct {
	raw string `parser:"'NAME'"`
}

type Parameter struct {
}

type Value struct {
}

func (grammar calendarGrammer) Calendar() (*Calendar, error) {
	return nil, nil
}

var names = []string{}

func getCalendarLexer() (lexer.Definition, error) {
	calendarLexerInit.Do(func() {
		calendarLexer, calendarLexerErr = lexer.NewSimple([]lexer.SimpleRule{
			{Name: "NAME", Pattern: strings.Join(names, "|")},
			{Name: ":", Pattern: `:`},
			{Name: ";", Pattern: `;`},
			{Name: "=", Pattern: `=`},
			// {Name: "Time", Pattern: _TIME_REGEX},
			// {Name: "Frequency", Pattern: `MINUTELY|HOURLY|DAILY|WEEKLY|MONTHLY|YEARLY`},
			// {Name: "WEEKDAY", Pattern: `SU|MO|TU|WE|TH|FR|SA`},
			// {Name: "Int", Pattern: `-?\d+`},
		})
	})
	return calendarLexer, calendarLexerErr
}

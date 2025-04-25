package icalendar

import (
	"errors"
	"fmt"
	"strings"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var ErrBadValue = errors.New(`invalid value`)

func Parse(input string) (*Calendar, error) {
	calLexer := lexer.MustSimple([]lexer.SimpleRule{
		{Name: "NAME", Pattern: `[A-Za-z0-9\-_]+`},
		{Name: "VALUE", Pattern: `[^\r\n;:]*`},
		{Name: ":", Pattern: `:`},
		{Name: ";", Pattern: `;`},
		{Name: "=", Pattern: `=`},
		{Name: "NEWLINE", Pattern: `\r?\n`},
		{Name: "WHITESPACE", Pattern: `[ \t]+`},
	})
	parser, err := participle.Build[calendarGrammer](
		participle.Lexer(calLexer),
		participle.UseLookahead(participle.MaxLookahead),
	)
	if err != nil {
		return nil, fmt.Errorf("grammar: %w", err)
	}

	tokens, err := parser.Lex("", strings.NewReader(input))
	if err != nil {
		fmt.Println("tokens:", tokens)
		return nil, fmt.Errorf("lexing: %w", err)
	}

	parsed, err := parser.ParseString("", input)
	if err != nil {
		for _, token := range tokens {
			fmt.Println(token)
		}
		return nil, fmt.Errorf("parsing: %w", err)
	}
	return parsed.Calendar()
}

// Line example
// PRODID:-//RDU Software//NONSGML HandCal//EN
type calendarGrammer struct {
	Begin      string      `parser:"'BEGIN:VCALENDAR'"`
	Properties []Property  `parser:"@@*"`
	Components []Component `parser:"@@*"`
	End        string      `parser:"'END:VCALENDAR'"`
}

type Property struct {
	Name       string      `parser:"@NAME"`
	Parameters []Parameter `parser:"(';' @@)*"`
	Value      string      `parser:"':' @VALUE"`
}

type Component struct {
	Begin      string      `parser:"'BEGIN:' @NAME"`
	Properties []Property  `parser:"@@*"`
	Components []Component `parser:"@@*"`
	End        string      `parser:"'END:' @NAME"`
}

type Parameter struct {
	Name  string `parser:"@NAME"`
	Value string `parser:"@NAME|@VALUE"`
}

func (grammar calendarGrammer) Calendar() (*Calendar, error) {
	return nil, nil
}

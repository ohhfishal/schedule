package icalendar

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/alecthomas/participle/v2/lexer"
)

func Lex(input string) ([]lexer.Token, error) {
	var tokens []lexer.Token
	for line := range strings.Lines(input) {
		lineTokens, err := LexLine(line)
		if err != nil {
			return tokens, err
		}
		for _, token := range lineTokens {
			tokens = append(tokens, token)
		}
	}
	return tokens, errors.New("DONE: TODO: LABEL TOKENS")
}
func LexLine(line string) ([]lexer.Token, error) {
	var tokens []lexer.Token
	if !strings.Contains(line, ":") {
		return tokens, fmt.Errorf(`missing: ":" in (%s)`, line)
	}
	s := strings.Split(line, ":")
	left := s[0]
	value := strings.TrimSpace(strings.Join(s[1:], ":"))

	if !strings.Contains(left, ";") {
		tokens = append(tokens, lexer.Token{Value: left})
		tokens = append(tokens, lexer.Token{Value: ":"})
	} else {
		sections := strings.Split(left, ";")
		tokens = append(tokens, lexer.Token{Value: sections[0]})
		for _, param := range sections[1:] {
			tokens = append(tokens, lexer.Token{Value: ";"})
			if strings.Count(param, "=") != 1 {
				return tokens, fmt.Errorf(`param missing "=" (%s)`, param)
			}
			pair := strings.Split(param, "=")
			tokens = append(tokens, lexer.Token{Value: pair[0]})
			tokens = append(tokens, lexer.Token{Value: "="})
			tokens = append(tokens, lexer.Token{Value: pair[1]})

		}
		tokens = append(tokens, lexer.Token{Value: ":"})
	}
	tokens = append(tokens, lexer.Token{Value: value})
	tokens = append(tokens, lexer.Token{Value: "EOL"})
	return tokens, nil

}

const (
	TokenEOF = iota - 1
	TokenSemicolon
	TokenEqual
	TokenBegin
	TokenName
	TokenColon
	TokenBeginLiteral
	TokenParameterValue
)

type State int
type StateFunc func(*CalendarLexer) (lexer.Token, error)

const (
	StateInvalid State = iota - 1
	StateRoot
	StateName
	StateValue
	StateBegin
	StateSemicolon
	StateColon
	StateEqual
	StateDone
)

// Implements lexer.Definition
type CalendarLexer struct {
	reader   io.Reader
	scanner  *bufio.Scanner
	position lexer.Position
}

func (l *CalendarLexer) MatchString(literal string) error {
	return nil
}

func (l *CalendarLexer) Err(err error) (lexer.Token, error) {
	return lexer.Token{}, err
}

func (l *CalendarLexer) Next() (lexer.Token, error) {
	for l.scanner.Scan() {
		fmt.Println("HERE")
		if err := l.scanner.Err(); err != nil {
			return l.Err(err)
		}
		line := l.scanner.Text()
		fmt.Println(line)
	}
	return l.Err(nil)
}

func (l *CalendarLexer) Pos() lexer.Position {
	// TODO: Implement
	return l.position
}

func NewLexer(str string) (lexer.Lexer, error) {
	return CalendarLexerInit{}.Lex(``, strings.NewReader(str))
}

// Implements lexer.Definition
type CalendarLexerInit struct{}

func (_ CalendarLexerInit) Lex(filename string, reader io.Reader) (lexer.Lexer, error) {
	return &CalendarLexer{
		reader:  reader,
		scanner: bufio.NewScanner(reader),
		position: lexer.Position{
			Filename: filename,
		},
	}, nil
}
func (_ CalendarLexerInit) Symbols() map[string]lexer.TokenType {
	// TODO: Add them here
	return map[string]lexer.TokenType{}
}

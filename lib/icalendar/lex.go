package icalendar

import (
	"fmt"
	"io"

	"github.com/alecthomas/participle/v2/lexer"
)

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
	text     string
	state    State
	position lexer.Position
}

func (l *CalendarLexer) MatchString(literal string) error {
	return nil
}

func (l *CalendarLexer) Err(err error) (lexer.Token, error) {
	return lexer.Token{}, err
}

func (l *CalendarLexer) Next() (lexer.Token, error) {
	switch l.state {
	case StateRoot:
		l.state = StateBegin
	case StateBegin:
		// TODO: Do we even what this as its own state?
		// TODO: I think we should:
		//       At root parse until the ';' or ':'
		//       Then if it's a known property, make it a token, otherwise a property
		//       For now that is just `BEGIN` and `END` since they are the predict sets for
		// BEGIN = Start looking for NAME / END
		// END = Look for an END/BEGIN or EOF
		// Name: Stop util ; or : then return the right token type based on content
		if err := l.MatchString("BEGIN"); err != nil {
			return l.Err(err)
		}
		l.state = StateColon
		return lexer.Token{
			Type:  TokenBegin,
			Value: `BEGIN`,
			Pos:   l.Pos(),
		}, nil
	case StateName:
		// TODO: Implement see notes above
	case StateValue:
		// TODO: Implement: Consume until the end of line or EOF
	case StateColon:
		if err := l.MatchString(":"); err != nil {
			return l.Err(err)
		}
		l.state = StateValue
		return lexer.Token{
			Type:  TokenColon,
			Value: `:`,
			Pos:   l.Pos(),
		}, nil
	case StateEqual:
		if err := l.MatchString("="); err != nil {
			return l.Err(err)
		}
		l.state = StateName // TODO: Confirm this is the right state
		return lexer.Token{
			Type:  TokenEqual,
			Value: `=`,
			Pos:   l.Pos(),
		}, nil
	case StateDone:
		l.state = StateInvalid
		return lexer.EOFToken(l.Pos()), nil
	default:
		return lexer.Token{}, fmt.Errorf(`invalid state: %v`, l.state)
	}
	return l.Next()
}

func (l *CalendarLexer) Pos() lexer.Position {
	// TODO: Implement
	return l.position
}

// Implements lexer.Definition
type CalendarLexerInit struct{}

func (_ CalendarLexerInit) Lex(filename string, reader io.Reader) (lexer.Lexer, error) {
	if input, err := io.ReadAll(reader); err != nil {
		return nil, fmt.Errorf(`reading: %w`, err)
	} else {
		return &CalendarLexer{
			text: string(input),
			position: lexer.Position{
				Filename: filename,
			},
		}, nil
	}
}
func (_ CalendarLexerInit) Symbols() map[string]lexer.TokenType {
	// TODO: Add them here
	return map[string]lexer.TokenType{}
}

package lexer

import (
	"fmt"
)

type LexerErrorType string

const (
	NumberFormatError LexerErrorType = "Number Format"
	UnexpectedToken   LexerErrorType = "Unexpected Token"
)

type LexerError struct {
	lexerErrorType LexerErrorType
	line           int
	col            int
	value          string
}

func (lexErr *LexerError) Error() string {
	return fmt.Sprintf("Lexer Error: %s @ line: %d, col: %d, value: '%s'", lexErr.lexerErrorType, lexErr.line, lexErr.col, lexErr.value)
}

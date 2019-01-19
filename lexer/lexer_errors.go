package lexer

import (
	"fmt"
)

type LexerErrorType string

const (
	NumberFormatError LexerErrorType = "Number Format"
)

type LexerError struct {
	lexerErrorType LexerErrorType
	row            int
	col            int
	value          string
}

func (lexErr *LexerError) Error() string {
	return fmt.Sprintf("Lexer Error: %s @ line: %d, col: %d, value: '%s'", lexErr.lexerErrorType, lexErr.row, lexErr.col, lexErr.value)
}

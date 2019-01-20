package parser

import (
	"fmt"
)

type ParserErrorType string

const (
	MISSING_IDENT ParserErrorType = "MISSING_IDENTIFIER"
)

type ParserError struct {
	errorType ParserErrorType
	line      int
	message   string
}

func (e *ParserError) Error() string {
	return fmt.Sprintf("Parser Error(%s) - %s @ Line: %d", e.errorType, e.message, e.line)
}

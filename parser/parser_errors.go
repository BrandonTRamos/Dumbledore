package parser

type ParserErrorType string

const ()

type ParserError struct {
	errorType ParserErrorType
	cow       int
	col       int
	message   string
}

func (parserError *ParserError) Error() string {
	return "Parser error!"
}

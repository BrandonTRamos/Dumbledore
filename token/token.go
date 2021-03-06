package token

import (
	"fmt"
	"strings"
)

type TokenType string

const (
	ILLEGAL      TokenType = "ILLEGAL"
	ERROR        TokenType = "ERROR"
	EOF          TokenType = "EOF"
	IDENTIFIER   TokenType = "INDENTIFIER"
	INT          TokenType = "INT"
	DOUBLE       TokenType = "DOUBLE"
	ASSIGN       TokenType = "ASSIGN"
	PLUS         TokenType = "PLUS"
	ASTERIK      TokenType = "ASTERIK"
	SLASH        TokenType = "SLASH"
	COMMA        TokenType = "COMMA"
	SEMICOLON    TokenType = "SEMICOLON"
	LPAREN       TokenType = "LPAREN"
	RPAREN       TokenType = "RPAREN"
	LBRACE       TokenType = "LBRACE"
	RBRACE       TokenType = "RBRACE"
	DOT          TokenType = "DOT"
	GT           TokenType = "GREATER_THAN"
	LT           TokenType = "LESS_THAN"
	EXCLAIMATION TokenType = "EXCLAIMATION"
	MINUS        TokenType = "MINUS"
	EQUAL        TokenType = "EQUAL"
	NOTEQUAL     TokenType = "NOT_EQUAL"

	//keywords
	FUNCTION TokenType = "KEYWORD: FUNCTION"
	VAR      TokenType = "KEYWORD: VAR"
	IF       TokenType = "KEYWORD: IF"
	ELSE     TokenType = "KEYWORD: ESLE"
	ELIF     TokenType = "KEYWORD: ELIF"
	TRUE     TokenType = "KEYWORD: TRUE"
	FALSE    TokenType = "KEYWORD: FALSE"
	RETURN   TokenType = "KEYWORD: RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"var":    VAR,
	"if":     IF,
	"elif":   ELIF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

type Token struct {
	Type    TokenType
	Literal string
}

func (t Token) ToString(indentLevel int) string {
	return fmt.Sprintf("Token {Type: %s , Literal: '%s'}", t.Type, t.Literal)
}

//static functions
func LookupKeywords(indentifier string) TokenType {
	tokenType, foundMatch := keywords[indentifier]
	if foundMatch {
		return tokenType
	}
	return IDENTIFIER
}

func CheckNumberType(number string) TokenType {
	if strings.Contains(number, ".") {
		return DOUBLE
	}
	return INT
}

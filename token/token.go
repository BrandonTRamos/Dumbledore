package token

import (
	"fmt"
	"strings"
)

type TokenType string

const (
	ILLEGAL    TokenType = "ILLEGAL"
	EOF        TokenType = "EOF"
	IDENTIFIER TokenType = "INDENTIFIER"
	INT        TokenType = "INT"
	DOUBLE     TokenType = "DOUBLE"
	ASSIGN     TokenType = "ASSIGN"
	PLUS       TokenType = "PLUS"
	COMMA      TokenType = "COMMA"
	SEMICOLON  TokenType = "SEMICOLON"
	LPAREN     TokenType = "LPAREN"
	RPAREN     TokenType = "RPAREN"
	LBRACE     TokenType = "LBRACE"
	RBRACE     TokenType = "RBRACE"
	DOT        TokenType = "DOT"
	//keywords
	FUNCTION = "KEYWORD: FUNCTION"
	VAR      = "KEYWORD: VAR"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"var": VAR,
}

type Token struct {
	Type    TokenType
	Literal string
}

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

func (t Token) ToString() string {
	return fmt.Sprintf("Token { Type: %s ,  Literal: '%s'}", t.Type, t.Literal)
}

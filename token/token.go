package token

import(
	"fmt"
	
)

const(
ILLEGAL = "ILLEGAL"
EOF = "EOF"
IDENTIFIER = "INDENTIFIER"
INT = "INT"
ASSIGN = "="
PLUS = "+"
COMMA = ","
SEMICOLON = ";"
LPAREN="("
RPAREN=")"
LBRACE = "{"
RBRACE = "}"

FUNCTION = "FUNCTION"
VAR = "VAR"

)

type TokenType string;

type Token struct {
	Type TokenType
	Literal string
}

func (t Token) ToString() string{
	return fmt.Sprintf("Type: %s , Literal: %s", t.Type,t.Literal)
}
package lexer

import (
	"Dumbledore/token"
	"fmt"
	"io/ioutil"
)

type Lexer struct {
	input           string
	currentPosition int
	nextPosition    int
	ch              byte
}

func NewLexerFromString(input string) *Lexer {
	return &Lexer{input: input}
}

func NewLexerFromFile(fileName string) *Lexer {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	return &Lexer{input: string(bytes)}
}

func (lexer *Lexer) HasNext() bool {
	if lexer.currentPosition >= len(lexer.input) {
		lexer.ch = 0
		return false
	}

	return true

}

func (lexer *Lexer) ReadChar() {
	if lexer.nextPosition >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.nextPosition]
	}
	lexer.currentPosition = lexer.nextPosition
	lexer.nextPosition += 1
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token
	lexer.skipWhiteSpace()
	switch lexer.ch {
	case '=':
		tok = newTokenFromChar(token.ASSIGN, lexer.ch)
	case '+':
		tok = newTokenFromChar(token.PLUS, lexer.ch)
	default:
		if isLetter(lexer.ch) {
			tok = newTokenFromChar(token.IDENTIFIER, lexer.ch)
		} else {
			fmt.Printf("Illegal char byte (hex notation): %x\n", lexer.ch)
			tok = newTokenFromChar(token.ILLEGAL, lexer.ch)
		}

	}
	lexer.ReadChar()
	return tok
}

func (lexer *Lexer) skipWhiteSpace() {
	for lexer.ch == ' ' || lexer.ch == '\n' || lexer.ch == '\r' || lexer.ch == '\t' {
		lexer.ReadChar()
	}
}

func (lexer *Lexer) ReadTokens() {
	lexer.ReadChar()
	for lexer.HasNext() {
		fmt.Println(lexer.NextToken().ToString())
	}
}

func newTokenFromChar(tokenType token.TokenType, charLiteral byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(charLiteral)}
}

func newTokenFromString(tokenType token.TokenType, strLiteral string) token.Token {
	return token.Token{Type: tokenType, Literal: strLiteral}
}

func isLetter(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == '_'
}

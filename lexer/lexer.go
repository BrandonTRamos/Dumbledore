package lexer

import (
	"Dumbledore/token"
	"fmt"
	"io/ioutil"
	"strings"
)

type Lexer struct {
	input           []byte
	currentPosition int
	nextPosition    int
	ch              byte
	line            int
	col             int
}

func NewLexerFromString(input []byte) *Lexer {
	return &Lexer{input: input, line: 1, col: 0}
}

func NewLexerFromFile(fileName string) *Lexer {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bytes)

	return &Lexer{input: bytes, line: 1, col: 0}
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
	lexer.col += 1
}

func (lexer *Lexer) readIdentifier() (token.TokenType, string) {
	beginIndex := lexer.currentPosition
	for isLetter(lexer.ch) {
		lexer.ReadChar()
	}
	identifier := string(lexer.input[beginIndex:lexer.currentPosition])
	tokenType := token.LookupKeywords(identifier)
	return tokenType, identifier
}

func (lexer *Lexer) readNumber() (token.Token, error) {
	beginCol := lexer.col
	beginIndex := lexer.currentPosition
	var tok token.Token
	for isNumber(lexer.ch) {
		lexer.ReadChar()
	}
	numberString := string(lexer.input[beginIndex:lexer.currentPosition])
	tokenType := token.CheckNumberType(numberString)

	if strings.Count(numberString, ".") > 1 {
		return token.Token{token.ERROR, numberString}, &LexerError{NumberFormatError, lexer.line, beginCol, numberString}
	}

	tok = newTokenFromString(tokenType, numberString)
	return tok, nil
}

func (lexer *Lexer) NextToken() (token.Token, error) {
	var tok token.Token
	var err error
	lexer.skipWhiteSpace()
	switch lexer.ch {
	case '=':
		tok = newTokenFromChar(token.ASSIGN, lexer.ch)
	case '+':
		tok = newTokenFromChar(token.PLUS, lexer.ch)
	case '.':
		if isNumber(lexer.input[lexer.nextPosition]) {
			tok, numErr := lexer.readNumber()
			return tok, numErr
		}
		tok = newTokenFromChar(token.DOT, lexer.ch)
	case ';':
		tok = newTokenFromChar(token.SEMICOLON, lexer.ch)
	default:
		if isLetter(lexer.ch) {
			identifier, tokenType := lexer.readIdentifier()
			return newTokenFromString(identifier, tokenType), nil
		} else if isNumber(lexer.ch) {
			tok, numErr := lexer.readNumber()
			return tok, numErr
		} else {
			fmt.Printf("Illegal char byte (ascii number): %d\n", lexer.ch)
			tok = newTokenFromChar(token.ILLEGAL, lexer.ch)
		}

	}
	lexer.ReadChar()
	return tok, err
}

func (lexer *Lexer) skipWhiteSpace() {
	for lexer.ch == ' ' || lexer.ch == '\n' || lexer.ch == '\r' || lexer.ch == '\t' {
		if lexer.ch == '\n' {
			lexer.line += 1
			lexer.col = 0
		}
		lexer.ReadChar()
	}
}

func (lexer *Lexer) ReadTokens() {
	lexer.ReadChar()
	for lexer.HasNext() {
		tok, err := lexer.NextToken()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(tok.ToString())
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

func isNumber(char byte) bool {
	return (char >= '0' && char <= '9') || char == '.'
}

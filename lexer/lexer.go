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
	Line            int
	Col             int
}

func NewLexerFromString(input []byte) *Lexer {
	return &Lexer{input: input, Line: 1, Col: 0}
}

func NewLexerFromFile(fileName string) *Lexer {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	lex := &Lexer{input: bytes, Line: 1, Col: 0}
	lex.readChar()
	return lex
}

func (lexer *Lexer) HasNext() bool {
	if lexer.currentPosition >= len(lexer.input) {
		lexer.ch = 0
		return false
	}

	return true

}

func (lexer *Lexer) readChar() {

	if lexer.nextPosition >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.nextPosition]
	}
	lexer.currentPosition = lexer.nextPosition
	lexer.nextPosition += 1
	lexer.Col += 1
}

func (lexer *Lexer) readIdentifier() (token.Token, error) {
	beginIndex := lexer.currentPosition
	for isLetter(lexer.ch) {
		lexer.readChar()
	}
	identifier := string(lexer.input[beginIndex:lexer.currentPosition])
	tokenType := token.LookupKeywords(identifier)
	return token.Token{tokenType, identifier}, nil
}

func (lexer *Lexer) readNumber() (token.Token, error) {
	beginCol := lexer.Col
	beginIndex := lexer.currentPosition
	var tok token.Token
	for isNumber(lexer.ch) {
		lexer.readChar()
	}
	numberString := string(lexer.input[beginIndex:lexer.currentPosition])
	tokenType := token.CheckNumberType(numberString)

	if strings.Count(numberString, ".") > 1 {
		return token.Token{token.ERROR, numberString}, &LexerError{NumberFormatError, lexer.Line, beginCol, numberString}
	}

	tok = token.Token{tokenType, numberString}
	return tok, nil
}

func (lexer *Lexer) nextChar() byte {
	if lexer.nextPosition >= len(lexer.input) {
		return 0
	} else {
		return lexer.input[lexer.nextPosition]
	}
}

func (lexer *Lexer) NextToken() (token.Token, error) {
	var tok token.Token
	lexer.skipWhiteSpace()
	switch lexer.ch {
	case 0:
		tok = newTokenFromChar(token.EOF, lexer.ch)
	case '}':
		tok = newTokenFromChar(token.RBRACE, lexer.ch)
	case '{':
		tok = newTokenFromChar(token.LBRACE, lexer.ch)
	case ')':
		tok = newTokenFromChar(token.RPAREN, lexer.ch)
	case '(':
		tok = newTokenFromChar(token.LPAREN, lexer.ch)
	case '/':
		tok = newTokenFromChar(token.SLASH, lexer.ch)
	case '*':
		tok = newTokenFromChar(token.ASTERIK, lexer.ch)
	case '<':
		tok = newTokenFromChar(token.LT, lexer.ch)
	case '>':
		tok = newTokenFromChar(token.GT, lexer.ch)
	case '!':
		if lexer.nextChar() == '=' {
			tok = token.Token{token.NOTEQUAL, "!="}
			lexer.readChar()
		} else {
			tok = newTokenFromChar(token.EXCLAIMATION, lexer.ch)
		}

	case '=':
		if lexer.nextChar() == '=' {
			tok = token.Token{token.EQUAL, "=="}
			lexer.readChar()
		} else {
			tok = newTokenFromChar(token.ASSIGN, lexer.ch)
		}
	case '+':
		tok = newTokenFromChar(token.PLUS, lexer.ch)
	case '-':
		tok = newTokenFromChar(token.MINUS, lexer.ch)
	case ';':
		tok = newTokenFromChar(token.SEMICOLON, lexer.ch)
	case '.':
		if isNumber(lexer.input[lexer.nextPosition]) {
			numTok, numErr := lexer.readNumber()
			return numTok, numErr
		}
		tok = newTokenFromChar(token.DOT, lexer.ch)
	default:
		if isLetter(lexer.ch) {
			indentToken, indentErr := lexer.readIdentifier()
			return indentToken, indentErr
		} else if isNumber(lexer.ch) {
			tok, numErr := lexer.readNumber()
			return tok, numErr
		} else {
			fmt.Printf("Illegal char byte (ascii number): %d\n", lexer.ch)
			tok = newTokenFromChar(token.ILLEGAL, lexer.ch)
		}

	}
	lexer.readChar()
	return tok, nil
}

func (lexer *Lexer) skipWhiteSpace() {
	for lexer.ch == ' ' || lexer.ch == '\n' || lexer.ch == '\r' || lexer.ch == '\t' {
		if lexer.ch == '\n' {
			lexer.Line += 1
			lexer.Col = 0
		}
		lexer.readChar()
	}
}

func (lexer *Lexer) ReadTokens() {
	lexer.readChar()
	for lexer.HasNext() {
		tok, err := lexer.NextToken()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(tok.ToString())
	}
}

//static functions

func newTokenFromChar(tokenType token.TokenType, charLiteral byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(charLiteral)}
}

func isLetter(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == '_'
}

func isNumber(char byte) bool {
	return (char >= '0' && char <= '9') || char == '.'
}

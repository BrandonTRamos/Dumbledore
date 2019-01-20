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
	CH              byte
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
	lex := &Lexer{input: bytes, line: 1, col: 0}
	lex.readChar()
	return lex
}

func (lexer *Lexer) HasNext() bool {
	if lexer.currentPosition >= len(lexer.input) {
		lexer.CH = 0
		return false
	}

	return true

}

func (lexer *Lexer) readChar() {

	if lexer.nextPosition >= len(lexer.input) {
		lexer.CH = 0
	} else {
		lexer.CH = lexer.input[lexer.nextPosition]
	}
	lexer.currentPosition = lexer.nextPosition
	lexer.nextPosition += 1
	lexer.col += 1
}

func (lexer *Lexer) readIdentifier() (token.Token, error) {
	beginIndex := lexer.currentPosition
	for isLetter(lexer.CH) {
		lexer.readChar()
	}
	identifier := string(lexer.input[beginIndex:lexer.currentPosition])
	tokenType := token.LookupKeywords(identifier)
	return token.Token{tokenType, identifier}, nil
}

func (lexer *Lexer) readNumber() (token.Token, error) {
	beginCol := lexer.col
	beginIndex := lexer.currentPosition
	var tok token.Token
	for isNumber(lexer.CH) {
		lexer.readChar()
	}
	numberString := string(lexer.input[beginIndex:lexer.currentPosition])
	tokenType := token.CheckNumberType(numberString)

	if strings.Count(numberString, ".") > 1 {
		return token.Token{token.ERROR, numberString}, &LexerError{NumberFormatError, lexer.line, beginCol, numberString}
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
	switch lexer.CH {
	case 0:
		tok = newTokenFromChar(token.EOF, lexer.CH)
	case '}':
		tok = newTokenFromChar(token.RBRACE, lexer.CH)
	case '{':
		tok = newTokenFromChar(token.LBRACE, lexer.CH)
	case ')':
		tok = newTokenFromChar(token.RPAREN, lexer.CH)
	case '(':
		tok = newTokenFromChar(token.LPAREN, lexer.CH)
	case '/':
		tok = newTokenFromChar(token.SLASH, lexer.CH)
	case '*':
		tok = newTokenFromChar(token.ASTERIK, lexer.CH)
	case '<':
		tok = newTokenFromChar(token.LT, lexer.CH)
	case '>':
		tok = newTokenFromChar(token.GT, lexer.CH)
	case '!':
		if lexer.nextChar() == '=' {
			tok = token.Token{token.NOTEQUAL, "!="}
			lexer.readChar()
		} else {
			tok = newTokenFromChar(token.EXCLAIMATION, lexer.CH)
		}

	case '=':
		if lexer.nextChar() == '=' {
			tok = token.Token{token.EQUAL, "=="}
			lexer.readChar()
		} else {
			tok = newTokenFromChar(token.ASSIGN, lexer.CH)
		}
	case '+':
		tok = newTokenFromChar(token.PLUS, lexer.CH)
	case '-':
		tok = newTokenFromChar(token.MINUS, lexer.CH)
	case ';':
		tok = newTokenFromChar(token.SEMICOLON, lexer.CH)
	case '.':
		if isNumber(lexer.input[lexer.nextPosition]) {
			numTok, numErr := lexer.readNumber()
			return numTok, numErr
		}
		tok = newTokenFromChar(token.DOT, lexer.CH)
	default:
		if isLetter(lexer.CH) {
			indentToken, indentErr := lexer.readIdentifier()
			return indentToken, indentErr
		} else if isNumber(lexer.CH) {
			tok, numErr := lexer.readNumber()
			return tok, numErr
		} else {
			fmt.Printf("Illegal char byte (ascii number): %d\n", lexer.CH)
			tok = newTokenFromChar(token.ILLEGAL, lexer.CH)
		}

	}
	lexer.readChar()
	return tok, nil
}

func (lexer *Lexer) skipWhiteSpace() {
	for lexer.CH == ' ' || lexer.CH == '\n' || lexer.CH == '\r' || lexer.CH == '\t' {
		if lexer.CH == '\n' {
			lexer.line += 1
			lexer.col = 0
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

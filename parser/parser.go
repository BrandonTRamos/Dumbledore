package parser

import (
	"Dumbledore/ast"
	"Dumbledore/lexer"
	"Dumbledore/token"
	"log"
)

type Parser struct {
	Lexer        *lexer.Lexer
	CurrentToken token.Token
	PeekToken    token.Token
}

func New(lexer *lexer.Lexer) *Parser {
	parser := &Parser{Lexer: lexer}
	parser.getNextToken()
	parser.getNextToken()
	return parser
}

func (parser *Parser) getNextToken() {
	parser.CurrentToken = parser.PeekToken
	nextTok, _ := parser.Lexer.NextToken()
	parser.PeekToken = nextTok

}

func (parser *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for parser.CurrentToken.Type != token.EOF {
		stmt, err := parser.parseStatement()
		if err != nil {
			log.Fatal(err)
		}

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		parser.getNextToken()
	}

	return program
}

func (parser *Parser) parseStatement() (ast.Statement, error) {
	switch parser.CurrentToken.Type {
	case token.VAR:
		return parser.parseVarStatement()
	default:
		return nil, nil
	}
}

func (parser *Parser) parseVarStatement() (*ast.VarStatement, error) {
	varStatement := &ast.VarStatement{VarToken: parser.CurrentToken}

	if parser.PeekToken.Type != token.IDENTIFIER {
		return varStatement, &ParserError{}
	}

	for parser.CurrentToken.Type != token.SEMICOLON {
		parser.getNextToken()
	}
	return varStatement, nil
}

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
	Errors       []error
}

func New(lexer *lexer.Lexer) *Parser {
	parser := &Parser{Lexer: lexer, Errors: []error{}}
	parser.getNextToken()
	parser.getNextToken()
	return parser
}

func (parser *Parser) getNextToken() {
	parser.CurrentToken = parser.PeekToken
	nextTok, err := parser.Lexer.NextToken()
	if err != nil {
		log.Fatal(err)
	}
	parser.PeekToken = nextTok

}

func (parser *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for parser.CurrentToken.Type != token.EOF {
		stmt, err := parser.parseStatement()
		if err != nil {
			parser.Errors = append(parser.Errors, err)
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
	case token.RETURN:
		return parser.parseReturnStatement()
	default:
		return nil, nil
	}
}

func (parser *Parser) parseVarStatement() (*ast.VarStatement, error) {
	varStatement := &ast.VarStatement{VarToken: parser.CurrentToken}
	if parser.PeekToken.Type != token.IDENTIFIER {
		return nil, &ParserError{errorType: MISSING_IDENT, message: "Var declaration missing variable name", line: parser.Lexer.Line}
	}
	parser.getNextToken()
	varStatement.Name = &ast.Identifier{IdentToken: parser.CurrentToken, Value: parser.CurrentToken.Literal}

	for parser.CurrentToken.Type != token.SEMICOLON && parser.CurrentToken.Type != token.EOF {
		parser.getNextToken()
	}
	return varStatement, nil
}

func (parser *Parser) parseReturnStatement() (*ast.ReturnStatement, error) {
	returnStatement := &ast.ReturnStatement{ReturnToken: parser.CurrentToken}

	parser.getNextToken()
	for parser.CurrentToken.Type != token.SEMICOLON && parser.CurrentToken.Type != token.EOF {
		parser.getNextToken()
	}
	return returnStatement, nil
}

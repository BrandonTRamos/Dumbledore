package parser

import (
	"Dumbledore/ast"
	"Dumbledore/lexer"
	"Dumbledore/token"
	"log"
	"strconv"
)

type OperatorPrecedence int

const (
	LOWEST OperatorPrecedence = iota
	EQUALS
	LESSGREATER
	SUMSUBTRACT
	PRODUCTDIVIDE
	PREFIX
	CALL
)

var precedenses = map[token.TokenType]OperatorPrecedence{
	token.EQUAL:    EQUALS,
	token.NOTEQUAL: EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.PLUS:     SUMSUBTRACT,
	token.MINUS:    SUMSUBTRACT,
	token.SLASH:    PRODUCTDIVIDE,
	token.ASTERIK:  PRODUCTDIVIDE,
}

type (
	prefixParserFn func() (ast.Expression, error)
	infixParserFn  func(ast.Expression) (ast.Expression, error)
)

type Parser struct {
	Lexer *lexer.Lexer

	CurrentToken token.Token
	PeekToken    token.Token

	prefixParserFns map[token.TokenType]prefixParserFn
	infixParserFns  map[token.TokenType]infixParserFn

	Errors []error
}

func New(lexer *lexer.Lexer) *Parser {
	parser := &Parser{Lexer: lexer, Errors: []error{}}

	parser.prefixParserFns = make(map[token.TokenType]prefixParserFn)
	parser.prefixParserFns[token.IDENTIFIER] = parser.parseIdentifier
	parser.prefixParserFns[token.INT] = parser.parseIntegerLiteral
	parser.prefixParserFns[token.DOUBLE] = parser.parseDoubleLiteral
	parser.prefixParserFns[token.EXCLAIMATION] = parser.parsePrefixExpression
	parser.prefixParserFns[token.MINUS] = parser.parsePrefixExpression
	parser.prefixParserFns[token.TRUE] = parser.parseBooleanLiteral
	parser.prefixParserFns[token.FALSE] = parser.parseBooleanLiteral

	parser.infixParserFns = make(map[token.TokenType]infixParserFn)
	parser.infixParserFns[token.PLUS] = parser.parseInfixExpression
	parser.infixParserFns[token.MINUS] = parser.parseInfixExpression
	parser.infixParserFns[token.SLASH] = parser.parseInfixExpression
	parser.infixParserFns[token.ASTERIK] = parser.parseInfixExpression
	parser.infixParserFns[token.EQUAL] = parser.parseInfixExpression
	parser.infixParserFns[token.NOTEQUAL] = parser.parseInfixExpression
	parser.infixParserFns[token.LT] = parser.parseInfixExpression
	parser.infixParserFns[token.GT] = parser.parseInfixExpression

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
		return parser.parseExpressionStatement()
	}
}

func (parser *Parser) parseVarStatement() (*ast.VarStatement, error) {
	varStatement := &ast.VarStatement{VarToken: parser.CurrentToken}
	if parser.PeekToken.Type != token.IDENTIFIER {
		return nil, &ParserError{errorType: MISSING_IDENT, message: "Var declaration missing variable name", line: parser.Lexer.Line}
	}
	parser.getNextToken()
	varStatement.Name = &ast.Identifier{IdentToken: parser.CurrentToken, Value: parser.CurrentToken.Literal}

	if parser.PeekToken.Type != token.ASSIGN {
		return nil, &ParserError{errorType: MISSING_ASSIGNMENT_OPERATOR, message: "Var declaration missing assignment operator '='", line: parser.Lexer.Line}
	}

	for parser.CurrentToken.Type != token.SEMICOLON && parser.CurrentToken.Type != token.EOF {
		parser.getNextToken()
		stmt, _ := parser.parseExpressionStatement()
		varStatement.Value = stmt
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

func (parser *Parser) parseExpressionStatement() (*ast.ExpressionStatement, error) {
	stmt := &ast.ExpressionStatement{ExpressionToken: parser.CurrentToken}

	stmt.Expression = parser.parseExpression(LOWEST)

	if parser.PeekToken.Type == token.SEMICOLON {
		parser.getNextToken()
	}

	return stmt, nil
}

func (parser *Parser) parseExpression(precedence OperatorPrecedence) ast.Expression {

	prefix := parser.prefixParserFns[parser.CurrentToken.Type]

	if prefix == nil {
		return nil
	}

	leftExpression, _ := prefix()

	for parser.PeekToken.Type != token.SEMICOLON && parser.PeekToken.Type != token.EOF && precedence < parser.checkNextPrecedence() {
		infixFn := parser.infixParserFns[parser.PeekToken.Type]

		if infixFn == nil {
			return leftExpression
		}
		parser.getNextToken()
		leftExpression, _ = infixFn(leftExpression)
	}

	return leftExpression

}

func (parser *Parser) parseIdentifier() (ast.Expression, error) {
	return &ast.Identifier{IdentToken: parser.CurrentToken, Value: parser.CurrentToken.Literal}, nil
}

func (parser *Parser) parseIntegerLiteral() (ast.Expression, error) {
	intLiteral := &ast.IntegerLiteral{IntegerToken: parser.CurrentToken}
	value, err := strconv.ParseInt(parser.CurrentToken.Literal, 0, 64)
	if err != nil {
		return nil, err
	}
	intLiteral.Value = value
	return intLiteral, nil
}

func (parser *Parser) parseDoubleLiteral() (ast.Expression, error) {
	doubleLiteral := &ast.DoubleLiteral{DoubleToken: parser.CurrentToken}

	value, err := strconv.ParseFloat(parser.CurrentToken.Literal, 64)

	if err != nil {
		return nil, err
	}

	doubleLiteral.Value = value

	return doubleLiteral, nil
}

func (parser *Parser) parseBooleanLiteral() (ast.Expression, error) {
	booleanLiteral := &ast.BooleanLiteral{BoolToken: parser.CurrentToken}
	value, err := strconv.ParseBool(parser.CurrentToken.Literal)
	if err != nil {
		return nil, err
	}
	booleanLiteral.Value = value
	return booleanLiteral, nil
}

func (parser *Parser) parsePrefixExpression() (ast.Expression, error) {
	expression := &ast.PrefixExpression{
		PrefixToken: parser.CurrentToken,
		Operator:    parser.CurrentToken.Literal,
	}

	parser.getNextToken()
	expression.Right = parser.parseExpression(PREFIX)
	return expression, nil

}

func (parser *Parser) currentPrecedence() OperatorPrecedence {
	precedence, found := precedenses[parser.CurrentToken.Type]
	if found {
		return precedence
	}
	return LOWEST
}

func (parser *Parser) checkNextPrecedence() OperatorPrecedence {
	precedence, found := precedenses[parser.PeekToken.Type]
	if found {
		return precedence
	}
	return LOWEST
}

func (parser *Parser) parseInfixExpression(left ast.Expression) (ast.Expression, error) {
	expression := &ast.InfixExpression{
		InfixToken: parser.CurrentToken,
		Operator:   parser.CurrentToken.Literal,
		Left:       left,
	}

	precedence := parser.currentPrecedence()
	parser.getNextToken()
	expression.Right = parser.parseExpression(precedence)

	return expression, nil
}

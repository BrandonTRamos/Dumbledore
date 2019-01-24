package ast

import (
	"Dumbledore/token"
	"fmt"
)

type Node interface {
	TokenLiteral() string
	ToString() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (prog *Program) TokenLiteral() string {
	if len(prog.Statements) > 0 {
		return prog.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type Identifier struct {
	IdentToken token.Token
	Value      string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.IdentToken.Literal }
func (i *Identifier) ToString() string {
	return fmt.Sprintf("Idenfifier {IdentToken: %s}", i.IdentToken.ToString())
}

// statements

type VarStatement struct {
	VarToken token.Token
	Name     *Identifier
	Value    Expression
}

func (vs *VarStatement) statementNode()       {}
func (vs *VarStatement) TokenLiteral() string { return vs.VarToken.Literal }
func (vs *VarStatement) ToString() string {

	return fmt.Sprintf("VarStatement{%s,%s,Expression {}}", vs.VarToken.ToString(), vs.Name.ToString())
}

type ReturnStatement struct {
	ReturnToken token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.ReturnToken.Literal }
func (rs *ReturnStatement) ToString() string {

	return fmt.Sprintf("ReturnStatement{%s,Expression {}}", rs.ReturnToken.ToString())
}

type ExpressionStatement struct {
	ExpressionToken token.Token
	Expression      Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.ExpressionToken.Literal }
func (es *ExpressionStatement) ToString() string {
	return fmt.Sprintf("Expression {%s}", es.ExpressionToken.ToString())
}
func (es *ExpressionStatement) expressionNode() {

}

type IntegerLiteral struct {
	IntegerToken token.Token
	Value        int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.IntegerToken.Literal }
func (il *IntegerLiteral) ToString() string     { return il.IntegerToken.Literal }

type DoubleLiteral struct {
	DoubleToken token.Token
	Value       float64
}

func (dl *DoubleLiteral) expressionNode()      {}
func (dl *DoubleLiteral) TokenLiteral() string { return dl.DoubleToken.Literal }
func (dl *DoubleLiteral) ToString() string     { return dl.DoubleToken.Literal }

package ast

import (
	"Dumbledore/token"
	"fmt"
	"strings"
)

type Node interface {
	TokenLiteral() string
	ToString(indentLevel int) string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Infix interface {
	infixinterface()
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

type IntegerLiteral struct {
	IntegerToken token.Token
	Value        int64
}

func (il *IntegerLiteral) expressionNode()                 {}
func (il *IntegerLiteral) TokenLiteral() string            { return il.IntegerToken.Literal }
func (il *IntegerLiteral) ToString(indentLevel int) string { return il.IntegerToken.Literal }

type DoubleLiteral struct {
	DoubleToken token.Token
	Value       float64
}

func (dl *DoubleLiteral) expressionNode()                 {}
func (dl *DoubleLiteral) TokenLiteral() string            { return dl.DoubleToken.Literal }
func (dl *DoubleLiteral) ToString(indentLevel int) string { return dl.DoubleToken.Literal }

type BooleanLiteral struct {
	BoolToken token.Token
	Value     bool
}

func (bl *BooleanLiteral) expressionNode()                 {}
func (bl *BooleanLiteral) TokenLiteral() string            { return bl.BoolToken.Literal }
func (bl *BooleanLiteral) ToString(indentLevel int) string { return bl.BoolToken.Literal }

type Identifier struct {
	IdentToken token.Token
	Value      string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.IdentToken.Literal }
func (i *Identifier) ToString(indentLevel int) string {
	return fmt.Sprintf("Idenfifier {IdentToken: %s}", i.IdentToken.ToString(0))
}

// statements

type VarStatement struct {
	VarToken token.Token
	Name     *Identifier
	Value    Expression
}

func (vs *VarStatement) statementNode()       {}
func (vs *VarStatement) TokenLiteral() string { return vs.VarToken.Literal }
func (vs *VarStatement) ToString(indentLevel int) string {

	return fmt.Sprintf("VarStatement{\n\t%s,\n\t%s,\n%s}", vs.VarToken.ToString(0), vs.Name.ToString(0), vs.Value.ToString(0))
}

type ReturnStatement struct {
	ReturnToken token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.ReturnToken.Literal }
func (rs *ReturnStatement) ToString(indentLevel int) string {

	return fmt.Sprintf("ReturnStatement{%s,Expression {}}", rs.ReturnToken.ToString(0))
}

type ExpressionStatement struct {
	ExpressionToken token.Token
	Expression      Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.ExpressionToken.Literal }
func (es *ExpressionStatement) ToString(indentLevel int) string {
	return fmt.Sprintf("%sExpression{%s}", strings.Repeat("\t", indentLevel+1), es.Expression.ToString(0))
}
func (es *ExpressionStatement) expressionNode() {

}

type PrefixExpression struct {
	PrefixToken token.Token
	Operator    string
	Right       Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.PrefixToken.Literal }
func (pe *PrefixExpression) ToString(indentLevel int) string {
	return fmt.Sprintf("(%s%s)", pe.Operator, pe.Right.ToString(0))
}

type InfixExpression struct {
	InfixToken token.Token
	Left       Expression
	Operator   string
	Right      Expression
}

func (ie *InfixExpression) infixinterface()        {}
func (ie *InfixExpression) expressionNode()        {}
func (ie *InfixExpression) TokenLiteral() string   { return ie.InfixToken.Literal }
func (ie *InfixExpression) Token() token.TokenType { return ie.InfixToken.Type }
func (ie *InfixExpression) ToString(indentLevel int) string {
	// indentLevel := 0
	return "\n" + strings.Repeat("\t", indentLevel+2) + fmt.Sprintf("InfixToken{%s},", ie.InfixToken.ToString(indentLevel+2)) + "\n" + strings.Repeat("\t", indentLevel+3) + fmt.Sprintf("Left(%s),", ie.Left.ToString(indentLevel+2)) + "\n" + strings.Repeat("\t", indentLevel+3) + fmt.Sprintf("Operator(%s),", ie.Operator) + "\n" + strings.Repeat("\t", indentLevel+3) + fmt.Sprintf("Right(%s)", ie.Right.ToString(indentLevel+2))
	// return "\n" + strings.Repeat("\t", indentLevel) + fmt.Sprintf("InfixToken{ %s}, \n\tLeft(%s),\n \tOperator(%s),\n \tRight(%s),\n", ie.InfixToken.ToString(indentLevel+1), ie.Left.ToString(indentLevel+1), ie.Operator, ie.Right.ToString(indentLevel+1))

}

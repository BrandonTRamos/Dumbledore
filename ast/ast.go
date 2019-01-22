package ast

import (
	"Dumbledore/token"
	"fmt"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
	ToString() string
}

type Expression interface {
	Node
	experssionNode()
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
func (i *Identifier) toString() string {
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

	return fmt.Sprintf("VarStatement{%s,%s,Expression {}}", vs.VarToken.ToString(), vs.Name.toString())
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

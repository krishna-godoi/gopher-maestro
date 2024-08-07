package ast

import "github.com/krishna-godoi/gopher-maestro/token"

type Node interface {
	Literal() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression string

type Program struct {
	Statements []Statement
}

func (p *Program) Literal() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].Literal()
	} else {
		return ""
	}
}

type VarStatement struct {
	Token      token.Token
	Type       string
	Context    string
	Identifier Identifier
	Value      string
}

func (vs *VarStatement) statementNode() {}
func (vs *VarStatement) Literal() string {
	return vs.Token.Literal
}

type Identifier struct {
	Token string
	Name  string
}

func (i *Identifier) Literal() string {
	return i.Name
}

type FuncStatement struct {
	Token    token.Token
	Context  string
	Name     Identifier
	Children []Statement
}

func (vs *FuncStatement) statementNode() {}
func (vs *FuncStatement) Literal() string {
	return vs.Token.Literal
}

type ForStatement struct {
	Token     token.Token
	Context   string
	Variable  *VarStatement
	Condition string
	Increment string
	Scope     []Statement
}

func (vs *ForStatement) statementNode() {}
func (vs *ForStatement) Literal() string {
	return vs.Token.Literal
}

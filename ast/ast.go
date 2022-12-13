package ast

import "jpg/tokentype"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	experssionNode()
}

type Program struct {
	Body []Statement
}

func (p Program) TokenLiteral() string {
	if len(p.Body) > 0 {
		return p.Body[0].TokenLiteral()
	} else {
		return ""
	}
}

// let statement、const statement、var statement
type VariableDeclaration struct {
	Declarations []Node
	Kind         tokentype.TokenType // specify let、var、const
}

func (vd *VariableDeclaration) TokenLiteral() string {
	return "VariableDeclaration"
}

func (vd *VariableDeclaration) statementNode() {}

type VariableDeclarator struct {
	Id   *Identifier
	Init Node
}

func (vdr *VariableDeclarator) TokenLiteral() string {
	return "VariableDeclarator"
}

func (vdr *VariableDeclarator) experssionNode() {

}

type Identifier struct {
	Token *tokentype.Token
	Name  string // the variable bindIdentfier name, for example, let a = 1, then a is the name
}

func (i *Identifier) expressionNode() {

}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

type Literal struct {
	Token tokentype.Token
}

func (l *Literal) TokenLiteral() string {
	return l.Token.Literal
}

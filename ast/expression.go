package ast

import "github.com/devfojo/ph/token"

type Expression interface {
	Node
	expressionNode()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) expressionNode() {}

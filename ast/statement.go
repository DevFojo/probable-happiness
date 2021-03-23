package ast

import "github.com/devfojo/ph/token"

type Statement interface {
	Node
	statementNode()
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
}

func (l *LetStatement) statementNode() {}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (r *ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}

func (r *ReturnStatement) statementNode() {}

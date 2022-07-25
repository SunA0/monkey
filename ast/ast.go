package ast

import (
	"monkey/token"
)

type Node interface {
	TokenLiteral() string // 返回词法单元的字面量
}

type Statement interface {
	Node
	StatementNode()
}

type Expression interface {
	Node
	ExpressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

var _ Statement = new(LetStatement)

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

//func (ls *LetStatement) StatementNode() {
//	panic("implement me")
//}

func (ls *LetStatement) StatementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}

func (l *Identifier) statementNode() {}
func (l *Identifier) TokenLiteral() string {
	return l.Token.Literal
}

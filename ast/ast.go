package ast

import (
	"bytes"
	"github.com/SunA0/monkey/token"
)

type Node interface {
	TokenLiteral() string // 返回词法单元的字面量
	String() string
}

// Statement 语句
type Statement interface {
	Node
	StatementNode()
}

// Expression 表达式
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

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (l *Identifier) ExpressionNode() {
}

func (l *Identifier) statementNode() {}
func (l *Identifier) TokenLiteral() string {
	return l.Token.Literal
}
func (l *Identifier) String() string {
	return l.Value
}

// let Let语句
var _ Statement = new(LetStatement)

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	return out.String()
}
func (ls *LetStatement) StatementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// return Return语句
var _ Statement = new(ReturnStatement)

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (ls *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	if ls.ReturnValue != nil {
		out.WriteString(ls.ReturnValue.String())
	}

	out.WriteString(";")
	return out.String()
}

func (ls *ReturnStatement) StatementNode() {}
func (ls *ReturnStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// expression 表达式
var _ Statement = new(ExpressionStatement)

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (e *ExpressionStatement) String() string {
	if e.Expression != nil {
		return e.Expression.String()
	}
	return ""
}

func (e *ExpressionStatement) TokenLiteral() string {
	return e.Token.Literal
}

func (e *ExpressionStatement) StatementNode() {
}

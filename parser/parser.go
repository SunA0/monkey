package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

// Parser 语法分析器
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token

	error []string
}

func NewParser(l *lexer.Lexer) *Parser {

	p := &Parser{
		l:     l,
		error: []string{},
	}
	p.NextToken()
	p.NextToken()

	return p
}

func (p *Parser) NextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram 解析AST根节点
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseLetStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.NextToken()
	}
	return program
}

// ParseStatement 解析语句
func (p *Parser) ParseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.curTokenIs(token.SEMICOLON) {
		p.NextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.NextToken()

	// TODO
	for !p.curTokenIs(token.SEMICOLON) {
		p.NextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.NextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) Errors() []string {
	return p.error
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s ,got %s instead", t, p.peekToken.Type)
	p.error = append(p.error, msg)
}

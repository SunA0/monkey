package ast

import (
	"github.com/SunA0/monkey/token"
	"testing"
)

func TestString(t *testing.T) {
	// AST
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "test"},
					Value: "test",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "val"},
					Value: "val",
				},
			},
		},
	}
	if program.String() != "let test = val" {
		t.Errorf("program.string() wrong. got %q", program.String())
	}
}

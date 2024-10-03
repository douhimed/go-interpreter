package ast

import (
	"go-interpreter/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherValue"},
					Value: "anotherValue",
				},
			},
		},
	}
	actualValue := program.String()
	if actualValue != "let myVar = anotherValue;" {
		t.Errorf("program.String() wrong. got=%q", actualValue)
	}
}

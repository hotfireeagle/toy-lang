package ast

import (
	"testing"
	"toy/tokentype"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: tokentype.New(tokentype.LET, "let"),
				Name: &Identifier{
					Token: tokentype.New(tokentype.IDENTIFIER, "myVar"),
					Value: "myVar",
				},
				Value: &Identifier{
					Token: tokentype.New(tokentype.IDENTIFIER, "anotherVar"),
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong, got=%q", program.String())
	}
}

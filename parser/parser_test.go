package parser

import (
	"jpg/ast"
	"jpg/lexer"
	"jpg/reader"
	"jpg/tokentype"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := "const a = 3;"
	r := reader.New(reader.TextMode, input)
	l := lexer.New(r)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() retuned nil")
	}

	if len(program.Body) != 1 {
		t.Fatalf("expected program.Body's length to be 1, but got %v", len(program.Body))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"a"},
	}

	for i, tt := range tests {
		stmt := program.Body[i]
		if !testVariableDeclaration(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testVariableDeclaration(t *testing.T, s ast.Node, name string) bool {
	// if s.TokenLiteral() != "let" {
	// 	t.Errorf("testVariableDeclaration err")
	// 	return false
	// }

	variableDeclarationStmt, ok := s.(*ast.VariableDeclaration)
	if !ok {
		t.Errorf("wrong type")
		return false
	}

	if variableDeclarationStmt.Kind != tokentype.CONST {
		t.Errorf("wrong token kind type")
		return false
	}

	for _, declarator := range variableDeclarationStmt.Declarations {
		dobj, ok := declarator.(*ast.VariableDeclarator)
		if !ok {
			t.Errorf("wrong Declarations type")
		}
		if dobj.Id.Name != "a" {
			t.Errorf("wrong variable name")
		}
	}

	return true
}

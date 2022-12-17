package parser

import (
	"jpg/ast"
	"jpg/lexer"
	"jpg/reader"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 838383;
	`
	r := reader.New(reader.TextMode, input)
	l := lexer.New(r)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() retuned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("expected program.Body's length to be 3, but got %v", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testVariableDeclaration(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testVariableDeclaration(t *testing.T, s ast.Node, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("testVariableDeclaration err")
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("wrong type")
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("wrong token kind type")
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("wrong token kind type2")
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

package parser

import (
	"testing"
	"toy/ast"
	"toy/lexer"
	"toy/reader"
	"toy/tokentype"
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

func TestReturnStatments(t *testing.T) {
	input := `
		return 5;
		return 10;
		return 993322;
	`

	r := reader.New(reader.TextMode, input)
	l := lexer.New(r)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements, got %d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)

		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement, got %T", stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
		}
	}
}

func TestExpressStatement(t *testing.T) {
	input := "foobar;"

	r := reader.New(reader.TextMode, input)
	l := lexer.New(r)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements, got %d", len(program.Statements))
	}

	expressionStatement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("TestExpressStatement should just contain ExpressionStatement, but got %T", expressionStatement)
	}

	ident, ok := expressionStatement.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("expression wrong type")
	}

	if ident.Value != "foobar" {
		t.Fatalf("wrong value")
	}

	if ident.TokenLiteral() != "foobar" {
		t.Fatalf("wrong literal")
	}
}

func TestIngeterStatement(t *testing.T) {
	input := "30;"
	r := reader.New(reader.TextMode, input)
	l := lexer.New(r)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements should be 1, but got %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0]'s Type not ExpressionStatement, got %T", program.Statements[0])
	}

	testIntegerLiteral(t, stmt.Expression, 30)
}

// test this expression is intergerLiteral
func testIntegerLiteral(t *testing.T, il ast.Expression, num int64) bool {
	integerExpression, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("integerExpression type should be intergerLiteral, but got %T", integerExpression)
		return false
	}

	if integerExpression.Token.Type != tokentype.NUM {
		t.Fatalf("interExpression.Token.Type should be num, but got %v", integerExpression.Token.Type)
		return false
	}

	if integerExpression.Value != num {
		t.Fatalf("interExpression.Value does's match, got %v, expected to be %v", integerExpression.Value, num)
		return false
	}

	return true
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}

	for _, tt := range prefixTests {
		r := reader.New(reader.TextMode, tt.input)
		l := lexer.New(r)
		p := New(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n", 1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)

		if !ok {
			t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s", tt.operator, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, tt.integerValue) {
			return
		}
	}
}

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input    string
		left     int64
		operator string
		right    int64
	}{
		{"5 + 5", 5, "+", 5},
		{"5 - 5", 5, "-", 5},
		{"5 * 5", 5, "*", 5},
		{"5 / 5", 5, "/", 5},
		{"5 > 5", 5, ">", 5},
		{"5 < 5", 5, "<", 5},
		{"5 == 5", 5, "==", 5},
		{"5 != 5", 5, "!=", 5},
	}

	for _, infixTest := range infixTests {
		r := reader.New(reader.TextMode, infixTest.input)
		l := lexer.New(r)
		p := New(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Errorf("program.Statements length not 1, got %d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("exp is not ast.InfixExpression. got=%T", stmt.Expression)
		}

		if !testIntegerLiteral(t, exp.Left, infixTest.left) {
			return
		}

		if exp.Operator != infixTest.operator {
			t.Fatalf("exp.operator is not '%s', got=%s", infixTest.operator, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, infixTest.right) {
			return
		}
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a+b+c",
			"((a + b) + c)",
		},
		{
			"a+b-c",
			"((a + b) - c)",
		},
		{
			"a*b*c",
			"((a * b) * c)",
		},
		{
			"a*b/c",
			"((a * b) / c)",
		},
		{
			"a+b/c",
			"(a + (b / c))",
		},
		{
			"a+b*c+d/e-f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3+4;-5*5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5>4==3<4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
	}

	for _, tt := range tests {
		r := reader.New(reader.TextMode, tt.input)
		l := lexer.New(r)
		p := New(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		actual := program.String()
		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}

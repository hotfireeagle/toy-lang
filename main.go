package main

import (
	"fmt"
	"toy/lexer"
	"toy/parser"
	"toy/reader"
)

func TestOperatorPrecedenceParsing() {
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
		p := parser.New(l)

		program := p.ParseProgram()

		actual := program.String()
		if actual != tt.expected {
			fmt.Printf("expected=%q, got=%q\n", tt.expected, actual)
		}
	}
}

func main() {
	filePath := "/Users/smallhai/learn/gitRepo/toy/test.js"
	r := reader.New(reader.FileMode, filePath)
	l := lexer.New(r)
	c := 1
	for tok := l.NextToken(); !tok.IsEof(); {
		c++
		fmt.Println(tok)
		tok = l.NextToken()
	}
	fmt.Println(c)

	TestOperatorPrecedenceParsing()

	// u, err := user.Current()

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Hello %s, Try this!\n", u.Username)

	// repl.Start(os.Stdin, os.Stdout)
}

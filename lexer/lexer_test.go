package lexer

import (
	"acorn/reader"
	"testing"
)

func TestNextToken(t *testing.T) {
	codeString := "20 30 40"
	reader := reader.New(reader.TextMode, codeString)
	lexer := New(reader)

	token := lexer.NextToken()

	if token.Literal != "20" {
		t.Fatal("error num")
	}

	t2 := lexer.NextToken()

	if t2.Literal != "30" {
		t.Fatal("error num2")
	}
}

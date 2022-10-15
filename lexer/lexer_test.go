package lexer

import (
	"acorn/reader"
	"fmt"
	"testing"
)

func TestNextToken(t *testing.T) {
	codeString := "20 30 40"
	reader := reader.New(reader.TextMode, codeString)
	lexer := New(reader)

	token := lexer.NextToken()

	fmt.Println("token is >>>", token)

	if token.Literal != "20" {
		t.Fatal("error num")
	}

	t2 := lexer.NextToken()

	if t2.Literal != "30" {
		t.Fatal("error num2")
	}
}

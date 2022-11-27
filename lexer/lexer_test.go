package lexer

import (
	"jpg/reader"
	"testing"
)

func TestNextToken(t *testing.T) {
	codeString := "20 30 shjdhsjdhjsdhjshdjshdjshdsjdhjshdjsdsjdsj"
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

	t3 := lexer.NextToken()

	if t3.Literal != "shjdhsjdhjsdhjshdjshdjshdsjdhjshdjsdsjdsj" {
		t.Fatal("error t3")
	}
}

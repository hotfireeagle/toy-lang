package tokentype

import "fmt"

type Token struct {
	Type    TokenType // token类型
	Literal string    // token的字面值
}

func (t *Token) Print() {
	fmt.Printf("[Type: %s Literal: %s]", tokenTypeLiteral[t.Type], t.Literal)
}

func New(t TokenType, s string) *Token {
	return &Token{
		t,
		s,
	}
}

func (t *Token) IsEof() bool {
	return t.Type == EOF
}

func (t *Token) IsEnter() bool {
	return t.Type == ENTER
}

func (t *Token) IsComma() bool {
	return t.Type == COMMA
}

func (t *Token) IsSemi() bool {
	return t.Type == SEMI
}

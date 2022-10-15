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

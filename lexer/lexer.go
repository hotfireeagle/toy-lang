package lexer

import "acorn/tokentype"

type Lexer struct {
	input string

	position int

	lineStart int

	curLine int

	currentTokenType tokentype.TokenType

	start int

	end int
}

func New() *Lexer {
	return &Lexer{
		position:         0,
		lineStart:        0,
		curLine:          1,
		currentTokenType: tokentype.EOF,
		start:            0,
		end:              0,
	}
}

func (l *Lexer) NextToken() {

}

// func (l *Lexer) Skip

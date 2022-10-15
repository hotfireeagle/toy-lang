package lexer

import (
	"acorn/constant"
	"acorn/languagespec"
	"acorn/reader"
	"acorn/tokentype"
	"strings"
)

type Lexer struct {
	reader reader.InputReader
}

func New(r reader.InputReader) *Lexer {
	return &Lexer{
		reader: r,
	}
}

func (l *Lexer) ConsumerWord(str strings.Builder) (string, bool) {
	rune := l.reader.NextRune()

	str.WriteRune(rune)

	if languagespec.LanguageNFA.Match(str.String()) {
		// 继续贪心下一个字符是否match，如果不match的话，那么返回当前匹配结果，否则消费它
		var greedMatch func() (string, bool)
		greedMatch = func() (string, bool) {
			nextRune, err := l.reader.PeekNextNRune(1)
			if err != nil {
				panic("TODO:")
			}
			if nextRune == constant.EOF {
				return str.String(), false
			} else {
				oneMoreStr := str.String() + string(nextRune)
				if languagespec.LanguageNFA.Match(oneMoreStr) {
					// 继续贪
					r := l.reader.NextRune()
					str.WriteRune(r)
					return greedMatch()
				} else {
					return str.String(), false
				}
			}
		}
		return greedMatch()
	} else {
		if rune == constant.EOF {
			return "", true
		}
		return l.ConsumerWord(str)
	}
}

func (l *Lexer) NextToken() *tokentype.Token {
	var str strings.Builder

	word, isUnvalid := l.ConsumerWord(str)

	if isUnvalid {
		panic("不接受的语法")
	} else {
		word = strings.TrimSpace(word)
	}

	switch word {
	case "[":
		return tokentype.New(tokentype.IDENTIFIER, "[")
	default:
		if languagespec.CheckIsNormalNum(word) {
			return tokentype.New(tokentype.NUM, word)
		}
		return tokentype.New(tokentype.INVALID, "err")
	}
}

// func (l *Lexer) isWhiteSpaceOrBreak(r rune) bool {
// 	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
// }

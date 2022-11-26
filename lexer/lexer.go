package lexer

import (
	"jpg/constant"
	"jpg/languagespec"
	"jpg/reader"
	"jpg/tokentype"
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

func (l *Lexer) ConsumerWord(str *strings.Builder) (string, bool) {
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
	word, isUnvalid := l.ConsumerWord(&str)

	if isUnvalid {
		panic("不接受的语法")
	} else {
		word = strings.TrimSpace(word)
	}

	switch word {
	case "var":
		return tokentype.New(tokentype.VAR, word)
	case "let":
		return tokentype.New(tokentype.LET, word)
	case "const":
		return tokentype.New(tokentype.CONST, word)
	case "true":
		return tokentype.New(tokentype.TRUE, word)
	case "false":
		return tokentype.New(tokentype.FALSE, word)
	case "undefined":
		return tokentype.New(tokentype.UNDEFINED, word)
	case "null":
		return tokentype.New(tokentype.NULL, word)
	case "if":
		return tokentype.New(tokentype.IF, word)
	case "else":
		return tokentype.New(tokentype.ELSE, word)
	case "elseif":
		return tokentype.New(tokentype.ELSEIF, word)
	case "for":
		return tokentype.New(tokentype.FOR, word)
	case "while":
		return tokentype.New(tokentype.WHILE, word)
	case "do":
		return tokentype.New(tokentype.DO, word)
	case "[":
		return tokentype.New(tokentype.BRACKETL, word)
	case "]":
		return tokentype.New(tokentype.BRACKETR, word)
	case "{":
		return tokentype.New(tokentype.BRACEL, word)
	case "}":
		return tokentype.New(tokentype.BRACER, word)
	case "(":
		return tokentype.New(tokentype.PARENL, word)
	case ")":
		return tokentype.New(tokentype.PARENR, word)
	case ",":
		return tokentype.New(tokentype.COMMA, word)
	case ";":
		return tokentype.New(tokentype.SEMI, word)
	case ":":
		return tokentype.New(tokentype.COLON, word)
	case ".":
		return tokentype.New(tokentype.DOT, word)
	case "?":
		return tokentype.New(tokentype.QUESTION, word)
	case "?.":
		return tokentype.New(tokentype.QUESTIONDOT, word)
	case "=>":
		return tokentype.New(tokentype.ARROW, word)
	case "...":
		return tokentype.New(tokentype.ELLIPSIS, word)
	case "=":
		return tokentype.New(tokentype.EQ, word)
	case "==":
		return tokentype.New(tokentype.EQUALITY, word)
	case "|":
		return tokentype.New(tokentype.BITWISEOR, word)
	case "^":
		return tokentype.New(tokentype.BITWISEXOR, word)
	case "&":
		return tokentype.New(tokentype.BITWISEAND, word)
	case "||":
		return tokentype.New(tokentype.LOGICOR, word)
	case "&&":
		return tokentype.New(tokentype.LOGICAND, word)
	case "+":
		return tokentype.New(tokentype.PLUS, word)
	case "-":
		return tokentype.New(tokentype.MIN, word)
	case "%":
		return tokentype.New(tokentype.MODULO, word)
	case "<<":
		return tokentype.New(tokentype.BITLEFTSHIFT, word)
	case ">>":
		return tokentype.New(tokentype.BITRIGHTSHIFT, word)
	case ">>>":
		return tokentype.New(tokentype.BITRIGHTSHIFT3, word)
	case "break":
		return tokentype.New(tokentype.BREAK, word)
	case "case":
		return tokentype.New(tokentype.CASE, word)
	case "catch":
		return tokentype.New(tokentype.CATCH, word)
	case "continue":
		return tokentype.New(tokentype.CONTINUE, word)
	case "default":
		return tokentype.New(tokentype.DEFAULT, word)
	case "finally":
		return tokentype.New(tokentype.FINALLY, word)
	case "function":
		return tokentype.New(tokentype.FUNCTION, word)
	case "return":
		return tokentype.New(tokentype.RETURN, word)
	case "switch":
		return tokentype.New(tokentype.SWITCH, word)
	case "throw":
		return tokentype.New(tokentype.THROW, word)
	case "try":
		return tokentype.New(tokentype.TRY, word)
	case "with":
		return tokentype.New(tokentype.WITH, word)
	case "new":
		return tokentype.New(tokentype.NEW, word)
	case "this":
		return tokentype.New(tokentype.THIS, word)
	case "super":
		return tokentype.New(tokentype.SUPER, word)
	case "class":
		return tokentype.New(tokentype.CLASS, word)
	case "extends":
		return tokentype.New(tokentype.EXTENDS, word)
	case "export":
		return tokentype.New(tokentype.EXPORT, word)
	case "import":
		return tokentype.New(tokentype.IMPORT, word)
	case "in":
		return tokentype.New(tokentype.IN, word)
	case "instanceof":
		return tokentype.New(tokentype.INSTANCEOF, word)
	case "typeof":
		return tokentype.New(tokentype.TYPEOF, word)
	case "void":
		return tokentype.New(tokentype.VOID, word)
	case "delete":
		return tokentype.New(tokentype.DELETE, word)
	default:
		if languagespec.CheckIsNormalNum(word) {
			return tokentype.New(tokentype.NUM, word)
		} else if languagespec.CheckIsIdentfier(word) {
			return tokentype.New(tokentype.IDENTIFIER, word)
		} else if languagespec.CheckIsString(word) {
			return tokentype.New(tokentype.STRING, word)
		} else if languagespec.CheckIsComment(word) {
			// ignore comment
			return l.NextToken()
		}
		return tokentype.New(tokentype.INVALID, "err")
	}
}

// func (l *Lexer) isWhiteSpaceOrBreak(r rune) bool {
// 	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
// }

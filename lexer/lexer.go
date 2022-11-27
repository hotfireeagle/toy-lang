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

// func (l *Lexer) ConsumerWord(str *strings.Builder) (string, bool) {
// 	rune := l.reader.NextRune()

// 	str.WriteRune(rune)

// 	if languagespec.LanguageNFA.Match(str.String()) {
// 		// 继续贪心下一个字符是否match，如果不match的话，那么返回当前匹配结果，否则消费它
// 		var greedMatch func() (string, bool)
// 		greedMatch = func() (string, bool) {
// 			nextRune, err := l.reader.PeekNextNRune(1)
// 			if err != nil {
// 				panic("TODO:")
// 			}
// 			if nextRune == constant.EOF {
// 				return str.String(), false
// 			} else {
// 				oneMoreStr := str.String() + string(nextRune)
// 				if languagespec.LanguageNFA.Match(oneMoreStr) {
// 					// 继续贪
// 					r := l.reader.NextRune()
// 					str.WriteRune(r)
// 					return greedMatch()
// 				} else {
// 					return str.String(), false
// 				}
// 			}
// 		}
// 		return greedMatch()
// 	} else {
// 		if rune == constant.EOF {
// 			return "", true
// 		}
// 		return l.ConsumerWord(str)
// 	}
// }

func (l *Lexer) NextToken() (*tokentype.Token, bool) {
	var lastMatchTokenType tokentype.TokenType
	var sb strings.Builder

	var greed func() (*tokentype.Token, bool)

	greed = func() (*tokentype.Token, bool) {
		ru := l.reader.NextRune()

		// TODO: not correct when we plan to support the template string
		if ru == 10 {
			return l.NextToken()
		}

		if ru == constant.EOF {
			if lastMatchTokenType == 0 {
				panic("Invalid syntax")
			} else {
				return tokentype.New(lastMatchTokenType, strings.TrimSpace(sb.String())), true
			}
		} else {
			sb.WriteRune(ru)
			str := sb.String()

			if languagespec.Num10DFA.Match(str) || languagespec.NumB2DFA.Match(str) || languagespec.Num16DFA.Match(str) {
				lastMatchTokenType = tokentype.NUM
			} else if languagespec.StringDoubleDFA.Match(str) || languagespec.StringSingleDFA.Match(str) {
				lastMatchTokenType = tokentype.STRING
			} else if languagespec.VarDFA.Match(str) {
				lastMatchTokenType = tokentype.VAR
			} else if languagespec.LetDFA.Match(str) {
				lastMatchTokenType = tokentype.LET
			} else if languagespec.ConstDFA.Match(str) {
				lastMatchTokenType = tokentype.CONST
			} else if languagespec.TrueDFA.Match(str) {
				lastMatchTokenType = tokentype.TRUE
			} else if languagespec.FalseDFA.Match(str) {
				lastMatchTokenType = tokentype.FALSE
			} else if languagespec.UndefinedDFA.Match(str) {
				lastMatchTokenType = tokentype.UNDEFINED
			} else if languagespec.NullDFA.Match(str) {
				lastMatchTokenType = tokentype.NULL
			} else if languagespec.IfDFA.Match(str) {
				lastMatchTokenType = tokentype.IF
			} else if languagespec.ElseDFA.Match(str) {
				lastMatchTokenType = tokentype.ELSE
			} else if languagespec.ElseIfDFA.Match(str) {
				lastMatchTokenType = tokentype.ELSEIF
			} else if languagespec.ForDFA.Match(str) {
				lastMatchTokenType = tokentype.FOR
			} else if languagespec.WhileDFA.Match(str) {
				lastMatchTokenType = tokentype.WHILE
			} else if languagespec.DoDFA.Match(str) {
				lastMatchTokenType = tokentype.DO
			} else if languagespec.BracketLDFA.Match(str) {
				lastMatchTokenType = tokentype.BRACKETL
			} else if languagespec.BracketRDFA.Match(str) {
				lastMatchTokenType = tokentype.BRACKETR
			} else if languagespec.BracelLDFA.Match(str) {
				lastMatchTokenType = tokentype.BRACEL
			} else if languagespec.BracelRDFA.Match(str) {
				lastMatchTokenType = tokentype.BRACER
			} else if languagespec.ParenlLDFA.Match(str) {
				lastMatchTokenType = tokentype.PARENL
			} else if languagespec.ParenlRDFA.Match(str) {
				lastMatchTokenType = tokentype.PARENR
			} else if languagespec.CommaDFA.Match(str) {
				lastMatchTokenType = tokentype.COMMA
			} else if languagespec.SemiDFA.Match(str) {
				lastMatchTokenType = tokentype.SEMI
			} else if languagespec.ColonDFA.Match(str) {
				lastMatchTokenType = tokentype.COLON
			} else if languagespec.DotDFA.Match(str) {
				lastMatchTokenType = tokentype.DOT
			} else if languagespec.QuestionDFA.Match(str) {
				lastMatchTokenType = tokentype.QUESTION
			} else if languagespec.QuestionDotDFA.Match(str) {
				lastMatchTokenType = tokentype.QUESTIONDOT
			} else if languagespec.ArrowDFA.Match(str) {
				lastMatchTokenType = tokentype.ARROW
			} else if languagespec.EllipsisDFA.Match(str) {
				lastMatchTokenType = tokentype.ELLIPSIS
			} else if languagespec.EqualDFA.Match(str) {
				lastMatchTokenType = tokentype.EQ
			} else if languagespec.EqualityDFA.Match(str) {
				lastMatchTokenType = tokentype.EQUALITY
			} else if languagespec.BitwiseorDFA.Match(str) {
				lastMatchTokenType = tokentype.BITWISEOR
			} else if languagespec.BitwisexorDFA.Match(str) {
				lastMatchTokenType = tokentype.BITWISEXOR
			} else if languagespec.BitwiseandDFA.Match(str) {
				lastMatchTokenType = tokentype.BITWISEAND
			} else if languagespec.LogicorDFA.Match(str) {
				lastMatchTokenType = tokentype.LOGICOR
			} else if languagespec.LogicandDFA.Match(str) {
				lastMatchTokenType = tokentype.LOGICAND
			} else if languagespec.PlusDFA.Match(str) {
				lastMatchTokenType = tokentype.PLUS
			} else if languagespec.MinDFA.Match(str) {
				lastMatchTokenType = tokentype.MIN
			} else if languagespec.ModuloDFA.Match(str) {
				lastMatchTokenType = tokentype.MODULO
			} else if languagespec.BitleftshiftDFA.Match(str) {
				lastMatchTokenType = tokentype.BITLEFTSHIFT
			} else if languagespec.BitrightshiftDFA.Match(str) {
				lastMatchTokenType = tokentype.BITRIGHTSHIFT
			} else if languagespec.Bitrightshift3DFA.Match(str) {
				lastMatchTokenType = tokentype.BITRIGHTSHIFT3
			} else if languagespec.BreakDFA.Match(str) {
				lastMatchTokenType = tokentype.BREAK
			} else if languagespec.CaseDFA.Match(str) {
				lastMatchTokenType = tokentype.CASE
			} else if languagespec.CatchDFA.Match(str) {
				lastMatchTokenType = tokentype.CATCH
			} else if languagespec.ContinueDFA.Match(str) {
				lastMatchTokenType = tokentype.CONTINUE
			} else if languagespec.DefaultDFA.Match(str) {
				lastMatchTokenType = tokentype.DEFAULT
			} else if languagespec.FinallyDFA.Match(str) {
				lastMatchTokenType = tokentype.FINALLY
			} else if languagespec.FunctionDFA.Match(str) {
				lastMatchTokenType = tokentype.FUNCTION
			} else if languagespec.ReturnDFA.Match(str) {
				lastMatchTokenType = tokentype.RETURN
			} else if languagespec.SwitchDFA.Match(str) {
				lastMatchTokenType = tokentype.SWITCH
			} else if languagespec.ThrowDFA.Match(str) {
				lastMatchTokenType = tokentype.THROW
			} else if languagespec.TryDFA.Match(str) {
				lastMatchTokenType = tokentype.TRY
			} else if languagespec.WithDFA.Match(str) {
				lastMatchTokenType = tokentype.WITH
			} else if languagespec.NewDFA.Match(str) {
				lastMatchTokenType = tokentype.NEW
			} else if languagespec.ThisDFA.Match(str) {
				lastMatchTokenType = tokentype.THIS
			} else if languagespec.SuperDFA.Match(str) {
				lastMatchTokenType = tokentype.SUPER
			} else if languagespec.ClassDFA.Match(str) {
				lastMatchTokenType = tokentype.CLASS
			} else if languagespec.ExtendsDFA.Match(str) {
				lastMatchTokenType = tokentype.EXTENDS
			} else if languagespec.ExportDFA.Match(str) {
				lastMatchTokenType = tokentype.EXPORT
			} else if languagespec.ImportDFA.Match(str) {
				lastMatchTokenType = tokentype.IMPORT
			} else if languagespec.InDFA.Match(str) {
				lastMatchTokenType = tokentype.IN
			} else if languagespec.InstanceofDFA.Match(str) {
				lastMatchTokenType = tokentype.INSTANCEOF
			} else if languagespec.TypeofDFA.Match(str) {
				lastMatchTokenType = tokentype.TYPEOF
			} else if languagespec.VoidDFA.Match(str) {
				lastMatchTokenType = tokentype.VOID
			} else if languagespec.DeleteDFA.Match(str) {
				lastMatchTokenType = tokentype.DELETE
			} else {
				if languagespec.IdentfierDFA.Match(str) {
					lastMatchTokenType = tokentype.IDENTIFIER
				} else {
					if lastMatchTokenType != 0 {
						l.reader.Backtrack()
						s2 := str[0 : len(str)-1]
						return tokentype.New(lastMatchTokenType, strings.TrimSpace(s2)), false
					} else {
						return greed()
					}
				}
			}

			return greed()
		}
	}

	return greed()
}

// func (l *Lexer) NextToken() *tokentype.Token {
// 	var str strings.Builder
// 	word, isUnvalid := l.ConsumerWord(&str)

// 	if isUnvalid {
// 		panic("不接受的语法")
// 	} else {
// 		word = strings.TrimSpace(word)
// 	}

// 	switch word {
// 	case "var":
// 		return tokentype.New(tokentype.VAR, word)
// 	case "let":
// 		return tokentype.New(tokentype.LET, word)
// 	case "const":
// 		return tokentype.New(tokentype.CONST, word)
// 	case "true":
// 		return tokentype.New(tokentype.TRUE, word)
// 	case "false":
// 		return tokentype.New(tokentype.FALSE, word)
// 	case "undefined":
// 		return tokentype.New(tokentype.UNDEFINED, word)
// 	case "null":
// 		return tokentype.New(tokentype.NULL, word)
// 	case "if":
// 		return tokentype.New(tokentype.IF, word)
// 	case "else":
// 		return tokentype.New(tokentype.ELSE, word)
// 	case "elseif":
// 		return tokentype.New(tokentype.ELSEIF, word)
// 	case "for":
// 		return tokentype.New(tokentype.FOR, word)
// 	case "while":
// 		return tokentype.New(tokentype.WHILE, word)
// 	case "do":
// 		return tokentype.New(tokentype.DO, word)
// 	case "[":
// 		return tokentype.New(tokentype.BRACKETL, word)
// 	case "]":
// 		return tokentype.New(tokentype.BRACKETR, word)
// 	case "{":
// 		return tokentype.New(tokentype.BRACEL, word)
// 	case "}":
// 		return tokentype.New(tokentype.BRACER, word)
// 	case "(":
// 		return tokentype.New(tokentype.PARENL, word)
// 	case ")":
// 		return tokentype.New(tokentype.PARENR, word)
// 	case ",":
// 		return tokentype.New(tokentype.COMMA, word)
// 	case ";":
// 		return tokentype.New(tokentype.SEMI, word)
// 	case ":":
// 		return tokentype.New(tokentype.COLON, word)
// 	case ".":
// 		return tokentype.New(tokentype.DOT, word)
// 	case "?":
// 		return tokentype.New(tokentype.QUESTION, word)
// 	case "?.":
// 		return tokentype.New(tokentype.QUESTIONDOT, word)
// 	case "=>":
// 		return tokentype.New(tokentype.ARROW, word)
// 	case "...":
// 		return tokentype.New(tokentype.ELLIPSIS, word)
// 	case "=":
// 		return tokentype.New(tokentype.EQ, word)
// 	case "==":
// 		return tokentype.New(tokentype.EQUALITY, word)
// 	case "|":
// 		return tokentype.New(tokentype.BITWISEOR, word)
// 	case "^":
// 		return tokentype.New(tokentype.BITWISEXOR, word)
// 	case "&":
// 		return tokentype.New(tokentype.BITWISEAND, word)
// 	case "||":
// 		return tokentype.New(tokentype.LOGICOR, word)
// 	case "&&":
// 		return tokentype.New(tokentype.LOGICAND, word)
// 	case "+":
// 		return tokentype.New(tokentype.PLUS, word)
// 	case "-":
// 		return tokentype.New(tokentype.MIN, word)
// 	case "%":
// 		return tokentype.New(tokentype.MODULO, word)
// 	case "<<":
// 		return tokentype.New(tokentype.BITLEFTSHIFT, word)
// 	case ">>":
// 		return tokentype.New(tokentype.BITRIGHTSHIFT, word)
// 	case ">>>":
// 		return tokentype.New(tokentype.BITRIGHTSHIFT3, word)
// 	case "break":
// 		return tokentype.New(tokentype.BREAK, word)
// 	case "case":
// 		return tokentype.New(tokentype.CASE, word)
// 	case "catch":
// 		return tokentype.New(tokentype.CATCH, word)
// 	case "continue":
// 		return tokentype.New(tokentype.CONTINUE, word)
// 	case "default":
// 		return tokentype.New(tokentype.DEFAULT, word)
// 	case "finally":
// 		return tokentype.New(tokentype.FINALLY, word)
// 	case "function":
// 		return tokentype.New(tokentype.FUNCTION, word)
// 	case "return":
// 		return tokentype.New(tokentype.RETURN, word)
// 	case "switch":
// 		return tokentype.New(tokentype.SWITCH, word)
// 	case "throw":
// 		return tokentype.New(tokentype.THROW, word)
// 	case "try":
// 		return tokentype.New(tokentype.TRY, word)
// 	case "with":
// 		return tokentype.New(tokentype.WITH, word)
// 	case "new":
// 		return tokentype.New(tokentype.NEW, word)
// 	case "this":
// 		return tokentype.New(tokentype.THIS, word)
// 	case "super":
// 		return tokentype.New(tokentype.SUPER, word)
// 	case "class":
// 		return tokentype.New(tokentype.CLASS, word)
// 	case "extends":
// 		return tokentype.New(tokentype.EXTENDS, word)
// 	case "export":
// 		return tokentype.New(tokentype.EXPORT, word)
// 	case "import":
// 		return tokentype.New(tokentype.IMPORT, word)
// 	case "in":
// 		return tokentype.New(tokentype.IN, word)
// 	case "instanceof":
// 		return tokentype.New(tokentype.INSTANCEOF, word)
// 	case "typeof":
// 		return tokentype.New(tokentype.TYPEOF, word)
// 	case "void":
// 		return tokentype.New(tokentype.VOID, word)
// 	case "delete":
// 		return tokentype.New(tokentype.DELETE, word)
// 	default:
// 		if languagespec.CheckIsNormalNum(word) {
// 			return tokentype.New(tokentype.NUM, word)
// 		} else if languagespec.CheckIsIdentfier(word) {
// 			return tokentype.New(tokentype.IDENTIFIER, word)
// 		} else if languagespec.CheckIsString(word) {
// 			return tokentype.New(tokentype.STRING, word)
// 		} else if languagespec.CheckIsComment(word) {
// 			// ignore comment
// 			return l.NextToken()
// 		}
// 		return tokentype.New(tokentype.INVALID, "err")
// 	}
// }

// func (l *Lexer) isWhiteSpaceOrBreak(r rune) bool {
// 	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
// }

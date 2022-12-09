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
	isEnd  bool
}

func New(r reader.InputReader) *Lexer {
	return &Lexer{
		reader: r,
		isEnd:  false,
	}
}

func (l *Lexer) NextToken() *tokentype.Token {
	if l.isEnd {
		return tokentype.New(tokentype.EOF, "")
	}

	var lastMatchTokenType tokentype.TokenType
	var sb strings.Builder

	var greed func() *tokentype.Token

	greed = func() *tokentype.Token {
		ru := l.reader.NextRune()

		if ru == constant.Tab || ru == constant.Whitespace {
			currentStr := sb.String()
			if strings.TrimSpace(currentStr) == "" {
				sb.Reset()
				return greed()
			}
		}

		if ru == constant.Enter {
			hitedType, _ := l.checkIsFixedType(sb.String())

			if hitedType != tokentype.INVALID {
				return tokentype.New(hitedType, strings.TrimSpace(sb.String()))
			} else {
				// TODO: 确定一下这个换行符是否需要抛弃
				// sb.WriteRune(ru)
				return greed()
			}
		}

		if ru == constant.EOF {
			if lastMatchTokenType == tokentype.INVALID {
				invalidStr := sb.String()
				nakedInvalidStr := strings.TrimSpace(invalidStr)
				if len(nakedInvalidStr) > 0 {
					panic("Invalid syntax")
				} else {
					return tokentype.New(tokentype.EOF, "")
				}
			} else {
				l.isEnd = true
				return tokentype.New(lastMatchTokenType, strings.TrimSpace(sb.String()))
			}
		} else {
			sb.WriteRune(ru)
			str := sb.String()

			nextTokenType, matched := l.checkIsFixedType(str)

			if matched {
				lastMatchTokenType = nextTokenType
			}

			if !matched {
				if lastMatchTokenType != tokentype.INVALID {
					l.reader.Backtrack()
					s2 := str[0 : len(str)-1]
					return tokentype.New(lastMatchTokenType, strings.TrimSpace(s2))
				} else {
					return greed()
				}
			} else {
				return greed()
			}
		}
	}

	return greed()
}

func (l *Lexer) checkIsFixedType(str string) (tokentype.TokenType, bool) {
	var hitType tokentype.TokenType

	if languagespec.Num10DFA.Match(str) || languagespec.NumB2DFA.Match(str) || languagespec.Num16DFA.Match(str) || languagespec.FloatDFA.Match(str) {
		hitType = tokentype.NUM
	} else if languagespec.StringDoubleDFA.Match(str) || languagespec.StringSingleDFA.Match(str) {
		hitType = tokentype.STRING
	} else if languagespec.VarDFA.Match(str) {
		hitType = tokentype.VAR
	} else if languagespec.LetDFA.Match(str) {
		hitType = tokentype.LET
	} else if languagespec.ConstDFA.Match(str) {
		hitType = tokentype.CONST
	} else if languagespec.TrueDFA.Match(str) {
		hitType = tokentype.TRUE
	} else if languagespec.FalseDFA.Match(str) {
		hitType = tokentype.FALSE
	} else if languagespec.UndefinedDFA.Match(str) {
		hitType = tokentype.UNDEFINED
	} else if languagespec.NullDFA.Match(str) {
		hitType = tokentype.NULL
	} else if languagespec.IfDFA.Match(str) {
		hitType = tokentype.IF
	} else if languagespec.ElseDFA.Match(str) {
		hitType = tokentype.ELSE
	} else if languagespec.ElseIfDFA.Match(str) {
		hitType = tokentype.ELSEIF
	} else if languagespec.ForDFA.Match(str) {
		hitType = tokentype.FOR
	} else if languagespec.WhileDFA.Match(str) {
		hitType = tokentype.WHILE
	} else if languagespec.DoDFA.Match(str) {
		hitType = tokentype.DO
	} else if languagespec.BracketLDFA.Match(str) {
		hitType = tokentype.BRACKETL
	} else if languagespec.BracketRDFA.Match(str) {
		hitType = tokentype.BRACKETR
	} else if languagespec.BracelLDFA.Match(str) {
		hitType = tokentype.BRACEL
	} else if languagespec.BracelRDFA.Match(str) {
		hitType = tokentype.BRACER
	} else if languagespec.ParenlLDFA.Match(str) {
		hitType = tokentype.PARENL
	} else if languagespec.ParenlRDFA.Match(str) {
		hitType = tokentype.PARENR
	} else if languagespec.CommaDFA.Match(str) {
		hitType = tokentype.COMMA
	} else if languagespec.SemiDFA.Match(str) {
		hitType = tokentype.SEMI
	} else if languagespec.ColonDFA.Match(str) {
		hitType = tokentype.COLON
	} else if languagespec.DotDFA.Match(str) {
		hitType = tokentype.DOT
	} else if languagespec.QuestionDFA.Match(str) {
		hitType = tokentype.QUESTION
	} else if languagespec.QuestionDotDFA.Match(str) {
		hitType = tokentype.QUESTIONDOT
	} else if languagespec.ArrowDFA.Match(str) {
		hitType = tokentype.ARROW
	} else if languagespec.EllipsisDFA.Match(str) {
		hitType = tokentype.ELLIPSIS
	} else if languagespec.EqualDFA.Match(str) {
		hitType = tokentype.EQ
	} else if languagespec.GreaterDFA.Match(str) {
		hitType = tokentype.GREATER
	} else if languagespec.LessDFA.Match(str) {
		hitType = tokentype.LESS
	} else if languagespec.EqualityDFA.Match(str) {
		hitType = tokentype.EQUALITY
	} else if languagespec.BitwiseorDFA.Match(str) {
		hitType = tokentype.BITWISEOR
	} else if languagespec.BitwisexorDFA.Match(str) {
		hitType = tokentype.BITWISEXOR
	} else if languagespec.BitwiseandDFA.Match(str) {
		hitType = tokentype.BITWISEAND
	} else if languagespec.LogicorDFA.Match(str) {
		hitType = tokentype.LOGICOR
	} else if languagespec.LogicandDFA.Match(str) {
		hitType = tokentype.LOGICAND
	} else if languagespec.PlusDFA.Match(str) {
		hitType = tokentype.PLUS
	} else if languagespec.MinDFA.Match(str) {
		hitType = tokentype.MIN
	} else if languagespec.ModuloDFA.Match(str) {
		hitType = tokentype.MODULO
	} else if languagespec.BitleftshiftDFA.Match(str) {
		hitType = tokentype.BITLEFTSHIFT
	} else if languagespec.BitrightshiftDFA.Match(str) {
		hitType = tokentype.BITRIGHTSHIFT
	} else if languagespec.Bitrightshift3DFA.Match(str) {
		hitType = tokentype.BITRIGHTSHIFT3
	} else if languagespec.BreakDFA.Match(str) {
		hitType = tokentype.BREAK
	} else if languagespec.CaseDFA.Match(str) {
		hitType = tokentype.CASE
	} else if languagespec.CatchDFA.Match(str) {
		hitType = tokentype.CATCH
	} else if languagespec.ContinueDFA.Match(str) {
		hitType = tokentype.CONTINUE
	} else if languagespec.DefaultDFA.Match(str) {
		hitType = tokentype.DEFAULT
	} else if languagespec.FinallyDFA.Match(str) {
		hitType = tokentype.FINALLY
	} else if languagespec.FunctionDFA.Match(str) {
		hitType = tokentype.FUNCTION
	} else if languagespec.ReturnDFA.Match(str) {
		hitType = tokentype.RETURN
	} else if languagespec.SwitchDFA.Match(str) {
		hitType = tokentype.SWITCH
	} else if languagespec.ThrowDFA.Match(str) {
		hitType = tokentype.THROW
	} else if languagespec.TryDFA.Match(str) {
		hitType = tokentype.TRY
	} else if languagespec.WithDFA.Match(str) {
		hitType = tokentype.WITH
	} else if languagespec.NewDFA.Match(str) {
		hitType = tokentype.NEW
	} else if languagespec.ThisDFA.Match(str) {
		hitType = tokentype.THIS
	} else if languagespec.SuperDFA.Match(str) {
		hitType = tokentype.SUPER
	} else if languagespec.ClassDFA.Match(str) {
		hitType = tokentype.CLASS
	} else if languagespec.ExtendsDFA.Match(str) {
		hitType = tokentype.EXTENDS
	} else if languagespec.ExportDFA.Match(str) {
		hitType = tokentype.EXPORT
	} else if languagespec.ImportDFA.Match(str) {
		hitType = tokentype.IMPORT
	} else if languagespec.InDFA.Match(str) {
		hitType = tokentype.IN
	} else if languagespec.InstanceofDFA.Match(str) {
		hitType = tokentype.INSTANCEOF
	} else if languagespec.TypeofDFA.Match(str) {
		hitType = tokentype.TYPEOF
	} else if languagespec.VoidDFA.Match(str) {
		hitType = tokentype.VOID
	} else if languagespec.DeleteDFA.Match(str) {
		hitType = tokentype.DELETE
	} else if languagespec.SingleRowComment.Match(str) {
		hitType = tokentype.COMMENT
	} else if languagespec.NotDFA.Match(str) {
		hitType = tokentype.NOT
	} else if languagespec.Eq3DFA.Match(str) {
		hitType = tokentype.EQ3
	} else if languagespec.IdentfierDFA.Match(str) {
		hitType = tokentype.IDENTIFIER
	}

	return hitType, hitType != tokentype.INVALID
}

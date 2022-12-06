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

		// TODO: not correct when we plan to support the template string
		if ru == 10 {
			return l.NextToken()
		}

		if ru == constant.EOF {
			if lastMatchTokenType == 0 {
				panic("Invalid syntax")
			} else {
				l.isEnd = true
				return tokentype.New(lastMatchTokenType, strings.TrimSpace(sb.String()))
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
						return tokentype.New(lastMatchTokenType, strings.TrimSpace(s2))
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

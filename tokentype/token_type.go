package tokentype

type TokenType int

const (
	INVALID TokenType = iota
	NUM
	REGEXP
	STRING
	NAME
	EOF
	BRACKETL       // [
	BRACKETR       // ]
	BRACEL         // {
	BRACER         // }
	PARENL         // (
	PARENR         // )
	COMMA          // ,
	SEMI           // ;
	COLON          // :
	DOT            // .
	QUESTION       // ?
	QUESTIONDOT    // ?.
	ARROW          // =>
	ELLIPSIS       // ...
	BACKQUOTE      // `
	EQ             // =
	EQUALITY       // ==
	ASSIGN         // _=
	BITWISEOR      // |
	BITWISEXOR     // ^
	BITWISEAND     // &
	LOGICOR        // ||
	LOGICAND       // &&
	PLUS           // +
	MIN            // -
	MODULO         // %
	BITLEFTSHIFT   // <<
	BITRIGHTSHIFT  // >>
	BITRIGHTSHIFT3 // >>>
	BREAK          // break
	CASE           // case
	CATCH          // catch
	CONTINUE       // continue
	DEBUGGER       // debugger
	DEFAULT        // default
	DO             // do
	ELSE           // else
	FINALLY        // finally
	FOR            // for
	FUNCTION       // function
	IF             // if
	RETURN         // return
	SWITCH         // switch
	THROW          // throw
	TRY            // try
	VAR            // var
	CONST          // const
	WHILE          // while
	WITH           // with
	NEW            // new
	THIS           // this
	SUPER          // super
	CLASS          // class
	EXTENDS        // extends
	EXPORT         // export
	IMPORT         // import
	NULL           // null
	TRUE           // true
	FALSE          // false
	IN             // in
	INSTANCEOF     // instanceof
	TYPEOF         // typeof
	VOID           // void
	DELETE         // delete
)

var tokenTypeLiteral map[TokenType]string = map[TokenType]string{
	INVALID:        "invalid token",
	NUM:            "number",
	REGEXP:         "regexp",
	STRING:         "string",
	NAME:           "name",
	EOF:            "eof",
	BRACKETL:       "[",
	BRACKETR:       "]",
	BRACEL:         "{",
	BRACER:         "}",
	PARENL:         "(",
	PARENR:         ")",
	COMMA:          ",",
	SEMI:           ";",
	COLON:          ":",
	DOT:            ".",
	QUESTION:       "?",
	QUESTIONDOT:    "?.",
	ARROW:          "=>",
	ELLIPSIS:       "...",
	BACKQUOTE:      "`",
	EQ:             "=",
	EQUALITY:       "==",
	ASSIGN:         "_=",
	BITWISEOR:      "|",
	BITWISEXOR:     "^",
	BITWISEAND:     "&",
	LOGICOR:        "||",
	LOGICAND:       "&&",
	PLUS:           "+",
	MIN:            "-",
	MODULO:         "%",
	BITLEFTSHIFT:   "<<",
	BITRIGHTSHIFT:  ">>",
	BITRIGHTSHIFT3: ">>>",
	BREAK:          "break",
	CASE:           "case",
	CATCH:          "catch",
	CONTINUE:       "continue",
	DEBUGGER:       "debugger",
	DEFAULT:        "default",
	DO:             "do",
	ELSE:           "else",
	FINALLY:        "finally",
	FOR:            "for",
	FUNCTION:       "function",
	IF:             "if",
	RETURN:         "return",
	SWITCH:         "switch",
	THROW:          "throw",
	TRY:            "try",
	VAR:            "var",
	CONST:          "const",
	WHILE:          "while",
	WITH:           "with",
	NEW:            "new",
	THIS:           "this",
	SUPER:          "super",
	CLASS:          "class",
	EXTENDS:        "extends",
	EXPORT:         "export",
	IMPORT:         "import",
	NULL:           "null",
	TRUE:           "true",
	FALSE:          "false",
	IN:             "in",
	INSTANCEOF:     "instanceof",
	TYPEOF:         "typeof",
	VOID:           "void",
	DELETE:         "delete",
}

var keyWord map[TokenType]bool = map[TokenType]bool{
	BREAK:      true,
	CASE:       true,
	CATCH:      true,
	CONTINUE:   true,
	DEBUGGER:   true,
	DEFAULT:    true,
	DO:         true,
	ELSE:       true,
	FINALLY:    true,
	FOR:        true,
	FUNCTION:   true,
	IF:         true,
	RETURN:     true,
	SWITCH:     true,
	THROW:      true,
	TRY:        true,
	VAR:        true,
	CONST:      true,
	WHILE:      true,
	WITH:       true,
	NEW:        true,
	THIS:       true,
	SUPER:      true,
	CLASS:      true,
	EXTENDS:    true,
	EXPORT:     true,
	IMPORT:     true,
	NULL:       true,
	TRUE:       true,
	FALSE:      true,
	IN:         true,
	INSTANCEOF: true,
	TYPEOF:     true,
	VOID:       true,
	DELETE:     true,
}

func CheckTokenIsKeyword(t TokenType) bool {
	return keyWord[t]
}

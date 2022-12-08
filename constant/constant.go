package constant

import "errors"

const EOF = -1

const Enter rune = 10 // 换行符

var ErrNextRune = errors.New("YOU NEED CALL NEXTRUNE() FIRST")

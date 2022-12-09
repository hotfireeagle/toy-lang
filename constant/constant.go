package constant

import "errors"

const EOF = -1

const Whitespace rune = ' '
const Tab rune = '	'
const Enter rune = 10

var ErrNextRune = errors.New("YOU NEED CALL NEXTRUNE() FIRST")

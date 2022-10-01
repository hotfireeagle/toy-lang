package constant

import "errors"

const EOF = -1

var ErrNextRune = errors.New("YOU NEED CALL NEXTRUNE() FIRST")

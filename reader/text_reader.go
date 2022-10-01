package reader

import (
	"acorn/constant"
)

type TextReader struct {
	inputSourceCodes []rune
	rowNumber        int
	runePosition     int
	currentRune      rune
}

func NewTextReader(input string) *TextReader {
	return &TextReader{
		inputSourceCodes: []rune(input),
		rowNumber:        1,
		runePosition:     -1,
		currentRune:      constant.EOF,
	}
}

func (tr *TextReader) NextRune() rune {
	nextRunePosition := tr.runePosition + 1
	if nextRunePosition >= len(tr.inputSourceCodes) {
		return constant.EOF
	}

	nextRune := tr.inputSourceCodes[nextRunePosition]
	tr.runePosition = nextRunePosition
	tr.currentRune = nextRune

	return nextRune
}

func (tr *TextReader) PeekCurrentRune() (rune, error) {
	if tr.runePosition == -1 {
		return constant.EOF, constant.ErrNextRune
	}
	return tr.currentRune, nil
}

// 获取下几个位置的字符
func (tr *TextReader) PeekNextNRune(n int) (rune, error) {
	if tr.runePosition == -1 {
		return constant.EOF, constant.ErrNextRune
	}

	nPostion := tr.runePosition + n
	if nPostion >= len(tr.inputSourceCodes) {
		return constant.EOF, nil
	}

	return tr.inputSourceCodes[nPostion], nil
}

package reader

import "acorn/constant"

type TextReader struct {
	inputSourceCodes    []rune
	rowNumber           int
	currentRunePosition int
	nextRunePosition    int
	currentRune         rune
}

func NewTextReader(input string) *TextReader {
	return &TextReader{
		inputSourceCodes: []rune(input),
		rowNumber:        1,
	}
}

// 返回当前position所在位置的token字符
// 同时两个位置向前加1
func (tr *TextReader) ReadRune() rune {
	if tr.currentRunePosition >= len(tr.inputSourceCodes) {
		return constant.EOF
	}
	currentRune := tr.inputSourceCodes[tr.currentRunePosition]

	tr.currentRunePosition = tr.nextRunePosition
	tr.nextRunePosition += 1

	return currentRune
}

// 获取当前位置的字符
func (tr *TextReader) PeekCurrentRune() rune {
	if tr.currentRunePosition >= len(tr.inputSourceCodes) {
		return constant.EOF
	}

	return tr.inputSourceCodes[tr.currentRunePosition]
}

// 获取下几个位置的字符
func (tr *TextReader) PeekNextNRune(n int) rune {
	nPostion := tr.currentRunePosition + n
	if nPostion >= len(tr.inputSourceCodes) {
		return constant.EOF
	}

	return tr.inputSourceCodes[nPostion]
}

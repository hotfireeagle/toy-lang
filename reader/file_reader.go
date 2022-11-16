package reader

import (
	"bufio"
	"errors"
	"io"
	"jpg/constant"
	"os"
)

type FileReader struct {
	path      string
	reader    *bufio.Reader
	isEOF     bool
	line      string
	rowNumber int
	tr        *TextReader
}

func NewFileReader(p string) *FileReader {
	fileObj, openFileErr := os.Open(p)

	if openFileErr != nil {
		panic(openFileErr)
	}

	fileReader := bufio.NewReader(fileObj)

	fr := &FileReader{
		path:      p,
		rowNumber: 0,
		reader:    fileReader,
	}

	fr.readLine()

	return fr
}

func (fr *FileReader) readLine() {
	isEndOfFile := false

	fr.rowNumber += 1

	lineContent, err := fr.reader.ReadString('\n')

	if err != nil {
		if errors.Is(err, io.EOF) {
			isEndOfFile = true
		} else {
			panic(err)
		}
	}

	fr.line = lineContent
	fr.isEOF = isEndOfFile
	fr.tr = NewTextReader(lineContent)
}

func (fr *FileReader) NextRune() rune {
	if fr.isEOF && fr.line == "" {
		return constant.EOF
	}

	answer := fr.tr.NextRune()

	if answer == constant.EOF {
		fr.readLine()
		return fr.NextRune()
	} else {
		return answer
	}
}

func (fr *FileReader) PeekCurrentRune() (rune, error) {
	return fr.tr.PeekCurrentRune()
}

func (fr *FileReader) PeekNextNRune(n int) (rune, error) {
	return fr.tr.PeekNextNRune(n)
}

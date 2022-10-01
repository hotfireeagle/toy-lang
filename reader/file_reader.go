package reader

import "os"

type FileReader struct {
	path             string
	file             *os.File
	rowNumber        int
	columnNumber     int
	currentPosition  int
	nextReadPosition int
	currentRune      rune
}

func (fr *FileReader) NextRune() rune {
	return 0
}

func (fr *FileReader) PeekRune() rune {
	return 0
}

// func NewFileReader() *FileReader {
// 	return 0
// }

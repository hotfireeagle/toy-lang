package reader

type ReaderMode int

const (
	FileMode ReaderMode = iota
	TextMode
)

type InputReader interface {
	NextRune() rune
	PeekCurrentRune() (rune, error)
	PeekNextNRune(n int) (rune, error)
	Backtrack()
}

func New(mode ReaderMode, filePathOrContent string) InputReader {
	if mode == FileMode {
		return NewFileReader(filePathOrContent)
	} else if mode == TextMode {
		return NewTextReader(filePathOrContent)
	} else {
		panic("wrong mode")
	}
}

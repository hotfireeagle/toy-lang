package reader

import "testing"

func TestNextRune(t *testing.T) {
	filePath := "/Users/smallhai/Desktop/index.js"

	fileReader := NewFileReader(filePath)

	tests := []rune{
		' ',
		'a',
		'b',
		' ',
		'\n',
		's',
		's',
		'j',
		'd',
		's',
		'd',
		's',
		'd',
		'\n',
	}

	for i, r := range tests {
		nr := fileReader.NextRune()
		if r != nr {
			t.Fatalf("Expected to be %s, but got %s in case %d", string(r), string(nr), i)
		}
	}
}

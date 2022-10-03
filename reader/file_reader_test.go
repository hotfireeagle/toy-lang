package reader

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestNextRune(t *testing.T) {
	tempFile, err := ioutil.TempFile(os.TempDir(), "testcase.*.js")

	if err != nil {
		t.Fatal(err)
	}

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

	_, err = tempFile.WriteString(string(tests))
	defer os.Remove(tempFile.Name())

	if err != nil {
		t.Fatal(err)
	}

	fileReader := NewFileReader(tempFile.Name())

	for i, r := range tests {
		nr := fileReader.NextRune()
		if r != nr {
			t.Fatalf("Expected to be %s, but got %s in case %d", string(r), string(nr), i)
		}
	}
}

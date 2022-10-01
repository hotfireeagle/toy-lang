package reader

import (
	"acorn/constant"
	"testing"
)

func TestEmptyInputReader(t *testing.T) {
	tr := NewTextReader("")

	nr := tr.NextRune()

	if nr != constant.EOF {
		t.Fatalf("expected to be eof, but got %s", string(nr))
	}
}

func TestNormalInputReader(t *testing.T) {
	tr := NewTextReader(" abcde ")

	tests := []struct {
		expected rune
	}{
		{' '},
		{'a'},
		{'b'},
		{'c'},
		{'d'},
		{'e'},
	}

	nr := tr.NextRune()

	i := 0

	for nr != constant.EOF {
		if nr != tests[i].expected {
			t.Fatalf("expecd to be %s, but got %s", string(tests[i].expected), string(nr))
		}

		i += 1
		nr = tr.NextRune()
	}
}

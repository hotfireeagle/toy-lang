package reader

import (
	"jpg/constant"
	"testing"
)

func TestEmptyInputReader(t *testing.T) {
	tr := NewTextReader("")

	nr := tr.NextRune()

	if nr != constant.EOF {
		t.Fatalf("expected to be eof, but got %s.", string(nr))
	}
}

func TestNormalInputReader(t *testing.T) {
	tr := NewTextReader(" abcde ")

	tests := []struct {
		expected                rune
		expectedPeekCurrentRune rune
		n                       int
		expectedPeekNextNRune   rune
	}{
		{' ', ' ', 1, 'a'},
		{'a', 'a', 1, 'b'},
		{'b', 'b', 2, 'd'},
		{'c', 'c', 2, 'e'},
		{'d', 'd', 1, 'e'},
		{'e', 'e', 1, ' '},
	}

	nr := tr.NextRune()

	i := 0

	for nr != constant.EOF && i > len(tests) {
		if nr != tests[i].expected {
			t.Fatalf("expecd to be %s, but got %s.", string(tests[i].expected), string(nr))
		}

		pcr, _ := tr.PeekCurrentRune()

		if tests[i].expectedPeekCurrentRune != pcr {
			t.Fatalf("expected peekCurrentRune return %s, but got %s.", string(tests[i].expectedPeekCurrentRune), string(pcr))
		}

		nextResult, _ := tr.PeekNextNRune(tests[i].n)

		if tests[i].expectedPeekNextNRune != nextResult {
			t.Fatalf("expected peekNextNRune(%d) return %s, but got %s.", tests[i].n, string(tests[i].expectedPeekNextNRune), string(nextResult))
		}

		i += 1
		nr = tr.NextRune()
	}
}

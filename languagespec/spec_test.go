package languagespec

import "testing"

func Test_re_num_10(t *testing.T) {
	dfaObj := newDFA(re_num_10)
	testCases := []struct {
		input  string
		output bool
	}{
		{"0", true},
		{"11", true},
		{"123", true},
		{"99999", true},
		{"1111", true},
		{"0011", false},
		{"0b10", false},
		{"011", false},
		{"0x1", false},
	}

	for _, obj := range testCases {
		result := dfaObj.Match(obj.input)
		if result != obj.output {
			t.Errorf("Test_re_num_10 accurs error, in %s, got %v, but expected %v", obj.input, result, obj.output)
		}
	}
}

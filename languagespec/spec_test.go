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

func Test_re_num_binary2(t *testing.T) {
	dfaObj := newDFA(re_num_binary2)
	testCases := []struct {
		input  string
		output bool
	}{
		{"1", false},
		{"0", false},
		{"00", false},
		{"0b1", true},
		{"0b0", true},
		{"0b00", true},
		{"0b001", true},
		{"0b100", true},
		{"0b111", true},
	}

	for _, tObj := range testCases {
		result := dfaObj.Match(tObj.input)
		if result != tObj.output {
			t.Errorf("Test_re_num_binary2 error in %s, expected to be %v, but got %v", tObj.input, tObj.output, result)
		}
	}
}

func Test_re_num_16(t *testing.T) {
	dfa16Obj := newDFA(re_num_16)

	testCases := []struct {
		input  string
		output bool
	}{
		{"1				", false},
		{"0", false},
		{"111", false},
		{"34", false},
		{"0x0", true},
		{"0x1a", true},
		{"0x1j", false},
		{"0x323a1a121a32b", true},
	}

	for _, testObj := range testCases {
		result := dfa16Obj.Match(testObj.input)
		if result != testObj.output {
			t.Errorf("Test_re_num_16 error in %s, expect to be %v, but got %v", testObj.input, testObj.output, result)
		}
	}
}

func Test_re_identfier(t *testing.T) {
	dfaObj := newDFA(re_identfier)

	testCases := []struct {
		input  string
		output bool
	}{
		{"j1      ", true},
		{"			$1", true},
		{"中文", false}, // true in javascript, but we do't plan to support it now
		{"。1", false},
		{".1", false},
		{"_1", true},
		{"sjdksjdkshsd", true},
		{"_sjkdskd9032iosods", true},
	}

	for _, tobj := range testCases {
		result := dfaObj.Match(tobj.input)
		if result != tobj.output {
			t.Errorf("Test_re_identfier error in %s, expected to be %v, but got %v", tobj.input, tobj.output, result)
		}
	}
}

func Test_re_double_string(t *testing.T) {
	doubleStringDFA := newDFA(re_double_string)

	testCases := []struct {
		input  string
		output bool
	}{
		{`			"abc"`, true},
		{`"sdksdjks"`, true},
		{`"sjksds901010.sdsjd121jskds;la;d=1-10212jskd"`, true},
		{`"sjkdskjdksdks'sjksjdksnxmcx'sjdksjsdskds,121-210"`, true},
		{`'sdjksd'`, false},
		{`"sdsjdsk"sdsd"`, false},
	}

	for _, tobj := range testCases {
		result := doubleStringDFA.Match(tobj.input)
		if result != tobj.output {
			t.Errorf("Test_re_double_string error in %s, expected to be %v, but got %v", tobj.input, tobj.output, result)
		}
	}
}

func Test_re_single_string(t *testing.T) {
	singleStringDFA := newDFA(re_single_string)

	testCases := []struct {
		input  string
		output bool
	}{
		{`			'112122302-.sds'`, true},
		{`'112122'302-.sds'`, false},
		{`'sjdksdksjsdsksds.sowio11ksjdskz0iosds'`, true},
	}

	for _, tcase := range testCases {
		result := singleStringDFA.Match(tcase.input)
		if result != tcase.output {
			t.Errorf("Test_re_single_string error in %s, expected to be %v, but got %v", tcase.input, tcase.output, result)
		}
	}
}

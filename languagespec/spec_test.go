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
		{`			'112122302-.sds'      `, true},
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

func Test_re_var(t *testing.T) {
	varDFA := newDFA(re_var)

	varTestCases := []struct {
		input  string
		output bool
	}{
		{"var", true},
		{"	 var		  ", true},
		{"var ", true},
		{" var", true},
	}

	for _, testObj := range varTestCases {
		result := varDFA.Match(testObj.input)
		if result != testObj.output {
			t.Errorf("Test_re_var error in %s, expected to be %v, but got %v", testObj.input, testObj.output, result)
		}
	}
}

func Test_re_let(t *testing.T) {
	letDFA := newDFA(re_let)

	testCases := []struct {
		input  string
		output bool
	}{
		{"		let ", true},
		{"let", true},
		{" let", true},
	}

	for _, testCase := range testCases {
		result := letDFA.Match(testCase.input)
		if result != testCase.output {
			t.Errorf("Test_re_let error in %s, expected to be %v but got %v", testCase.input, testCase.output, result)
		}
	}
}

func Test_re_const(t *testing.T) {
	constDFA := newDFA(re_const)

	testCases := []struct {
		input  string
		output bool
	}{
		{"    const    ", true},
		{"			const  ", true},
		{"const ", true},
		{"  const", true},
		{"const", true},
	}

	for _, testObj := range testCases {
		result := constDFA.Match(testObj.input)
		if result != testObj.output {
			t.Errorf("Test_re_const error in %s, expected to be %v but got %v", testObj.input, testObj.output, result)
		}
	}
}

func Test_re_true(t *testing.T) {
	trueDFA := newDFA(re_true)

	testCaseList := []struct {
		input  string
		output bool
	}{
		{"true  ", true},
		{" true", true},
		{"	true 	", true},
		{"true", true},
	}

	for _, testObj := range testCaseList {
		result := trueDFA.Match(testObj.input)
		if result != testObj.output {
			t.Errorf("Test_re_true error in %s, expectd to be %v but got %v", testObj.input, testObj.output, result)
		}
	}
}

func Test_re_false(t *testing.T) {
	falseDFA := newDFA(re_false)

	testCases := []struct {
		input  string
		output bool
	}{
		{"false", true},
		{" false", true},
		{"false	  ", true},
		{" 	false 	", true},
	}

	for _, testObj := range testCases {
		result := falseDFA.Match(testObj.input)
		if result != testObj.output {
			t.Errorf("Test_re_false error in %s, expected to be %v but got %v", testObj.input, result, testObj.output)
		}
	}
}

func Test_re_undefined(t *testing.T) {
	undefinedDFA := newDFA(re_undefined)

	testCases := []struct {
		input  string
		output bool
	}{
		{"undefined", true},
		{" undefined", true},
		{"undefined  ", true},
		{"	undefined", true},
		{"undefined	", true},
		{"	 undefined	 ", true},
	}

	for _, to := range testCases {
		result := undefinedDFA.Match(to.input)
		if result != to.output {
			t.Errorf("Test_re_undefined error in %s, expected to be %v but got %v", to.input, to.output, result)
		}
	}
}

func Test_re_null(te *testing.T) {
	nullDFA := newDFA(re_null)

	tests := []struct {
		input  string
		output bool
	}{
		{"null", true},
		{" null", true},
		{"	null", true},
		{"null ", true},
		{"null	", true},
		{"	 null 	", true},
	}

	for _, t := range tests {
		result := nullDFA.Match(t.input)
		if result != t.output {
			te.Errorf("Test_re_null error in %s, expected to be %v but got %v", t.input, t.output, result)
		}
	}
}

func Test_re_if(t *testing.T) {
	ifDFA := newDFA(re_if)

	cases := []struct {
		input  string
		output bool
	}{
		{"if", true},
		{" if", true},
		{"if ", true},
		{"	if", true},
		{"if	", true},
		{"	 if 	", true},
	}

	for _, casev := range cases {
		result := ifDFA.Match(casev.input)
		if result != casev.output {
			t.Errorf("Test_re_if got error in %s, expected to be %v but got %v", casev.input, casev.output, result)
		}
	}
}

func Test_re_else(t *testing.T) {
	elseDFA := newDFA(re_else)

	cases := []struct {
		input  string
		output bool
	}{
		{"else", true},
		{" else", true},
		{"	else", true},
		{"else ", true},
		{"else	", true},
		{"	 else 	", true},
	}

	for _, casev := range cases {
		result := elseDFA.Match(casev.input)
		if result != casev.output {
			t.Errorf("Test_re_else got error in %s, expected to be %v, but got %v", casev.input, casev.output, result)
		}
	}
}

func Test_re_elseif(t *testing.T) {
	elseifDFA := newDFA(re_elseif)

	cases := []struct {
		input  string
		output bool
	}{
		{"else if", true},
		{" else  if ", true},
		{"	else	if	", true},
		{"	else	if", true},
		{"else	if	", true},
		{"	 else 	if	 ", true},
	}

	for _, cobj := range cases {
		result := elseifDFA.Match(cobj.input)
		if result != cobj.output {
			t.Errorf("Test_re_elseif got error in %s, expected to be %v but got %v", cobj.input, cobj.output, result)
		}
	}
}

func Test_re_for(t *testing.T) {
	fordfa := newDFA(re_for)

	testcases := []struct {
		input  string
		output bool
	}{
		{"for", true},
		{" for", true},
		{"	for", true},
		{"for ", true},
		{"for	", true},
		{"	 for 	", true},
	}

	for _, tobj := range testcases {
		result := fordfa.Match(tobj.input)
		if result != tobj.output {
			t.Errorf("Test_re_for got error in %s, expected to be %v but got %v", tobj.input, tobj.output, result)
		}
	}
}

func Test_re_while(t *testing.T) {
	whiledfa := newDFA(re_while)

	testcases := []struct {
		input  string
		output bool
	}{
		{"while", true},
		{" while", true},
		{"while ", true},
		{"	while", true},
		{"while	", true},
		{"	 while	 ", true},
	}

	for _, tc := range testcases {
		result := whiledfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_while got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_do(t *testing.T) {
	dodfa := newDFA(re_do)

	testcases := []struct {
		input  string
		output bool
	}{
		{"do", true},
		{"	do", true},
		{"do	", true},
		{"	 do 	", true},
	}

	for _, tc := range testcases {
		result := dodfa.Match(tc.input)

		if result != tc.output {
			t.Errorf("Test_re_do got")
		}
	}
}

func Test_re_single_row_comment(t *testing.T) {
	scDFA := newDFA(re_single_row_comment)
	cases := []struct {
		input  string
		output bool
	}{
		{"//sdshjds", true},
		{" // skksd ", true},
		{"	// ksdjksds ", true},
	}

	for _, co := range cases {
		result := scDFA.Match(co.input)
		if result != co.output {
			t.Errorf("Test_re_single_row_comment got err in %s, expected to be %v, but got %v", co.input, co.output, result)
		}
	}
}

func Test_re_left_bracket(t *testing.T) {
	lbdfa := newDFA(re_left_bracket)

	testCases := []struct {
		input  string
		output bool
	}{
		{" [", true},
		{"[ ", true},
		{"[", true},
		{" [ 		", true},
		{"]", false},
		{"/[", false},
		{"a", false},
	}

	for _, tc := range testCases {
		result := lbdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_left_bracket got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

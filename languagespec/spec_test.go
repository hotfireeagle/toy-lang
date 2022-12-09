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

func Test_re_num_float(t *testing.T) {
	dfaFloatObj := newDFA(re_num_float)

	testCases := []struct {
		input  string
		output bool
	}{
		{"121.2", true},
		{"12", false},
		{"000.23", true},
		{"1233.22", true},
		{"sds.sd", false},
		{"  212.23330 ", true},
	}

	for _, to := range testCases {
		result := dfaFloatObj.Match(to.input)
		if result != to.output {
			t.Errorf("Test_re_num_float got error in %s, expected to be %v, but got %v", to.input, to.output, result)
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

func Test_re_right_bracket(t *testing.T) {
	rrdfa := newDFA(re_right_bracket)

	testCases := []struct {
		input  string
		output bool
	}{
		{" ]", true},
		{"] ", true},
		{"]", true},
		{" ] 		", true},
		{"[", false},
		{"/]", false},
		{"a", false},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_right_bracket got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_left_bracel(t *testing.T) {
	rrdfa := newDFA(re_left_bracel)

	testCases := []struct {
		input  string
		output bool
	}{
		{" {", true},
		{"{ ", true},
		{"{", true},
		{" { 		", true},
		{"}", false},
		{"/{", false},
		{"a{", false},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_left_bracel got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_right_bracel(t *testing.T) {
	rrdfa := newDFA(re_right_bracel)

	testCases := []struct {
		input  string
		output bool
	}{
		{" }", true},
		{"} ", true},
		{"}", true},
		{" } 		", true},
		{"{", false},
		{"/}", false},
		{"a}", false},
		{".}", false},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_right_bracel got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_left_parenl(t *testing.T) {
	rrdfa := newDFA(re_left_parenl)

	testCases := []struct {
		input  string
		output bool
	}{
		{" (", true},
		{"( ", true},
		{"(", true},
		{" ( 		", true},
		{"(", true},
		{"/(", false},
		{"a(", false},
		{".(", false},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_left_parenl got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_right_parenl(t *testing.T) {
	rrdfa := newDFA(re_right_parenl)

	testCases := []struct {
		input  string
		output bool
	}{
		{" )", true},
		{") ", true},
		{")", true},
		{" ) 		", true},
		{")", true},
		{"/)", false},
		{"a)", false},
		{".)", false},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_right_parenl got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_comma(t *testing.T) {
	rrdfa := newDFA(re_comma)

	testCases := []struct {
		input  string
		output bool
	}{
		{" ,", true},
		{", ", true},
		{",", true},
		{" , 		", true},
		{",", true},
		{"/,", false},
		{"a,", false},
		{".,", false},
		{",.", false},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_comma got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_semi(t *testing.T) {
	rrdfa := newDFA(re_semi)

	testCases := []struct {
		input  string
		output bool
	}{
		{" ;", true},
		{"; ", true},
		{";", true},
		{" ; 		", true},
		{";", true},
		{"/;", false},
		{"a;", false},
		{".;", false},
		{";.", false},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_semi got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_colon(t *testing.T) {
	rrdfa := newDFA(re_colon)

	testCases := []struct {
		input  string
		output bool
	}{
		{" :", true},
		{": ", true},
		{":", true},
		{" : 		", true},
		{":", true},
		{"/:", false},
		{"a:", false},
		{".:", false},
		{":.", false},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_colon got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_dot(t *testing.T) {
	rrdfa := newDFA(re_dot)

	testCases := []struct {
		input  string
		output bool
	}{
		{" .", true},
		{". ", true},
		{".", true},
		{" . 		", true},
		{".", true},
		{"/.", false},
		{"a.", false},
		{"..", false},
		{"a.", false},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_dot got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_question(t *testing.T) {
	rrdfa := newDFA(re_question)

	testCases := []struct {
		input  string
		output bool
	}{
		{" ?", true},
		{"? ", true},
		{"?", true},
		{" ? 		", true},
		{"?", true},
		{"/?", false},
		{"a?", false},
		{"?.", false},
		{"a?", false},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_question got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_question_dot(t *testing.T) {
	rrdfa := newDFA(re_question_dot)

	testCases := []struct {
		input  string
		output bool
	}{
		{" ?.", true},
		{"?. ", true},
		{"?.", true},
		{" ?. 		", true},
		{"?.", true},
		{"/?.", false},
		{"a?.", false},
		{"/?.", false},
		{"a?.", false},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_question_dot got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_arrow(t *testing.T) {
	rrdfa := newDFA(re_arrow)

	testCases := []struct {
		input  string
		output bool
	}{
		{" =>", true},
		{"=> ", true},
		{"=>", true},
		{" => 		", true},
		{"=>", true},
		{"/=>", false},
		{"a=>", false},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_arrow got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_ellipsis(t *testing.T) {
	rrdfa := newDFA(re_ellipsis)

	testCases := []struct {
		input  string
		output bool
	}{
		{" ...", true},
		{"... ", true},
		{"...", true},
		{" ... 		", true},
		{"...", true},
		{"/...", false},
		{".../", false},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_ellipsis got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_equal(t *testing.T) {
	rrdfa := newDFA(re_equal)

	testCases := []struct {
		input  string
		output bool
	}{
		{" =", true},
		{"= ", true},
		{"=", true},
		{" = 		", true},
		{"=", true},
		{"/=", false},
		{"=/", false},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_equal got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_equality(t *testing.T) {
	rrdfa := newDFA(re_equality)

	testCases := []struct {
		input  string
		output bool
	}{
		{" ==", true},
		{"== ", true},
		{"==", true},
		{" == 		", true},
		{"/==", false},
		{"==/", false},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_equality got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_bitwiseor(t *testing.T) {
	rrdfa := newDFA(re_bitwiseor)

	testCases := []struct {
		input  string
		output bool
	}{
		{" |", true},
		{"| ", true},
		{"|", true},
		{" | 		", true},
		{"/|", false},
		{"|/", false},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_bitwiseor got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_bitwisexor(t *testing.T) {
	rrdfa := newDFA(re_bitwisexor)

	testCases := []struct {
		input  string
		output bool
	}{
		{" ^", true},
		{"^ ", true},
		{"^", true},
		{" ^ 		", true},
		{"/^", false},
		{"^/", false},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_bitwisexor got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_bitwiseand(t *testing.T) {
	rrdfa := newDFA(re_bitwiseand)

	testCases := []struct {
		input  string
		output bool
	}{
		{" &", true},
		{"& ", true},
		{"&", true},
		{" & 		", true},
		{"/&", false},
		{"&/", false},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_bitwiseand got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_logicor(t *testing.T) {
	rrdfa := newDFA(re_logicor)

	testCases := []struct {
		input  string
		output bool
	}{
		{" ||", true},
		{"|| ", true},
		{"||", true},
		{" || 		", true},
		{"/|/|", false},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_logicor got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_logicand(t *testing.T) {
	rrdfa := newDFA(re_logicand)

	testCases := []struct {
		input  string
		output bool
	}{
		{" &&", true},
		{"&& ", true},
		{"&&", true},
		{" && 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_logicand got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_plus(t *testing.T) {
	rrdfa := newDFA(re_plus)

	testCases := []struct {
		input  string
		output bool
	}{
		{" +", true},
		{"+ ", true},
		{"+", true},
		{" + 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_plus got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_min(t *testing.T) {
	rrdfa := newDFA(re_min)

	testCases := []struct {
		input  string
		output bool
	}{
		{" -", true},
		{"- ", true},
		{"-", true},
		{" - 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_min got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_modulo(t *testing.T) {
	rrdfa := newDFA(re_modulo)

	testCases := []struct {
		input  string
		output bool
	}{
		{" %", true},
		{"% ", true},
		{"%", true},
		{" % 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_modulo got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_bitleftshift(t *testing.T) {
	rrdfa := newDFA(re_bitleftshift)

	testCases := []struct {
		input  string
		output bool
	}{
		{" <<", true},
		{"<< ", true},
		{"<<", true},
		{" << 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_bitleftshift got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_bitrightshift(t *testing.T) {
	rrdfa := newDFA(re_bitrightshift)

	testCases := []struct {
		input  string
		output bool
	}{
		{" >>", true},
		{">> ", true},
		{">>", true},
		{" >> 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_bitrightshift got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_bitrightshift3(t *testing.T) {
	rrdfa := newDFA(re_bitrightshift3)

	testCases := []struct {
		input  string
		output bool
	}{
		{" >>>", true},
		{">>> ", true},
		{">>>", true},
		{" >>> 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_bitrightshift3 got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_break(t *testing.T) {
	rrdfa := newDFA(re_break)

	testCases := []struct {
		input  string
		output bool
	}{
		{" break", true},
		{"break ", true},
		{"break", true},
		{" break 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_break got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_case(t *testing.T) {
	rrdfa := newDFA(re_case)

	testCases := []struct {
		input  string
		output bool
	}{
		{" case", true},
		{"case ", true},
		{"case", true},
		{" case 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_case got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_catch(t *testing.T) {
	rrdfa := newDFA(re_catch)

	testCases := []struct {
		input  string
		output bool
	}{
		{" catch", true},
		{"catch ", true},
		{"catch", true},
		{" catch 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_catch got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_continue(t *testing.T) {
	rrdfa := newDFA(re_continue)

	testCases := []struct {
		input  string
		output bool
	}{
		{" continue", true},
		{"continue ", true},
		{"continue", true},
		{" continue 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_continue got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_default(t *testing.T) {
	rrdfa := newDFA(re_default)

	testCases := []struct {
		input  string
		output bool
	}{
		{" default", true},
		{"default ", true},
		{"default", true},
		{" default 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_default got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_finally(t *testing.T) {
	rrdfa := newDFA(re_finally)

	testCases := []struct {
		input  string
		output bool
	}{
		{" finally", true},
		{"finally ", true},
		{" finally 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_finally got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_function(t *testing.T) {
	rrdfa := newDFA(re_function)

	testCases := []struct {
		input  string
		output bool
	}{
		{" function", true},
		{"function ", true},
		{" function 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_function got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_return(t *testing.T) {
	rrdfa := newDFA(re_return)

	testCases := []struct {
		input  string
		output bool
	}{
		{" return", true},
		{"return ", true},
		{" return 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_return got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_switch(t *testing.T) {
	rrdfa := newDFA(re_switch)

	testCases := []struct {
		input  string
		output bool
	}{
		{" switch", true},
		{"switch ", true},
		{" switch 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_switch got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_throw(t *testing.T) {
	rrdfa := newDFA(re_throw)

	testCases := []struct {
		input  string
		output bool
	}{
		{" throw", true},
		{"throw ", true},
		{" throw 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_throw got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_try(t *testing.T) {
	rrdfa := newDFA(re_try)

	testCases := []struct {
		input  string
		output bool
	}{
		{" try", true},
		{"try ", true},
		{" try 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_try got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_with(t *testing.T) {
	rrdfa := newDFA(re_with)

	testCases := []struct {
		input  string
		output bool
	}{
		{" with", true},
		{"with ", true},
		{" with 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_with got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_new(t *testing.T) {
	rrdfa := newDFA(re_new)

	testCases := []struct {
		input  string
		output bool
	}{
		{" new", true},
		{"new ", true},
		{" new 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_new got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_this(t *testing.T) {
	rrdfa := newDFA(re_this)

	testCases := []struct {
		input  string
		output bool
	}{
		{" this", true},
		{"this ", true},
		{" this 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_this got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_super(t *testing.T) {
	rrdfa := newDFA(re_super)

	testCases := []struct {
		input  string
		output bool
	}{
		{" super", true},
		{"super ", true},
		{" super 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_super got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_class(t *testing.T) {
	rrdfa := newDFA(re_class)

	testCases := []struct {
		input  string
		output bool
	}{
		{" class", true},
		{"class ", true},
		{" class 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_class got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_extends(t *testing.T) {
	rrdfa := newDFA(re_extends)

	testCases := []struct {
		input  string
		output bool
	}{
		{" extends", true},
		{"extends ", true},
		{" extends 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_extends got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_export(t *testing.T) {
	rrdfa := newDFA(re_export)

	testCases := []struct {
		input  string
		output bool
	}{
		{" export", true},
		{"export ", true},
		{" export 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_export got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_import(t *testing.T) {
	rrdfa := newDFA(re_import)

	testCases := []struct {
		input  string
		output bool
	}{
		{" import", true},
		{"import ", true},
		{" import 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_import got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_in(t *testing.T) {
	rrdfa := newDFA(re_in)

	testCases := []struct {
		input  string
		output bool
	}{
		{" in", true},
		{"in ", true},
		{" in 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_in got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_instanceof(t *testing.T) {
	rrdfa := newDFA(re_instanceof)

	testCases := []struct {
		input  string
		output bool
	}{
		{" instanceof", true},
		{"instanceof ", true},
		{" instanceof 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_instanceof got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_typeof(t *testing.T) {
	rrdfa := newDFA(re_typeof)

	testCases := []struct {
		input  string
		output bool
	}{
		{" typeof", true},
		{"typeof ", true},
		{" typeof 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_typeof got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_void(t *testing.T) {
	rrdfa := newDFA(re_void)

	testCases := []struct {
		input  string
		output bool
	}{
		{" void", true},
		{"void ", true},
		{" void 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_void got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_delete(t *testing.T) {
	rrdfa := newDFA(re_delete)

	testCases := []struct {
		input  string
		output bool
	}{
		{" delete", true},
		{"delete ", true},
		{" delete 		", true},
	}

	for _, tc := range testCases {
		result := rrdfa.Match(tc.input)
		if result != tc.output {
			t.Errorf("Test_re_delete got error in %s, expected to be %v but got %v", tc.input, tc.output, result)
		}
	}
}

func Test_re_not(t *testing.T) {
	notDFA := newDFA(re_not)

	testCases := []struct {
		input  string
		output bool
	}{
		{"", false},
		{"!", true},
		{"!!", true},
		{" !  ", true},
		{"!! !", false},
		{"  !!!!", true},
		{"!!!!!!", true},
	}

	for _, tobj := range testCases {
		result := notDFA.Match(tobj.input)
		if result != tobj.output {
			t.Errorf("Test_re_not got error in %s, expected to be %v, but got %v", tobj.input, tobj.output, result)
		}
	}
}

func Test_re_eq3(t *testing.T) {
	notDFA := newDFA(re_eq3)

	testCases := []struct {
		input  string
		output bool
	}{
		{"", false},
		{"===", true},
		{"=", false},
		{" === ", true},
		{"== =", false},
	}

	for _, tobj := range testCases {
		result := notDFA.Match(tobj.input)
		if result != tobj.output {
			t.Errorf("Test_re_eq3 got error in %s, expected to be %v, but got %v", tobj.input, tobj.output, result)
		}
	}
}

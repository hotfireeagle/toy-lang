package languagespec

import "testing"

func TestPreProcessForSugar(t *testing.T) {
	testCases := []struct {
		caseValue   string
		expectValue []rune
	}{
		{"[1-5]", []rune("(1|2|3|4|5)")},
		{"[0-3]", []rune("(0|1|2|3)")},
		{"[a-c]", []rune("(a|b|c)")},
		{"[A-D]", []rune("(A|B|C|D)")},
		{"[1-3][3-5]", []rune("(1|2|3)(3|4|5)")},
		{"ashj[1-2][a-c]sjd", []rune("ashj(1|2)(a|b|c)sjd")},
		{"ashj[1-1][a-c]sjd", []rune("ashj(1)(a|b|c)sjd")},
		{"ashj[1-1][a-c]", []rune("ashj(1)(a|b|c)")},
		{"[5-1]", []rune("[5-1]")},
		{"$any$", []rune{anyInputSymbol}},
		{"$alphabet$", []rune{alphabetInputSymbol}},
		{"$not$(1)", []rune{lastNotInputSymbol}},
		{"[1-2]$any$", []rune{'(', '1', '|', '2', ')', anyInputSymbol}},
	}

	for _, testCase := range testCases {
		result := preProcessForSugar(testCase.caseValue)

		if len(result) != len(testCase.expectValue) {
			t.Error("preProcessForSugar error, wrong length")
		}

		for i := 0; i < len(result); i++ {
			if result[i] != testCase.expectValue[i] {
				t.Errorf("preProcessForSugar error, expect to be %s in %s, but got %s", string(testCase.expectValue), testCase.caseValue, string(result))
			}
		}
	}
}

func TestInfix2postfix(t *testing.T) {
	testcases := []struct {
		casev  []rune
		answer []rune
	}{
		{preProcessForSugar("1|2"), []rune{'1', '2', unionOperator}},
		{preProcessForSugar("1|2*"), []rune{'1', '2', starOperator, unionOperator}},
		{preProcessForSugar("1*|2"), []rune{'1', starOperator, '2', unionOperator}},
	}

	for _, obj := range testcases {
		result := infix2postfix(obj.casev)
		if len(result) != len(obj.answer) {
			t.Errorf("infix2postfix err %d, %d", len(result), len(obj.answer))
			return
		}
		for i := 0; i < len(result); i++ {
			if result[i] != obj.answer[i] {
				t.Errorf("infix2postfix err in %v, %v", result[i], obj.answer[i])
			}
		}
	}
}

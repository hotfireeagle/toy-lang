package languagespec

type InputSymbolType int

const (
	NormalInputSymbolType InputSymbolType = iota
	EpsilonInputSymbolType
	AnyInputSymbolType
	AlphabetInputSymbolType
	NotInputSymbolType
)

const starOperator rune = -1
const zeroOrOneOperator rune = -2
const oneOrMoreOperator rune = -3
const concatOperator rune = -4
const unionOperator rune = -5
const leftBracketOperator rune = -6
const epslilonInputSymbol rune = -7
const anyInputSymbol rune = -8
const alphabetInputSymbol rune = -9

var lastNotInputSymbol rune = -10

var notInputSymbolMap map[rune]bool
var notInputSymbol2IgnoreAlphabet map[rune]string

const starOperatorSymbolRE rune = '*'
const zeroOrOneOperatorSymbolRE rune = '?'
const oneOrMoreOperatorSymbolRE rune = '+'
const unionOperatorSymbolRE = '|'
const leftBracketOperatorSymbolRE = '('
const anySymbolRE = "$any$"
const alphabetSymbolRE = "$alphabet$"
const notSymbolRE = "$not$"

const anySymbolRELen = len(anySymbolRE)
const alphabetSymbolRELen = len(alphabetSymbolRE)
const notSymbolRELen = len(notSymbolRE)

var operatorPriority = map[rune]int{
	starOperatorSymbolRE:        4,
	zeroOrOneOperatorSymbolRE:   4,
	oneOrMoreOperatorSymbolRE:   4,
	concatOperator:              3,
	unionOperatorSymbolRE:       2,
	leftBracketOperatorSymbolRE: 1,
}

type InputSymbol struct {
	SymbolType        InputSymbolType
	SymbolLiteral     rune
	NotSymbolLiterals []rune
}

type DFA struct {
}

// func NewDFA(infixStr string) *DFA {
// 	infixStrAfterPreProcess := preProcessForSugar(infixStr)
// 	postfixRunes, notSymbolIdx2Runes := infix2postfix(infixStrAfterPreProcess)
// 	return nil
// }

// func infix2postfix(infix string) ([]rune, map[int][]rune) {
// 	operatorStack := NewRuneStack()
// 	postfixResult := make([]rune, 0)
// 	notSymbolIdx2Values := make(map[int][]rune)
// 	shouldAddConcat := false

// 	pushOperatorInRightPriority := func(operator rune) {

// 	}
// }

func preProcessForSugar(str string) []rune {
	// var sb strings.Builder
	notInputSymbolMap = make(map[rune]bool)
	notInputSymbol2IgnoreAlphabet = make(map[rune]string)
	needJumpIdxMap := make(map[int]bool)
	answer := make([]rune, 0)

	setNeedJumpIdx := func(begin int, end int) {
		for i := begin; i <= end; i++ {
			needJumpIdxMap[i] = true
		}
	}

	for idx, literal := range str {
		if needJumpIdxMap[idx] {
			continue
		}

		if literal == starOperatorSymbolRE {
			answer = append(answer, starOperator)
		} else if literal == zeroOrOneOperatorSymbolRE {
			answer = append(answer, zeroOrOneOperator)
		} else if literal == oneOrMoreOperatorSymbolRE {
			answer = append(answer, oneOrMoreOperator)
		} else if literal == unionOperatorSymbolRE {
			answer = append(answer, unionOperator)
		} else if literal == leftBracketOperatorSymbolRE {
			answer = append(answer, leftBracketOperator)
		} else if literal == '[' {
			if str[idx+2] == '-' && str[idx+4] == ']' {
				beginValIdx := idx + 1
				endValIdx := idx + 3
				rightBracketIdx := idx + 4

				beginLiteral := str[beginValIdx]
				endLiteral := str[endValIdx]

				isValid := false
				convertResult := make([]rune, 0)

				if (beginLiteral <= '9' && endLiteral <= '9' && beginLiteral <= endLiteral) || (isAlphabet(beginLiteral) && isAlphabet(endLiteral) && beginLiteral <= endLiteral) {
					isValid = true

					k := beginLiteral

					for k <= endLiteral {
						convertResult = append(convertResult, rune(k))
						if k != endLiteral {
							convertResult = append(convertResult, '|')
						}
						k++
					}
				}

				if isValid {
					setNeedJumpIdx(beginValIdx, rightBracketIdx)

					answer = append(answer, '(')
					answer = append(answer, convertResult...)
					answer = append(answer, ')')
					// sb.WriteRune('(')
					// sb.WriteString(convertResult)
					// sb.WriteRune(')')
				} else {
					answer = append(answer, literal)
					// sb.WriteRune(literal)
				}
			} else {
				answer = append(answer, literal)
				// sb.WriteRune(literal)
			}
		} else if literal == '$' {
			strLen := len(str)
			if idx+anySymbolRELen <= strLen && str[idx:idx+anySymbolRELen] == anySymbolRE {
				setNeedJumpIdx(idx, idx+anySymbolRELen-1)
				answer = append(answer, anyInputSymbol)
			} else if idx+alphabetSymbolRELen <= strLen && str[idx:idx+alphabetSymbolRELen] == alphabetSymbolRE {
				setNeedJumpIdx(idx, idx+alphabetSymbolRELen-1)
				answer = append(answer, alphabetInputSymbol)
			} else if idx+notSymbolRELen <= strLen && str[idx:idx+notSymbolRELen] == notSymbolRE {
				// we trust the builder self, so ignore the check process
				leftBracketIdx := idx + notSymbolRELen
				rightBracketIdx := leftBracketIdx

				for str[rightBracketIdx] != ')' {
					rightBracketIdx++
				}

				ignoreAlphabetStr := str[leftBracketIdx+1 : rightBracketIdx]

				setNeedJumpIdx(idx, rightBracketIdx)
				answer = append(answer, lastNotInputSymbol)
				notInputSymbolMap[lastNotInputSymbol] = true
				notInputSymbol2IgnoreAlphabet[lastNotInputSymbol] = ignoreAlphabetStr
				lastNotInputSymbol++
			} else {
				answer = append(answer, literal)
			}
		} else {
			answer = append(answer, literal)
			// sb.WriteRune(literal)
		}
	}

	return answer
	// return sb.String()
}

func isAlphabet(a byte) bool {
	return (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z')
}

type RuneStack struct {
	vals []rune
}

func NewRuneStack() *RuneStack {
	return &RuneStack{
		vals: make([]rune, 0),
	}
}

func (r *RuneStack) in(v rune) {
	r.vals = append(r.vals, v)
}

func (r *RuneStack) out() rune {
	result := r.vals[len(r.vals)-1]

	r.vals = r.vals[:len(r.vals)-1]

	return result
}

package languagespec

import (
	"sort"
	"strconv"
	"strings"
)

const starOperator rune = -1
const zeroOrOneOperator rune = -2
const oneOrMoreOperator rune = -3
const concatOperator rune = -4
const unionOperator rune = -5
const leftBracketOperator rune = -6
const rightBracketOperator rune = -7
const epslilonInputSymbol rune = -8
const anyInputSymbol rune = -9
const alphabetInputSymbol rune = -10

var lastNotInputSymbol rune = -11

var notInputSymbolMap map[rune]bool
var notInputSymbol2IgnoreAlphabet map[rune]string

const starOperatorSymbolRE rune = '*'
const zeroOrOneOperatorSymbolRE rune = '?'
const oneOrMoreOperatorSymbolRE rune = '+'
const unionOperatorSymbolRE = '|'
const leftBracketOperatorSymbolRE = '('
const rightBracketOperatorSymbolRE = ')'
const anySymbolRE = "$any$"
const alphabetSymbolRE = "$alphabet$"
const notSymbolRE = "$not$"

const anySymbolRELen = len(anySymbolRE)
const alphabetSymbolRELen = len(alphabetSymbolRE)
const notSymbolRELen = len(notSymbolRE)

var inputSymbolCacheMap map[rune]*inputSymbol

var operatorPriority = map[rune]int{
	starOperator:        4,
	zeroOrOneOperator:   4,
	oneOrMoreOperator:   4,
	concatOperator:      3,
	unionOperator:       2,
	leftBracketOperator: 1,
}

type inputSymbol struct {
	symbolLiteral          rune
	notSymbolLiteralString string
}

func newInputSymbol(symbol rune, ns string) *inputSymbol {
	if inputSymbolCacheMap[symbol] != nil {
		return inputSymbolCacheMap[symbol]
	}
	ips := &inputSymbol{
		symbolLiteral:          symbol,
		notSymbolLiteralString: ns,
	}
	inputSymbolCacheMap[symbol] = ips
	return ips
}

func newEpsilonInputSymbol() *inputSymbol {
	return &inputSymbol{
		symbolLiteral:          epslilonInputSymbol,
		notSymbolLiteralString: "",
	}
}

func newNotInputSymbol(symbol rune, str string) *inputSymbol {
	return &inputSymbol{
		symbolLiteral:          symbol,
		notSymbolLiteralString: str,
	}
}

func checkIsNotSymbol(symbol rune) bool {
	return notInputSymbolMap[symbol]
}

func checkIsEpsilonInputSymbol(ips *inputSymbol) bool {
	return ips.symbolLiteral == epslilonInputSymbol
}

type nfa struct {
	states              []int
	inputSymbols        []*inputSymbol
	transitionMap       map[int]map[*inputSymbol][]int
	startState          int
	acceptStates        []int
	beginEndStatePairs  map[int]int
	iDCount             int
	inputSymbolAddedMap map[*inputSymbol]bool
}

func newNFA(infixStr string) *nfa {
	nfaObj := &nfa{
		states:              make([]int, 0),
		inputSymbols:        make([]*inputSymbol, 0),
		transitionMap:       make(map[int]map[*inputSymbol][]int),
		acceptStates:        make([]int, 0),
		beginEndStatePairs:  make(map[int]int),
		inputSymbolAddedMap: make(map[*inputSymbol]bool),
	}
	infixStrAfterPreProcess := preProcessForSugar(infixStr)
	postfixRunes := infix2postfix(infixStrAfterPreProcess)
	nfaObj.postfix2NFA(postfixRunes)
	return nfaObj
}

func (n *nfa) addState() int {
	stateId := n.iDCount
	n.states = append(n.states, stateId)
	n.iDCount++
	return stateId
}

func (n *nfa) addInputSymbol(ips *inputSymbol) {
	if !n.inputSymbolAddedMap[ips] {
		n.inputSymbols = append(n.inputSymbols, ips)
		n.inputSymbolAddedMap[ips] = true
	}
}

func (n *nfa) addTransition(ips *inputSymbol, fromStateId int, toStateId int) {
	n.addInputSymbol(ips)

	if _, ok := n.transitionMap[fromStateId]; !ok {
		n.transitionMap[fromStateId] = make(map[*inputSymbol][]int)
	}

	if n.transitionMap[fromStateId][ips] == nil {
		n.transitionMap[fromStateId][ips] = make([]int, 0)
	}

	n.transitionMap[fromStateId][ips] = append(n.transitionMap[fromStateId][ips], toStateId)
}

func (n *nfa) setBeginEndPairs(beginState int, endState int) {
	n.beginEndStatePairs[beginState] = endState
}

func (n *nfa) getEndState(beginState int) int {
	return n.beginEndStatePairs[beginState]
}

func (n *nfa) setStartState(stateId int) {
	n.startState = stateId
}

func (n *nfa) addAcceptState(stateId int) {
	n.acceptStates = append(n.acceptStates, stateId)
}

func (n *nfa) postfix2NFA(runes []rune) {
	stateStack := newStateState()

	for _, symbol := range runes {
		if symbol == concatOperator {
			rightState := stateStack.out()
			leftState := stateStack.out()

			n.addTransition(newEpsilonInputSymbol(), n.getEndState(leftState), rightState)
			n.setBeginEndPairs(leftState, n.getEndState(rightState))

			stateStack.in(leftState)
		} else if symbol == unionOperator {
			rightState := stateStack.out()
			leftState := stateStack.out()

			newBegin := n.addState()
			newEnd := n.addState()

			n.setBeginEndPairs(newBegin, newEnd)

			n.addTransition(newEpsilonInputSymbol(), newBegin, leftState)
			n.addTransition(newEpsilonInputSymbol(), newBegin, rightState)

			rightStateEnd := n.getEndState(rightState)
			leftStateEnd := n.getEndState(leftState)

			n.addTransition(newEpsilonInputSymbol(), rightStateEnd, newEnd)
			n.addTransition(newEpsilonInputSymbol(), leftStateEnd, newEnd)

			stateStack.in(newBegin)
		} else if symbol == starOperator {
			state := stateStack.out()

			newBegin := n.addState()
			newEnd := n.addState()
			n.setBeginEndPairs(newBegin, newEnd)

			stateEnd := n.getEndState(state)

			n.addTransition(newEpsilonInputSymbol(), newBegin, state)
			n.addTransition(newEpsilonInputSymbol(), stateEnd, state)
			n.addTransition(newEpsilonInputSymbol(), stateEnd, newEnd)
			n.addTransition(newEpsilonInputSymbol(), newBegin, newEnd)

			stateStack.in(newBegin)
		} else if symbol == zeroOrOneOperator {
			state := stateStack.out()

			newBegin := n.addState()
			newEnd := n.addState()
			n.setBeginEndPairs(newBegin, newEnd)

			stateEnd := n.getEndState(state)

			n.addTransition(newEpsilonInputSymbol(), newBegin, state)
			n.addTransition(newEpsilonInputSymbol(), stateEnd, newEnd)
			n.addTransition(newEpsilonInputSymbol(), newBegin, newEnd)

			stateStack.in(newBegin)
		} else if symbol == oneOrMoreOperator {
			state := stateStack.out()

			newBegin := n.addState()
			newEnd := n.addState()
			n.setBeginEndPairs(newBegin, newEnd)

			stateEnd := n.getEndState(state)

			n.addTransition(newEpsilonInputSymbol(), newBegin, state)
			n.addTransition(newEpsilonInputSymbol(), stateEnd, state)
			n.addTransition(newEpsilonInputSymbol(), stateEnd, newEnd)

			stateStack.in(newBegin)
		} else {
			beginStateId := n.addState()
			endStateId := n.addState()

			n.setBeginEndPairs(beginStateId, endStateId)
			stateStack.in(beginStateId)

			var inp *inputSymbol
			if checkIsNotSymbol(symbol) {
				inp = newNotInputSymbol(symbol, notInputSymbol2IgnoreAlphabet[symbol])
			} else {
				inp = newInputSymbol(symbol, "")
			}

			n.addTransition(inp, beginStateId, endStateId)
		}
	}

	startState := stateStack.out()
	n.setStartState(startState)
	n.addAcceptState(n.getEndState(startState))
}

type dfa struct {
	states          []int
	inputSymbols    []*inputSymbol
	transitionMap   map[int]map[*inputSymbol]int
	startState      int
	acceptStates    []int
	stateIdToSetMap map[string]int
	deadStateId     int
	dfaStateCount   int
}

func (d *dfa) addTransition(ips *inputSymbol, fromStateId int, toStateId int) {
	if _, ok := d.transitionMap[fromStateId]; !ok {
		d.transitionMap[fromStateId] = make(map[*inputSymbol]int)
	}
	d.transitionMap[fromStateId][ips] = toStateId
}

func (d *dfa) addState(nfaStates []int) (int, bool) {
	if len(nfaStates) == 0 {
		return d.addDeadState()
	}

	idStr := slice2str(nfaStates)
	if _, ok := d.stateIdToSetMap[idStr]; ok {
		return d.stateIdToSetMap[idStr], true
	}

	stateId := d.dfaStateCount
	d.states = append(d.states, stateId)
	d.addStateIdToSetMap(stateId, idStr)

	d.dfaStateCount++
	return stateId, false
}

func (d *dfa) addDeadState() (int, bool) {
	sId := d.dfaStateCount

	if d.deadStateId == -1 {
		d.deadStateId = sId
		d.states = append(d.states, sId)
		d.addTransitionForDeadState()
		d.dfaStateCount++
		return sId, false
	} else {
		return d.deadStateId, true
	}
}

func (d *dfa) addTransitionForDeadState() {
	for _, inputSymbol := range d.inputSymbols {
		d.addTransition(inputSymbol, d.deadStateId, d.deadStateId)
	}
}

func (d *dfa) addStateIdToSetMap(stateId int, set string) {
	d.stateIdToSetMap[set] = stateId
}

func (d *dfa) getStateIdByStr(stateStr string) int {
	if stateStr == "" {
		return d.deadStateId
	}
	return d.stateIdToSetMap[stateStr]
}

func (d *dfa) setStartState(stateId int) {
	d.startState = stateId
}

func (d *dfa) setInputSymbols(inputSymbols []*inputSymbol) {
	ism := make([]*inputSymbol, 0)
	for _, ips := range inputSymbols {
		if checkIsEpsilonInputSymbol(ips) {
			continue
		}
		ism = append(ism, ips)
	}
	d.inputSymbols = ism
}

func newDFA(regexp string) *dfa {
	inputSymbolCacheMap = make(map[rune]*inputSymbol)
	nfaObj := newNFA(regexp)
	dfaObj := &dfa{
		states:          make([]int, 0),
		acceptStates:    make([]int, 0),
		inputSymbols:    make([]*inputSymbol, 0),
		transitionMap:   make(map[int]map[*inputSymbol]int),
		stateIdToSetMap: make(map[string]int),
		deadStateId:     -1,
	}

	dfaObj.setInputSymbols(nfaObj.inputSymbols)

	var findCurrentStateCanGoAnyStateByEpsilon = func(state int) []int {
		canGoStates := make([]int, 0)
		visited := make(map[int]bool)

		var dfs func(s int)

		dfs = func(currentState int) {
			if visited[currentState] {
				return
			}
			visited[currentState] = true

			transitions := nfaObj.transitionMap[currentState]
			for ipsymbol, toStates := range transitions {
				if checkIsEpsilonInputSymbol(ipsymbol) {
					for _, stateId := range toStates {
						canGoStates = append(canGoStates, stateId)
						dfs(stateId)
					}
				}
			}
		}

		dfs(state)

		return canGoStates
	}

	// 开始节点state
	startStates := findCurrentStateCanGoAnyStateByEpsilon(nfaObj.startState)
	startStates = append(startStates, nfaObj.startState)

	startStateId, _ := dfaObj.addState(startStates)
	dfaObj.setStartState(startStateId)

	needBeSettle := make([][]int, 0)
	needBeSettle = append(needBeSettle, startStates)

	for len(needBeSettle) > 0 {
		nextNeedBeSettle := make([][]int, 0)

		for _, states := range needBeSettle {

			for _, ips := range dfaObj.inputSymbols {

				nextCanGoStates := make(map[int]bool)

				for _, fromStateId := range states {

					canGoStateList := nfaObj.transitionMap[fromStateId][ips]
					for _, canGoStateId := range canGoStateList {
						nextCanGoStates[canGoStateId] = true
						// we also need find the epsilon move
						thisStateCanGoByEpsilonMove := findCurrentStateCanGoAnyStateByEpsilon(canGoStateId)
						for _, id := range thisStateCanGoByEpsilonMove {
							nextCanGoStates[id] = true
						}
					}

				}

				nextCanGoStateIds := getKeys(nextCanGoStates)
				dfaStateId, hasExist := dfaObj.addState(nextCanGoStateIds)

				fromStateStr := slice2str(states)
				dfaFromStateId := dfaObj.getStateIdByStr(fromStateStr)
				dfaObj.addTransition(ips, dfaFromStateId, dfaStateId)

				if !hasExist {
					nextNeedBeSettle = append(nextNeedBeSettle, nextCanGoStateIds)
				}
			}
		}

		needBeSettle = nextNeedBeSettle
	}

	endStates := make([]int, 0)
	for stateStr, stateId := range dfaObj.stateIdToSetMap {
		nfaFinalState := nfaObj.acceptStates[0]
		if strings.Contains(stateStr, strconv.Itoa(nfaFinalState)) {
			endStates = append(endStates, stateId)
		}
	}

	dfaObj.acceptStates = endStates

	return dfaObj
}

func (d *dfa) Match(str string) bool {
	currentState := d.startState

	checkInputSymbolIsMatch := func(ips *inputSymbol, character rune) bool {
		if ips.symbolLiteral == anyInputSymbol {
			return true
		} else if ips.symbolLiteral == alphabetInputSymbol {
			return isAlphabet(byte(character))
		} else if checkIsNotSymbol(ips.symbolLiteral) {
			notStr := ips.notSymbolLiteralString

			hit := false

			for _, r := range notStr {
				if r == character {
					hit = true
					break
				}
			}

			return !hit
		} else {
			return ips.symbolLiteral == character
		}
	}

	for _, character := range str {
		currentStateTransitions := d.transitionMap[currentState]
		currentState = -1
		for isp, nextStateId := range currentStateTransitions {
			if checkInputSymbolIsMatch(isp, character) {
				currentState = nextStateId
				break
			}
		}
	}

	answer := false

	for _, acceptState := range d.acceptStates {
		if acceptState == currentState {
			answer = true
			break
		}
	}

	return answer
}

// TODO: minimize and rebuild dfa
// func (d *dfa) minimize() *dfa {
// 	dfaAllStates := d.states
// 	dfaAcceptStates := d.acceptStates
// 	dfaAcceptStateMapObj := make(map[int]bool)

// 	for _, s := range dfaAcceptStates {
// 		dfaAcceptStateMapObj[s] = true
// 	}

// 	acceptStates := make([]int, len(dfaAcceptStates))
// 	copy(acceptStates, dfaAcceptStates)
// 	unacceptStates := make([]int, 0)

// 	for _, stateId := range dfaAllStates {
// 		if !dfaAcceptStateMapObj[stateId] {
// 			unacceptStates = append(unacceptStates, stateId)
// 		}
// 	}

// 	setPool := [][]int{
// 		acceptStates,
// 		unacceptStates,
// 	}

// 	// 某个set沿着某个inputSymbol所到达的下一个set
// 	set2setByInputSymbol := func(sets []int, ips *inputSymbol) []int {
// 		newSet := make([]int, 0)
// 		addedSetId := make(map[int]bool)

// 		for _, fromSetId := range sets {
// 			nextStateId := d.transitionMap[fromSetId][ips]
// 			if !addedSetId[nextStateId] {
// 				addedSetId[nextStateId] = true
// 				newSet = append(newSet, nextStateId)
// 			}
// 		}

// 		return newSet
// 	}

// 	// 判断当前set是否是setPool的的子集
// 	checkSetIsSubSetOfSetPool := func(someSet []int) bool {
// 		setPoolMaps := make([]map[int]bool, 0)

// 		for _, setlist := range setPool {
// 			setmap := make(map[int]bool)
// 			for _, setid := range setlist {
// 				setmap[setid] = true
// 			}
// 			setPoolMaps = append(setPoolMaps, setmap)
// 		}

// 		for _, setPoolMap := range setPoolMaps {
// 			isCurrentSetPoolMapSubset := true
// 			for _, currentSetId := range someSet {
// 				if setPoolMap[currentSetId] == false {
// 					isCurrentSetPoolMapSubset = false
// 					break
// 				}
// 			}
// 			if isCurrentSetPoolMapSubset {
// 				return true
// 			}
// 		}

// 		return false
// 	}

// 	checkSomeSetNeedDivide := func(someSet []int) bool {
// 		inputSymbolList := d.inputSymbols

// 		// 但凡某个set针对某个输入symbol不是现存set池所有set的子集，那么就认为它需要分裂

// 		shouldDivide := false

// 		for _, ips := range inputSymbolList {
// 			// 当前set针对这个输入symbol所能到达的下一个状态set
// 			nextStateSet := set2setByInputSymbol(someSet, ips)

// 			if !checkSetIsSubSetOfSetPool(nextStateSet) {
// 				// 说明这个nextStateSet集合不是当前的集合池子中任意一个集合的子集
// 				// 既然如此，那么表明，需要分裂，且无需在判断其它输入了
// 				shouldDivide = true
// 				break
// 			}
// 		}

// 		return shouldDivide
// 	}

// 	return nil
// }

func infix2postfix(infix []rune) []rune {
	operatorStack := newRuneStack()
	postfixResult := make([]rune, 0)
	shouldAddConcat := false

	pushOperatorInRightPriority := func(operator rune) {
		currentOperatorPriority := operatorPriority[operator]

		for operatorStack.notEmpty() {
			beforeOperator := operatorStack.out()
			beforeOperatorPriority := operatorPriority[beforeOperator]

			if beforeOperatorPriority >= currentOperatorPriority {
				postfixResult = append(postfixResult, beforeOperator)
			} else {
				operatorStack.in(beforeOperator)
				break
			}
		}

		operatorStack.in(operator)
	}

	for _, word := range infix {
		if word == starOperator || word == zeroOrOneOperator || word == oneOrMoreOperator {
			shouldAddConcat = true
			pushOperatorInRightPriority(word)
		} else if word == unionOperator {
			shouldAddConcat = false
			pushOperatorInRightPriority(word)
		} else if word == leftBracketOperator {
			if shouldAddConcat {
				pushOperatorInRightPriority(concatOperator)
			}
			operatorStack.in(word)
			shouldAddConcat = false
		} else if word == rightBracketOperator {
			var operator rune
			for operatorStack.notEmpty() {
				operator = operatorStack.out()
				if operator == leftBracketOperator {
					break
				}
				postfixResult = append(postfixResult, operator)
			}
			if operator != leftBracketOperator {
				panic("unmatched )")
			}
			shouldAddConcat = true
		} else {
			if shouldAddConcat {
				pushOperatorInRightPriority(concatOperator)
			}

			postfixResult = append(postfixResult, word)
			shouldAddConcat = true
		}
	}

	for operatorStack.notEmpty() {
		oper := operatorStack.out()
		if oper == leftBracketOperator {
			panic("unexpected single (")
		}
		postfixResult = append(postfixResult, oper)
	}

	return postfixResult
}

func preProcessForSugar(str string) []rune {
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
		} else if literal == rightBracketOperatorSymbolRE {
			answer = append(answer, rightBracketOperator)
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
							convertResult = append(convertResult, unionOperator)
						}
						k++
					}
				}

				if isValid {
					setNeedJumpIdx(beginValIdx, rightBracketIdx)

					answer = append(answer, leftBracketOperator)
					answer = append(answer, convertResult...)
					answer = append(answer, rightBracketOperator)
				} else {
					answer = append(answer, literal)
				}
			} else {
				answer = append(answer, literal)
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

				for str[rightBracketIdx] != rightBracketOperatorSymbolRE {
					rightBracketIdx++
				}

				ignoreAlphabetStr := str[leftBracketIdx+1 : rightBracketIdx]

				setNeedJumpIdx(idx, rightBracketIdx)
				answer = append(answer, lastNotInputSymbol)
				notInputSymbolMap[lastNotInputSymbol] = true
				notInputSymbol2IgnoreAlphabet[lastNotInputSymbol] = ignoreAlphabetStr
				lastNotInputSymbol--
			} else {
				answer = append(answer, literal)
			}
		} else {
			answer = append(answer, literal)
		}
	}

	return answer
}

func isAlphabet(a byte) bool {
	return (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z')
}

type runeStack struct {
	vals []rune
}

func newRuneStack() *runeStack {
	return &runeStack{
		vals: make([]rune, 0),
	}
}

func (r *runeStack) in(v rune) {
	r.vals = append(r.vals, v)
}

func (r *runeStack) out() rune {
	result := r.vals[len(r.vals)-1]

	r.vals = r.vals[:len(r.vals)-1]

	return result
}

func (r *runeStack) notEmpty() bool {
	return len(r.vals) > 0
}

type stateStack struct {
	vals []int
}

func newStateState() *stateStack {
	return &stateStack{
		vals: make([]int, 0),
	}
}

func (s *stateStack) in(v int) {
	s.vals = append(s.vals, v)
}

func (s *stateStack) out() int {
	st := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return st
}

func slice2str(arr []int) string {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	var result string

	for idx, v := range arr {
		if idx != 0 {
			result += ","
		}
		s := strconv.Itoa(v)
		result += s
	}

	return result
}

func getKeys(m map[int]bool) []int {
	keys := make([]int, 0)
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

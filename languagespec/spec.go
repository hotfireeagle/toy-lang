package languagespec

import (
	"strings"
)

// TODO: 支持模版字符串

// 10进制数字
const re_num_10 = "($whitespace$)*(0+)|([1-9][0-9]*)($whitespace$)*"

// 二进制数字
const re_num_binary2 = "($whitespace$)*0b[0-1]+($whitespace$)*"

// 16进制数字
const re_num_16 = "($whitespace$)*0x(0|1|2|3|4|5|6|7|8|9|a|b|c|d|e|f)+($whitespace$)*"

// 变量名
const re_identfier = "($whitespace$)*($alphabet$|_|$)($alphabet$|_|$|[0-9])*($whitespace$)*"

// 双引号字符串
const re_double_string = `($whitespace$)*"($not$("))*"($whitespace$)*`

// 单引号字符串
const re_single_string = `($whitespace$)*'($not$('))*'($whitespace$)*`

// var
const re_var = "($whitespace$)*var($whitespace$)*"

// let
const re_let = "($whitespace$)*let($whitespace$)*"

// const
const re_const = "($whitespace$)*const($whitespace$)*"

// true
const re_true = "($whitespace$)*true($whitespace$)*"

// false
const re_false = "($whitespace$)*false($whitespace$)*"

// undefined
const re_undefined = "($whitespace$)*undefined($whitespace$)*"

// null
const re_null = "($whitespace$)*null($whitespace$)*"

// if
const re_if = "($whitespace$)*if($whitespace$)*"

// else
const re_else = "($whitespace$)*else($whitespace$)*"

// else if
const re_elseif = "($whitespace$)*else($whitespace$)+if($whitespace$)*"

// for
const re_for = "($whitespace$)*for($whitespace$)*"

// while
const re_while = "($whitespace$)*while($whitespace$)*"

// do
const re_do = "($whitespace$)*do($whitespace$)*"

// 单行注释
const re_single_row_comment = "($whitespace$)*//($any$)*($whitespace$)*"

var languageSpecs = []string{
	re_num_10,
	re_num_binary2,
	re_num_16,
	re_identfier,
	re_double_string,
	re_single_string,
	re_var,
	re_let,
	re_const,
	re_true,
	re_false,
	re_undefined,
	re_null,
	re_if,
	re_else,
	re_elseif,
	re_for,
	re_while,
	re_do,
	re_single_row_comment,
}

func combineSpecsRegularLanguage(specs []string) string {
	var result strings.Builder

	for idx, v := range specs {
		if idx != 0 {
			result.WriteRune('|')
		}
		result.WriteString(v)
	}

	return result.String()
}

var spec = combineSpecsRegularLanguage(languageSpecs)
var LanguageNFA = newDFA(spec)

func CheckIsNormalNum(str string) bool {
	dfaobj := newDFA(re_num_10)
	return dfaobj.Match(str)
	// normalNumNFA := Re2nfaConstructor(re_num_10)
	// return normalNumNFA.Match(str)
}

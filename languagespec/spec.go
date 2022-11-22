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

// 模板字符串
// const re_template_string = "( )*`()*`( )*"
const re_var = "($whitespace$)*var($whitespace$)*"
const re_let = "( )*let( )*"
const re_const = "( )*const( )*"

const re_true = "( )*true( )*"
const re_false = "( )*false( )*"

const re_undefined = "( )*undefined( )*"
const re_null = "( )*null( )*"

const re_if = "( )*if( )*"
const re_else = "( )*else( )*"
const re_elseif = "( )*else ( )*if( )*"
const re_for = "( )*for( )*"
const re_while = "( )*while( )*"
const re_do = "( )*do( )*"

// 单行注释
// const re_single_row_comment = "( )*//(( )*|([0-9])*|()*)\n( )*"

var languageSpecs = []string{
	re_num_10,
	// re_num_binary2,
	// re_num_16,
	// re_identfier,
	// re_var,
	// re_let,
	// re_const,
	// re_true,
	// re_false,
	// re_undefined,
	// re_null,
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

// var spec = combineSpecsRegularLanguage(languageSpecs)
// var LanguageNFA = newDFA(re_num_10)

func CheckIsNormalNum(str string) bool {
	dfaobj := newDFA(re_num_10)
	return dfaobj.Match(str)
	// normalNumNFA := Re2nfaConstructor(re_num_10)
	// return normalNumNFA.Match(str)
}

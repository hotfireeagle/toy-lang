package languagespec

import (
	"strings"
)

// 平常的数字
const re_normal_num = "( )*[1-9][0-9]*( )*"

// 数字0
const re_num_0 = "( )*0*( )*"

// 二进制数字
const re_num_binary2 = "( )*0b[0-1]*( )*"

// 16进制数字
const re_num_16 = "( )*0x(0|1|2|3|4|5|6|7|8|9|a|b|c|d|e|f)*( )*"

// 变量名
const re_identfier = "( )*([a-z]|[A-Z]|$|_)([a-z]|[A-Z]|$|_|.)*( )*"

const re_var = "( )*var( )*"
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
	re_normal_num,
	re_num_0,
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
// var LanguageNFA = Re2nfaConstructor(spec)

func CheckIsNormalNum(str string) bool {
	normalNumNFA := Re2nfaConstructor(re_normal_num)
	return normalNumNFA.Match(str)
}

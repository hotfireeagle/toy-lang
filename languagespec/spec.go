package languagespec

import "strings"

// TODO: 需要支持中文变量名，中文注释

// 平常的数字
const re_normal_num = "( )*[1-9]+[0-9]*( )*"

// 数字0
const re_num_0 = "( )*0*( )*"

// 二进制数字
const re_num_binary2 = "( )*0b[0-1]*( )*"

// 16进制数字
const re_num_16 = "( )*0x[0|1|2|3|4|5|6|7|8|9|a|b|c|d|e|f]*( )*"

// 单行注释
// const re_single_row_comment = "( )*//(( )*|([0-9])*|()*)\n( )*"

var languageSpecs = []string{
	re_normal_num,
	re_num_0,
	re_num_binary2,
	re_num_16,
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
var LanguageNFA = Re2nfaConstructor(spec)

func CheckIsNormalNum(str string) bool {
	normalNumNFA := Re2nfaConstructor(re_normal_num)
	return normalNumNFA.Match(str)
}

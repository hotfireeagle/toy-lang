package languagespec

import "strings"

const re_normal_num = "(\t| )*[1-9]+[0-9]*(\t| )*" // 平常的数字
const re_num_0 = "(\t| )*0*(\t| )*"                // 数字0
const re_num_binary2 = "(\t| )*0b[0-1]*(\t| )*"    // 二进制数字

var languageSpecs = []string{
	re_normal_num,
	re_num_0,
	re_num_binary2,
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

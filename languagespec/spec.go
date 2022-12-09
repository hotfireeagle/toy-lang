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

// 浮点数
const re_num_float = "($whitespace$)*[0-9]*.[0-9]*($whitespace$)*"

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
const re_single_row_comment = `($whitespace$)*//($not$(/enter))*($whitespace$)*`

// [
const re_left_bracket = "($whitespace$)*/[($whitespace$)*"

// ]
const re_right_bracket = "($whitespace$)*/]($whitespace$)*"

// {
const re_left_bracel = "($whitespace$)*{($whitespace$)*"

// }
const re_right_bracel = "($whitespace$)*}($whitespace$)*"

// (
const re_left_parenl = "($whitespace$)*/(($whitespace$)*"

// )
const re_right_parenl = "($whitespace$)*/)($whitespace$)*"

// ,
const re_comma = "($whitespace$)*,($whitespace$)*"

// ;
const re_semi = "($whitespace$)*;($whitespace$)*"

// :
const re_colon = "($whitespace$)*:($whitespace$)*"

// .
const re_dot = "($whitespace$)*.($whitespace$)*"

// ?
const re_question = "($whitespace$)*/?($whitespace$)*"

// ?.
const re_question_dot = "($whitespace$)*/?.($whitespace$)*"

// =>
const re_arrow = "($whitespace$)*=>($whitespace$)*"

// ...
const re_ellipsis = "($whitespace$)*...($whitespace$)*"

// =
const re_equal = "($whitespace$)*=($whitespace$)*"

// >
const re_greater = "($whitespace$)*>($whitespace$)*"

// <
const re_less = "($whitespace$)*<($whitespace$)*"

// ==
const re_equality = "($whitespace$)*==($whitespace$)*"

// |
const re_bitwiseor = "($whitespace$)*/|($whitespace$)*"

// ^
const re_bitwisexor = "($whitespace$)*^($whitespace$)*"

// &
const re_bitwiseand = "($whitespace$)*&($whitespace$)*"

// ||
const re_logicor = "($whitespace$)*/|/|($whitespace$)*"

// &&
const re_logicand = "($whitespace$)*&&($whitespace$)*"

// +
const re_plus = "($whitespace$)*/+($whitespace$)*"

// -
const re_min = "($whitespace$)*-($whitespace$)*"

// %
const re_modulo = "($whitespace$)*%($whitespace$)*"

// <<
const re_bitleftshift = "($whitespace$)*<<($whitespace$)*"

// >>
const re_bitrightshift = "($whitespace$)*>>($whitespace$)*"

// >>>
const re_bitrightshift3 = "($whitespace$)*>>>($whitespace$)*"

// break
const re_break = "($whitespace$)*break($whitespace$)*"

// case
const re_case = "($whitespace$)*case($whitespace$)*"

// catch
const re_catch = "($whitespace$)*catch($whitespace$)*"

// continue
const re_continue = "($whitespace$)*continue($whitespace$)*"

// default
const re_default = "($whitespace$)*default($whitespace$)*"

// finally
const re_finally = "($whitespace$)*finally($whitespace$)*"

// function
const re_function = "($whitespace$)*function($whitespace$)*"

// return
const re_return = "($whitespace$)*return($whitespace$)*"

// switch
const re_switch = "($whitespace$)*switch($whitespace$)*"

// throw
const re_throw = "($whitespace$)*throw($whitespace$)*"

// try
const re_try = "($whitespace$)*try($whitespace$)*"

// with
const re_with = "($whitespace$)*with($whitespace$)*"

// new
const re_new = "($whitespace$)*new($whitespace$)*"

// this
const re_this = "($whitespace$)*this($whitespace$)*"

// super
const re_super = "($whitespace$)*super($whitespace$)*"

// class
const re_class = "($whitespace$)*class($whitespace$)*"

// extends
const re_extends = "($whitespace$)*extends($whitespace$)*"

// export
const re_export = "($whitespace$)*export($whitespace$)*"

// import
const re_import = "($whitespace$)*import($whitespace$)*"

// in
const re_in = "($whitespace$)*in($whitespace$)*"

// instanceof
const re_instanceof = "($whitespace$)*instanceof($whitespace$)*"

// typeof
const re_typeof = "($whitespace$)*typeof($whitespace$)*"

// void
const re_void = "($whitespace$)*void($whitespace$)*"

// delete
const re_delete = "($whitespace$)*delete($whitespace$)*"

// var languageSpecs = []string{
// 	re_num_10,
// 	re_num_binary2,
// 	re_num_16,
// 	re_identfier,
// 	re_double_string,
// 	re_single_string,
// 	re_var,
// 	re_let,
// 	re_const,
// 	re_true,
// 	re_false,
// 	re_undefined,
// 	re_null,
// 	re_if,
// 	re_else,
// 	re_elseif,
// 	re_for,
// 	re_while,
// 	re_do,
// 	re_single_row_comment,
// 	re_left_bracket,
// 	re_right_bracket,
// 	re_left_bracel,
// 	re_right_bracel,
// 	re_left_parenl,
// 	re_right_parenl,
// 	re_comma,
// 	re_semi,
// 	re_colon,
// 	re_dot,
// 	re_question,
// 	re_question_dot,
// 	re_arrow,
// 	re_ellipsis,
// 	re_equal,
// 	re_equality,
// 	re_bitwiseor,
// 	re_bitwisexor,
// 	re_bitwiseand,
// 	re_logicor,
// 	re_logicand,
// 	re_plus,
// 	re_min,
// 	re_modulo,
// 	re_bitleftshift,
// 	re_bitrightshift,
// 	re_bitrightshift3,
// 	re_break,
// 	re_case,
// 	re_catch,
// 	re_continue,
// 	re_default,
// 	re_finally,
// 	re_function,
// 	re_return,
// 	re_switch,
// 	re_throw,
// 	re_try,
// 	re_with,
// 	re_new,
// 	re_this,
// 	re_super,
// 	re_class,
// 	re_extends,
// 	re_export,
// 	re_import,
// 	re_in,
// 	re_instanceof,
// 	re_typeof,
// 	re_void,
// 	re_delete,
// }

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

var Num10DFA = newDFA(re_num_10)
var NumB2DFA = newDFA(re_num_binary2)
var Num16DFA = newDFA(re_num_16)
var FloatDFA = newDFA(re_num_float)
var IdentfierDFA = newDFA(re_identfier)
var StringDoubleDFA = newDFA(re_double_string)
var StringSingleDFA = newDFA(re_single_string)
var VarDFA = newDFA(re_var)
var LetDFA = newDFA(re_let)
var ConstDFA = newDFA(re_const)
var TrueDFA = newDFA(re_true)
var FalseDFA = newDFA(re_false)
var UndefinedDFA = newDFA(re_undefined)
var NullDFA = newDFA(re_null)
var IfDFA = newDFA(re_if)
var ElseDFA = newDFA(re_else)
var ElseIfDFA = newDFA(re_elseif)
var ForDFA = newDFA(re_for)
var WhileDFA = newDFA(re_while)
var DoDFA = newDFA(re_do)
var BracketLDFA = newDFA(re_left_bracket)
var BracketRDFA = newDFA(re_right_bracket)
var BracelLDFA = newDFA(re_left_bracel)
var BracelRDFA = newDFA(re_right_bracel)
var ParenlLDFA = newDFA(re_left_parenl)
var ParenlRDFA = newDFA(re_right_parenl)
var CommaDFA = newDFA(re_comma)
var SemiDFA = newDFA(re_semi)
var ColonDFA = newDFA(re_colon)
var DotDFA = newDFA(re_dot)
var QuestionDFA = newDFA(re_question)
var QuestionDotDFA = newDFA(re_question_dot)
var ArrowDFA = newDFA(re_arrow)
var EllipsisDFA = newDFA(re_ellipsis)
var EqualDFA = newDFA(re_equal)
var GreaterDFA = newDFA(re_greater)
var LessDFA = newDFA(re_less)
var EqualityDFA = newDFA(re_equality)
var BitwiseorDFA = newDFA(re_bitwiseor)
var BitwisexorDFA = newDFA(re_bitwisexor)
var BitwiseandDFA = newDFA(re_bitwiseand)
var LogicorDFA = newDFA(re_logicor)
var LogicandDFA = newDFA(re_logicand)
var PlusDFA = newDFA(re_plus)
var MinDFA = newDFA(re_min)
var ModuloDFA = newDFA(re_modulo)
var BitleftshiftDFA = newDFA(re_bitleftshift)
var BitrightshiftDFA = newDFA(re_bitrightshift)
var Bitrightshift3DFA = newDFA(re_bitrightshift3)
var BreakDFA = newDFA(re_break)
var CaseDFA = newDFA(re_case)
var CatchDFA = newDFA(re_catch)
var ContinueDFA = newDFA(re_continue)
var DefaultDFA = newDFA(re_default)
var FinallyDFA = newDFA(re_finally)
var FunctionDFA = newDFA(re_function)
var ReturnDFA = newDFA(re_return)
var SwitchDFA = newDFA(re_switch)
var ThrowDFA = newDFA(re_throw)
var TryDFA = newDFA(re_try)
var WithDFA = newDFA(re_with)
var NewDFA = newDFA(re_new)
var ThisDFA = newDFA(re_this)
var SuperDFA = newDFA(re_super)
var ClassDFA = newDFA(re_class)
var ExtendsDFA = newDFA(re_extends)
var ExportDFA = newDFA(re_export)
var ImportDFA = newDFA(re_import)
var InDFA = newDFA(re_in)
var InstanceofDFA = newDFA(re_instanceof)
var TypeofDFA = newDFA(re_typeof)
var VoidDFA = newDFA(re_void)
var DeleteDFA = newDFA(re_delete)
var SingleRowComment = newDFA(re_single_row_comment)

package main

import (
	"fmt"
	"jpg/languagespec"
)

func main() {
	// codeString := "20 30 40"
	// reader := reader.New(reader.TextMode, codeString)
	// lexer := lexer.New(reader)

	// token := lexer.NextToken()

	// fmt.Println("token", token.Literal, token.Type)

	// t2 := lexer.NextToken()

	// fmt.Println("token", t2.Literal, t2.Type)

	fmt.Println(languagespec.CheckIsNormalNum("11"))
}

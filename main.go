package main

import (
	"fmt"
	"jpg/lexer"
	"jpg/reader"
)

func main() {
	codeString := "sjdsjkdjskdjskdjkskdsjdjsjk"
	reader := reader.New(reader.TextMode, codeString)
	lexer := lexer.New(reader)

	token := lexer.NextToken()

	fmt.Println("token", token.Literal, token.Type)
}

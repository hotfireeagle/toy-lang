package main

import (
	"fmt"
	"jpg/lexer"
	"jpg/reader"
)

func main() {
	// codeString := "sjdsjkdjskdjskdjkskdsjdjsjk"
	filePath := "/Users/smallhai/learn/gitRepo/jpg/test.js"
	reader := reader.New(reader.FileMode, filePath)
	lexer := lexer.New(reader)

	token := lexer.NextToken()

	for !token.IsEof() {
		fmt.Println("token", token.Literal, token.Type)
		token = lexer.NextToken()
	}
}

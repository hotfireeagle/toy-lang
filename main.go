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

	token, end := lexer.NextToken()

	for !end {
		fmt.Println("token", token.Literal, token.Type)
		token, end = lexer.NextToken()
	}

	fmt.Println("token", token.Literal, token.Type)
}

package main

import (
	"fmt"
	"toy/lexer"
	"toy/reader"
)

func main() {
	filePath := "/Users/smallhai/learn/gitRepo/toy/test.js"
	r := reader.New(reader.FileMode, filePath)
	l := lexer.New(r)
	c := 1
	for tok := l.NextToken(); !tok.IsEof(); {
		c++
		fmt.Println(tok)
		tok = l.NextToken()
	}
	fmt.Println(c)

	// u, err := user.Current()

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Hello %s, Try this!\n", u.Username)

	// repl.Start(os.Stdin, os.Stdout)
}

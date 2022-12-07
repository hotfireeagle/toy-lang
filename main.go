package main

import (
	"fmt"
	"jpg/languagespec"
	"jpg/lexer"
	"jpg/reader"
)

func main() {
	fmt.Println(languagespec.CheckIsString(`"fuck"`))
	filePath := "/Users/smallhai/learn/gitRepo/jpg/test.js"
	r := reader.New(reader.FileMode, filePath)
	l := lexer.New(r)

	fmt.Println(l.NextToken())
	fmt.Println(l.NextToken())
	fmt.Println(l.NextToken())
	fmt.Println(l.NextToken())
	// fmt.Println(l.NextToken())
	// fmt.Println(l.NextToken())
	// fmt.Println(l.NextToken())
	// fmt.Println(l.NextToken())
	// for tok := l.NextToken(); !tok.IsEof(); {
	// 	fmt.Println(tok)
	// }

	// u, err := user.Current()

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Hello %s, Try this!\n", u.Username)

	// repl.Start(os.Stdin, os.Stdout)
}

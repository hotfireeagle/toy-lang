package repl

import (
	"bufio"
	"fmt"
	"io"
	"toy/lexer"
	"toy/reader"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lr := reader.New(reader.TextMode, line)
		l := lexer.New(lr)

		for tok := l.NextToken(); !tok.IsEof(); tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}

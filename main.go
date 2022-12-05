package main

import (
	"fmt"
	"jpg/repl"
	"os"
	"os/user"
)

func main() {
	u, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s, Try this!\n", u.Username)

	repl.Start(os.Stdin, os.Stdout)
}

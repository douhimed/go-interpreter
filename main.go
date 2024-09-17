package main

import (
	"fmt"
	"go-interpreter/repl"
	"os"
	"os/user"
)

/**
*
* 1. The lexer
* 2. The parser
* 3. AST (the Abstract Syntax Tree)
* 4. The internal Object system
* 5. The evaluator
*
 */

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s into Monkey prog language! \n", user.Username)
	fmt.Printf("Type a commands \n\n")
	repl.Start(os.Stdin, os.Stdout)
}

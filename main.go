package main

import (
	"Dumbledore/lexer"
	"Dumbledore/repl"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the Dumbledore Interpreter...")
	fmt.Println("It has known magic.")
	fmt.Println("-------------------\n")
	var lex *lexer.Lexer
	if len(os.Args) == 1 {
		repl.Run()

	} else {
		lex = lexer.NewLexerFromFile(os.Args[1])
		lex.ReadTokens()
	}

}

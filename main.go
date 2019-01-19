package main

import (
	"Dumbledore/lexer"
	"fmt"
	"os"
)

func main() {
	fmt.Println("It has known magic.")
	fmt.Println("-------------------\n")
	var lex *lexer.Lexer
	if len(os.Args) == 1 {
		lex = lexer.NewLexerFromString("var x = blah")
	} else {
		lex = lexer.NewLexerFromFile(os.Args[1])
	}

	lex.ReadTokens()
}

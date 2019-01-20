package main

import (
	"Dumbledore/lexer"
	"Dumbledore/parser"
	"Dumbledore/repl"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		repl.Run()
	} else {
		lex := lexer.NewLexerFromFile(os.Args[1])
		p := parser.New(lex)
		program := p.ParseProgram()
		for _, statement := range program.Statements {
			fmt.Println(statement.ToString())
		}

	}

}

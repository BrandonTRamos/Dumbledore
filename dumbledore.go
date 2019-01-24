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
		if len(p.Errors) > 0 {
			for _, err := range p.Errors {
				fmt.Println(err)
			}
			os.Exit(1)
		}
		for _, statement := range program.Statements {
			fmt.Println(statement.ToString())
		}

	}

}

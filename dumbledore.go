package main

import (
	"Dumbledore/lexer"
	"Dumbledore/parser"
	"Dumbledore/repl"
	"fmt"
	"os"
)

func main() {
	switch len(os.Args) {
	case 1:
		repl.Run()
	case 2:
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
	case 3:
		flag := os.Args[2]
		switch flag {
		case "-l":
			lex := lexer.NewLexerFromFile(os.Args[1])
			lex.ReadTokens()
		case "-p":
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
		default:
			fmt.Println("Not a recognized command.")

		}

	}

}

package main

import (
	"Dumbledore/lexer"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the Dumbledore Interpreter...")
	fmt.Println("It has known magic.")
	fmt.Println("-------------------\n")
	var lex *lexer.Lexer
	if len(os.Args) == 1 {
		repl()

	} else {
		lex = lexer.NewLexerFromFile(os.Args[1])
		lex.ReadTokens()
	}

}

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(">>> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		input := scanner.Text()
		lex := lexer.NewLexerFromString(input)
		fmt.Println("")
		lex.ReadTokens()
		fmt.Println("")
	}
}

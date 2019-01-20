package repl

import (
	"Dumbledore/lexer"
	"bufio"
	"fmt"
	"os"
)

func Run() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Dumbledore Interpreter...")
	fmt.Println("It has known magic.")
	fmt.Println("-------------------\n")
	for {
		fmt.Printf(">>> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		input := scanner.Bytes()
		inputStr := string(input)
		if inputStr == "exit" || inputStr == "quit" || inputStr == "exit()" || inputStr == "quit()" {
			os.Exit(0)
		}
		lex := lexer.NewLexerFromString(input)
		fmt.Println("")
		lex.ReadTokens()
		fmt.Println("")
	}
}

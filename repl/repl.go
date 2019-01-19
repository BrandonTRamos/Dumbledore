package repl

import (
	"Dumbledore/lexer"
	"bufio"
	"fmt"
	"os"
)

func Run() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(">>> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		input := scanner.Text()
		if input == "exit" || input == "quit" || input == "exit()" || input == "quit()" {
			os.Exit(0)
		}
		lex := lexer.NewLexerFromString(input)
		fmt.Println("")
		lex.ReadTokens()
		fmt.Println("")
	}
}

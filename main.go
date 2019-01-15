package main

import (
	"os"
  	"fmt"
  	"Dumbledore/lexer"
)

func main(){
  fmt.Println("It has known magic.")
  fmt.Println("-------------------\n\n")
  var lex *lexer.Lexer
  if len(os.Args)==1{
  		lex=lexer.NewLexerFromString("This is a test.=+_")
  	} else {
  		lex=lexer.NewLexerFromFile(os.Args[1])
  	}
  
  lex.ReadTokens()

}

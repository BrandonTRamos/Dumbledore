package lexer
import (

	// "Dumbledore/token"
	"fmt"
	"io/ioutil"

)

type Lexer struct {
	input string
	currentPosition int
	nextPosition int
	ch byte
}


func NewLexerFromString (input string) *Lexer{
	return &Lexer {input:input}
}

func NewLexerFromFile(fileName string) *Lexer {
	bytes,err := ioutil.ReadFile(fileName)
	if (err!= nil){
		fmt.Println(err)
	}

	return &Lexer{input: string(bytes)}
}

func (lexer *Lexer) HasNext() bool{
		if lexer.nextPosition >= len(lexer.input){
		lexer.ch =0
		return false;
	}

	return true;

}

func (lexer *Lexer) ReadChar() {
	if lexer.nextPosition >= len(lexer.input){
		lexer.ch =0
	}else{
		lexer.ch=lexer.input[lexer.nextPosition]
	}
	lexer.currentPosition =lexer.nextPosition 	
	lexer.nextPosition+=1
	fmt.Println(string(lexer.ch))
}


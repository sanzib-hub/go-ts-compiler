package main

import (
	"fmt"
	"go-ts-compiler/interpreter"
	"go-ts-compiler/lexer"
	"go-ts-compiler/parser"
	"log"
	"os"
)

func main() {
	source, err := os.ReadFile("examples/basic.ts")
	if err != nil {
		log.Fatal(err)
	}

	lex := lexer.New(string(source))
	var tokens []lexer.Token
	for tok := lex.NextToken(); tok.Type != lexer.EOF; tok = lex.NextToken() {
		tokens = append(tokens, tok)
	}

	p := parser.New(tokens)
	env := interpreter.NewEnvironment()

	for p.Peek().Type != "EOF" {
		stmt := p.ParseStatement()
		value := interpreter.Eval(stmt, env)
		fmt.Println("Result:", value)
	}
}

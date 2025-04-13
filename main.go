package main

import (
	"fmt"
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

	fmt.Println("Tokens:")
	for _, tok := range tokens {
		fmt.Printf("%+v\n", tok)
	}

	p := parser.New(tokens)

	fmt.Println("Parsed Statements:")
	for stmt := p.ParseStatement(); stmt != nil; stmt = p.ParseStatement() {
		fmt.Println(stmt.String())
	}
	
}

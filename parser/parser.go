package parser

import (
	"fmt"
	"go-ts-compiler/lexer"
	"log"
)

// Node is the base interface for all AST nodes
type Node interface {
	String() string
}

// ValueNode represents simple values like integers or identifiers
type ValueNode struct {
	Token lexer.Token
}

func (v *ValueNode) String() string {
	return v.Token.Literal
}

// ExpressionNode represents an expression like `5 + 3`
type ExpressionNode struct {
	Left     Node
	Operator string
	Right    Node
}

func (e *ExpressionNode) String() string {
	return fmt.Sprintf("(%s %s %s)", e.Left.String(), e.Operator, e.Right.String())
}

// StatementNode represents something like: let x = expression;
type StatementNode struct {
	Variable string
	Value    Node
}

func (s *StatementNode) String() string {
	return fmt.Sprintf("let %s = %s;", s.Variable, s.Value.String())
}

type Parser struct {
	tokens  []lexer.Token
	current int
}

func New(tokens []lexer.Token) *Parser {
	return &Parser{tokens: tokens, current: 0}
}

func (p *Parser) Peek() lexer.Token {
	if p.current >= len(p.tokens) {
		return lexer.Token{Type: "EOF", Literal: ""}
	}
	return p.tokens[p.current]
}

func (p *Parser) advance() lexer.Token {
	tok := p.Peek()
	p.current++
	return tok
}
func (p *Parser) ParseStatement() Node {
	tok := p.Peek()
	if tok.Type == "EOF" {
		return nil
	}
	// expect "let"
	tok = p.advance()
	fmt.Printf("DEBUG: Current token at start of ParseStatement: %+v\n", tok)

	if tok.Type != "LET" {
		log.Fatalf("Expected 'let', got %s", tok.Type)
	}

	// variable name
	varName := p.advance()
	if varName.Type != "IDENT" {
		log.Fatalf("Expected 'identifier', got %s", varName.Type)
	}


	// expect "="
	eq := p.advance()

	if eq.Type != "ASSIGN" {
		log.Fatalf("Expected '=', got %s", eq.Type)
	}

	right:= p.parseExpression()

	// expression
	semi:= p.advance()
	if semi.Type != "SEMICOLON" {
		log.Fatalf("Expected ';', got %s", semi.Type)
	}
	return &StatementNode{
		Variable: varName.Literal,
		Value:    right,
	}
}
func (p *Parser) parseExpression() Node {
	left := p.advance()

	// Check if the next token is an operator
	if p.Peek().Type == "PLUS" || p.Peek().Type == "MINUS" {
		operator := p.advance()

		// Parse the right-hand side
		right := p.advance()

		return &ExpressionNode{
			Left:     &ValueNode{Token: left},
			Operator: operator.Literal,
			Right:    &ValueNode{Token: right},
		}
	}

	return &ValueNode{Token: left} // Return just the number or identifier
}
func (p *Parser) HasMoreTokens() bool {
	return p.current < len(p.tokens)
}

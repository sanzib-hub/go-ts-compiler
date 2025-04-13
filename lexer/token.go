package lexer

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT" 
	INT   = "INT"

	ASSIGN    = "ASSIGN"
	PLUS      = "PLUS"
	MINUS     = "MINUS"
	LPAREN    = "LPAREN"
	RPAREN    = "RPAREN"
	SEMICOLON = "SEMICOLON"

	FUNCTION = "FUNCTION"
	LET      = "LET"
	STRING   = "STRING"
	CONSOLE  = "CONSOLE"
	DOT      = "DOT"
	LOG      = "LOG"
)

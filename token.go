package token

type TokenType string

type Token struct {
Type TokenType
Literal String
}

const (
	ILLEGAL = "ILLEGAL"
	EOF		= "EOF"

	// Indentifiers and literals
	IDENT = "IDENT" // add, foobar, x, y
	INT	  = "INT"   // 1234567890

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA	  = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords

	FUNCTION = "FUNCTION"
	LET		 = "LET
)

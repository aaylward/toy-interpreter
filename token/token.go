package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// Identifiers and literals
	IDENT = "IDENT"
	INT = "INT"

	// Operators
	ASSIGN = "="
	NOT = "!"
	EQUALS = "=="
	NOT_EQUALS = "!="
	GREATER = ">"
	GREATER_EQUALS = ">="
	LESSER = "<"
	LESSER_EQUALS = "<="
	PLUS = "+"
	MINUS = "-"
	MULTIPLY = "*"
	DIVIDE = "/"

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)

type Token struct {
	Type TokenType
	Literal string
	Filename string
	LineNumber int
	ColumnNumber int
}

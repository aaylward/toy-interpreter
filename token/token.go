package token

type TokenType string

const (
	UNSET   = "UNSET"
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN              = "="
	NOT                 = "!"
	EQUALS              = "=="
	NOT_EQUALS          = "!="
	GREATER_THAN        = ">"
	GREATER_THAN_EQUALS = ">="
	LESS_THAN           = "<"
	LESS_THAN_EQUALS    = "<="
	PLUS                = "+"
	MINUS               = "-"
	MULTIPLY            = "*"
	DIVIDE              = "/"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	FUNCTION = "FUNCTION"
	LET      = "LET"
	RETURN   = "RETURN"
)

type Token struct {
	Type     TokenType
	Literal  string
	Filename string
	Line     int
	Column   int
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
}

func LookupIdent(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}
	return IDENT
}
